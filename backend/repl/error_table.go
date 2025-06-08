package repl

import "github.com/antlr4-go/antlr/v4"

/*
ErrorTable es una estructura que almacena errores encontrados durante el análisis de código.
Esta tabla permite registrar errores léxicos, sintácticos, semánticos y de tiempo de ejecución.
*/
const (
	LexicalError  = "Error léxico"
	SyntaxError   = "Error sintáctico"
	SemanticError = "Error semántico"
	RuntimeError  = "Error en tiempo de ejecución"
)

// Definimos una tabla de errores
type Error struct {
	Line   int
	Column int
	Msg    string
	Type   string
}

// La tabla de errores
type ErrorTable struct {
	Errors []Error
}

// AddError agrega un error a la tabla de errores con la información proporcionada.
func (et *ErrorTable) AddError(line int, column int, msg string, errorType string) {
	et.Errors = append(et.Errors, Error{line, column, msg, errorType})
}

// NewLexicalError crea un nuevo error léxico y lo agrega a la tabla de errores.
func (et *ErrorTable) NewLexicalError(line int, column int, msg string) {
	et.AddError(line, column, msg, LexicalError)
}

// NewSyntaxError crea un nuevo error sintáctico y lo agrega a la tabla de errores.
func (et *ErrorTable) NewSyntaxError(line int, column int, msg string) {
	et.AddError(line, column, msg, SyntaxError)
}

// NewSemanticError crea un nuevo error semántico y lo agrega a la tabla de errores.
func (et *ErrorTable) NewSemanticError(token antlr.Token, msg string) {
	et.AddError(token.GetLine(), token.GetColumn(), msg, SemanticError)
}

// NewRuntimeError crea un nuevo error de tiempo de ejecución y lo agrega a la tabla de errores.
func (et *ErrorTable) NewRuntimeError(line int, column int, msg string) {
	et.AddError(line, column, msg, RuntimeError)
}

// NewErrorTable crea una nueva instancia de ErrorTable.
func NewErrorTable() *ErrorTable {
	return &ErrorTable{}
}
