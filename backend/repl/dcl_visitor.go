package repl

import (
	"log"

	"github.com/antlr4-go/antlr/v4"
	compiler "main.go/grammar"
)

type DclVisitor struct {
	compiler.BaseVLangGrammarVisitor
	ScopeTrace  *ScopeTrace
	ErrorTable  *ErrorTable
	StructNames []string
}

func NewDclVisitor(errorTable *ErrorTable) *DclVisitor {
	return &DclVisitor{
		ScopeTrace:  NewScopeTrace(),
		ErrorTable:  errorTable,
		StructNames: []string{},
	}
}

func (v *DclVisitor) Visit(tree antlr.ParseTree) interface{} {

	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		return tree.Accept(v)
	}

}

func (v *DclVisitor) VisitProgram(ctx *compiler.ProgContext) interface{} {

	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}

	return nil
}

func (v *DclVisitor) VisitStmt(ctx *compiler.StmtContext) interface{} {
	return nil
}
