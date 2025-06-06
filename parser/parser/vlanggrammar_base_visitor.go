// Code generated from C:/Users/72358/Documents/Compi_2/Semana_01/awesomeProject/grammar/VLangGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // VLangGrammar
import "github.com/antlr4-go/antlr/v4"

type BaseVLangGrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseVLangGrammarVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitMutableVarDecl(ctx *MutableVarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitTypedVarDecl(ctx *TypedVarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitInferredVarDecl(ctx *InferredVarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitType_annotation(ctx *Type_annotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitSlice_type(ctx *Slice_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitFunction_declaration(ctx *Function_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitReceiver(ctx *ReceiverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitParameter_list(ctx *Parameter_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitStruct_declaration(ctx *Struct_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitStruct_field(ctx *Struct_fieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitSimpleAssignment(ctx *SimpleAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitPlusAssignment(ctx *PlusAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitMinusAssignment(ctx *MinusAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitIndexAssignment(ctx *IndexAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitFieldAssignment(ctx *FieldAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitIf_statement(ctx *If_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitForCondition(ctx *ForConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitForLoop(ctx *ForLoopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitForIn(ctx *ForInContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitSwitch_statement(ctx *Switch_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitCase_clause(ctx *Case_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitDefault_clause(ctx *Default_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitBreak_statement(ctx *Break_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitContinue_statement(ctx *Continue_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitReturn_statement(ctx *Return_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitExpression_statement(ctx *Expression_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitBlock_statement(ctx *Block_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitUnaryNot(ctx *UnaryNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitMulDivMod(ctx *MulDivModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitIndexAccess(ctx *IndexAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitLogicalAnd(ctx *LogicalAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitRelational(ctx *RelationalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitUnaryMinus(ctx *UnaryMinusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitEquality(ctx *EqualityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitFieldAccess(ctx *FieldAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitLogicalOr(ctx *LogicalOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitIdentifierExpr(ctx *IdentifierExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitIntLiteral(ctx *IntLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitFloatLiteral(ctx *FloatLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitRuneLiteral(ctx *RuneLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitTrueLiteral(ctx *TrueLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitFalseLiteral(ctx *FalseLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitNilLiteral(ctx *NilLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitSliceLiteralExpr(ctx *SliceLiteralExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitStructLiteralExpr(ctx *StructLiteralExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitBuiltinCall(ctx *BuiltinCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitSlice_literal(ctx *Slice_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitStruct_literal(ctx *Struct_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitField_assignment_list(ctx *Field_assignment_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitField_assignment(ctx *Field_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitPrintCall(ctx *PrintCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitAtoiCall(ctx *AtoiCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitParseFloatCall(ctx *ParseFloatCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitTypeOfCall(ctx *TypeOfCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitIndexOfCall(ctx *IndexOfCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitJoinCall(ctx *JoinCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitLenCall(ctx *LenCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitAppendCall(ctx *AppendCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitArgument_list(ctx *Argument_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVLangGrammarVisitor) VisitExpression_list(ctx *Expression_listContext) interface{} {
	return v.VisitChildren(ctx)
}
