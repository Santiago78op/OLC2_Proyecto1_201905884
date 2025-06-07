// Code generated from grammar/Calculator.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // Calculator
import "github.com/antlr4-go/antlr/v4"

// CalculatorListener is a complete listener for a parse tree produced by CalculatorParser.
type CalculatorListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterParentesis is called when entering the Parentesis production.
	EnterParentesis(c *ParentesisContext)

	// EnterEntero is called when entering the entero production.
	EnterEntero(c *EnteroContext)

	// EnterMultipliacacion is called when entering the Multipliacacion production.
	EnterMultipliacacion(c *MultipliacacionContext)

	// EnterIdentificador is called when entering the identificador production.
	EnterIdentificador(c *IdentificadorContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitParentesis is called when exiting the Parentesis production.
	ExitParentesis(c *ParentesisContext)

	// ExitEntero is called when exiting the entero production.
	ExitEntero(c *EnteroContext)

	// ExitMultipliacacion is called when exiting the Multipliacacion production.
	ExitMultipliacacion(c *MultipliacacionContext)

	// ExitIdentificador is called when exiting the identificador production.
	ExitIdentificador(c *IdentificadorContext)
}
