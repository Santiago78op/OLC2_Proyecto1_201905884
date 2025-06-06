// Code generated from C:/Users/72358/Documents/Compi_2/Semana_01/awesomeProject/grammar/VLangGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // VLangGrammar
import "github.com/antlr4-go/antlr/v4"

// VLangGrammarListener is a complete listener for a parse tree produced by VLangGrammarParser.
type VLangGrammarListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterMutableVarDecl is called when entering the MutableVarDecl production.
	EnterMutableVarDecl(c *MutableVarDeclContext)

	// EnterTypedVarDecl is called when entering the TypedVarDecl production.
	EnterTypedVarDecl(c *TypedVarDeclContext)

	// EnterInferredVarDecl is called when entering the InferredVarDecl production.
	EnterInferredVarDecl(c *InferredVarDeclContext)

	// EnterType_annotation is called when entering the type_annotation production.
	EnterType_annotation(c *Type_annotationContext)

	// EnterSlice_type is called when entering the slice_type production.
	EnterSlice_type(c *Slice_typeContext)

	// EnterFunction_declaration is called when entering the function_declaration production.
	EnterFunction_declaration(c *Function_declarationContext)

	// EnterReceiver is called when entering the receiver production.
	EnterReceiver(c *ReceiverContext)

	// EnterParameter_list is called when entering the parameter_list production.
	EnterParameter_list(c *Parameter_listContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterStruct_declaration is called when entering the struct_declaration production.
	EnterStruct_declaration(c *Struct_declarationContext)

	// EnterStruct_field is called when entering the struct_field production.
	EnterStruct_field(c *Struct_fieldContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterSimpleAssignment is called when entering the SimpleAssignment production.
	EnterSimpleAssignment(c *SimpleAssignmentContext)

	// EnterPlusAssignment is called when entering the PlusAssignment production.
	EnterPlusAssignment(c *PlusAssignmentContext)

	// EnterMinusAssignment is called when entering the MinusAssignment production.
	EnterMinusAssignment(c *MinusAssignmentContext)

	// EnterIndexAssignment is called when entering the IndexAssignment production.
	EnterIndexAssignment(c *IndexAssignmentContext)

	// EnterFieldAssignment is called when entering the FieldAssignment production.
	EnterFieldAssignment(c *FieldAssignmentContext)

	// EnterIf_statement is called when entering the if_statement production.
	EnterIf_statement(c *If_statementContext)

	// EnterForCondition is called when entering the ForCondition production.
	EnterForCondition(c *ForConditionContext)

	// EnterForLoop is called when entering the ForLoop production.
	EnterForLoop(c *ForLoopContext)

	// EnterForIn is called when entering the ForIn production.
	EnterForIn(c *ForInContext)

	// EnterSwitch_statement is called when entering the switch_statement production.
	EnterSwitch_statement(c *Switch_statementContext)

	// EnterCase_clause is called when entering the case_clause production.
	EnterCase_clause(c *Case_clauseContext)

	// EnterDefault_clause is called when entering the default_clause production.
	EnterDefault_clause(c *Default_clauseContext)

	// EnterBreak_statement is called when entering the break_statement production.
	EnterBreak_statement(c *Break_statementContext)

	// EnterContinue_statement is called when entering the continue_statement production.
	EnterContinue_statement(c *Continue_statementContext)

	// EnterReturn_statement is called when entering the return_statement production.
	EnterReturn_statement(c *Return_statementContext)

	// EnterExpression_statement is called when entering the expression_statement production.
	EnterExpression_statement(c *Expression_statementContext)

	// EnterBlock_statement is called when entering the block_statement production.
	EnterBlock_statement(c *Block_statementContext)

	// EnterUnaryNot is called when entering the UnaryNot production.
	EnterUnaryNot(c *UnaryNotContext)

	// EnterMulDivMod is called when entering the MulDivMod production.
	EnterMulDivMod(c *MulDivModContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterIndexAccess is called when entering the IndexAccess production.
	EnterIndexAccess(c *IndexAccessContext)

	// EnterPrimary is called when entering the Primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterLogicalAnd is called when entering the LogicalAnd production.
	EnterLogicalAnd(c *LogicalAndContext)

	// EnterRelational is called when entering the Relational production.
	EnterRelational(c *RelationalContext)

	// EnterUnaryMinus is called when entering the UnaryMinus production.
	EnterUnaryMinus(c *UnaryMinusContext)

	// EnterEquality is called when entering the Equality production.
	EnterEquality(c *EqualityContext)

	// EnterFunctionCall is called when entering the FunctionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterFieldAccess is called when entering the FieldAccess production.
	EnterFieldAccess(c *FieldAccessContext)

	// EnterLogicalOr is called when entering the LogicalOr production.
	EnterLogicalOr(c *LogicalOrContext)

	// EnterIdentifierExpr is called when entering the IdentifierExpr production.
	EnterIdentifierExpr(c *IdentifierExprContext)

	// EnterIntLiteral is called when entering the IntLiteral production.
	EnterIntLiteral(c *IntLiteralContext)

	// EnterFloatLiteral is called when entering the FloatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// EnterStringLiteral is called when entering the StringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterRuneLiteral is called when entering the RuneLiteral production.
	EnterRuneLiteral(c *RuneLiteralContext)

	// EnterTrueLiteral is called when entering the TrueLiteral production.
	EnterTrueLiteral(c *TrueLiteralContext)

	// EnterFalseLiteral is called when entering the FalseLiteral production.
	EnterFalseLiteral(c *FalseLiteralContext)

	// EnterNilLiteral is called when entering the NilLiteral production.
	EnterNilLiteral(c *NilLiteralContext)

	// EnterSliceLiteralExpr is called when entering the SliceLiteralExpr production.
	EnterSliceLiteralExpr(c *SliceLiteralExprContext)

	// EnterStructLiteralExpr is called when entering the StructLiteralExpr production.
	EnterStructLiteralExpr(c *StructLiteralExprContext)

	// EnterBuiltinCall is called when entering the BuiltinCall production.
	EnterBuiltinCall(c *BuiltinCallContext)

	// EnterParenExpr is called when entering the ParenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterSlice_literal is called when entering the slice_literal production.
	EnterSlice_literal(c *Slice_literalContext)

	// EnterStruct_literal is called when entering the struct_literal production.
	EnterStruct_literal(c *Struct_literalContext)

	// EnterField_assignment_list is called when entering the field_assignment_list production.
	EnterField_assignment_list(c *Field_assignment_listContext)

	// EnterField_assignment is called when entering the field_assignment production.
	EnterField_assignment(c *Field_assignmentContext)

	// EnterPrintCall is called when entering the PrintCall production.
	EnterPrintCall(c *PrintCallContext)

	// EnterAtoiCall is called when entering the AtoiCall production.
	EnterAtoiCall(c *AtoiCallContext)

	// EnterParseFloatCall is called when entering the ParseFloatCall production.
	EnterParseFloatCall(c *ParseFloatCallContext)

	// EnterTypeOfCall is called when entering the TypeOfCall production.
	EnterTypeOfCall(c *TypeOfCallContext)

	// EnterIndexOfCall is called when entering the IndexOfCall production.
	EnterIndexOfCall(c *IndexOfCallContext)

	// EnterJoinCall is called when entering the JoinCall production.
	EnterJoinCall(c *JoinCallContext)

	// EnterLenCall is called when entering the LenCall production.
	EnterLenCall(c *LenCallContext)

	// EnterAppendCall is called when entering the AppendCall production.
	EnterAppendCall(c *AppendCallContext)

	// EnterArgument_list is called when entering the argument_list production.
	EnterArgument_list(c *Argument_listContext)

	// EnterExpression_list is called when entering the expression_list production.
	EnterExpression_list(c *Expression_listContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitMutableVarDecl is called when exiting the MutableVarDecl production.
	ExitMutableVarDecl(c *MutableVarDeclContext)

	// ExitTypedVarDecl is called when exiting the TypedVarDecl production.
	ExitTypedVarDecl(c *TypedVarDeclContext)

	// ExitInferredVarDecl is called when exiting the InferredVarDecl production.
	ExitInferredVarDecl(c *InferredVarDeclContext)

	// ExitType_annotation is called when exiting the type_annotation production.
	ExitType_annotation(c *Type_annotationContext)

	// ExitSlice_type is called when exiting the slice_type production.
	ExitSlice_type(c *Slice_typeContext)

	// ExitFunction_declaration is called when exiting the function_declaration production.
	ExitFunction_declaration(c *Function_declarationContext)

	// ExitReceiver is called when exiting the receiver production.
	ExitReceiver(c *ReceiverContext)

	// ExitParameter_list is called when exiting the parameter_list production.
	ExitParameter_list(c *Parameter_listContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitStruct_declaration is called when exiting the struct_declaration production.
	ExitStruct_declaration(c *Struct_declarationContext)

	// ExitStruct_field is called when exiting the struct_field production.
	ExitStruct_field(c *Struct_fieldContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitSimpleAssignment is called when exiting the SimpleAssignment production.
	ExitSimpleAssignment(c *SimpleAssignmentContext)

	// ExitPlusAssignment is called when exiting the PlusAssignment production.
	ExitPlusAssignment(c *PlusAssignmentContext)

	// ExitMinusAssignment is called when exiting the MinusAssignment production.
	ExitMinusAssignment(c *MinusAssignmentContext)

	// ExitIndexAssignment is called when exiting the IndexAssignment production.
	ExitIndexAssignment(c *IndexAssignmentContext)

	// ExitFieldAssignment is called when exiting the FieldAssignment production.
	ExitFieldAssignment(c *FieldAssignmentContext)

	// ExitIf_statement is called when exiting the if_statement production.
	ExitIf_statement(c *If_statementContext)

	// ExitForCondition is called when exiting the ForCondition production.
	ExitForCondition(c *ForConditionContext)

	// ExitForLoop is called when exiting the ForLoop production.
	ExitForLoop(c *ForLoopContext)

	// ExitForIn is called when exiting the ForIn production.
	ExitForIn(c *ForInContext)

	// ExitSwitch_statement is called when exiting the switch_statement production.
	ExitSwitch_statement(c *Switch_statementContext)

	// ExitCase_clause is called when exiting the case_clause production.
	ExitCase_clause(c *Case_clauseContext)

	// ExitDefault_clause is called when exiting the default_clause production.
	ExitDefault_clause(c *Default_clauseContext)

	// ExitBreak_statement is called when exiting the break_statement production.
	ExitBreak_statement(c *Break_statementContext)

	// ExitContinue_statement is called when exiting the continue_statement production.
	ExitContinue_statement(c *Continue_statementContext)

	// ExitReturn_statement is called when exiting the return_statement production.
	ExitReturn_statement(c *Return_statementContext)

	// ExitExpression_statement is called when exiting the expression_statement production.
	ExitExpression_statement(c *Expression_statementContext)

	// ExitBlock_statement is called when exiting the block_statement production.
	ExitBlock_statement(c *Block_statementContext)

	// ExitUnaryNot is called when exiting the UnaryNot production.
	ExitUnaryNot(c *UnaryNotContext)

	// ExitMulDivMod is called when exiting the MulDivMod production.
	ExitMulDivMod(c *MulDivModContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitIndexAccess is called when exiting the IndexAccess production.
	ExitIndexAccess(c *IndexAccessContext)

	// ExitPrimary is called when exiting the Primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitLogicalAnd is called when exiting the LogicalAnd production.
	ExitLogicalAnd(c *LogicalAndContext)

	// ExitRelational is called when exiting the Relational production.
	ExitRelational(c *RelationalContext)

	// ExitUnaryMinus is called when exiting the UnaryMinus production.
	ExitUnaryMinus(c *UnaryMinusContext)

	// ExitEquality is called when exiting the Equality production.
	ExitEquality(c *EqualityContext)

	// ExitFunctionCall is called when exiting the FunctionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitFieldAccess is called when exiting the FieldAccess production.
	ExitFieldAccess(c *FieldAccessContext)

	// ExitLogicalOr is called when exiting the LogicalOr production.
	ExitLogicalOr(c *LogicalOrContext)

	// ExitIdentifierExpr is called when exiting the IdentifierExpr production.
	ExitIdentifierExpr(c *IdentifierExprContext)

	// ExitIntLiteral is called when exiting the IntLiteral production.
	ExitIntLiteral(c *IntLiteralContext)

	// ExitFloatLiteral is called when exiting the FloatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)

	// ExitStringLiteral is called when exiting the StringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitRuneLiteral is called when exiting the RuneLiteral production.
	ExitRuneLiteral(c *RuneLiteralContext)

	// ExitTrueLiteral is called when exiting the TrueLiteral production.
	ExitTrueLiteral(c *TrueLiteralContext)

	// ExitFalseLiteral is called when exiting the FalseLiteral production.
	ExitFalseLiteral(c *FalseLiteralContext)

	// ExitNilLiteral is called when exiting the NilLiteral production.
	ExitNilLiteral(c *NilLiteralContext)

	// ExitSliceLiteralExpr is called when exiting the SliceLiteralExpr production.
	ExitSliceLiteralExpr(c *SliceLiteralExprContext)

	// ExitStructLiteralExpr is called when exiting the StructLiteralExpr production.
	ExitStructLiteralExpr(c *StructLiteralExprContext)

	// ExitBuiltinCall is called when exiting the BuiltinCall production.
	ExitBuiltinCall(c *BuiltinCallContext)

	// ExitParenExpr is called when exiting the ParenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitSlice_literal is called when exiting the slice_literal production.
	ExitSlice_literal(c *Slice_literalContext)

	// ExitStruct_literal is called when exiting the struct_literal production.
	ExitStruct_literal(c *Struct_literalContext)

	// ExitField_assignment_list is called when exiting the field_assignment_list production.
	ExitField_assignment_list(c *Field_assignment_listContext)

	// ExitField_assignment is called when exiting the field_assignment production.
	ExitField_assignment(c *Field_assignmentContext)

	// ExitPrintCall is called when exiting the PrintCall production.
	ExitPrintCall(c *PrintCallContext)

	// ExitAtoiCall is called when exiting the AtoiCall production.
	ExitAtoiCall(c *AtoiCallContext)

	// ExitParseFloatCall is called when exiting the ParseFloatCall production.
	ExitParseFloatCall(c *ParseFloatCallContext)

	// ExitTypeOfCall is called when exiting the TypeOfCall production.
	ExitTypeOfCall(c *TypeOfCallContext)

	// ExitIndexOfCall is called when exiting the IndexOfCall production.
	ExitIndexOfCall(c *IndexOfCallContext)

	// ExitJoinCall is called when exiting the JoinCall production.
	ExitJoinCall(c *JoinCallContext)

	// ExitLenCall is called when exiting the LenCall production.
	ExitLenCall(c *LenCallContext)

	// ExitAppendCall is called when exiting the AppendCall production.
	ExitAppendCall(c *AppendCallContext)

	// ExitArgument_list is called when exiting the argument_list production.
	ExitArgument_list(c *Argument_listContext)

	// ExitExpression_list is called when exiting the expression_list production.
	ExitExpression_list(c *Expression_listContext)
}
