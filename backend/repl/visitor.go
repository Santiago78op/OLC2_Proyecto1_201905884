package repl

import (
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
	*compiler.BaseVLangGrammarVisitor
	ScopeTrace *ScopeTrace
	Console    *Console
	CallStack  *CallStack
	ErrorTable *ErrorTable
	// Structs map[string]*Struct
}

/*
NewReplVisitor crea un nuevo ReplVisitor con el contexto proporcionado.
*/
func NewReplVisitor(ctx *ReplContext) *ReplVisitor {
	return &ReplVisitor{
		BaseVLangGrammarVisitor: &compiler.BaseVLangGrammarVisitor{},
		Console:                 ctx.Console,
		ErrorTable:              ctx.ErrorTable,
		// Structs:            make(map[string]*Struct),
	}
}

func NewVisitor(dclVisitor *DclVisitor) *ReplVisitor {
	return &ReplVisitor{
		ErrorTable: dclVisitor.ErrorTable,
		Console:    NewConsole(),
	}
}

func (v *ReplVisitor) GetReplContext() *ReplContext {
	return &ReplContext{
		Console:    v.Console,
		ErrorTable: v.ErrorTable,
	}
}

func (v *ReplVisitor) ValidType(_type string) bool {
	return v.ScopeTrace.GlobalScope.ValidType(_type)
}

func (v *ReplVisitor) Visit(tree antlr.ParseTree) interface{} {

	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		return tree.Accept(v)
	}

}

func (v *ReplVisitor) VisitProgram(ctx *compiler.ProgContext) interface{} {

	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}

	return nil
}

func (v *ReplVisitor) VisitStmt(ctx *compiler.StmtContext) interface{} {

	if ctx.Decl_stmt() != nil {
		v.Visit(ctx.Decl_stmt())
	} else if ctx.Assign_stmt() != nil {
		v.Visit(ctx.Assign_stmt())
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
func (v *ReplVisitor) VisitMulVarDecl(ctx *compiler.MutVarDeclContext) interface{} {

	// Si hubiera constantes se validan aquí
	// isConst := isDeclConst(ctx.Var_type().GetText())
	isConst := false

	// Obtenemos el context de la declaración MutVarDecl
	exprName := ctx.ID().GetText()
	exprType := v.Visit(ctx.Type_annotation()).(string)

	// Validar expresión si existe
	if ctx.Expression() != nil {

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

	} else {
		// Si no hay expresión, se crea una variable sin valor
		variable, msg := v.ScopeTrace.AddVariable(exprName, exprType, value.DefaultUnInitializedValue, isConst, false, ctx.GetStart())

		// Variable already exists
		if variable == nil {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		}
	}

	// Si es una variable mut, se agrega al scope actual
	return nil
}

// Ejemplo: variable_1 int = 10
func (v *ReplVisitor) VisitVarDecl(ctx *compiler.VarAssDeclContext) interface{} {

	// Si hubiera constantes se validan aquí
	// isConst := isDeclConst(ctx.Var_type().GetText())
	isConst := true

	// Obtenemos el context de la declaración VarAssDecñ
	exprName := ctx.ID().GetText()
	exprType := v.Visit(ctx.Type_annotation()).(string)

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

		// Si es un vector, se copia el valor
		/*
			if IsVectorType(varValue.Type()) {
				varValue = varValue.Copy()
			}
		*/

		// Si es una variable mut, se puede mutar
		canMutate := true

		// Si la variable es una propiedad, no se puede mutar
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

func (v *ReplVisitor) VisitArithmeticAssign(ctx *compiler.AugmentedAssignmentDeclContext) interface{} {
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

func (v *ReplVisitor) VisitFloatLiteral(ctx *compiler.FloatLiteralContext) interface{} {

	floatVal, _ := strconv.ParseFloat(ctx.GetText(), 64)

	return &value.FloatValue{
		InternalValue: floatVal,
	}

}

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

func (v *ReplVisitor) VisitBoolLiteral(ctx *compiler.BoolLiteralContext) interface{} {

	boolVal, _ := strconv.ParseBool(ctx.GetText())

	return &value.BoolValue{
		InternalValue: boolVal,
	}

}

func (v *ReplVisitor) VisitNilLiteral(ctx *compiler.NilLiteralContext) interface{} {
	return value.DefaultNilValue
}
