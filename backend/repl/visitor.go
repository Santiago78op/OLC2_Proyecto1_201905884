package repl

import (
	"log"

	"github.com/antlr4-go/antlr/v4"
	compiler "main.go/grammar"
)

/*
ReplVisitor es una estructura que implementa el visitor para el REPL (Read-Eval-Print Loop).
*/
type ReplVisitor struct {
	*compiler.BaseVLangParserVisitor
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
		BaseVLangParserVisitor: &compiler.BaseVLangParserVisitor{},
		Console:                ctx.Console,
		ErrorTable:             ctx.ErrorTable,
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

	if ctx.Declaration() != nil {
		v.Visit(ctx.Declaration())
	} else {
		log.Fatal("Statement not recognized: ", ctx.GetText())
	}

	return nil
}

func isDeclConst(lexval string) bool {
	return lexval == "const"
}

// Validar declaración de variables - Pendiente
