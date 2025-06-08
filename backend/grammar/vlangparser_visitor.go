// Code generated from grammar/VLangParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // VLangParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by VLangParserParser.
type VLangParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by VLangParserParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by VLangParserParser#delim.
	VisitDelim(ctx *DelimContext) interface{}

	// Visit a parse tree produced by VLangParserParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by VLangParserParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by VLangParserParser#MutableVarDecl.
	VisitMutableVarDecl(ctx *MutableVarDeclContext) interface{}

	// Visit a parse tree produced by VLangParserParser#TypedVarDecl.
	VisitTypedVarDecl(ctx *TypedVarDeclContext) interface{}

	// Visit a parse tree produced by VLangParserParser#type_annotation.
	VisitType_annotation(ctx *Type_annotationContext) interface{}

	// Visit a parse tree produced by VLangParserParser#AssignmentDecl.
	VisitAssignmentDecl(ctx *AssignmentDeclContext) interface{}

	// Visit a parse tree produced by VLangParserParser#PlusAssignmentDecl.
	VisitPlusAssignmentDecl(ctx *PlusAssignmentDeclContext) interface{}

	// Visit a parse tree produced by VLangParserParser#MinusAssignmentDecl.
	VisitMinusAssignmentDecl(ctx *MinusAssignmentDeclContext) interface{}

	// Visit a parse tree produced by VLangParserParser#BinaryExpr.
	VisitBinaryExpr(ctx *BinaryExprContext) interface{}

	// Visit a parse tree produced by VLangParserParser#PrimaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by VLangParserParser#ParenthesizedExpr.
	VisitParenthesizedExpr(ctx *ParenthesizedExprContext) interface{}

	// Visit a parse tree produced by VLangParserParser#UnaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by VLangParserParser#IdentifierExpr.
	VisitIdentifierExpr(ctx *IdentifierExprContext) interface{}

	// Visit a parse tree produced by VLangParserParser#IntLiteral.
	VisitIntLiteral(ctx *IntLiteralContext) interface{}

	// Visit a parse tree produced by VLangParserParser#FloatLiteral.
	VisitFloatLiteral(ctx *FloatLiteralContext) interface{}

	// Visit a parse tree produced by VLangParserParser#StringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by VLangParserParser#TrueLiteral.
	VisitTrueLiteral(ctx *TrueLiteralContext) interface{}

	// Visit a parse tree produced by VLangParserParser#FalseLiteral.
	VisitFalseLiteral(ctx *FalseLiteralContext) interface{}

	// Visit a parse tree produced by VLangParserParser#NilLiteral.
	VisitNilLiteral(ctx *NilLiteralContext) interface{}

	// Visit a parse tree produced by VLangParserParser#BuiltinCall.
	VisitBuiltinCall(ctx *BuiltinCallContext) interface{}

	// Visit a parse tree produced by VLangParserParser#ParenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by VLangParserParser#PrintCall.
	VisitPrintCall(ctx *PrintCallContext) interface{}

	// Visit a parse tree produced by VLangParserParser#PrintlnCall.
	VisitPrintlnCall(ctx *PrintlnCallContext) interface{}

	// Visit a parse tree produced by VLangParserParser#ArgList.
	VisitArgList(ctx *ArgListContext) interface{}
}
