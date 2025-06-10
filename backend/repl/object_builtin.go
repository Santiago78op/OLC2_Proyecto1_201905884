package repl

import (
	"github.com/antlr4-go/antlr/v4"
	"main.go/value"
)

type ObjectBuiltInFunction struct {
	*Function
	Object     *ObjectValue
	CustomExec func(builtinRef *ObjectBuiltInFunction, visitor *ReplVisitor, args map[string]*Argument, token antlr.Token)
}

// implementing ivor
func (b ObjectBuiltInFunction) Type() string {
	return value.IVOR_FUNCTION
}

func (b ObjectBuiltInFunction) Value() interface{} {
	return b
}

func (b ObjectBuiltInFunction) Copy() value.IVOR {
	return b
}

func (f *ObjectBuiltInFunction) Exec(visitor *ReplVisitor, args []*Argument, token antlr.Token) {

	context := visitor.GetReplContext()

	// validate args
	argsOk, argsMap := f.ValidateArgs(context, args, token)

	if !argsOk {
		f.ReturnValue = value.DefaultNilValue
		return
	}

	f.CustomExec(f, visitor, argsMap, token)

}

// Falta lo de vectores
