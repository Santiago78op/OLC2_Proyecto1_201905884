// Code generated from grammar/VLangParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // VLangParser
import "github.com/antlr4-go/antlr/v4"

type BaseVLangParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseVLangParserVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitDelim(ctx *DelimContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitMutableVarDecl(ctx *MutableVarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitTypedVarDecl(ctx *TypedVarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitType_annotation(ctx *Type_annotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitAssignmentDecl(ctx *AssignmentDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitPlusAssignmentDecl(ctx *PlusAssignmentDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitMinusAssignmentDecl(ctx *MinusAssignmentDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitBinaryExpr(ctx *BinaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitParenthesizedExpr(ctx *ParenthesizedExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitIdentifierExpr(ctx *IdentifierExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitIntLiteral(ctx *IntLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitFloatLiteral(ctx *FloatLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitTrueLiteral(ctx *TrueLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitFalseLiteral(ctx *FalseLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitNilLiteral(ctx *NilLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitBuiltinCall(ctx *BuiltinCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitPrintCall(ctx *PrintCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitPrintlnCall(ctx *PrintlnCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangParserVisitor) VisitArgList(ctx *ArgListContext) interface{} {
	return v.VisitChildren(ctx)
}
