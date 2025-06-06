// Code generated from C:/Users/72358/Documents/Compi_2/Semana_01/awesomeProject/grammar/VLangGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // VLangGrammar
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by VLangGrammarParser.
type VLangGrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by VLangGrammarParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#MutableVarDecl.
	VisitMutableVarDecl(ctx *MutableVarDeclContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#TypedVarDecl.
	VisitTypedVarDecl(ctx *TypedVarDeclContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#InferredVarDecl.
	VisitInferredVarDecl(ctx *InferredVarDeclContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#type_annotation.
	VisitType_annotation(ctx *Type_annotationContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#slice_type.
	VisitSlice_type(ctx *Slice_typeContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#function_declaration.
	VisitFunction_declaration(ctx *Function_declarationContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#receiver.
	VisitReceiver(ctx *ReceiverContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#parameter_list.
	VisitParameter_list(ctx *Parameter_listContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#struct_declaration.
	VisitStruct_declaration(ctx *Struct_declarationContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#struct_field.
	VisitStruct_field(ctx *Struct_fieldContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#SimpleAssignment.
	VisitSimpleAssignment(ctx *SimpleAssignmentContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#PlusAssignment.
	VisitPlusAssignment(ctx *PlusAssignmentContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#MinusAssignment.
	VisitMinusAssignment(ctx *MinusAssignmentContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#IndexAssignment.
	VisitIndexAssignment(ctx *IndexAssignmentContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#FieldAssignment.
	VisitFieldAssignment(ctx *FieldAssignmentContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#if_statement.
	VisitIf_statement(ctx *If_statementContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#ForCondition.
	VisitForCondition(ctx *ForConditionContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#ForLoop.
	VisitForLoop(ctx *ForLoopContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#ForIn.
	VisitForIn(ctx *ForInContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#switch_statement.
	VisitSwitch_statement(ctx *Switch_statementContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#case_clause.
	VisitCase_clause(ctx *Case_clauseContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#default_clause.
	VisitDefault_clause(ctx *Default_clauseContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#break_statement.
	VisitBreak_statement(ctx *Break_statementContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#continue_statement.
	VisitContinue_statement(ctx *Continue_statementContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#return_statement.
	VisitReturn_statement(ctx *Return_statementContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#expression_statement.
	VisitExpression_statement(ctx *Expression_statementContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#block_statement.
	VisitBlock_statement(ctx *Block_statementContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#UnaryNot.
	VisitUnaryNot(ctx *UnaryNotContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#MulDivMod.
	VisitMulDivMod(ctx *MulDivModContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#IndexAccess.
	VisitIndexAccess(ctx *IndexAccessContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#Primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#LogicalAnd.
	VisitLogicalAnd(ctx *LogicalAndContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#Relational.
	VisitRelational(ctx *RelationalContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#UnaryMinus.
	VisitUnaryMinus(ctx *UnaryMinusContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#Equality.
	VisitEquality(ctx *EqualityContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#FunctionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#FieldAccess.
	VisitFieldAccess(ctx *FieldAccessContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#LogicalOr.
	VisitLogicalOr(ctx *LogicalOrContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#IdentifierExpr.
	VisitIdentifierExpr(ctx *IdentifierExprContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#IntLiteral.
	VisitIntLiteral(ctx *IntLiteralContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#FloatLiteral.
	VisitFloatLiteral(ctx *FloatLiteralContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#StringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#RuneLiteral.
	VisitRuneLiteral(ctx *RuneLiteralContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#TrueLiteral.
	VisitTrueLiteral(ctx *TrueLiteralContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#FalseLiteral.
	VisitFalseLiteral(ctx *FalseLiteralContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#NilLiteral.
	VisitNilLiteral(ctx *NilLiteralContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#SliceLiteralExpr.
	VisitSliceLiteralExpr(ctx *SliceLiteralExprContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#StructLiteralExpr.
	VisitStructLiteralExpr(ctx *StructLiteralExprContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#BuiltinCall.
	VisitBuiltinCall(ctx *BuiltinCallContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#ParenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#slice_literal.
	VisitSlice_literal(ctx *Slice_literalContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#struct_literal.
	VisitStruct_literal(ctx *Struct_literalContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#field_assignment_list.
	VisitField_assignment_list(ctx *Field_assignment_listContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#field_assignment.
	VisitField_assignment(ctx *Field_assignmentContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#PrintCall.
	VisitPrintCall(ctx *PrintCallContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#AtoiCall.
	VisitAtoiCall(ctx *AtoiCallContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#ParseFloatCall.
	VisitParseFloatCall(ctx *ParseFloatCallContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#TypeOfCall.
	VisitTypeOfCall(ctx *TypeOfCallContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#IndexOfCall.
	VisitIndexOfCall(ctx *IndexOfCallContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#JoinCall.
	VisitJoinCall(ctx *JoinCallContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#LenCall.
	VisitLenCall(ctx *LenCallContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#AppendCall.
	VisitAppendCall(ctx *AppendCallContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#argument_list.
	VisitArgument_list(ctx *Argument_listContext) interface{}

	// Visit a parse tree produced by VLangGrammarParser#expression_list.
	VisitExpression_list(ctx *Expression_listContext) interface{}
}
