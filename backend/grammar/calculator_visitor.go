// Code generated from grammar/Calculator.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // Calculator
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by CalculatorParser.
type CalculatorVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CalculatorParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by CalculatorParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by CalculatorParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by CalculatorParser#Parentesis.
	VisitParentesis(ctx *ParentesisContext) interface{}

	// Visit a parse tree produced by CalculatorParser#entero.
	VisitEntero(ctx *EnteroContext) interface{}

	// Visit a parse tree produced by CalculatorParser#Multipliacacion.
	VisitMultipliacacion(ctx *MultipliacacionContext) interface{}

	// Visit a parse tree produced by CalculatorParser#identificador.
	VisitIdentificador(ctx *IdentificadorContext) interface{}
}
