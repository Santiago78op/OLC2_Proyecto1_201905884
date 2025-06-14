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
	log.Println("🎯 ¡ENTRANDO A ReplVisitor.VisitProgram!")

	for _, stmt := range ctx.AllStmt() {
		log.Println("🔹 Procesando statement:", stmt.GetText())
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
	} else if ctx.Func_dcl() != nil {
		v.Visit(ctx.Func_dcl())
	} else if ctx.If_stmt() != nil {
		v.Visit(ctx.If_stmt())
	} else if ctx.Switch_stmt() != nil {
		v.Visit(ctx.Switch_stmt())
	} else if ctx.For_stmt() != nil {
		v.Visit(ctx.For_stmt())
	} else if ctx.Strct_dcl() != nil {
		v.Visit(ctx.Strct_dcl())
	} else if ctx.Struct_instantiation() != nil {
		v.Visit(ctx.Struct_instantiation())
	} else {
		log.Println("⚠️ Statement no reconocido:", ctx.GetText())
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

	isConst := false
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

func (v *ReplVisitor) VisitAssignmentDecl(ctx *compiler.AssignmentDeclContext) interface{} {

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

func (v *ReplVisitor) VisitArgAddAssigDecl(ctx *compiler.ArgAddAssigDeclContext) interface{} {
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

// VisitIncredecr maneja las expresiones de incremento y decremento
// Ejemplo: i++ o i-- dentro de una expresión
func (v *ReplVisitor) VisitIncredecr(ctx *compiler.IncredecrContext) interface{} {
	return v.Visit(ctx.Incredecre())
}

// VisitIncremento maneja el incremento (ID++)
// Comportamiento: Post-incremento - retorna el valor actual, luego incrementa
func (v *ReplVisitor) VisitIncremento(ctx *compiler.IncrementoContext) interface{} {
	fmt.Printf("🔹 Visitando Incremento: %s\n", ctx.GetText())

	// Obtener el nombre de la variable
	varName := ctx.ID().GetText()

	// Buscar la variable en el scope
	variable := v.ScopeTrace.GetVariable(varName)
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable '"+varName+"' no encontrada")
		return value.DefaultNilValue
	}

	// Verificar que la variable sea de tipo entero
	if variable.Value.Type() != value.IVOR_INT {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "El operador ++ solo puede aplicarse a variables de tipo int")
		return value.DefaultNilValue
	}

	// Verificar que no sea constante
	if variable.IsConst {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "No se puede incrementar una variable constante")
		return value.DefaultNilValue
	}

	// Obtener el valor actual (para retornarlo - post-incremento)
	currentValue := variable.Value.(*value.IntValue).InternalValue

	// Crear el nuevo valor incrementado
	newValue := &value.IntValue{
		InternalValue: currentValue + 1,
	}

	// Verificar contexto de mutación (para propiedades de struct)
	canMutate := true
	if v.ScopeTrace.CurrentScope.isStruct {
		canMutate = v.ScopeTrace.IsMutatingEnvironment()
	}

	// Asignar el nuevo valor a la variable
	ok, msg := variable.AssignValue(newValue, canMutate)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		return value.DefaultNilValue
	}

	// Retornar el valor original (comportamiento post-incremento)
	return &value.IntValue{
		InternalValue: currentValue,
	}
}

// VisitDecremento maneja el decremento (ID--)
// Comportamiento: Post-decremento - retorna el valor actual, luego decrementa
func (v *ReplVisitor) VisitDecremento(ctx *compiler.DecrementoContext) interface{} {
	fmt.Printf("🔹 Visitando Decremento: %s\n", ctx.GetText())

	// Obtener el nombre de la variable
	varName := ctx.ID().GetText()

	// Buscar la variable en el scope
	variable := v.ScopeTrace.GetVariable(varName)
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable '"+varName+"' no encontrada")
		return value.DefaultNilValue
	}

	// Verificar que la variable sea de tipo entero
	if variable.Value.Type() != value.IVOR_INT {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "El operador -- solo puede aplicarse a variables de tipo int")
		return value.DefaultNilValue
	}

	// Verificar que no sea constante
	if variable.IsConst {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "No se puede decrementar una variable constante")
		return value.DefaultNilValue
	}

	// Obtener el valor actual (para retornarlo - post-decremento)
	currentValue := variable.Value.(*value.IntValue).InternalValue

	// Crear el nuevo valor decrementado
	newValue := &value.IntValue{
		InternalValue: currentValue - 1,
	}

	// Verificar contexto de mutación (para propiedades de struct)
	canMutate := true
	if v.ScopeTrace.CurrentScope.isStruct {
		canMutate = v.ScopeTrace.IsMutatingEnvironment()
	}

	// Asignar el nuevo valor a la variable
	ok, msg := variable.AssignValue(newValue, canMutate)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		return value.DefaultNilValue
	}

	// Retornar el valor original (comportamiento post-decremento)
	return &value.IntValue{
		InternalValue: currentValue,
	}
}

func (v *ReplVisitor) VisitIdPatternExpr(ctx *compiler.IdPatternExprContext) interface{} {
	idParts := strings.Split(ctx.GetText(), ".")
	base := idParts[0]

	variable := v.ScopeTrace.GetVariable(base)
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable "+base+" no encontrada")
		return nil
	}

	val := variable.Value

	for i := 1; i < len(idParts); i++ {
		if structVal, ok := val.(*value.StructValue); ok {
			field := idParts[i]
			val = structVal.Instance.Fields[field]
		} else {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "Acceso inválido a campo "+idParts[i])
			return nil
		}
	}

	return val
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
	// Obtener las tres partes del for: inicialización, condición e incremento
	initAssign := ctx.Assign_stmt()    // i = 0
	condition := ctx.Expression(0)     // i < 5  (primera expresión)
	incrementExpr := ctx.Expression(1) // i++    (segunda expresión)

	// Crear nuevo scope para el for
	v.ScopeTrace.PushScope("for_assignment")

	// Ejecutar la inicialización (i = 0)
	v.Visit(initAssign)

	// Crear item para manejo de break/continue
	forItem := &CallStackItem{
		ReturnValue: value.DefaultNilValue,
		Type:        []string{BreakItem, ContinueItem},
	}
	v.CallStack.Push(forItem)

	// Defer para cleanup y manejo de control de flujo
	defer func() {
		v.ScopeTrace.PopScope()    // Limpiar scope
		v.CallStack.Clean(forItem) // Limpiar call stack

		// Manejo de panic/recover para break/continue
		if item, ok := recover().(*CallStackItem); item != nil && ok {
			// Si no es nuestro forItem, propagar panic hacia arriba
			if item != forItem {
				panic(item)
			}

			// Si es continue, se resetea la acción y continúa
			if item.IsAction(ContinueItem) {
				item.ResetAction()
			}

			// Si es break, simplemente termina (no hace nada más)
			if item.IsAction(BreakItem) {
				// El bucle termina naturalmente
			}
		}
	}()

	// Bucle principal
	for {
		// Evaluar condición (i < 5)
		condValue := v.Visit(condition)
		if condValue == nil {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "Error evaluando la condición del for")
			break
		}

		condIVOR, ok := condValue.(value.IVOR)
		if !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "La condición del for debe evaluar a un valor IVOR")
			break
		}

		// Verificar que la condición sea booleana
		if condIVOR.Type() != value.IVOR_BOOL {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "La condición del for debe ser un booleano")
			break
		}

		// Obtener valor booleano
		boolVal := condIVOR.Value().(bool)
		if !boolVal {
			break // Condición falsa, salir del bucle
		}

		// Ejecutar cuerpo del bucle (println, suma = suma + i, etc.)
		for _, stmt := range ctx.AllStmt() {
			v.Visit(stmt)

			// Verificar si hubo break/continue durante la ejecución del cuerpo
			// Esto se maneja automáticamente por el panic/recover
		}

		// Ejecutar incremento (i++)
		v.Visit(incrementExpr)
	}

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

func (v *ReplVisitor) VisitFuncDecl(ctx *compiler.FuncDeclContext) interface{} {

	if v.ScopeTrace.CurrentScope == v.ScopeTrace.GlobalScope {
		// aready declared by dcl_visitor
		return nil
	}

	if v.ScopeTrace.CurrentScope != v.ScopeTrace.GlobalScope && !v.ScopeTrace.CurrentScope.isStruct {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Las funciones solo pueden ser declaradas en el scope global o en un struct")
	}

	funcName := ctx.ID().GetText()

	params := make([]*Param, 0)

	if ctx.Param_list() != nil {
		params = v.Visit(ctx.Param_list()).([]*Param)
	}

	if len(params) > 0 {

		baseParamType := params[0].ParamType()

		for _, param := range params {
			if param.ParamType() != baseParamType {
				v.ErrorTable.NewSemanticError(param.Token, "Todos los parametros de la funcion deben ser del mismo tipo")
				return nil
			}
		}
	}

	returnType := value.IVOR_NIL
	var returnTypeToken antlr.Token = nil

	if ctx.Type_() != nil {
		returnType = v.Visit(ctx.Type_()).(string)
		returnTypeToken = ctx.Type_().GetStart()
	}

	body := ctx.AllStmt()

	function := &Function{ // pointer ?
		Name:            funcName,
		Param:           params,
		ReturnType:      returnType,
		Body:            body,
		DeclScope:       v.ScopeTrace.CurrentScope,
		ReturnTypeToken: returnTypeToken,
		Token:           ctx.GetStart(),
	}

	ok, msg := v.ScopeTrace.AddFunction(funcName, function)

	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		return nil
	}

	return function
}

func (v *ReplVisitor) VisitParamList(ctx *compiler.ParamListContext) interface{} {

	params := make([]*Param, 0)

	for _, param := range ctx.AllFunc_param() {
		params = append(params, v.Visit(param).(*Param))
	}

	return params
}

func (v *ReplVisitor) VisitFuncParam(ctx *compiler.FuncParamContext) interface{} {

	externName := ""
	innerName := ""

	// at least ID(0) is defined
	// only 1 ID defined
	if ctx.ID() == nil {
		// innerName : type
		// _ : type
		innerName = ctx.ID().GetText()
	} else {
		// externName innerName : type
		externName = ctx.ID().GetText()
	}

	passByReference := false

	paramType := v.Visit(ctx.Type_()).(string)

	return &Param{
		ExternName:      externName,
		InnerName:       innerName,
		PassByReference: passByReference,
		Type:            paramType,
		Token:           ctx.GetStart(),
	}

}

func (v *ReplVisitor) VisitSwitchStmt(ctx *compiler.SwitchStmtContext) interface{} {

	mainValue := v.Visit(ctx.Expression()).(value.IVOR)

	v.ScopeTrace.PushScope("switch")

	// Push break switchItem to call stack [breakable]
	switchItem := &CallStackItem{
		ReturnValue: value.DefaultNilValue,
		Type: []string{
			BreakItem,
		},
	}

	v.CallStack.Push(switchItem)

	// handle break statements from call stack
	defer func() {

		v.ScopeTrace.PopScope()       // pop switch scope
		v.CallStack.Clean(switchItem) // clean item if it's still in call stack

		if item, ok := recover().(*CallStackItem); item != nil && ok {

			// Not a switch item, propagate panic
			if item != switchItem {
				panic(item)
			}

			return // break
		}
	}()

	visited := false

	// evaluate cases
	for _, switchCase := range ctx.AllSwitch_case() {

		caseValue := v.GetCaseValue(switchCase)

		// ? use binary strat
		if caseValue.Type() != mainValue.Type() {
			// warning
			continue
		}

		if caseValue.Value() == mainValue.Value() {
			v.Visit(switchCase)
			visited = true
			break // implicit break
		}

	}

	// evaluate default
	if ctx.Default_case() != nil && !visited {
		v.Visit(ctx.Default_case())
	}

	return nil
}

func (v *ReplVisitor) GetCaseValue(tree antlr.ParseTree) value.IVOR {

	switch val := tree.(type) {
	case *compiler.SwitchCaseContext:
		return v.Visit(val.Expression()).(value.IVOR)
	default:
		return nil
	}

}

func (v *ReplVisitor) VisitSwitchCase(ctx *compiler.SwitchCaseContext) interface{} {

	// * all cases inside switch case will share the same scope

	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}
	return nil
}

func (v *ReplVisitor) VisitDefaultCase(ctx *compiler.DefaultCaseContext) interface{} {
	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}
	return nil
}

/*
func (v *ReplVisitor) VisitForStmt(ctx *compiler.ForStmtContext) interface{} {

	varName := ctx.ID().GetText()
	var iterableItem *VectorValue = DefaultEmptyVectorValue

	if ctx.Range_() != nil {
		rangeItem, ok := v.Visit(ctx.Range_()).(*VectorValue)

		if !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "El valor del rango debe ser un vector")
			return nil
		}

		iterableItem = rangeItem
	}

	if ctx.Expr() != nil {
		iterableValue := v.Visit(ctx.Expr()).(value.IVOR)

		if IsVectorType(iterableValue.Type()) {
			iterableItem = iterableValue.(*VectorValue)
		} else if iterableValue.Type() == value.IVOR_STRING {
			iterableItem = StringToVector(iterableValue.(*value.StringValue))
		} else {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "El valor del rango debe ser un vector o una cadena")
			return nil
		}
	}

	if iterableItem.Size() == 0 {
		return nil
	}

	// Push scope outer scope
	outerForScope := v.ScopeTrace.PushScope("outer_for")

	// create the associated variable to the iterable
	iterableVariable, msg := outerForScope.AddVariable(varName, iterableItem.ItemType, iterableItem.Current(), true, false, ctx.ID().GetSymbol())

	if iterableVariable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		log.Fatal("This should not happen")
		return nil
	}

	// Push forItem to call stack [breakable, continuable]

	forItem := &CallStackItem{
		ReturnValue: value.DefaultNilValue,
		Type: []string{
			BreakItem,
			ContinueItem,
		},
	}

	v.CallStack.Push(forItem)

	// Push inner for scope
	innerForScope := v.ScopeTrace.PushScope("inner_for")

	v.VisitInnerFor(ctx, outerForScope, innerForScope, forItem, iterableItem, iterableVariable)

	iterableItem.Reset()
	v.ScopeTrace.PopScope()    // pop inner for scope
	v.ScopeTrace.PopScope()    // pop outer for scope
	v.CallStack.Clean(forItem) // ? clean item if it's still in call stack

	return nil
}

func (v *ReplVisitor) VisitInnerFor(ctx *compiler.ForStmtContext, outerForScope *BaseScope, innerForScope *BaseScope, forItem *CallStackItem, iterableItem *VectorValue, iterableVariable *Variable) {

	// handle break and continue statements from call stack
	defer func() {

		// reset scope
		innerForScope.Reset()
		if item, ok := recover().(*CallStackItem); item != nil && ok {

			// Not a for item, propagate panic
			if item != forItem {
				panic(item)
			}

			// Continue
			if item.IsAction(ContinueItem) {
				item.ResetAction()                                                                          // reset action, can be used again
				iterableItem.Next()                                                                         // next item
				v.VisitInnerFor(ctx, outerForScope, innerForScope, forItem, iterableItem, iterableVariable) // continue
			}

			// Break
			if item.IsAction(BreakItem) {
				return
			}

		}
	}()

	// iterableItem.Size()
	for iterableItem.CurrentIndex < iterableItem.Size() {

		// update variable value
		iterableVariable.Value = iterableItem.Current()

		for _, stmt := range ctx.AllStmt() {
			v.Visit(stmt)
		}

		iterableItem.Next()
		innerForScope.Reset()
	}
}

func (v *ReplVisitor) VisitNumericRange(ctx *compiler.NumericRangeContext) interface{} {

	leftExpr := v.Visit(ctx.Expr(0)).(value.IVOR)
	rightExpr := v.Visit(ctx.Expr(1)).(value.IVOR)

	if leftExpr.Type() != value.IVOR_INT || rightExpr.Type() != value.IVOR_INT {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Los valores de los rangos deben ser enteros")
		return value.DefaultNilValue
	}

	left := leftExpr.(*value.IntValue).InternalValue
	right := rightExpr.(*value.IntValue).InternalValue

	if left > right {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "El valor izquierdo del rango debe ser menor o igual al valor derecho")
	}

	var values []value.IVOR

	for i := left; i <= right; i++ {
		values = append(values, &value.IntValue{
			InternalValue: i,
		})
	}

	return &VectorValue{
		InternalValue: values,
		CurrentIndex:  0,
		ItemType:      value.IVOR_INT,
	}
}
*/

// Structs
func (v *ReplVisitor) VisitStructDecl(ctx *compiler.StructDeclContext) interface{} {
	if v.ScopeTrace.CurrentScope != v.ScopeTrace.GlobalScope {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Los structs solo pueden ser declaradas en el scope global")
		return nil
	}

	structAdded, msg := v.ScopeTrace.GlobalScope.AddStruct(ctx.ID().GetText(), &Struct{
		Name:   ctx.ID().GetText(),
		Fields: ctx.AllStruct_prop(),
		Token:  ctx.GetStart(),
	})

	if !structAdded {
		v.ErrorTable.NewSemanticError(ctx.ID().GetSymbol(), msg)
	}

	return nil
}

func (v *ReplVisitor) VisitStructAttr(ctx *compiler.StructAttrContext) interface{} {
	varName := ctx.ID().GetText()
	explicitType := ""
	finalType := ""
	var varValue value.IVOR = value.DefaultUnInitializedValue

	if ctx.Type_() != nil {
		explicitType = v.Visit(ctx.Type_()).(string)
		finalType = explicitType
	} else {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Los atributos de un struct deben tener tipo explícito")
		return nil
	}

	variable, msg := v.ScopeTrace.AddVariable(varName, finalType, varValue, true, true, ctx.ID().GetSymbol())

	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
	}

	return nil
}

type StructInstance struct {
	StructName string
	Fields     map[string]value.IVOR
}

func (v *ReplVisitor) VisitStructInstantiationExpr(ctx *compiler.Struct_instantiationContext) interface{} {
	idToken := ctx.ID()
	if idToken == nil {
		log.Println("❌ Error: no se pudo obtener el ID del struct")
		return nil
	}
	structName := idToken.GetText()
	log.Printf("Nombre del struct a instanciar: %s\n", structName)

	params := ctx.Struct_param_list()
	fieldsMap := make(map[string]value.IVOR)

	if params != nil {
		i := 0
		for {
			paramCtx := params.Struct_param(i)
			if paramCtx == nil {
				break
			}
			paramName := paramCtx.ID().GetText()
			exprValue := v.Visit(paramCtx.Expression()).(value.IVOR)
			log.Printf("Param: %s = %v\n", paramName, exprValue)
			fieldsMap[paramName] = exprValue
			i++
		}
	}

	structValue := &value.StructValue{
		Instance: &value.StructInstance{
			StructName: structName,
			Fields:     fieldsMap,
		},
	}

	log.Println("Instancia creada correctamente:", structValue.ToString())
	return structValue
}

func (v *ReplVisitor) VisitStruct_param(ctx *compiler.Struct_paramContext) interface{} {
	fieldName := ctx.ID().GetText()
	val := v.Visit(ctx.Expression()).(value.IVOR)

	return &StructFieldValue{
		Name:  fieldName,
		Value: val,
	}
}

type StructFieldValue struct {
	Name  string
	Value value.IVOR
}

// func (v *ReplVisitor) VisitStructFunc(ctx *compiler.StructFuncContext) interface{} {

// 	funcDcl := v.Visit(ctx.Func_dcl())

// 	if ctx.MUTATING_KW() != nil {
// 		structFunc, ok := funcDcl.(*Function)

// 		if !ok {
// 			return nil
// 		}
// 		structFunc.IsMutating = true
// 	}

// 	return nil
// }
