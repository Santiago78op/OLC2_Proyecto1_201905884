package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"main.go/repl"
	"main.go/value"
	"strconv"
	"strings"
)

// Interpreter implementa el visitor para evaluar el AST
type Interpreter struct {
	*BaseVLangGrammarVisitor
	Context        *repl.ReplContext
	ScopeTrace     *repl.ScopeTrace
	CallStack      *CallStack
	shouldReturn   bool
	shouldBreak    bool
	shouldContinue bool
	returnValue    value.IVOR
}

// CallStack maneja las llamadas de función y control de flujo
type CallStack struct {
	frames []CallFrame
}

type CallFrame struct {
	FunctionName string
	CanBreak     bool
	CanContinue  bool
	CanReturn    bool
}

func NewCallStack() *CallStack {
	return &CallStack{
		frames: make([]CallFrame, 0),
	}
}

func (cs *CallStack) PushFrame(functionName string, canBreak, canContinue, canReturn bool) {
	frame := CallFrame{
		FunctionName: functionName,
		CanBreak:     canBreak,
		CanContinue:  canContinue,
		CanReturn:    canReturn,
	}
	cs.frames = append(cs.frames, frame)
}

func (cs *CallStack) PopFrame() {
	if len(cs.frames) > 0 {
		cs.frames = cs.frames[:len(cs.frames)-1]
	}
}

func (cs *CallStack) CanBreak() bool {
	for i := len(cs.frames) - 1; i >= 0; i-- {
		if cs.frames[i].CanBreak {
			return true
		}
	}
	return false
}

func (cs *CallStack) CanContinue() bool {
	for i := len(cs.frames) - 1; i >= 0; i-- {
		if cs.frames[i].CanContinue {
			return true
		}
	}
	return false
}

func (cs *CallStack) CanReturn() bool {
	for i := len(cs.frames) - 1; i >= 0; i-- {
		if cs.frames[i].CanReturn {
			return true
		}
	}
	return false
}

// NewInterpreter crea una nueva instancia del intérprete
func NewInterpreter(context *repl.ReplContext) *Interpreter {
	return &Interpreter{
		BaseVLangGrammarVisitor: &BaseVLangGrammarVisitor{},
		Context:                 context,
		ScopeTrace:              context.ScopeTrace,
		CallStack:               NewCallStack(),
		shouldReturn:            false,
		shouldBreak:             false,
		shouldContinue:          false,
		returnValue:             nil,
	}
}

// VisitProgram visita el nodo del programa principal
func (i *Interpreter) VisitProgram(ctx *ProgramContext) interface{} {
	// Primero visitar todas las declaraciones de struct y funciones
	for _, child := range ctx.GetChildren() {
		switch child := child.(type) {
		case *StructDeclContext:
			i.Visit(child)
		case *FunctionDeclContext:
			i.Visit(child)
		}
	}

	// Luego ejecutar la función main
	if mainFunc := ctx.MainFunction(); mainFunc != nil {
		i.CallStack.PushFrame("main", false, false, true)
		i.ScopeTrace.PushScope("main")
		result := i.Visit(mainFunc)
		i.ScopeTrace.PopScope()
		i.CallStack.PopFrame()
		return result
	}

	return nil
}

// VisitMainFunction visita la función main
func (i *Interpreter) VisitMainFunction(ctx *MainFunctionContext) interface{} {
	return i.Visit(ctx.Block())
}

// VisitStructDecl visita una declaración de struct
func (i *Interpreter) VisitStructDecl(ctx *StructDeclContext) interface{} {
	structName := ctx.ID().GetText()
	fields := make(map[string]string)

	// Recolectar campos del struct
	for _, fieldCtx := range ctx.AllStructField() {
		fieldName := fieldCtx.ID().GetText()
		fieldType := fieldCtx.DataType().GetText()
		fields[fieldName] = fieldType
	}

	// Registrar el struct en la tabla de símbolos
	err := i.ScopeTrace.SetStruct(structName, fields,
		ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	if err != nil {
		i.Context.ErrorTable.NewSemanticError(ctx.GetStart(), err.Error())
	}

	return nil
}

// VisitFunctionDecl visita una declaración de función
func (i *Interpreter) VisitFunctionDecl(ctx *FunctionDeclContext) interface{} {
	funcName := ctx.ID().GetText()
	var parameters []repl.Parameter

	// Procesar parámetros
	if paramList := ctx.ParameterList(); paramList != nil {
		for _, paramCtx := range paramList.AllParameter() {
			paramName := paramCtx.ID().GetText()
			paramType := paramCtx.DataType().GetText()
			parameters = append(parameters, repl.Parameter{
				Name: paramName,
				Type: paramType,
			})
		}
	}

	// Determinar tipo de retorno
	returnType := "void"
	if ctx.DataType() != nil {
		returnType = ctx.DataType().GetText()
	}

	// Registrar función en la tabla de símbolos
	err := i.ScopeTrace.SetFunction(funcName, parameters, returnType,
		ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	if err != nil {
		i.Context.ErrorTable.NewSemanticError(ctx.GetStart(), err.Error())
	}

	return nil
}

// VisitBlock visita un bloque de código
func (i *Interpreter) VisitBlock(ctx *BlockContext) interface{} {
	i.ScopeTrace.PushScope("block")
	defer i.ScopeTrace.PopScope()

	for _, stmtCtx := range ctx.AllStatement() {
		if i.shouldReturn || i.shouldBreak || i.shouldContinue {
			break
		}
		i.Visit(stmtCtx)
	}

	return nil
}

// VisitExplicitVarDecl visita una declaración explícita de variable
func (i *Interpreter) VisitExplicitVarDecl(ctx *ExplicitVarDeclContext) interface{} {
	varName := ctx.ID().GetText()
	dataType := ctx.DataType().GetText()
	isMutable := ctx.MUT() != nil

	var val value.IVOR
	if ctx.Expression() != nil {
		val = i.Visit(ctx.Expression()).(value.IVOR)
		// Verificar compatibilidad de tipos
		if !i.isTypeCompatible(dataType, val.Type()) {
			i.Context.ErrorTable.NewSemanticError(ctx.GetStart(),
				fmt.Sprintf("Tipo incompatible: esperado %s, recibido %s", dataType, val.Type()))
			return nil
		}
		// Convertir si es necesario
		val = i.convertType(val, dataType)
	} else {
		val = value.GetDefaultValue(dataType)
	}

	err := i.ScopeTrace.SetVariable(varName, dataType, val,
		ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), isMutable)
	if err != nil {
		i.Context.ErrorTable.NewSemanticError(ctx.GetStart(), err.Error())
	}

	return nil
}

// VisitInferredVarDecl visita una declaración con inferencia de tipo
func (i *Interpreter) VisitInferredVarDecl(ctx *InferredVarDeclContext) interface{} {
	varName := ctx.ID().GetText()
	val := i.Visit(ctx.Expression()).(value.IVOR)
	dataType := val.Type()

	err := i.ScopeTrace.SetVariable(varName, dataType, val,
		ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), true)
	if err != nil {
		i.Context.ErrorTable.NewSemanticError(ctx.GetStart(), err.Error())
	}

	return nil
}

// VisitSimpleAssignment visita una asignación simple
func (i *Interpreter) VisitSimpleAssignment(ctx *SimpleAssignmentContext) interface{} {
	varName := ctx.ID().GetText()
	val := i.Visit(ctx.Expression()).(value.IVOR)

	err := i.ScopeTrace.UpdateVariable(varName, val)
	if err != nil {
		i.Context.ErrorTable.NewSemanticError(ctx.GetStart(), err.Error())
	}

	return nil
}

// VisitPlusAssignment visita una asignación con suma
func (i *Interpreter) VisitPlusAssignment(ctx *PlusAssignmentContext) interface{} {
	varName := ctx.ID().GetText()
	symbol, err := i.ScopeTrace.GetVariable(varName)
	if err != nil {
		i.Context.ErrorTable.NewSemanticError(ctx.GetStart(), err.Error())
		return nil
	}

	newVal := i.Visit(ctx.Expression()).(value.IVOR)
	result := i.performArithmetic(symbol.Value, newVal, "+")

	err = i.ScopeTrace.UpdateVariable(varName, result)
	if err != nil {
		i.Context.ErrorTable.NewSemanticError(ctx.GetStart(), err.Error())
	}

	return nil
}

// VisitMinusAssignment visita una asignación con resta
func (i *Interpreter) VisitMinusAssignment(ctx *MinusAssignmentContext) interface{} {
	varName := ctx.ID().GetText()
	symbol, err := i.ScopeTrace.GetVariable(varName)
	if err != nil {
		i.Context.ErrorTable.NewSemanticError(ctx.GetStart(), err.Error())
		return nil
	}

	newVal := i.Visit(ctx.Expression()).(value.IVOR)
	result := i.performArithmetic(symbol.Value, newVal, "-")

	err = i.ScopeTrace.UpdateVariable(varName, result)
	if err != nil {
		i.Context.ErrorTable.NewSemanticError(ctx.GetStart(), err.Error())
	}

	return nil
}

// VisitIfStatement visita una sentencia if
func (i *Interpreter) VisitIfStatement(ctx *IfStatementContext) interface{} {
	// Evaluar condición principal
	condition := i.Visit(ctx.Expression()).(value.IVOR)
	if i.isTruthy(condition) {
		return i.Visit(ctx.Block())
	}

	// Evaluar else if
	expressions := ctx.AllExpression()
	blocks := ctx.AllBlock()

	for j := 1; j < len(expressions); j++ {
		condition := i.Visit(expressions[j]).(value.IVOR)
		if i.isTruthy(condition) {
			return i.Visit(blocks[j])
		}
	}

	// Evaluar else
	if len(blocks) > len(expressions) {
		return i.Visit(blocks[len(blocks)-1])
	}

	return nil
}

// VisitForCondition visita un for con solo condición
func (i *Interpreter) VisitForCondition(ctx *ForConditionContext) interface{} {
	i.CallStack.PushFrame("for", true, true, false)
	defer i.CallStack.PopFrame()

	for {
		condition := i.Visit(ctx.Expression()).(value.IVOR)
		if !i.isTruthy(condition) {
			break
		}

		i.Visit(ctx.Block())

		if i.shouldReturn || i.shouldBreak || i.shouldContinue {
			if i.shouldBreak {
				i.shouldBreak = false
			}
			if i.shouldContinue {
				i.shouldContinue = false
				continue
			}
			break
		}
	}

	return nil
}

// VisitForTraditional visita un for tradicional
func (i *Interpreter) VisitForTraditional(ctx *ForTraditionalContext) interface{} {
	i.CallStack.PushFrame("for", true, true, false)
	i.ScopeTrace.PushScope("for")
	defer func() {
		i.ScopeTrace.PopScope()
		i.CallStack.PopFrame()
	}()

	// Inicialización
	varName := ctx.ID().GetText()
	initVal := i.Visit(ctx.Expression(0)).(value.IVOR)
	i.ScopeTrace.SetVariable(varName, initVal.Type(), initVal,
		ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), true)

	// Loop
	for {
		// Verificar condición
		condition := i.Visit(ctx.Expression(1)).(value.IVOR)
		if !i.isTruthy(condition) {
			break
		}

		// Ejecutar cuerpo
		i.Visit(ctx.Block())

		if i.shouldReturn || i.shouldBreak || i.shouldContinue {
			if i.shouldBreak {
				i.shouldBreak = false
			}
			if i.shouldContinue {
				i.shouldContinue = false
				// Ejecutar incremento antes de continuar
				i.Visit(ctx.Expression(2))
				continue
			}
			break
		}

		// Incremento
		i.Visit(ctx.Expression(2))
	}

	return nil
}

// VisitBreakStatement visita una sentencia break
func (i *Interpreter) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	if !i.CallStack.CanBreak() {
		i.Context.ErrorTable.NewSemanticError(ctx.GetStart(), "break fuera de contexto válido")
		return nil
	}
	i.shouldBreak = true
	return nil
}

// VisitContinueStatement visita una sentencia continue
func (i *Interpreter) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	if !i.CallStack.CanContinue() {
		i.Context.ErrorTable.NewSemanticError(ctx.GetStart(), "continue fuera de contexto válido")
		return nil
	}
	i.shouldContinue = true
	return nil
}

// VisitReturnStatement visita una sentencia return
func (i *Interpreter) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	if !i.CallStack.CanReturn() {
		i.Context.ErrorTable.NewSemanticError(ctx.GetStart(), "return fuera de función")
		return nil
	}

	if ctx.Expression() != nil {
		i.returnValue = i.Visit(ctx.Expression()).(value.IVOR)
	} else {
		i.returnValue = value.NewNilValue()
	}

	i.shouldReturn = true
	return nil
}

// VisitFunctionCall visita una llamada a función
func (i *Interpreter) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	funcName := ctx.ID().GetText()

	// Verificar si es función built-in
	if symbol, exists := i.ScopeTrace.GlobalScope.Symbols[funcName]; exists {
		if symbol.Type == repl.SYMBOL_FUNCTION {
			// Evaluar argumentos
			var args []value.IVOR
			if ctx.ArgumentList() != nil {
				for _, argCtx := range ctx.ArgumentList().AllExpression() {
					arg := i.Visit(argCtx).(value.IVOR)
					args = append(args, arg)
				}
			}

			// Llamar función built-in
			if builtin, ok := symbol.Value.(BuiltinFunction); ok {
				return builtin.Call(args)
			}

			// Llamar función definida por usuario (simplificado)
			i.Context.ErrorTable.NewSemanticError(ctx.GetStart(),
				"Llamadas a funciones definidas por usuario no implementadas aún")
		}
	}

	i.Context.ErrorTable.NewSemanticError(ctx.GetStart(),
		fmt.Sprintf("Función '%s' no definida", funcName))
	return value.NewNilValue()
}

// VisitAddSub visita una expresión de suma/resta
func (i *Interpreter) VisitAddSub(ctx *AddSubContext) interface{} {
	left := i.Visit(ctx.Expression(0)).(value.IVOR)
	right := i.Visit(ctx.Expression(1)).(value.IVOR)
	operator := ctx.GetOp().GetText()

	return i.performArithmetic(left, right, operator)
}

// VisitMulDivMod visita una expresión de multiplicación/división/módulo
func (i *Interpreter) VisitMulDivMod(ctx *MulDivModContext) interface{} {
	left := i.Visit(ctx.Expression(0)).(value.IVOR)
	right := i.Visit(ctx.Expression(1)).(value.IVOR)
	operator := ctx.GetOp().GetText()

	return i.performArithmetic(left, right, operator)
}

// VisitRelational visita una expresión relacional
func (i *Interpreter) VisitRelational(ctx *RelationalContext) interface{} {
	left := i.Visit(ctx.Expression(0)).(value.IVOR)
	right := i.Visit(ctx.Expression(1)).(value.IVOR)
	operator := ctx.GetOp().GetText()

	return i.performComparison(left, right, operator)
}

// VisitEquality visita una expresión de igualdad
func (i *Interpreter) VisitEquality(ctx *EqualityContext) interface{} {
	left := i.Visit(ctx.Expression(0)).(value.IVOR)
	right := i.Visit(ctx.Expression(1)).(value.IVOR)
	operator := ctx.GetOp().GetText()

	switch operator {
	case "==":
		return value.NewBoolValue(i.valuesEqual(left, right))
	case "!=":
		return value.NewBoolValue(!i.valuesEqual(left, right))
	}

	return value.NewBoolValue(false)
}

// VisitLogicalAnd visita una expresión AND lógica
func (i *Interpreter) VisitLogicalAnd(ctx *LogicalAndContext) interface{} {
	left := i.Visit(ctx.Expression(0)).(value.IVOR)

	// Evaluación perezosa - si left es false, no evaluar right
	if !i.isTruthy(left) {
		return value.NewBoolValue(false)
	}

	right := i.Visit(ctx.Expression(1)).(value.IVOR)
	return value.NewBoolValue(i.isTruthy(right))
}

// VisitLogicalOr visita una expresión OR lógica
func (i *Interpreter) VisitLogicalOr(ctx *LogicalOrContext) interface{} {
	left := i.Visit(ctx.Expression(0)).(value.IVOR)

	// Evaluación perezosa - si left es true, no evaluar right
	if i.isTruthy(left) {
		return value.NewBoolValue(true)
	}

	right := i.Visit(ctx.Expression(1)).(value.IVOR)
	return value.NewBoolValue(i.isTruthy(right))
}

// VisitUnaryMinus visita una negación unaria
func (i *Interpreter) VisitUnaryMinus(ctx *UnaryMinusContext) interface{} {
	val := i.Visit(ctx.Expression()).(value.IVOR)

	switch val.Type() {
	case value.IVOR_INT:
		return value.NewIntValue(-val.(*value.IntValue).InternalValue)
	case value.IVOR_FLOAT:
		return value.NewFloatValue(-val.(*value.FloatValue).InternalValue)
	default:
		i.Context.ErrorTable.NewSemanticError(ctx.GetStart(),
			"Operador unario - solo válido para números")
		return value.NewNilValue()
	}
}

// VisitUnaryNot visita una negación lógica
func (i *Interpreter) VisitUnaryNot(ctx *UnaryNotContext) interface{} {
	val := i.Visit(ctx.Expression()).(value.IVOR)
	return value.NewBoolValue(!i.isTruthy(val))
}

// VisitIntLiteral visita un literal entero
func (i *Interpreter) VisitIntLiteral(ctx *IntLiteralContext) interface{} {
	val, err := strconv.Atoi(ctx.INT().GetText())
	if err != nil {
		i.Context.ErrorTable.NewSemanticError(ctx.GetStart(), "Error al convertir entero")
		return value.NewIntValue(0)
	}
	return value.NewIntValue(val)
}

// VisitFloatLiteral visita un literal flotante
func (i *Interpreter) VisitFloatLiteral(ctx *FloatLiteralContext) interface{} {
	val, err := strconv.ParseFloat(ctx.FLOAT().GetText(), 64)
	if err != nil {
		i.Context.ErrorTable.NewSemanticError(ctx.GetStart(), "Error al convertir flotante")
		return value.NewFloatValue(0.0)
	}
	return value.NewFloatValue(val)
}

// VisitStringLiteral visita un literal de cadena
func (i *Interpreter) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	text := ctx.STRING().GetText()
	// Remover comillas y procesar escapes
	text = text[1 : len(text)-1] // Quitar comillas
	text = i.processEscapes(text)
	return value.NewStringValue(text)
}

// VisitBoolLiteral visita un literal booleano
func (i *Interpreter) VisitBoolLiteral(ctx *BoolLiteralContext) interface{} {
	if ctx.TRUE() != nil {
		return value.NewBoolValue(true)
	}
	return value.NewBoolValue(false)
}

// VisitRuneLiteral visita un literal rune
func (i *Interpreter) VisitRuneLiteral(ctx *RuneLiteralContext) interface{} {
	text := ctx.RUNE().GetText()
	// Remover comillas simples
	text = text[1 : len(text)-1]

	if len(text) == 1 {
		return value.NewRuneValue(rune(text[0]))
	} else if len(text) == 2 && text[0] == '\\' {
		// Procesar escape
		switch text[1] {
		case 'n':
			return value.NewRuneValue('\n')
		case 't':
			return value.NewRuneValue('\t')
		case 'r':
			return value.NewRuneValue('\r')
		case '\\':
			return value.NewRuneValue('\\')
		case '\'':
			return value.NewRuneValue('\'')
		default:
			return value.NewRuneValue(rune(text[1]))
		}
	}

	return value.NewRuneValue(0)
}

// VisitNilLiteral visita un literal nil
func (i *Interpreter) VisitNilLiteral(ctx *NilLiteralContext) interface{} {
	return value.NewNilValue()
}

// VisitIdentifier visita un identificador
func (i *Interpreter) VisitIdentifier(ctx *IdentifierContext) interface{} {
	varName := ctx.ID().GetText()
	symbol, err := i.ScopeTrace.GetVariable(varName)
	if err != nil {
		i.Context.ErrorTable.NewSemanticError(ctx.GetStart(), err.Error())
		return value.NewNilValue()
	}
	return symbol.Value
}

// VisitParenExpr visita una expresión entre paréntesis
func (i *Interpreter) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return i.Visit(ctx.Expression())
}

// VisitExpressionStatement visita una expresión como sentencia
func (i *Interpreter) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return i.Visit(ctx.Expression())
}

// VisitArrayLiteral visita un literal de array
func (i *Interpreter) VisitArrayLiteral(ctx *ArrayLiteralContext) interface{} {
	elementType := ctx.DataType().GetText()
	slice := value.NewSliceValue(elementType).(*value.SliceValue)

	// Procesar elementos si existen
	if len(ctx.AllExpression()) > 0 {
		for _, exprCtx := range ctx.AllExpression() {
			element := i.Visit(exprCtx).(value.IVOR)

			// Verificar tipo compatible
			if !i.isTypeCompatible(elementType, element.Type()) {
				i.Context.ErrorTable.NewSemanticError(ctx.GetStart(),
					fmt.Sprintf("Tipo incompatible: esperado %s, recibido %s",
						elementType, element.Type()))
				continue
			}

			slice.Elements = append(slice.Elements, i.convertType(element, elementType))
		}
	}

	return slice
}

// VisitArrayLiteralExpr visita un literal de array como expresión
func (i *Interpreter) VisitArrayLiteralExpr(ctx *ArrayLiteralExprContext) interface{} {
	return i.Visit(ctx.ArrayLiteral())
}

// Funciones auxiliares

func (i *Interpreter) performArithmetic(left, right value.IVOR, operator string) value.IVOR {
	switch operator {
	case "+":
		return i.performAddition(left, right)
	case "-":
		return i.performSubtraction(left, right)
	case "*":
		return i.performMultiplication(left, right)
	case "/":
		return i.performDivision(left, right)
	case "%":
		return i.performModulo(left, right)
	default:
		return value.NewNilValue()
	}
}

func (i *Interpreter) performAddition(left, right value.IVOR) value.IVOR {
	// Suma de enteros
	if left.Type() == value.IVOR_INT && right.Type() == value.IVOR_INT {
		l := left.(*value.IntValue).InternalValue
		r := right.(*value.IntValue).InternalValue
		return value.NewIntValue(l + r)
	}

	// Suma de flotantes o mixto
	if (left.Type() == value.IVOR_INT || left.Type() == value.IVOR_FLOAT) &&
		(right.Type() == value.IVOR_INT || right.Type() == value.IVOR_FLOAT) {

		var l, r float64

		if left.Type() == value.IVOR_INT {
			l = float64(left.(*value.IntValue).InternalValue)
		} else {
			l = left.(*value.FloatValue).InternalValue
		}

		if right.Type() == value.IVOR_INT {
			r = float64(right.(*value.IntValue).InternalValue)
		} else {
			r = right.(*value.FloatValue).InternalValue
		}

		return value.NewFloatValue(l + r)
	}

	// Concatenación de strings
	if left.Type() == value.IVOR_STRING && right.Type() == value.IVOR_STRING {
		l := left.(*value.StringValue).InternalValue
		r := right.(*value.StringValue).InternalValue
		return value.NewStringValue(l + r)
	}

	return value.NewNilValue()
}

func (i *Interpreter) performSubtraction(left, right value.IVOR) value.IVOR {
	if left.Type() == value.IVOR_INT && right.Type() == value.IVOR_INT {
		l := left.(*value.IntValue).InternalValue
		r := right.(*value.IntValue).InternalValue
		return value.NewIntValue(l - r)
	}

	if (left.Type() == value.IVOR_INT || left.Type() == value.IVOR_FLOAT) &&
		(right.Type() == value.IVOR_INT || right.Type() == value.IVOR_FLOAT) {

		var l, r float64

		if left.Type() == value.IVOR_INT {
			l = float64(left.(*value.IntValue).InternalValue)
		} else {
			l = left.(*value.FloatValue).InternalValue
		}

		if right.Type() == value.IVOR_INT {
			r = float64(right.(*value.IntValue).InternalValue)
		} else {
			r = right.(*value.FloatValue).InternalValue
		}

		return value.NewFloatValue(l - r)
	}

	return value.NewNilValue()
}

func (i *Interpreter) performMultiplication(left, right value.IVOR) value.IVOR {
	if left.Type() == value.IVOR_INT && right.Type() == value.IVOR_INT {
		l := left.(*value.IntValue).InternalValue
		r := right.(*value.IntValue).InternalValue
		return value.NewIntValue(l * r)
	}

	if (left.Type() == value.IVOR_INT || left.Type() == value.IVOR_FLOAT) &&
		(right.Type() == value.IVOR_INT || right.Type() == value.IVOR_FLOAT) {

		var l, r float64

		if left.Type() == value.IVOR_INT {
			l = float64(left.(*value.IntValue).InternalValue)
		} else {
			l = left.(*value.FloatValue).InternalValue
		}

		if right.Type() == value.IVOR_INT {
			r = float64(right.(*value.IntValue).InternalValue)
		} else {
			r = right.(*value.FloatValue).InternalValue
		}

		return value.NewFloatValue(l * r)
	}

	return value.NewNilValue()
}

func (i *Interpreter) performDivision(left, right value.IVOR) value.IVOR {
	// Verificar división por cero
	if right.Type() == value.IVOR_INT && right.(*value.IntValue).InternalValue == 0 {
		i.Context.ErrorTable.NewRuntimeError(0, 0, "División por cero")
		return value.NewNilValue()
	}
	if right.Type() == value.IVOR_FLOAT && right.(*value.FloatValue).InternalValue == 0.0 {
		i.Context.ErrorTable.NewRuntimeError(0, 0, "División por cero")
		return value.NewNilValue()
	}

	if left.Type() == value.IVOR_INT && right.Type() == value.IVOR_INT {
		l := left.(*value.IntValue).InternalValue
		r := right.(*value.IntValue).InternalValue
		return value.NewIntValue(l / r) // División entera
	}

	if (left.Type() == value.IVOR_INT || left.Type() == value.IVOR_FLOAT) &&
		(right.Type() == value.IVOR_INT || right.Type() == value.IVOR_FLOAT) {

		var l, r float64

		if left.Type() == value.IVOR_INT {
			l = float64(left.(*value.IntValue).InternalValue)
		} else {
			l = left.(*value.FloatValue).InternalValue
		}

		if right.Type() == value.IVOR_INT {
			r = float64(right.(*value.IntValue).InternalValue)
		} else {
			r = right.(*value.FloatValue).InternalValue
		}

		return value.NewFloatValue(l / r)
	}

	return value.NewNilValue()
}

func (i *Interpreter) performModulo(left, right value.IVOR) value.IVOR {
	// Módulo solo para enteros
	if left.Type() == value.IVOR_INT && right.Type() == value.IVOR_INT {
		l := left.(*value.IntValue).InternalValue
		r := right.(*value.IntValue).InternalValue

		if r == 0 {
			i.Context.ErrorTable.NewRuntimeError(0, 0, "División por cero en operación módulo")
			return value.NewNilValue()
		}

		return value.NewIntValue(l % r)
	}

	return value.NewNilValue()
}

func (i *Interpreter) performComparison(left, right value.IVOR, operator string) value.IVOR {
	// Comparaciones numéricas
	if (left.Type() == value.IVOR_INT || left.Type() == value.IVOR_FLOAT) &&
		(right.Type() == value.IVOR_INT || right.Type() == value.IVOR_FLOAT) {

		var l, r float64

		if left.Type() == value.IVOR_INT {
			l = float64(left.(*value.IntValue).InternalValue)
		} else {
			l = left.(*value.FloatValue).InternalValue
		}

		if right.Type() == value.IVOR_INT {
			r = float64(right.(*value.IntValue).InternalValue)
		} else {
			r = right.(*value.FloatValue).InternalValue
		}

		switch operator {
		case "<":
			return value.NewBoolValue(l < r)
		case "<=":
			return value.NewBoolValue(l <= r)
		case ">":
			return value.NewBoolValue(l > r)
		case ">=":
			return value.NewBoolValue(l >= r)
		}
	}

	// Comparaciones de strings
	if left.Type() == value.IVOR_STRING && right.Type() == value.IVOR_STRING {
		l := left.(*value.StringValue).InternalValue
		r := right.(*value.StringValue).InternalValue

		switch operator {
		case "<":
			return value.NewBoolValue(l < r)
		case "<=":
			return value.NewBoolValue(l <= r)
		case ">":
			return value.NewBoolValue(l > r)
		case ">=":
			return value.NewBoolValue(l >= r)
		}
	}

	return value.NewBoolValue(false)
}

func (i *Interpreter) valuesEqual(left, right value.IVOR) bool {
	if left.Type() != right.Type() {
		// Comparación especial int/float
		if (left.Type() == value.IVOR_INT && right.Type() == value.IVOR_FLOAT) ||
			(left.Type() == value.IVOR_FLOAT && right.Type() == value.IVOR_INT) {

			var l, r float64
			if left.Type() == value.IVOR_INT {
				l = float64(left.(*value.IntValue).InternalValue)
			} else {
				l = left.(*value.FloatValue).InternalValue
			}

			if right.Type() == value.IVOR_INT {
				r = float64(right.(*value.IntValue).InternalValue)
			} else {
				r = right.(*value.FloatValue).InternalValue
			}

			return l == r
		}
		return false
	}

	switch left.Type() {
	case value.IVOR_INT:
		return left.(*value.IntValue).InternalValue == right.(*value.IntValue).InternalValue
	case value.IVOR_FLOAT:
		return left.(*value.FloatValue).InternalValue == right.(*value.FloatValue).InternalValue
	case value.IVOR_STRING:
		return left.(*value.StringValue).InternalValue == right.(*value.StringValue).InternalValue
	case value.IVOR_BOOL:
		return left.(*value.BoolValue).InternalValue == right.(*value.BoolValue).InternalValue
	case value.IVOR_CHARACTER:
		return left.(*value.RuneValue).InternalValue == right.(*value.RuneValue).InternalValue
	case value.IVOR_NIL:
		return true
	default:
		return false
	}
}

func (i *Interpreter) isTruthy(val value.IVOR) bool {
	switch val.Type() {
	case value.IVOR_BOOL:
		return val.(*value.BoolValue).InternalValue
	case value.IVOR_NIL:
		return false
	case value.IVOR_INT:
		return val.(*value.IntValue).InternalValue != 0
	case value.IVOR_FLOAT:
		return val.(*value.FloatValue).InternalValue != 0.0
	case value.IVOR_STRING:
		return val.(*value.StringValue).InternalValue != ""
	default:
		return true
	}
}

func (i *Interpreter) processEscapes(text string) string {
	text = strings.ReplaceAll(text, "\\n", "\n")
	text = strings.ReplaceAll(text, "\\t", "\t")
	text = strings.ReplaceAll(text, "\\r", "\r")
	text = strings.ReplaceAll(text, "\\\"", "\"")
	text = strings.ReplaceAll(text, "\\'", "'")
	text = strings.ReplaceAll(text, "\\\\", "\\")
	return text
}

func (i *Interpreter) isTypeCompatible(expected, actual string) bool {
	if expected == actual {
		return true
	}
	// Permitir conversión automática int <-> float
	if (expected == value.IVOR_INT && actual == value.IVOR_FLOAT) ||
		(expected == value.IVOR_FLOAT && actual == value.IVOR_INT) {
		return true
	}
	return false
}

func (i *Interpreter) convertType(val value.IVOR, targetType string) value.IVOR {
	if val.Type() == targetType {
		return val
	}

	// Conversiones automáticas
	if val.Type() == value.IVOR_INT && targetType == value.IVOR_FLOAT {
		intVal := val.(*value.IntValue).InternalValue
		return value.NewFloatValue(float64(intVal))
	}

	if val.Type() == value.IVOR_FLOAT && targetType == value.IVOR_INT {
		floatVal := val.(*value.FloatValue).InternalValue
		return value.NewIntValue(int(floatVal))
	}

	return val
}
