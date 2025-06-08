// Code generated from grammar/VLangParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // VLangParser
import "github.com/antlr4-go/antlr/v4"

// VLangParserListener is a complete listener for a parse tree produced by VLangParserParser.
type VLangParserListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterDelim is called when entering the delim production.
	EnterDelim(c *DelimContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterMutableVarDecl is called when entering the MutableVarDecl production.
	EnterMutableVarDecl(c *MutableVarDeclContext)

	// EnterTypedVarDecl is called when entering the TypedVarDecl production.
	EnterTypedVarDecl(c *TypedVarDeclContext)

	// EnterType_annotation is called when entering the type_annotation production.
	EnterType_annotation(c *Type_annotationContext)

	// EnterAssignmentDecl is called when entering the AssignmentDecl production.
	EnterAssignmentDecl(c *AssignmentDeclContext)

	// EnterPlusAssignmentDecl is called when entering the PlusAssignmentDecl production.
	EnterPlusAssignmentDecl(c *PlusAssignmentDeclContext)

	// EnterMinusAssignmentDecl is called when entering the MinusAssignmentDecl production.
	EnterMinusAssignmentDecl(c *MinusAssignmentDeclContext)

	// EnterBinaryExpr is called when entering the BinaryExpr production.
	EnterBinaryExpr(c *BinaryExprContext)

	// EnterPrimaryExpr is called when entering the PrimaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterParenthesizedExpr is called when entering the ParenthesizedExpr production.
	EnterParenthesizedExpr(c *ParenthesizedExprContext)

	// EnterUnaryExpr is called when entering the UnaryExpr production.
	EnterUnaryExpr(c *UnaryExprContext)

	// EnterIdentifierExpr is called when entering the IdentifierExpr production.
	EnterIdentifierExpr(c *IdentifierExprContext)

	// EnterIntLiteral is called when entering the IntLiteral production.
	EnterIntLiteral(c *IntLiteralContext)

	// EnterFloatLiteral is called when entering the FloatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// EnterStringLiteral is called when entering the StringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterTrueLiteral is called when entering the TrueLiteral production.
	EnterTrueLiteral(c *TrueLiteralContext)

	// EnterFalseLiteral is called when entering the FalseLiteral production.
	EnterFalseLiteral(c *FalseLiteralContext)

	// EnterNilLiteral is called when entering the NilLiteral production.
	EnterNilLiteral(c *NilLiteralContext)

	// EnterBuiltinCall is called when entering the BuiltinCall production.
	EnterBuiltinCall(c *BuiltinCallContext)

	// EnterParenExpr is called when entering the ParenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterPrintCall is called when entering the PrintCall production.
	EnterPrintCall(c *PrintCallContext)

	// EnterPrintlnCall is called when entering the PrintlnCall production.
	EnterPrintlnCall(c *PrintlnCallContext)

	// EnterArgList is called when entering the ArgList production.
	EnterArgList(c *ArgListContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitDelim is called when exiting the delim production.
	ExitDelim(c *DelimContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitMutableVarDecl is called when exiting the MutableVarDecl production.
	ExitMutableVarDecl(c *MutableVarDeclContext)

	// ExitTypedVarDecl is called when exiting the TypedVarDecl production.
	ExitTypedVarDecl(c *TypedVarDeclContext)

	// ExitType_annotation is called when exiting the type_annotation production.
	ExitType_annotation(c *Type_annotationContext)

	// ExitAssignmentDecl is called when exiting the AssignmentDecl production.
	ExitAssignmentDecl(c *AssignmentDeclContext)

	// ExitPlusAssignmentDecl is called when exiting the PlusAssignmentDecl production.
	ExitPlusAssignmentDecl(c *PlusAssignmentDeclContext)

	// ExitMinusAssignmentDecl is called when exiting the MinusAssignmentDecl production.
	ExitMinusAssignmentDecl(c *MinusAssignmentDeclContext)

	// ExitBinaryExpr is called when exiting the BinaryExpr production.
	ExitBinaryExpr(c *BinaryExprContext)

	// ExitPrimaryExpr is called when exiting the PrimaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitParenthesizedExpr is called when exiting the ParenthesizedExpr production.
	ExitParenthesizedExpr(c *ParenthesizedExprContext)

	// ExitUnaryExpr is called when exiting the UnaryExpr production.
	ExitUnaryExpr(c *UnaryExprContext)

	// ExitIdentifierExpr is called when exiting the IdentifierExpr production.
	ExitIdentifierExpr(c *IdentifierExprContext)

	// ExitIntLiteral is called when exiting the IntLiteral production.
	ExitIntLiteral(c *IntLiteralContext)

	// ExitFloatLiteral is called when exiting the FloatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)

	// ExitStringLiteral is called when exiting the StringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitTrueLiteral is called when exiting the TrueLiteral production.
	ExitTrueLiteral(c *TrueLiteralContext)

	// ExitFalseLiteral is called when exiting the FalseLiteral production.
	ExitFalseLiteral(c *FalseLiteralContext)

	// ExitNilLiteral is called when exiting the NilLiteral production.
	ExitNilLiteral(c *NilLiteralContext)

	// ExitBuiltinCall is called when exiting the BuiltinCall production.
	ExitBuiltinCall(c *BuiltinCallContext)

	// ExitParenExpr is called when exiting the ParenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitPrintCall is called when exiting the PrintCall production.
	ExitPrintCall(c *PrintCallContext)

	// ExitPrintlnCall is called when exiting the PrintlnCall production.
	ExitPrintlnCall(c *PrintlnCallContext)

	// ExitArgList is called when exiting the ArgList production.
	ExitArgList(c *ArgListContext)
}
