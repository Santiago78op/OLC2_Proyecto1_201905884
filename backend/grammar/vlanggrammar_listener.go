// Code generated from grammar/VLangGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // VLangGrammar
import "github.com/antlr4-go/antlr/v4"

// VLangGrammarListener is a complete listener for a parse tree produced by VLangGrammar.
type VLangGrammarListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterMutVarDecl is called when entering the MutVarDecl production.
	EnterMutVarDecl(c *MutVarDeclContext)

	// EnterValueDecl is called when entering the ValueDecl production.
	EnterValueDecl(c *ValueDeclContext)

	// EnterValDeclVec is called when entering the ValDeclVec production.
	EnterValDeclVec(c *ValDeclVecContext)

	// EnterVarAssDecl is called when entering the VarAssDecl production.
	EnterVarAssDecl(c *VarAssDeclContext)

	// EnterVarVectDecl is called when entering the VarVectDecl production.
	EnterVarVectDecl(c *VarVectDeclContext)

	// EnterVectorItemLis is called when entering the VectorItemLis production.
	EnterVectorItemLis(c *VectorItemLisContext)

	// EnterVectorItem is called when entering the VectorItem production.
	EnterVectorItem(c *VectorItemContext)

	// EnterVectorProperty is called when entering the VectorProperty production.
	EnterVectorProperty(c *VectorPropertyContext)

	// EnterVectorFuncCall is called when entering the VectorFuncCall production.
	EnterVectorFuncCall(c *VectorFuncCallContext)

	// EnterRepeatingDecl is called when entering the RepeatingDecl production.
	EnterRepeatingDecl(c *RepeatingDeclContext)

	// EnterVar_type is called when entering the var_type production.
	EnterVar_type(c *Var_typeContext)

	// EnterVector_type is called when entering the vector_type production.
	EnterVector_type(c *Vector_typeContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterMatrix_type is called when entering the matrix_type production.
	EnterMatrix_type(c *Matrix_typeContext)

	// EnterAux_matrix_type is called when entering the aux_matrix_type production.
	EnterAux_matrix_type(c *Aux_matrix_typeContext)

	// EnterAssignmentDecl is called when entering the AssignmentDecl production.
	EnterAssignmentDecl(c *AssignmentDeclContext)

	// EnterArgAddAssigDecl is called when entering the ArgAddAssigDecl production.
	EnterArgAddAssigDecl(c *ArgAddAssigDeclContext)

	// EnterVectorAssign is called when entering the VectorAssign production.
	EnterVectorAssign(c *VectorAssignContext)

	// EnterIdPattern is called when entering the IdPattern production.
	EnterIdPattern(c *IdPatternContext)

	// EnterIntLiteral is called when entering the IntLiteral production.
	EnterIntLiteral(c *IntLiteralContext)

	// EnterFloatLiteral is called when entering the FloatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// EnterStringLiteral is called when entering the StringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterBoolLiteral is called when entering the BoolLiteral production.
	EnterBoolLiteral(c *BoolLiteralContext)

	// EnterNilLiteral is called when entering the NilLiteral production.
	EnterNilLiteral(c *NilLiteralContext)

	// EnterIncremento is called when entering the incremento production.
	EnterIncremento(c *IncrementoContext)

	// EnterDecremento is called when entering the decremento production.
	EnterDecremento(c *DecrementoContext)

	// EnterRepeatingExpr is called when entering the RepeatingExpr production.
	EnterRepeatingExpr(c *RepeatingExprContext)

	// EnterVectorPropertyExpr is called when entering the VectorPropertyExpr production.
	EnterVectorPropertyExpr(c *VectorPropertyExprContext)

	// EnterIncredecr is called when entering the incredecr production.
	EnterIncredecr(c *IncredecrContext)

	// EnterVectorItemExpr is called when entering the VectorItemExpr production.
	EnterVectorItemExpr(c *VectorItemExprContext)

	// EnterBinaryExpr is called when entering the BinaryExpr production.
	EnterBinaryExpr(c *BinaryExprContext)

	// EnterParensExpr is called when entering the ParensExpr production.
	EnterParensExpr(c *ParensExprContext)

	// EnterLiteralExpr is called when entering the LiteralExpr production.
	EnterLiteralExpr(c *LiteralExprContext)

	// EnterVectorFuncCallExpr is called when entering the VectorFuncCallExpr production.
	EnterVectorFuncCallExpr(c *VectorFuncCallExprContext)

	// EnterVectorExpr is called when entering the VectorExpr production.
	EnterVectorExpr(c *VectorExprContext)

	// EnterUnaryExpr is called when entering the UnaryExpr production.
	EnterUnaryExpr(c *UnaryExprContext)

	// EnterIdPatternExpr is called when entering the IdPatternExpr production.
	EnterIdPatternExpr(c *IdPatternExprContext)

	// EnterFuncCallExpr is called when entering the FuncCallExpr production.
	EnterFuncCallExpr(c *FuncCallExprContext)

	// EnterIfStmt is called when entering the IfStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterIfChain is called when entering the IfChain production.
	EnterIfChain(c *IfChainContext)

	// EnterElseStmt is called when entering the ElseStmt production.
	EnterElseStmt(c *ElseStmtContext)

	// EnterSwitchStmt is called when entering the SwitchStmt production.
	EnterSwitchStmt(c *SwitchStmtContext)

	// EnterSwitchCase is called when entering the SwitchCase production.
	EnterSwitchCase(c *SwitchCaseContext)

	// EnterDefaultCase is called when entering the DefaultCase production.
	EnterDefaultCase(c *DefaultCaseContext)

	// EnterWhileStmt is called when entering the WhileStmt production.
	EnterWhileStmt(c *WhileStmtContext)

	// EnterForStmtCond is called when entering the ForStmtCond production.
	EnterForStmtCond(c *ForStmtCondContext)

	// EnterForAssCond is called when entering the ForAssCond production.
	EnterForAssCond(c *ForAssCondContext)

	// EnterForStmt is called when entering the ForStmt production.
	EnterForStmt(c *ForStmtContext)

	// EnterNumericRange is called when entering the NumericRange production.
	EnterNumericRange(c *NumericRangeContext)

	// EnterReturnStmt is called when entering the ReturnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterBreakStmt is called when entering the BreakStmt production.
	EnterBreakStmt(c *BreakStmtContext)

	// EnterContinueStmt is called when entering the ContinueStmt production.
	EnterContinueStmt(c *ContinueStmtContext)

	// EnterFuncCall is called when entering the FuncCall production.
	EnterFuncCall(c *FuncCallContext)

	// EnterBlockInd is called when entering the BlockInd production.
	EnterBlockInd(c *BlockIndContext)

	// EnterArgList is called when entering the ArgList production.
	EnterArgList(c *ArgListContext)

	// EnterFuncArg is called when entering the FuncArg production.
	EnterFuncArg(c *FuncArgContext)

	// EnterFuncDecl is called when entering the FuncDecl production.
	EnterFuncDecl(c *FuncDeclContext)

	// EnterParamList is called when entering the ParamList production.
	EnterParamList(c *ParamListContext)

	// EnterFuncParam is called when entering the FuncParam production.
	EnterFuncParam(c *FuncParamContext)

	// EnterStructDecl is called when entering the StructDecl production.
	EnterStructDecl(c *StructDeclContext)

	// EnterStructAttr is called when entering the StructAttr production.
	EnterStructAttr(c *StructAttrContext)

	// EnterStructVector is called when entering the StructVector production.
	EnterStructVector(c *StructVectorContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitMutVarDecl is called when exiting the MutVarDecl production.
	ExitMutVarDecl(c *MutVarDeclContext)

	// ExitValueDecl is called when exiting the ValueDecl production.
	ExitValueDecl(c *ValueDeclContext)

	// ExitValDeclVec is called when exiting the ValDeclVec production.
	ExitValDeclVec(c *ValDeclVecContext)

	// ExitVarAssDecl is called when exiting the VarAssDecl production.
	ExitVarAssDecl(c *VarAssDeclContext)

	// ExitVarVectDecl is called when exiting the VarVectDecl production.
	ExitVarVectDecl(c *VarVectDeclContext)

	// ExitVectorItemLis is called when exiting the VectorItemLis production.
	ExitVectorItemLis(c *VectorItemLisContext)

	// ExitVectorItem is called when exiting the VectorItem production.
	ExitVectorItem(c *VectorItemContext)

	// ExitVectorProperty is called when exiting the VectorProperty production.
	ExitVectorProperty(c *VectorPropertyContext)

	// ExitVectorFuncCall is called when exiting the VectorFuncCall production.
	ExitVectorFuncCall(c *VectorFuncCallContext)

	// ExitRepeatingDecl is called when exiting the RepeatingDecl production.
	ExitRepeatingDecl(c *RepeatingDeclContext)

	// ExitVar_type is called when exiting the var_type production.
	ExitVar_type(c *Var_typeContext)

	// ExitVector_type is called when exiting the vector_type production.
	ExitVector_type(c *Vector_typeContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitMatrix_type is called when exiting the matrix_type production.
	ExitMatrix_type(c *Matrix_typeContext)

	// ExitAux_matrix_type is called when exiting the aux_matrix_type production.
	ExitAux_matrix_type(c *Aux_matrix_typeContext)

	// ExitAssignmentDecl is called when exiting the AssignmentDecl production.
	ExitAssignmentDecl(c *AssignmentDeclContext)

	// ExitArgAddAssigDecl is called when exiting the ArgAddAssigDecl production.
	ExitArgAddAssigDecl(c *ArgAddAssigDeclContext)

	// ExitVectorAssign is called when exiting the VectorAssign production.
	ExitVectorAssign(c *VectorAssignContext)

	// ExitIdPattern is called when exiting the IdPattern production.
	ExitIdPattern(c *IdPatternContext)

	// ExitIntLiteral is called when exiting the IntLiteral production.
	ExitIntLiteral(c *IntLiteralContext)

	// ExitFloatLiteral is called when exiting the FloatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)

	// ExitStringLiteral is called when exiting the StringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitBoolLiteral is called when exiting the BoolLiteral production.
	ExitBoolLiteral(c *BoolLiteralContext)

	// ExitNilLiteral is called when exiting the NilLiteral production.
	ExitNilLiteral(c *NilLiteralContext)

	// ExitIncremento is called when exiting the incremento production.
	ExitIncremento(c *IncrementoContext)

	// ExitDecremento is called when exiting the decremento production.
	ExitDecremento(c *DecrementoContext)

	// ExitRepeatingExpr is called when exiting the RepeatingExpr production.
	ExitRepeatingExpr(c *RepeatingExprContext)

	// ExitVectorPropertyExpr is called when exiting the VectorPropertyExpr production.
	ExitVectorPropertyExpr(c *VectorPropertyExprContext)

	// ExitIncredecr is called when exiting the incredecr production.
	ExitIncredecr(c *IncredecrContext)

	// ExitVectorItemExpr is called when exiting the VectorItemExpr production.
	ExitVectorItemExpr(c *VectorItemExprContext)

	// ExitBinaryExpr is called when exiting the BinaryExpr production.
	ExitBinaryExpr(c *BinaryExprContext)

	// ExitParensExpr is called when exiting the ParensExpr production.
	ExitParensExpr(c *ParensExprContext)

	// ExitLiteralExpr is called when exiting the LiteralExpr production.
	ExitLiteralExpr(c *LiteralExprContext)

	// ExitVectorFuncCallExpr is called when exiting the VectorFuncCallExpr production.
	ExitVectorFuncCallExpr(c *VectorFuncCallExprContext)

	// ExitVectorExpr is called when exiting the VectorExpr production.
	ExitVectorExpr(c *VectorExprContext)

	// ExitUnaryExpr is called when exiting the UnaryExpr production.
	ExitUnaryExpr(c *UnaryExprContext)

	// ExitIdPatternExpr is called when exiting the IdPatternExpr production.
	ExitIdPatternExpr(c *IdPatternExprContext)

	// ExitFuncCallExpr is called when exiting the FuncCallExpr production.
	ExitFuncCallExpr(c *FuncCallExprContext)

	// ExitIfStmt is called when exiting the IfStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitIfChain is called when exiting the IfChain production.
	ExitIfChain(c *IfChainContext)

	// ExitElseStmt is called when exiting the ElseStmt production.
	ExitElseStmt(c *ElseStmtContext)

	// ExitSwitchStmt is called when exiting the SwitchStmt production.
	ExitSwitchStmt(c *SwitchStmtContext)

	// ExitSwitchCase is called when exiting the SwitchCase production.
	ExitSwitchCase(c *SwitchCaseContext)

	// ExitDefaultCase is called when exiting the DefaultCase production.
	ExitDefaultCase(c *DefaultCaseContext)

	// ExitWhileStmt is called when exiting the WhileStmt production.
	ExitWhileStmt(c *WhileStmtContext)

	// ExitForStmtCond is called when exiting the ForStmtCond production.
	ExitForStmtCond(c *ForStmtCondContext)

	// ExitForAssCond is called when exiting the ForAssCond production.
	ExitForAssCond(c *ForAssCondContext)

	// ExitForStmt is called when exiting the ForStmt production.
	ExitForStmt(c *ForStmtContext)

	// ExitNumericRange is called when exiting the NumericRange production.
	ExitNumericRange(c *NumericRangeContext)

	// ExitReturnStmt is called when exiting the ReturnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitBreakStmt is called when exiting the BreakStmt production.
	ExitBreakStmt(c *BreakStmtContext)

	// ExitContinueStmt is called when exiting the ContinueStmt production.
	ExitContinueStmt(c *ContinueStmtContext)

	// ExitFuncCall is called when exiting the FuncCall production.
	ExitFuncCall(c *FuncCallContext)

	// ExitBlockInd is called when exiting the BlockInd production.
	ExitBlockInd(c *BlockIndContext)

	// ExitArgList is called when exiting the ArgList production.
	ExitArgList(c *ArgListContext)

	// ExitFuncArg is called when exiting the FuncArg production.
	ExitFuncArg(c *FuncArgContext)

	// ExitFuncDecl is called when exiting the FuncDecl production.
	ExitFuncDecl(c *FuncDeclContext)

	// ExitParamList is called when exiting the ParamList production.
	ExitParamList(c *ParamListContext)

	// ExitFuncParam is called when exiting the FuncParam production.
	ExitFuncParam(c *FuncParamContext)

	// ExitStructDecl is called when exiting the StructDecl production.
	ExitStructDecl(c *StructDeclContext)

	// ExitStructAttr is called when exiting the StructAttr production.
	ExitStructAttr(c *StructAttrContext)

	// ExitStructVector is called when exiting the StructVector production.
	ExitStructVector(c *StructVectorContext)
}
