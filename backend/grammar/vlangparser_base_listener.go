// Code generated from grammar/VLangParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // VLangParser
import "github.com/antlr4-go/antlr/v4"

// BaseVLangParserListener is a complete listener for a parse tree produced by VLangParserParser.
type BaseVLangParserListener struct{}

var _ VLangParserListener = &BaseVLangParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVLangParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVLangParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVLangParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVLangParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseVLangParserListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseVLangParserListener) ExitProg(ctx *ProgContext) {}

// EnterDelim is called when production delim is entered.
func (s *BaseVLangParserListener) EnterDelim(ctx *DelimContext) {}

// ExitDelim is called when production delim is exited.
func (s *BaseVLangParserListener) ExitDelim(ctx *DelimContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseVLangParserListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseVLangParserListener) ExitStmt(ctx *StmtContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseVLangParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseVLangParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterMutableVarDecl is called when production MutableVarDecl is entered.
func (s *BaseVLangParserListener) EnterMutableVarDecl(ctx *MutableVarDeclContext) {}

// ExitMutableVarDecl is called when production MutableVarDecl is exited.
func (s *BaseVLangParserListener) ExitMutableVarDecl(ctx *MutableVarDeclContext) {}

// EnterTypedVarDecl is called when production TypedVarDecl is entered.
func (s *BaseVLangParserListener) EnterTypedVarDecl(ctx *TypedVarDeclContext) {}

// ExitTypedVarDecl is called when production TypedVarDecl is exited.
func (s *BaseVLangParserListener) ExitTypedVarDecl(ctx *TypedVarDeclContext) {}

// EnterType_annotation is called when production type_annotation is entered.
func (s *BaseVLangParserListener) EnterType_annotation(ctx *Type_annotationContext) {}

// ExitType_annotation is called when production type_annotation is exited.
func (s *BaseVLangParserListener) ExitType_annotation(ctx *Type_annotationContext) {}

// EnterAssignmentDecl is called when production AssignmentDecl is entered.
func (s *BaseVLangParserListener) EnterAssignmentDecl(ctx *AssignmentDeclContext) {}

// ExitAssignmentDecl is called when production AssignmentDecl is exited.
func (s *BaseVLangParserListener) ExitAssignmentDecl(ctx *AssignmentDeclContext) {}

// EnterPlusAssignmentDecl is called when production PlusAssignmentDecl is entered.
func (s *BaseVLangParserListener) EnterPlusAssignmentDecl(ctx *PlusAssignmentDeclContext) {}

// ExitPlusAssignmentDecl is called when production PlusAssignmentDecl is exited.
func (s *BaseVLangParserListener) ExitPlusAssignmentDecl(ctx *PlusAssignmentDeclContext) {}

// EnterMinusAssignmentDecl is called when production MinusAssignmentDecl is entered.
func (s *BaseVLangParserListener) EnterMinusAssignmentDecl(ctx *MinusAssignmentDeclContext) {}

// ExitMinusAssignmentDecl is called when production MinusAssignmentDecl is exited.
func (s *BaseVLangParserListener) ExitMinusAssignmentDecl(ctx *MinusAssignmentDeclContext) {}

// EnterBinaryExpr is called when production BinaryExpr is entered.
func (s *BaseVLangParserListener) EnterBinaryExpr(ctx *BinaryExprContext) {}

// ExitBinaryExpr is called when production BinaryExpr is exited.
func (s *BaseVLangParserListener) ExitBinaryExpr(ctx *BinaryExprContext) {}

// EnterPrimaryExpr is called when production PrimaryExpr is entered.
func (s *BaseVLangParserListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production PrimaryExpr is exited.
func (s *BaseVLangParserListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterParenthesizedExpr is called when production ParenthesizedExpr is entered.
func (s *BaseVLangParserListener) EnterParenthesizedExpr(ctx *ParenthesizedExprContext) {}

// ExitParenthesizedExpr is called when production ParenthesizedExpr is exited.
func (s *BaseVLangParserListener) ExitParenthesizedExpr(ctx *ParenthesizedExprContext) {}

// EnterUnaryExpr is called when production UnaryExpr is entered.
func (s *BaseVLangParserListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production UnaryExpr is exited.
func (s *BaseVLangParserListener) ExitUnaryExpr(ctx *UnaryExprContext) {}

// EnterIdentifierExpr is called when production IdentifierExpr is entered.
func (s *BaseVLangParserListener) EnterIdentifierExpr(ctx *IdentifierExprContext) {}

// ExitIdentifierExpr is called when production IdentifierExpr is exited.
func (s *BaseVLangParserListener) ExitIdentifierExpr(ctx *IdentifierExprContext) {}

// EnterIntLiteral is called when production IntLiteral is entered.
func (s *BaseVLangParserListener) EnterIntLiteral(ctx *IntLiteralContext) {}

// ExitIntLiteral is called when production IntLiteral is exited.
func (s *BaseVLangParserListener) ExitIntLiteral(ctx *IntLiteralContext) {}

// EnterFloatLiteral is called when production FloatLiteral is entered.
func (s *BaseVLangParserListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production FloatLiteral is exited.
func (s *BaseVLangParserListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterStringLiteral is called when production StringLiteral is entered.
func (s *BaseVLangParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production StringLiteral is exited.
func (s *BaseVLangParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterTrueLiteral is called when production TrueLiteral is entered.
func (s *BaseVLangParserListener) EnterTrueLiteral(ctx *TrueLiteralContext) {}

// ExitTrueLiteral is called when production TrueLiteral is exited.
func (s *BaseVLangParserListener) ExitTrueLiteral(ctx *TrueLiteralContext) {}

// EnterFalseLiteral is called when production FalseLiteral is entered.
func (s *BaseVLangParserListener) EnterFalseLiteral(ctx *FalseLiteralContext) {}

// ExitFalseLiteral is called when production FalseLiteral is exited.
func (s *BaseVLangParserListener) ExitFalseLiteral(ctx *FalseLiteralContext) {}

// EnterNilLiteral is called when production NilLiteral is entered.
func (s *BaseVLangParserListener) EnterNilLiteral(ctx *NilLiteralContext) {}

// ExitNilLiteral is called when production NilLiteral is exited.
func (s *BaseVLangParserListener) ExitNilLiteral(ctx *NilLiteralContext) {}

// EnterBuiltinCall is called when production BuiltinCall is entered.
func (s *BaseVLangParserListener) EnterBuiltinCall(ctx *BuiltinCallContext) {}

// ExitBuiltinCall is called when production BuiltinCall is exited.
func (s *BaseVLangParserListener) ExitBuiltinCall(ctx *BuiltinCallContext) {}

// EnterParenExpr is called when production ParenExpr is entered.
func (s *BaseVLangParserListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production ParenExpr is exited.
func (s *BaseVLangParserListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterPrintCall is called when production PrintCall is entered.
func (s *BaseVLangParserListener) EnterPrintCall(ctx *PrintCallContext) {}

// ExitPrintCall is called when production PrintCall is exited.
func (s *BaseVLangParserListener) ExitPrintCall(ctx *PrintCallContext) {}

// EnterPrintlnCall is called when production PrintlnCall is entered.
func (s *BaseVLangParserListener) EnterPrintlnCall(ctx *PrintlnCallContext) {}

// ExitPrintlnCall is called when production PrintlnCall is exited.
func (s *BaseVLangParserListener) ExitPrintlnCall(ctx *PrintlnCallContext) {}

// EnterArgList is called when production ArgList is entered.
func (s *BaseVLangParserListener) EnterArgList(ctx *ArgListContext) {}

// ExitArgList is called when production ArgList is exited.
func (s *BaseVLangParserListener) ExitArgList(ctx *ArgListContext) {}
