package repl

import (
	"log"

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

func isDeclMut(lexval string) bool {
	return lexval == "mut"
}

func (v *ReplVisitor) VisitMulVarDecl(ctx *compiler.MutVarDeclContext) interface{} {
	// mut num
	isMut := isDeclMut(ctx.Var_type().GetText())

	// Validar tipo de variable
	if !isMut {
		v.ErrorTable.NewSemanticError(ctx.Var_type().GetStart(), "Tipo de expresión no válido: "+ctx.Var_type().GetText())
	}

	exprName := ctx.ID().GetText()
	exprType := ctx.Type_annotation().GetText()

	// Validar expresión si existe
	if ctx.Expression() != nil {

		exprValue := v.Visit(ctx.Expression()).(value.IVOR)

		// Validar tipo de expresión
		if obj, ok := exprValue.(*ObjectValue); ok {
			exprValue = obj.Copy()
		}

		// Valida si el tipo de expresión es igual al tipo de la variable
		if !v.ValidType(exprType) {
			v.ErrorTable.NewSemanticError(ctx.Type_annotation().GetStart(), "Tipo de expresión no válido: "+exprType)
		}

		variable, msg := v.ScopeTrace.AddVariable(exprName, exprType, exprValue, false, false, ctx.GetStart())

		// Variable already exists
		if variable == nil {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		}

	} else {
		// Si no hay expresión, se crea una variable sin valor
		variable, msg := v.ScopeTrace.AddVariable(exprName, exprType, nil, isMut, false, ctx.GetStart())

		// Variable already exists
		if variable == nil {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		}
	}

	// Si es una variable mut, se agrega al scope actual
	return nil
}

// Validar declaración de variables - Pendiente
