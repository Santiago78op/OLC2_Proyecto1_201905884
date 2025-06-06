// Code generated from C:/Users/72358/Documents/Compi_2/Semana_01/awesomeProject/grammar/VLangGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // VLangGrammar
import "github.com/antlr4-go/antlr/v4"

// BaseVLangGrammarListener is a complete listener for a parse tree produced by VLangGrammarParser.
type BaseVLangGrammarListener struct{}

var _ VLangGrammarListener = &BaseVLangGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVLangGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVLangGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVLangGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVLangGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseVLangGrammarListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseVLangGrammarListener) ExitProgram(ctx *ProgramContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseVLangGrammarListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseVLangGrammarListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterMutableVarDecl is called when production MutableVarDecl is entered.
func (s *BaseVLangGrammarListener) EnterMutableVarDecl(ctx *MutableVarDeclContext) {}

// ExitMutableVarDecl is called when production MutableVarDecl is exited.
func (s *BaseVLangGrammarListener) ExitMutableVarDecl(ctx *MutableVarDeclContext) {}

// EnterTypedVarDecl is called when production TypedVarDecl is entered.
func (s *BaseVLangGrammarListener) EnterTypedVarDecl(ctx *TypedVarDeclContext) {}

// ExitTypedVarDecl is called when production TypedVarDecl is exited.
func (s *BaseVLangGrammarListener) ExitTypedVarDecl(ctx *TypedVarDeclContext) {}

// EnterInferredVarDecl is called when production InferredVarDecl is entered.
func (s *BaseVLangGrammarListener) EnterInferredVarDecl(ctx *InferredVarDeclContext) {}

// ExitInferredVarDecl is called when production InferredVarDecl is exited.
func (s *BaseVLangGrammarListener) ExitInferredVarDecl(ctx *InferredVarDeclContext) {}

// EnterType_annotation is called when production type_annotation is entered.
func (s *BaseVLangGrammarListener) EnterType_annotation(ctx *Type_annotationContext) {}

// ExitType_annotation is called when production type_annotation is exited.
func (s *BaseVLangGrammarListener) ExitType_annotation(ctx *Type_annotationContext) {}

// EnterSlice_type is called when production slice_type is entered.
func (s *BaseVLangGrammarListener) EnterSlice_type(ctx *Slice_typeContext) {}

// ExitSlice_type is called when production slice_type is exited.
func (s *BaseVLangGrammarListener) ExitSlice_type(ctx *Slice_typeContext) {}

// EnterFunction_declaration is called when production function_declaration is entered.
func (s *BaseVLangGrammarListener) EnterFunction_declaration(ctx *Function_declarationContext) {}

// ExitFunction_declaration is called when production function_declaration is exited.
func (s *BaseVLangGrammarListener) ExitFunction_declaration(ctx *Function_declarationContext) {}

// EnterReceiver is called when production receiver is entered.
func (s *BaseVLangGrammarListener) EnterReceiver(ctx *ReceiverContext) {}

// ExitReceiver is called when production receiver is exited.
func (s *BaseVLangGrammarListener) ExitReceiver(ctx *ReceiverContext) {}

// EnterParameter_list is called when production parameter_list is entered.
func (s *BaseVLangGrammarListener) EnterParameter_list(ctx *Parameter_listContext) {}

// ExitParameter_list is called when production parameter_list is exited.
func (s *BaseVLangGrammarListener) ExitParameter_list(ctx *Parameter_listContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseVLangGrammarListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseVLangGrammarListener) ExitParameter(ctx *ParameterContext) {}

// EnterStruct_declaration is called when production struct_declaration is entered.
func (s *BaseVLangGrammarListener) EnterStruct_declaration(ctx *Struct_declarationContext) {}

// ExitStruct_declaration is called when production struct_declaration is exited.
func (s *BaseVLangGrammarListener) ExitStruct_declaration(ctx *Struct_declarationContext) {}

// EnterStruct_field is called when production struct_field is entered.
func (s *BaseVLangGrammarListener) EnterStruct_field(ctx *Struct_fieldContext) {}

// ExitStruct_field is called when production struct_field is exited.
func (s *BaseVLangGrammarListener) ExitStruct_field(ctx *Struct_fieldContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseVLangGrammarListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseVLangGrammarListener) ExitStatement(ctx *StatementContext) {}

// EnterSimpleAssignment is called when production SimpleAssignment is entered.
func (s *BaseVLangGrammarListener) EnterSimpleAssignment(ctx *SimpleAssignmentContext) {}

// ExitSimpleAssignment is called when production SimpleAssignment is exited.
func (s *BaseVLangGrammarListener) ExitSimpleAssignment(ctx *SimpleAssignmentContext) {}

// EnterPlusAssignment is called when production PlusAssignment is entered.
func (s *BaseVLangGrammarListener) EnterPlusAssignment(ctx *PlusAssignmentContext) {}

// ExitPlusAssignment is called when production PlusAssignment is exited.
func (s *BaseVLangGrammarListener) ExitPlusAssignment(ctx *PlusAssignmentContext) {}

// EnterMinusAssignment is called when production MinusAssignment is entered.
func (s *BaseVLangGrammarListener) EnterMinusAssignment(ctx *MinusAssignmentContext) {}

// ExitMinusAssignment is called when production MinusAssignment is exited.
func (s *BaseVLangGrammarListener) ExitMinusAssignment(ctx *MinusAssignmentContext) {}

// EnterIndexAssignment is called when production IndexAssignment is entered.
func (s *BaseVLangGrammarListener) EnterIndexAssignment(ctx *IndexAssignmentContext) {}

// ExitIndexAssignment is called when production IndexAssignment is exited.
func (s *BaseVLangGrammarListener) ExitIndexAssignment(ctx *IndexAssignmentContext) {}

// EnterFieldAssignment is called when production FieldAssignment is entered.
func (s *BaseVLangGrammarListener) EnterFieldAssignment(ctx *FieldAssignmentContext) {}

// ExitFieldAssignment is called when production FieldAssignment is exited.
func (s *BaseVLangGrammarListener) ExitFieldAssignment(ctx *FieldAssignmentContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *BaseVLangGrammarListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BaseVLangGrammarListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterForCondition is called when production ForCondition is entered.
func (s *BaseVLangGrammarListener) EnterForCondition(ctx *ForConditionContext) {}

// ExitForCondition is called when production ForCondition is exited.
func (s *BaseVLangGrammarListener) ExitForCondition(ctx *ForConditionContext) {}

// EnterForLoop is called when production ForLoop is entered.
func (s *BaseVLangGrammarListener) EnterForLoop(ctx *ForLoopContext) {}

// ExitForLoop is called when production ForLoop is exited.
func (s *BaseVLangGrammarListener) ExitForLoop(ctx *ForLoopContext) {}

// EnterForIn is called when production ForIn is entered.
func (s *BaseVLangGrammarListener) EnterForIn(ctx *ForInContext) {}

// ExitForIn is called when production ForIn is exited.
func (s *BaseVLangGrammarListener) ExitForIn(ctx *ForInContext) {}

// EnterSwitch_statement is called when production switch_statement is entered.
func (s *BaseVLangGrammarListener) EnterSwitch_statement(ctx *Switch_statementContext) {}

// ExitSwitch_statement is called when production switch_statement is exited.
func (s *BaseVLangGrammarListener) ExitSwitch_statement(ctx *Switch_statementContext) {}

// EnterCase_clause is called when production case_clause is entered.
func (s *BaseVLangGrammarListener) EnterCase_clause(ctx *Case_clauseContext) {}

// ExitCase_clause is called when production case_clause is exited.
func (s *BaseVLangGrammarListener) ExitCase_clause(ctx *Case_clauseContext) {}

// EnterDefault_clause is called when production default_clause is entered.
func (s *BaseVLangGrammarListener) EnterDefault_clause(ctx *Default_clauseContext) {}

// ExitDefault_clause is called when production default_clause is exited.
func (s *BaseVLangGrammarListener) ExitDefault_clause(ctx *Default_clauseContext) {}

// EnterBreak_statement is called when production break_statement is entered.
func (s *BaseVLangGrammarListener) EnterBreak_statement(ctx *Break_statementContext) {}

// ExitBreak_statement is called when production break_statement is exited.
func (s *BaseVLangGrammarListener) ExitBreak_statement(ctx *Break_statementContext) {}

// EnterContinue_statement is called when production continue_statement is entered.
func (s *BaseVLangGrammarListener) EnterContinue_statement(ctx *Continue_statementContext) {}

// ExitContinue_statement is called when production continue_statement is exited.
func (s *BaseVLangGrammarListener) ExitContinue_statement(ctx *Continue_statementContext) {}

// EnterReturn_statement is called when production return_statement is entered.
func (s *BaseVLangGrammarListener) EnterReturn_statement(ctx *Return_statementContext) {}

// ExitReturn_statement is called when production return_statement is exited.
func (s *BaseVLangGrammarListener) ExitReturn_statement(ctx *Return_statementContext) {}

// EnterExpression_statement is called when production expression_statement is entered.
func (s *BaseVLangGrammarListener) EnterExpression_statement(ctx *Expression_statementContext) {}

// ExitExpression_statement is called when production expression_statement is exited.
func (s *BaseVLangGrammarListener) ExitExpression_statement(ctx *Expression_statementContext) {}

// EnterBlock_statement is called when production block_statement is entered.
func (s *BaseVLangGrammarListener) EnterBlock_statement(ctx *Block_statementContext) {}

// ExitBlock_statement is called when production block_statement is exited.
func (s *BaseVLangGrammarListener) ExitBlock_statement(ctx *Block_statementContext) {}

// EnterUnaryNot is called when production UnaryNot is entered.
func (s *BaseVLangGrammarListener) EnterUnaryNot(ctx *UnaryNotContext) {}

// ExitUnaryNot is called when production UnaryNot is exited.
func (s *BaseVLangGrammarListener) ExitUnaryNot(ctx *UnaryNotContext) {}

// EnterMulDivMod is called when production MulDivMod is entered.
func (s *BaseVLangGrammarListener) EnterMulDivMod(ctx *MulDivModContext) {}

// ExitMulDivMod is called when production MulDivMod is exited.
func (s *BaseVLangGrammarListener) ExitMulDivMod(ctx *MulDivModContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseVLangGrammarListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseVLangGrammarListener) ExitAddSub(ctx *AddSubContext) {}

// EnterIndexAccess is called when production IndexAccess is entered.
func (s *BaseVLangGrammarListener) EnterIndexAccess(ctx *IndexAccessContext) {}

// ExitIndexAccess is called when production IndexAccess is exited.
func (s *BaseVLangGrammarListener) ExitIndexAccess(ctx *IndexAccessContext) {}

// EnterPrimary is called when production Primary is entered.
func (s *BaseVLangGrammarListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production Primary is exited.
func (s *BaseVLangGrammarListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterLogicalAnd is called when production LogicalAnd is entered.
func (s *BaseVLangGrammarListener) EnterLogicalAnd(ctx *LogicalAndContext) {}

// ExitLogicalAnd is called when production LogicalAnd is exited.
func (s *BaseVLangGrammarListener) ExitLogicalAnd(ctx *LogicalAndContext) {}

// EnterRelational is called when production Relational is entered.
func (s *BaseVLangGrammarListener) EnterRelational(ctx *RelationalContext) {}

// ExitRelational is called when production Relational is exited.
func (s *BaseVLangGrammarListener) ExitRelational(ctx *RelationalContext) {}

// EnterUnaryMinus is called when production UnaryMinus is entered.
func (s *BaseVLangGrammarListener) EnterUnaryMinus(ctx *UnaryMinusContext) {}

// ExitUnaryMinus is called when production UnaryMinus is exited.
func (s *BaseVLangGrammarListener) ExitUnaryMinus(ctx *UnaryMinusContext) {}

// EnterEquality is called when production Equality is entered.
func (s *BaseVLangGrammarListener) EnterEquality(ctx *EqualityContext) {}

// ExitEquality is called when production Equality is exited.
func (s *BaseVLangGrammarListener) ExitEquality(ctx *EqualityContext) {}

// EnterFunctionCall is called when production FunctionCall is entered.
func (s *BaseVLangGrammarListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production FunctionCall is exited.
func (s *BaseVLangGrammarListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterFieldAccess is called when production FieldAccess is entered.
func (s *BaseVLangGrammarListener) EnterFieldAccess(ctx *FieldAccessContext) {}

// ExitFieldAccess is called when production FieldAccess is exited.
func (s *BaseVLangGrammarListener) ExitFieldAccess(ctx *FieldAccessContext) {}

// EnterLogicalOr is called when production LogicalOr is entered.
func (s *BaseVLangGrammarListener) EnterLogicalOr(ctx *LogicalOrContext) {}

// ExitLogicalOr is called when production LogicalOr is exited.
func (s *BaseVLangGrammarListener) ExitLogicalOr(ctx *LogicalOrContext) {}

// EnterIdentifierExpr is called when production IdentifierExpr is entered.
func (s *BaseVLangGrammarListener) EnterIdentifierExpr(ctx *IdentifierExprContext) {}

// ExitIdentifierExpr is called when production IdentifierExpr is exited.
func (s *BaseVLangGrammarListener) ExitIdentifierExpr(ctx *IdentifierExprContext) {}

// EnterIntLiteral is called when production IntLiteral is entered.
func (s *BaseVLangGrammarListener) EnterIntLiteral(ctx *IntLiteralContext) {}

// ExitIntLiteral is called when production IntLiteral is exited.
func (s *BaseVLangGrammarListener) ExitIntLiteral(ctx *IntLiteralContext) {}

// EnterFloatLiteral is called when production FloatLiteral is entered.
func (s *BaseVLangGrammarListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production FloatLiteral is exited.
func (s *BaseVLangGrammarListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterStringLiteral is called when production StringLiteral is entered.
func (s *BaseVLangGrammarListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production StringLiteral is exited.
func (s *BaseVLangGrammarListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterRuneLiteral is called when production RuneLiteral is entered.
func (s *BaseVLangGrammarListener) EnterRuneLiteral(ctx *RuneLiteralContext) {}

// ExitRuneLiteral is called when production RuneLiteral is exited.
func (s *BaseVLangGrammarListener) ExitRuneLiteral(ctx *RuneLiteralContext) {}

// EnterTrueLiteral is called when production TrueLiteral is entered.
func (s *BaseVLangGrammarListener) EnterTrueLiteral(ctx *TrueLiteralContext) {}

// ExitTrueLiteral is called when production TrueLiteral is exited.
func (s *BaseVLangGrammarListener) ExitTrueLiteral(ctx *TrueLiteralContext) {}

// EnterFalseLiteral is called when production FalseLiteral is entered.
func (s *BaseVLangGrammarListener) EnterFalseLiteral(ctx *FalseLiteralContext) {}

// ExitFalseLiteral is called when production FalseLiteral is exited.
func (s *BaseVLangGrammarListener) ExitFalseLiteral(ctx *FalseLiteralContext) {}

// EnterNilLiteral is called when production NilLiteral is entered.
func (s *BaseVLangGrammarListener) EnterNilLiteral(ctx *NilLiteralContext) {}

// ExitNilLiteral is called when production NilLiteral is exited.
func (s *BaseVLangGrammarListener) ExitNilLiteral(ctx *NilLiteralContext) {}

// EnterSliceLiteralExpr is called when production SliceLiteralExpr is entered.
func (s *BaseVLangGrammarListener) EnterSliceLiteralExpr(ctx *SliceLiteralExprContext) {}

// ExitSliceLiteralExpr is called when production SliceLiteralExpr is exited.
func (s *BaseVLangGrammarListener) ExitSliceLiteralExpr(ctx *SliceLiteralExprContext) {}

// EnterStructLiteralExpr is called when production StructLiteralExpr is entered.
func (s *BaseVLangGrammarListener) EnterStructLiteralExpr(ctx *StructLiteralExprContext) {}

// ExitStructLiteralExpr is called when production StructLiteralExpr is exited.
func (s *BaseVLangGrammarListener) ExitStructLiteralExpr(ctx *StructLiteralExprContext) {}

// EnterBuiltinCall is called when production BuiltinCall is entered.
func (s *BaseVLangGrammarListener) EnterBuiltinCall(ctx *BuiltinCallContext) {}

// ExitBuiltinCall is called when production BuiltinCall is exited.
func (s *BaseVLangGrammarListener) ExitBuiltinCall(ctx *BuiltinCallContext) {}

// EnterParenExpr is called when production ParenExpr is entered.
func (s *BaseVLangGrammarListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production ParenExpr is exited.
func (s *BaseVLangGrammarListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterSlice_literal is called when production slice_literal is entered.
func (s *BaseVLangGrammarListener) EnterSlice_literal(ctx *Slice_literalContext) {}

// ExitSlice_literal is called when production slice_literal is exited.
func (s *BaseVLangGrammarListener) ExitSlice_literal(ctx *Slice_literalContext) {}

// EnterStruct_literal is called when production struct_literal is entered.
func (s *BaseVLangGrammarListener) EnterStruct_literal(ctx *Struct_literalContext) {}

// ExitStruct_literal is called when production struct_literal is exited.
func (s *BaseVLangGrammarListener) ExitStruct_literal(ctx *Struct_literalContext) {}

// EnterField_assignment_list is called when production field_assignment_list is entered.
func (s *BaseVLangGrammarListener) EnterField_assignment_list(ctx *Field_assignment_listContext) {}

// ExitField_assignment_list is called when production field_assignment_list is exited.
func (s *BaseVLangGrammarListener) ExitField_assignment_list(ctx *Field_assignment_listContext) {}

// EnterField_assignment is called when production field_assignment is entered.
func (s *BaseVLangGrammarListener) EnterField_assignment(ctx *Field_assignmentContext) {}

// ExitField_assignment is called when production field_assignment is exited.
func (s *BaseVLangGrammarListener) ExitField_assignment(ctx *Field_assignmentContext) {}

// EnterPrintCall is called when production PrintCall is entered.
func (s *BaseVLangGrammarListener) EnterPrintCall(ctx *PrintCallContext) {}

// ExitPrintCall is called when production PrintCall is exited.
func (s *BaseVLangGrammarListener) ExitPrintCall(ctx *PrintCallContext) {}

// EnterAtoiCall is called when production AtoiCall is entered.
func (s *BaseVLangGrammarListener) EnterAtoiCall(ctx *AtoiCallContext) {}

// ExitAtoiCall is called when production AtoiCall is exited.
func (s *BaseVLangGrammarListener) ExitAtoiCall(ctx *AtoiCallContext) {}

// EnterParseFloatCall is called when production ParseFloatCall is entered.
func (s *BaseVLangGrammarListener) EnterParseFloatCall(ctx *ParseFloatCallContext) {}

// ExitParseFloatCall is called when production ParseFloatCall is exited.
func (s *BaseVLangGrammarListener) ExitParseFloatCall(ctx *ParseFloatCallContext) {}

// EnterTypeOfCall is called when production TypeOfCall is entered.
func (s *BaseVLangGrammarListener) EnterTypeOfCall(ctx *TypeOfCallContext) {}

// ExitTypeOfCall is called when production TypeOfCall is exited.
func (s *BaseVLangGrammarListener) ExitTypeOfCall(ctx *TypeOfCallContext) {}

// EnterIndexOfCall is called when production IndexOfCall is entered.
func (s *BaseVLangGrammarListener) EnterIndexOfCall(ctx *IndexOfCallContext) {}

// ExitIndexOfCall is called when production IndexOfCall is exited.
func (s *BaseVLangGrammarListener) ExitIndexOfCall(ctx *IndexOfCallContext) {}

// EnterJoinCall is called when production JoinCall is entered.
func (s *BaseVLangGrammarListener) EnterJoinCall(ctx *JoinCallContext) {}

// ExitJoinCall is called when production JoinCall is exited.
func (s *BaseVLangGrammarListener) ExitJoinCall(ctx *JoinCallContext) {}

// EnterLenCall is called when production LenCall is entered.
func (s *BaseVLangGrammarListener) EnterLenCall(ctx *LenCallContext) {}

// ExitLenCall is called when production LenCall is exited.
func (s *BaseVLangGrammarListener) ExitLenCall(ctx *LenCallContext) {}

// EnterAppendCall is called when production AppendCall is entered.
func (s *BaseVLangGrammarListener) EnterAppendCall(ctx *AppendCallContext) {}

// ExitAppendCall is called when production AppendCall is exited.
func (s *BaseVLangGrammarListener) ExitAppendCall(ctx *AppendCallContext) {}

// EnterArgument_list is called when production argument_list is entered.
func (s *BaseVLangGrammarListener) EnterArgument_list(ctx *Argument_listContext) {}

// ExitArgument_list is called when production argument_list is exited.
func (s *BaseVLangGrammarListener) ExitArgument_list(ctx *Argument_listContext) {}

// EnterExpression_list is called when production expression_list is entered.
func (s *BaseVLangGrammarListener) EnterExpression_list(ctx *Expression_listContext) {}

// ExitExpression_list is called when production expression_list is exited.
func (s *BaseVLangGrammarListener) ExitExpression_list(ctx *Expression_listContext) {}
