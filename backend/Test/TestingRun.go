/*
	TestingRun.go es un archivo que contiene la lógica para ejecutar pruebas de código en un entorno de backend.
*/

package Test

import (
	"fmt"
	"log"

	"github.com/antlr4-go/antlr/v4"
	compiler "main.go/grammar" // Importa tu paquete de gramática
)

// TestingRun es una función que simula la ejecución de pruebas de código.
func TestingRun(code string) (string, error) {
	// Aquí podrías agregar la lógica para ejecutar el código y devolver el resultado.
	// Por ahora, simplemente retornamos el código recibido como una simulación.
	if code == "" {
		return "", fmt.Errorf("el código no puede estar vacío")
	}

	// 1. Análisis Léxico
	// Para verificar errores
	//lexicalErrorListener := errors.NewLexicalErrorListener()
	//
	lexer := compiler.NewCalculatorLexer(antlr.NewInputStream(code))

	lexer.RemoveErrorListeners()
	//lexer.AddErrorListener(lexicalErrorListener)

	// 2. Tokens
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// 3. Parser + errores sintácticos
	// New<Nombre de mi gramatica>(Stream)
	parser := compiler.NewCalculatorParser(stream)
	parser.BuildParseTrees = true

	tree := parser.Prog()

	evaluator := NewCalculatorVisitor()
	resultado := evaluator.Visit(tree)

	fmt.Printf("Resultado de la evaluación: %v\n", resultado)

	return fmt.Sprintf("Código ejecutado: %s", code), nil
}

// Implementar visitor
type CalculatorVisitor struct {
	*compiler.BaseCalculatorVisitor
}

// Es como un constructor para el visitor
func NewCalculatorVisitor() *CalculatorVisitor {
	return &CalculatorVisitor{
		BaseCalculatorVisitor: &compiler.BaseCalculatorVisitor{},
	}
}

func (v *CalculatorVisitor) VisitMulDiv(ctx *compiler.MultipliacacionContext) interface{} {
	izq := v.Visit(ctx.Expr(0)).(int)
	der := v.Visit(ctx.Expr(1)).(int)
	op := ctx.GetText()

	fmt.Printf("Visitando operación: %s %s %s\n", izq, op, der)

	return op
}

func (v *CalculatorVisitor) Visit(tree antlr.ParseTree) interface{} {

	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		return tree.Accept(v) //<---- devolvemos el metodo recursivo que nos da el arbol
	}

}
