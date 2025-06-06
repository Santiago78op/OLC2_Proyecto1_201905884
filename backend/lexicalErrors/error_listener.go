package lexicalErrors

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

type SyntaxErrorListener struct {
	*antlr.DefaultErrorListener
}

func FuncionPrueba() {
	fmt.Println("hola")
}
