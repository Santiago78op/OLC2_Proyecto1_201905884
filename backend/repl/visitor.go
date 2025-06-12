package repl

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	compiler "main.go/grammar"
	"main.go/value"
)

/*
ReplVisitor es una estructura que implementa el visitor para el REPL (Read-Eval-Print Loop).
*/
type ReplVisitor struct {
	compiler.BaseVLangGrammarVisitor
	ScopeTrace  *ScopeTrace
	CallStack   *CallStack
	Console     *Console
	ErrorTable  *ErrorTable
	StructNames []string
}

func NewVisitor(dclVisitor *DclVisitor) *ReplVisitor {
	return &ReplVisitor{
		ScopeTrace:  dclVisitor.ScopeTrace,
		ErrorTable:  dclVisitor.ErrorTable,
		StructNames: dclVisitor.StructNames,
		CallStack:   NewCallStack(),
		Console:     NewConsole(),
	}
}

func (v *ReplVisitor) GetReplContext() *ReplContext {
	return &ReplContext{
		Console:    v.Console,
		ScopeTrace: v.ScopeTrace,
		CallStack:  v.CallStack,
		ErrorTable: v.ErrorTable,
	}
}

func (v *ReplVisitor) ValidType(_type string) bool {
	return v.ScopeTrace.GlobalScope.ValidType(_type)
}

func (v *ReplVisitor) Visit(tree antlr.ParseTree) interface{} {
	fmt.Printf("-------------------------------------------\n")
	fmt.Printf("🔹 ReplVisitor.Visit llamado con: %T\n", tree)

	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		fmt.Printf("❌ ERROR NODE en ReplVisitor: %s\n", val.GetText())
		log.Fatal(val.GetText())
		return nil
	default:
		fmt.Printf("🔹 ReplVisitor aceptando tree\n")
		return tree.Accept(v)
	}
}

func (v *ReplVisitor) VisitProgram(ctx *compiler.ProgramContext) interface{} {
	fmt.Printf("🎯 ¡ENTRANDO A ReplVisitor.VisitProgram!\n")
	fmt.Printf("🔹 Número de statements: %d\n", len(ctx.AllStmt()))

	for i, stmt := range ctx.AllStmt() {
		fmt.Printf("🔹 Procesando statement %d: %s\n", i, stmt.GetText())
		v.Visit(stmt)
	}
	return nil
}

func (v *ReplVisitor) VisitStmt(ctx *compiler.StmtContext) interface{} {

	if ctx.Decl_stmt() != nil {
		v.Visit(ctx.Decl_stmt())
	} else if ctx.Assign_stmt() != nil {
		v.Visit(ctx.Assign_stmt())
	} else if ctx.Func_call() != nil {
		v.Visit(ctx.Func_call())
	} else if ctx.If_stmt() != nil {
		v.Visit(ctx.If_stmt())
	} else if ctx.Switch_stmt() != nil {
		v.Visit(ctx.Switch_stmt())
	} else if ctx.For_stmt() != nil {
		v.Visit(ctx.For_stmt())
	} else {
		log.Fatal("Statement not recognized: ", ctx.GetText())
	}

	return nil
}

// En el enunciado no hay constantes, solo variables mut
func isDeclConst(lexval string) bool {
	return lexval == "let"
}

// Ejemplo: Mut variable_1 int = 10
// Ejemplo: Mut variable_2 int
func (v *ReplVisitor) VisitMutVarDecl(ctx *compiler.MutVarDeclContext) interface{} {

	// Si hubiera constantes se validan aquí
	// isConst := isDeclConst(ctx.Var_type().GetText())
	isConst := false

	// Obtenemos el context de la declaración MutVarDecl
	varName := ctx.ID().GetText()
	varType := v.Visit(ctx.Type_()).(string)
	varValue := v.Visit(ctx.Expression()).(value.IVOR)

	// copy object
	if obj, ok := varValue.(*ObjectValue); ok {
		varValue = obj.Copy()
	}

	variable, msg := v.ScopeTrace.AddVariable(varName, varType, varValue, isConst, false, ctx.GetStart())

	// Variable already exists
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
	}

	return nil
}

func (v *ReplVisitor) VisitValueDecl(ctx *compiler.ValueDeclContext) interface{} {

	isConst := isDeclConst(ctx.Var_type().GetText())
	varName := ctx.ID().GetText()
	varValue := v.Visit(ctx.Expression()).(value.IVOR)
	varType := varValue.Type()

	if varType == "[]" {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "No se puede inferir el tipo de un vector vacio '"+varName+"'")
		return nil
	}

	// copy object
	if obj, ok := varValue.(*ObjectValue); ok {
		varValue = obj.Copy()
	}

	variable, msg := v.ScopeTrace.AddVariable(varName, varType, varValue, isConst, false, ctx.GetStart())

	// Variable already exists
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
	}
	return nil
}

// Ejemplo: variable_1 int = 10
func (v *ReplVisitor) VisitVarDecl(ctx *compiler.VarAssDeclContext) interface{} {

	// Si hubiera constantes se validan aquí
	// isConst := isDeclConst(ctx.Var_type().GetText())
	isConst := false

	// Obtenemos el context de la declaración VarAssDecñ
	exprName := ctx.ID().GetText()
	exprType := v.Visit(ctx.Type_()).(string)

	exprValue := v.Visit(ctx.Expression()).(value.IVOR)

	// Validar tipo de expresión
	if obj, ok := exprValue.(*ObjectValue); ok {
		exprValue = obj.Copy()
	}

	variable, msg := v.ScopeTrace.AddVariable(exprName, exprType, exprValue, isConst, false, ctx.GetStart())

	// Si la variable ya existe, se lanza un error
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
	}

	return nil
}

// Faltan Vectores aca Vector Item - vect_expr

// VisitType es el visitor para el tipo de dato en la declaración de variables.
func (v *ReplVisitor) VisitType(ctx *compiler.TypeContext) interface{} {

	// remove white spaces
	_type := ctx.GetText()

	if v.ValidType(_type) {
		return _type
	}

	/*
		if IsVectorType(_type) {
			// remove [ ]
			internType := RemoveBrackets(_type)
			if v.ValidType(internType) {
				return _type
			}

			v.ErrorTable.NewSemanticError(ctx.GetStart(), "El tipo "+internType+" no es valido para un vector")
			return value.IVOR_NIL
		}

		if IsMatrixType(_type) {
			// remove [[]]
			internType := RemoveBrackets(_type)
			if value.IsPrimitiveType(internType) {
				return _type
			}

			v.ErrorTable.NewSemanticError(ctx.GetStart(), "Las matrices solo pueden contener tipos primitivos")
			return value.IVOR_NIL
		}
	*/

	v.ErrorTable.NewSemanticError(ctx.GetStart(), "Tipo "+ctx.GetText()+" no encontrado")
	return value.IVOR_NIL
}

// Falta el visit repeating
// Falta todo de Vectores

func (v *ReplVisitor) VisitDirectAssign(ctx *compiler.AssignmentDeclContext) interface{} {

	varName := v.Visit(ctx.Id_pattern()).(string)
	varValue := v.Visit(ctx.Expression()).(value.IVOR)

	variable := v.ScopeTrace.GetVariable(varName)

	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable "+varName+" no encontrada")
	} else {

		// copy object
		if obj, ok := varValue.(*ObjectValue); ok {
			varValue = obj.Copy()
		}

		canMutate := true

		if v.ScopeTrace.CurrentScope.isStruct {
			canMutate = v.ScopeTrace.IsMutatingEnvironment()
		}

		ok, msg := variable.AssignValue(varValue, canMutate)

		if !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		}
	}

	return nil

}

func (v *ReplVisitor) VisitArithmeticAssign(ctx *compiler.ArgAddAssigDeclContext) interface{} {
	varName := v.Visit(ctx.Id_pattern()).(string)

	variable := v.ScopeTrace.GetVariable(varName)

	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable "+varName+" no encontrada")
	} else {

		leftValue := variable.Value
		rightValue := v.Visit(ctx.Expression()).(value.IVOR)

		op := string(ctx.GetOp().GetText()[0])

		strat, ok := BinaryStrats[op]

		if !ok {
			log.Fatal("Binary operator not found")
		}

		ok, msg, varValue := strat.Validate(leftValue, rightValue)

		if !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
			return nil
		}

		canMutate := true

		if v.ScopeTrace.CurrentScope.isStruct {
			canMutate = v.ScopeTrace.IsMutatingEnvironment()
		}

		ok, msg = variable.AssignValue(varValue, canMutate)

		if !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		}
	}

	return nil
}

// Falta  VisitVectorAssign

// Id parents
func (v *ReplVisitor) VisitIdPattern(ctx *compiler.IdPatternContext) interface{} {
	fmt.Print("El valor de IdPattern es: " + ctx.GetText() + "\n")
	return ctx.GetText()
}

// literales
//
//	Literal Int
func (v *ReplVisitor) VisitIntLiteral(ctx *compiler.IntLiteralContext) interface{} {

	intVal, _ := strconv.Atoi(ctx.GetText())

	return &value.IntValue{
		InternalValue: intVal,
	}

}

// literal Float
func (v *ReplVisitor) VisitFloatLiteral(ctx *compiler.FloatLiteralContext) interface{} {

	floatVal, _ := strconv.ParseFloat(ctx.GetText(), 64)

	return &value.FloatValue{
		InternalValue: floatVal,
	}

}

// literal String
func (v *ReplVisitor) VisitStringLiteral(ctx *compiler.StringLiteralContext) interface{} {

	// remove quotes
	stringVal := ctx.GetText()[1 : len(ctx.GetText())-1]

	// \" \\ \n \r \
	stringVal = strings.ReplaceAll(stringVal, "\\\"", "\"")
	stringVal = strings.ReplaceAll(stringVal, "\\\\", "\\")
	stringVal = strings.ReplaceAll(stringVal, "\\n", "\n")
	stringVal = strings.ReplaceAll(stringVal, "\\r", "\r")

	// Character literal
	if len(stringVal) == 1 {
		return &value.CharacterValue{
			InternalValue: stringVal,
		}
	}

	// String literal
	return &value.StringValue{
		InternalValue: stringVal,
	}

}

// literal Bool
func (v *ReplVisitor) VisitBoolLiteral(ctx *compiler.BoolLiteralContext) interface{} {

	boolVal, _ := strconv.ParseBool(ctx.GetText())

	return &value.BoolValue{
		InternalValue: boolVal,
	}

}

// literal Nil
func (v *ReplVisitor) VisitNilLiteral(ctx *compiler.NilLiteralContext) interface{} {
	return value.DefaultNilValue
}

// literal en Exp
func (v *ReplVisitor) VisitLiteralExpr(ctx *compiler.LiteralExprContext) interface{} {
	fmt.Print("El valor de LiteralExp es: " + ctx.GetText() + "\n")
	return v.Visit(ctx.Literal())
}

func (v *ReplVisitor) VisitIdPatternExpr(ctx *compiler.IdPatternExprContext) interface{} {
	varName := ctx.Id_pattern().GetText()

	variable := v.ScopeTrace.GetVariable(varName)

	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable "+varName+" no encontrada")
		return value.DefaultNilValue
	}

	// ? pointer
	return variable.Value
}

// Expresiones con parentesis
func (v *ReplVisitor) VisitParensExpr(ctx *compiler.ParensExprContext) interface{} {
	fmt.Print("El valor de ParenExp es: " + ctx.GetText() + "\n")
	return v.Visit(ctx.Expression())
}

// Funciones con expresiones
func (v *ReplVisitor) VisitFuncCallExp(ctx *compiler.FuncCallExprContext) interface{} {
	return v.Visit(ctx.Func_call())
}

func (v *ReplVisitor) VisitUnaryExpr(ctx *compiler.UnaryExprContext) interface{} {

	exp := v.Visit(ctx.Expression()).(value.IVOR)

	strat, ok := UnaryStrats[ctx.GetOp().GetText()]

	if !ok {
		log.Fatal("Unary operator not found")
	}

	ok, msg, result := strat.Validate(exp)

	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetOp(), msg)
		return value.DefaultNilValue
	}

	return result

}

func (v *ReplVisitor) VisitBinaryExpr(ctx *compiler.BinaryExprContext) interface{} {

	op := ctx.GetOp().GetText()
	left := v.Visit(ctx.GetLeft()).(value.IVOR)

	earlyCheck, ok := EarlyReturnStrats[op]

	if ok {
		ok, _, result := earlyCheck.Validate(left)

		if ok {
			return result
		}
	}

	//
	right := v.Visit(ctx.GetRight()).(value.IVOR)

	// Si right es un IVOR, lo convertimos a IVOR

	strat, ok := BinaryStrats[op]

	if !ok {
		log.Fatal("Binary operator not found")
	}

	ok, msg, result := strat.Validate(left, right)

	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetOp(), msg)
		return value.DefaultNilValue
	}

	return result
}

func (v *ReplVisitor) VisitIfStmt(ctx *compiler.IfStmtContext) interface{} {

	runChain := true

	for _, ifStmt := range ctx.AllIf_chain() {

		runChain = !v.Visit(ifStmt).(bool)

		if !runChain {
			break
		}
	}

	if runChain && ctx.Else_stmt() != nil {
		v.Visit(ctx.Else_stmt())
	}

	return nil
}

func (v *ReplVisitor) VisitIfChain(ctx *compiler.IfChainContext) interface{} {

	condition := v.Visit(ctx.Expression()).(value.IVOR)

	if condition.Type() != value.IVOR_BOOL {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "La condicion del if debe ser un booleano")
		return false

	}

	if condition.(*value.BoolValue).InternalValue {

		// Push scope
		v.ScopeTrace.PushScope("if")

		for _, stmt := range ctx.AllStmt() {
			v.Visit(stmt)
		}

		// Pop scope
		v.ScopeTrace.PopScope()

		return true
	}

	return false
}

func (v *ReplVisitor) VisitElseStmt(ctx *compiler.ElseStmtContext) interface{} {

	// Push scope
	v.ScopeTrace.PushScope("else")

	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}

	// Pop scope
	v.ScopeTrace.PopScope()

	return nil
}

func (v *ReplVisitor) VisitForStmtCond(ctx *compiler.ForStmtCondContext) interface{} {
	condition := ctx.Expression()

	forItem := &CallStackItem{ReturnValue: value.DefaultNilValue, Type: []string{BreakItem, ContinueItem}}
	v.CallStack.Push(forItem)
	v.ScopeTrace.PushScope("for_cond")

	// Manejo de control de flujo con panic/recover
	defer func() {
		if item, ok := recover().(*CallStackItem); item != nil && ok {

			// Si no es el for actual, propaga el panic hacia arriba
			if item != forItem {
				panic(item)
			}

			// Si es un continue, simplemente dejamos que siga el ciclo
			if item.IsAction(ContinueItem) {
				item.ResetAction()

			}

			// Si es un break, terminamos el for
			if item.IsAction(BreakItem) {
				// terminar el ciclo asd
			}
		}
	}()

	for {
		condValue, ok := v.Visit(condition).(value.IVOR)
		if !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "Error evaluando la condición del for")
			return nil
		}

		if condValue.Type() != value.IVOR_BOOL {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "La condición del for debe ser un booleano")
			return nil
		}

		boolVal := condValue.Value().(bool)
		if !boolVal {
			break
		}

		for _, stmt := range ctx.AllStmt() {
			v.Visit(stmt)
			// no se necesita verificar break, continue ni return porque son manejados automáticamente por el panic/recover
		}
	}

	v.ScopeTrace.PopScope()
	v.CallStack.Clean(forItem)
	return nil
}

func (v *ReplVisitor) VisitForAssCond(ctx *compiler.ForAssCondContext) interface{} {
	initAssign := ctx.Assign_stmt(0)
	condition := ctx.Expression()
	updateAssign := ctx.Assign_stmt(1)

	v.ScopeTrace.PushScope("for_assignamet")
	v.Visit(initAssign)

	forItem := &CallStackItem{ReturnValue: value.DefaultNilValue, Type: []string{BreakItem, ContinueItem}}
	v.CallStack.Push(forItem)

	defer func() {
		if item, ok := recover().(*CallStackItem); item != nil && ok {
			if item != forItem {
				panic(item)
			}

			if item.IsAction(ContinueItem) {
				item.ResetAction()
				// Sigue normalmente en la siguiente iteración
			}

			if item.IsAction(BreakItem) {
				// Finaliza el bucle
			}
		}
	}()

	for {
		condValue, ok := v.Visit(condition).(value.IVOR)
		if !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "Error evaluando la condición del for")
			return nil
		}

		if condValue.Type() != value.IVOR_BOOL {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "La condición del for debe ser un booleano")
			return nil
		}

		boolVal := condValue.Value().(bool)
		if !boolVal {
			break
		}

		for _, stmt := range ctx.AllStmt() {
			v.Visit(stmt)
		}

		v.Visit(updateAssign)
	}

	v.ScopeTrace.PopScope()
	v.CallStack.Clean(forItem)
	return nil
}

func (v *ReplVisitor) VisitFuncCall(ctx *compiler.FuncCallContext) interface{} {

	// find if its a func or constructor of a struct

	canditateName := v.Visit(ctx.Id_pattern()).(string)
	funcObj, msg1 := v.ScopeTrace.GetFunction(canditateName)
	structObj, msg2 := v.ScopeTrace.GlobalScope.GetStruct(canditateName)

	if funcObj == nil && structObj == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg1+msg2)
		return value.DefaultNilValue
	}

	args := make([]*Argument, 0)
	if ctx.Arg_list() != nil {
		// Visualizar lo que tiene Arg_list
		fmt.Printf("🔹 Visitando Arg_list: %s\n", ctx.Arg_list().GetText())
		// Como buscaria Arg_list en el visitor -> Ejemplo VisitArgList(ctx.Arg_list())
		args = v.Visit(ctx.Arg_list()).([]*Argument)
	}

	/*
		// struct has priority over func
		if structObj != nil {
			if IsArgValidForStruct(args) {
				return NewObjectValue(v, canditateName, ctx.Id_pattern().GetStart(), args, false)
			} else {
				v.ErrorTable.NewSemanticError(ctx.GetStart(), "Si bien "+canditateName+" es un struct, no se puede llamar a su constructor con los argumentos especificados. Ni tampoco es una funcion.")
				return value.DefaultNilValue
			}
		}*/

	switch funcObj := funcObj.(type) {
	case *BuiltInFunction:
		returnValue, ok, msg := funcObj.Exec(v.GetReplContext(), args)

		if !ok {

			if msg != "" {
				v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
			}

			return value.DefaultNilValue

		}

		return returnValue

	case *Function:
		funcObj.Exec(v, args, ctx.GetStart())
		return funcObj.ReturnValue

	case *ObjectBuiltInFunction:
		funcObj.Exec(v, args, ctx.GetStart())
		return funcObj.ReturnValue

	default:
		log.Fatal("Function type not found")
	}

	return value.DefaultNilValue
}

func (v *ReplVisitor) VisitArgList(ctx *compiler.ArgListContext) interface{} {

	args := make([]*Argument, 0)

	for _, arg := range ctx.AllFunc_arg() {
		// Visualizar lo que tiene arg
		fmt.Printf("🔹 Visitando FuncArg: %s\n", arg.GetText())
		args = append(args, v.Visit(arg).(*Argument))
	}

	return args

}

func (v *ReplVisitor) VisitFuncArg(ctx *compiler.FuncArgContext) interface{} {

	argName := ""
	passByReference := false

	var argValue value.IVOR = value.DefaultNilValue
	var argVariableRef *Variable = nil

	if ctx.Id_pattern() != nil {
		// Because is a reference to a variable, the treatment is a bit different
		argName = ctx.Id_pattern().GetText()
		argVariableRef = v.ScopeTrace.GetVariable(argName)

		if argVariableRef != nil {
			argValue = argVariableRef.Value
		} else {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable "+argName+" no encontrada")
		}
	} else {
		val := v.Visit(ctx.Expression())
		fmt.Printf("Tipo retornado por v.Visit(ctx.Expression()): %T\n", val)
		ivor, ok := val.(value.IVOR)
		if !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "El argumento no es un valor válido")
			argValue = value.DefaultNilValue
		} else {
			argValue = ivor
		}
	}

	if ctx.ID() != nil {
		argName = ctx.ID().GetText()
	}

	return &Argument{
		Name:            argName,
		Value:           argValue,
		PassByReference: passByReference,
		Token:           ctx.GetStart(),
		VariableRef:     argVariableRef,
	}

}
