package repl

import {
	"fmt"
}

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
		Console:            ctx.Console,
		ErrorTable:         ctx.ErrorTable,
		// Structs:            make(map[string]*Struct),
	}
}

func NewVisitor(dclVisitor *DclVisitor) *ReplVisitor {
	return &ReplVisitor{
		ErrorTable:  dclVisitor.ErrorTable,
		Console:     NewConsole(),
	}
}

func (v *ReplVisitor) GetReplContext() *ReplContext {
	return &ReplContext{
		Console:    v.Console,
		ErrorTable: v.ErrorTable,
	}
}