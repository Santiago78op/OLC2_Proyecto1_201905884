package errors

import (
	"fmt"
	"strings"
	// Traer de parser/parser generado por ANTLR
	"github.com/antlr4-go/antlr/v4"
)

type CustomErrorStrategy struct {
	*antlr.DefaultErrorStrategy
}

func NewCustomErrorStrategy() *CustomErrorStrategy {
	return &CustomErrorStrategy{
		DefaultErrorStrategy: antlr.NewDefaultErrorStrategy(),
	}
}

// Traducción en español de mensajes de error
func (es *CustomErrorStrategy) ReportInputMismatch(recognizer antlr.Parser, e *antlr.InputMismatchException) {
	token := e.GetOffendingToken()
	expected := es.GetExpectedTokens(recognizer, e.GetExpectedTokens())
	msg := fmt.Sprintf("Se recibió '%s', se esperaba %s", token.GetText(), expected)
	recognizer.NotifyErrorListeners(msg, e.GetOffendingToken(), e)
}

func (es *CustomErrorStrategy) ReportNoViableAlternative(recognizer antlr.Parser, e *antlr.NoViableAltException) {
	token := e.GetOffendingToken()
	msg := fmt.Sprintf("No se encontró alternativa válida en '%s'", token.GetText())
	recognizer.NotifyErrorListeners(msg, e.GetOffendingToken(), e)
}

func (es *CustomErrorStrategy) GetExpectedTokens(recognizer antlr.Parser, expectedTokens *antlr.IntervalSet) string {
	if expectedTokens.Length() == 0 {
		return "fin de archivo"
	}

	tokenNames := recognizer.GetLiteralNames()
	symbolicNames := recognizer.GetSymbolicNames()

	var expected []string
	for i := 0; i < expectedTokens.Length(); i++ {
		tokenType := expectedTokens.Get(i)
		if tokenType < len(tokenNames) && tokenNames[tokenType] != "" {
			expected = append(expected, tokenNames[tokenType])
		} else if tokenType < len(symbolicNames) && symbolicNames[tokenType] != "" {
			expected = append(expected, symbolicNames[tokenType])
		} else {
			expected = append(expected, fmt.Sprintf("token_%d", tokenType))
		}
	}

	if len(expected) == 1 {
		return expected[0]
	}

	return strings.Join(expected[:len(expected)-1], ", ") + " o " + expected[len(expected)-1]
}
