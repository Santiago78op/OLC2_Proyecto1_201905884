// Code generated from grammar/Calculator.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // Calculator
import "github.com/antlr4-go/antlr/v4"

type BaseCalculatorVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCalculatorVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitParentesis(ctx *ParentesisContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitEntero(ctx *EnteroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitMultipliacacion(ctx *MultipliacacionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitIdentificador(ctx *IdentificadorContext) interface{} {
	return v.VisitChildren(ctx)
}
