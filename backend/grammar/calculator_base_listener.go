// Code generated from grammar/Calculator.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // Calculator
import "github.com/antlr4-go/antlr/v4"

// BaseCalculatorListener is a complete listener for a parse tree produced by CalculatorParser.
type BaseCalculatorListener struct{}

var _ CalculatorListener = &BaseCalculatorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalculatorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalculatorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalculatorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalculatorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseCalculatorListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseCalculatorListener) ExitProg(ctx *ProgContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseCalculatorListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseCalculatorListener) ExitStatement(ctx *StatementContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseCalculatorListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseCalculatorListener) ExitAddSub(ctx *AddSubContext) {}

// EnterParentesis is called when production Parentesis is entered.
func (s *BaseCalculatorListener) EnterParentesis(ctx *ParentesisContext) {}

// ExitParentesis is called when production Parentesis is exited.
func (s *BaseCalculatorListener) ExitParentesis(ctx *ParentesisContext) {}

// EnterEntero is called when production entero is entered.
func (s *BaseCalculatorListener) EnterEntero(ctx *EnteroContext) {}

// ExitEntero is called when production entero is exited.
func (s *BaseCalculatorListener) ExitEntero(ctx *EnteroContext) {}

// EnterMultipliacacion is called when production Multipliacacion is entered.
func (s *BaseCalculatorListener) EnterMultipliacacion(ctx *MultipliacacionContext) {}

// ExitMultipliacacion is called when production Multipliacacion is exited.
func (s *BaseCalculatorListener) ExitMultipliacacion(ctx *MultipliacacionContext) {}

// EnterIdentificador is called when production identificador is entered.
func (s *BaseCalculatorListener) EnterIdentificador(ctx *IdentificadorContext) {}

// ExitIdentificador is called when production identificador is exited.
func (s *BaseCalculatorListener) ExitIdentificador(ctx *IdentificadorContext) {}
