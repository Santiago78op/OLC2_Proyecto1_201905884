// Generated from /home/julian/Documents/OLC2_Proyecto1_201905884/backend/grammar/VLangParse.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class VLangParseParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, MUT=2, PRINT=3, PRINTLN=4, NIL=5, INT_TYPE=6, FLOAT_TYPE=7, STRING_TYPE=8, 
		BOOL_TYPE=9, SLICE_TYPE=10, TRUE=11, FALSE=12, PLUS=13, MINUS=14, MULT=15, 
		DIV=16, MOD=17, ASSIGN=18, PLUS_ASSIGN=19, MINUS_ASSIGN=20, EQ=21, NE=22, 
		LT=23, LE=24, GT=25, GE=26, AND=27, OR=28, NOT=29, LPAREN=30, RPAREN=31, 
		COMMA=32, INT_LITERAL=33, FLOAT_LITERAL=34, STRING_LITERAL=35, ID=36, 
		LINE_COMMENT=37, BLOCK_COMMENT=38, WS=39;
	public static final int
		RULE_prog = 0, RULE_delim = 1, RULE_stmt = 2, RULE_declaration = 3, RULE_variable_declaration = 4, 
		RULE_type_annotation = 5, RULE_assignment_declaration = 6, RULE_expression = 7, 
		RULE_primary_expression = 8, RULE_builtin_function_call = 9, RULE_argument_list = 10;
	private static String[] makeRuleNames() {
		return new String[] {
			"prog", "delim", "stmt", "declaration", "variable_declaration", "type_annotation", 
			"assignment_declaration", "expression", "primary_expression", "builtin_function_call", 
			"argument_list"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'\\n'", "'mut'", "'print'", "'println'", "'nil'", "'int'", "'float64'", 
			"'string'", "'bool'", "'slice'", "'true'", "'false'", "'+'", "'-'", "'*'", 
			"'/'", "'%'", "'='", "'+='", "'-='", "'=='", "'!='", "'<'", "'<='", "'>'", 
			"'>='", "'&&'", "'||'", "'!'", "'('", "')'", "','"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, "MUT", "PRINT", "PRINTLN", "NIL", "INT_TYPE", "FLOAT_TYPE", 
			"STRING_TYPE", "BOOL_TYPE", "SLICE_TYPE", "TRUE", "FALSE", "PLUS", "MINUS", 
			"MULT", "DIV", "MOD", "ASSIGN", "PLUS_ASSIGN", "MINUS_ASSIGN", "EQ", 
			"NE", "LT", "LE", "GT", "GE", "AND", "OR", "NOT", "LPAREN", "RPAREN", 
			"COMMA", "INT_LITERAL", "FLOAT_LITERAL", "STRING_LITERAL", "ID", "LINE_COMMENT", 
			"BLOCK_COMMENT", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "VLangParse.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public VLangParseParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgContext extends ParserRuleContext {
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public TerminalNode EOF() { return getToken(VLangParseParser.EOF, 0); }
		public ProgContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prog; }
	}

	public final ProgContext prog() throws RecognitionException {
		ProgContext _localctx = new ProgContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_prog);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(25);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==MUT || _la==ID) {
				{
				{
				setState(22);
				stmt();
				}
				}
				setState(27);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(29);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				{
				setState(28);
				match(EOF);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DelimContext extends ParserRuleContext {
		public DelimContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_delim; }
	}

	public final DelimContext delim() throws RecognitionException {
		DelimContext _localctx = new DelimContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_delim);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(31);
			match(T__0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StmtContext extends ParserRuleContext {
		public DeclarationContext declaration() {
			return getRuleContext(DeclarationContext.class,0);
		}
		public DelimContext delim() {
			return getRuleContext(DelimContext.class,0);
		}
		public StmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmt; }
	}

	public final StmtContext stmt() throws RecognitionException {
		StmtContext _localctx = new StmtContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_stmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(33);
			declaration();
			setState(34);
			delim();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DeclarationContext extends ParserRuleContext {
		public Variable_declarationContext variable_declaration() {
			return getRuleContext(Variable_declarationContext.class,0);
		}
		public Assignment_declarationContext assignment_declaration() {
			return getRuleContext(Assignment_declarationContext.class,0);
		}
		public DeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaration; }
	}

	public final DeclarationContext declaration() throws RecognitionException {
		DeclarationContext _localctx = new DeclarationContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_declaration);
		try {
			setState(38);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(36);
				variable_declaration();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(37);
				assignment_declaration();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Variable_declarationContext extends ParserRuleContext {
		public Variable_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variable_declaration; }
	 
		public Variable_declarationContext() { }
		public void copyFrom(Variable_declarationContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TypedVarDeclContext extends Variable_declarationContext {
		public TerminalNode ID() { return getToken(VLangParseParser.ID, 0); }
		public Type_annotationContext type_annotation() {
			return getRuleContext(Type_annotationContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(VLangParseParser.ASSIGN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TypedVarDeclContext(Variable_declarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MutableVarDeclContext extends Variable_declarationContext {
		public TerminalNode MUT() { return getToken(VLangParseParser.MUT, 0); }
		public TerminalNode ID() { return getToken(VLangParseParser.ID, 0); }
		public Type_annotationContext type_annotation() {
			return getRuleContext(Type_annotationContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(VLangParseParser.ASSIGN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public MutableVarDeclContext(Variable_declarationContext ctx) { copyFrom(ctx); }
	}

	public final Variable_declarationContext variable_declaration() throws RecognitionException {
		Variable_declarationContext _localctx = new Variable_declarationContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_variable_declaration);
		int _la;
		try {
			setState(52);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MUT:
				_localctx = new MutableVarDeclContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(40);
				match(MUT);
				setState(41);
				match(ID);
				setState(42);
				type_annotation();
				setState(45);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ASSIGN) {
					{
					setState(43);
					match(ASSIGN);
					setState(44);
					expression(0);
					}
				}

				}
				break;
			case ID:
				_localctx = new TypedVarDeclContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(47);
				match(ID);
				setState(48);
				type_annotation();
				setState(49);
				match(ASSIGN);
				setState(50);
				expression(0);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Type_annotationContext extends ParserRuleContext {
		public TerminalNode INT_TYPE() { return getToken(VLangParseParser.INT_TYPE, 0); }
		public TerminalNode FLOAT_TYPE() { return getToken(VLangParseParser.FLOAT_TYPE, 0); }
		public TerminalNode STRING_TYPE() { return getToken(VLangParseParser.STRING_TYPE, 0); }
		public TerminalNode BOOL_TYPE() { return getToken(VLangParseParser.BOOL_TYPE, 0); }
		public TerminalNode SLICE_TYPE() { return getToken(VLangParseParser.SLICE_TYPE, 0); }
		public TerminalNode ID() { return getToken(VLangParseParser.ID, 0); }
		public Type_annotationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_annotation; }
	}

	public final Type_annotationContext type_annotation() throws RecognitionException {
		Type_annotationContext _localctx = new Type_annotationContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_type_annotation);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(54);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 68719478720L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Assignment_declarationContext extends ParserRuleContext {
		public Assignment_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignment_declaration; }
	 
		public Assignment_declarationContext() { }
		public void copyFrom(Assignment_declarationContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MinusAssignmentDeclContext extends Assignment_declarationContext {
		public TerminalNode ID() { return getToken(VLangParseParser.ID, 0); }
		public TerminalNode MINUS_ASSIGN() { return getToken(VLangParseParser.MINUS_ASSIGN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public MinusAssignmentDeclContext(Assignment_declarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PlusAssignmentDeclContext extends Assignment_declarationContext {
		public TerminalNode ID() { return getToken(VLangParseParser.ID, 0); }
		public TerminalNode PLUS_ASSIGN() { return getToken(VLangParseParser.PLUS_ASSIGN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public PlusAssignmentDeclContext(Assignment_declarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentDeclContext extends Assignment_declarationContext {
		public TerminalNode ID() { return getToken(VLangParseParser.ID, 0); }
		public TerminalNode ASSIGN() { return getToken(VLangParseParser.ASSIGN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public AssignmentDeclContext(Assignment_declarationContext ctx) { copyFrom(ctx); }
	}

	public final Assignment_declarationContext assignment_declaration() throws RecognitionException {
		Assignment_declarationContext _localctx = new Assignment_declarationContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_assignment_declaration);
		try {
			setState(65);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				_localctx = new AssignmentDeclContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(56);
				match(ID);
				setState(57);
				match(ASSIGN);
				setState(58);
				expression(0);
				}
				break;
			case 2:
				_localctx = new PlusAssignmentDeclContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(59);
				match(ID);
				setState(60);
				match(PLUS_ASSIGN);
				setState(61);
				expression(0);
				}
				break;
			case 3:
				_localctx = new MinusAssignmentDeclContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(62);
				match(ID);
				setState(63);
				match(MINUS_ASSIGN);
				setState(64);
				expression(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionContext extends ParserRuleContext {
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	 
		public ExpressionContext() { }
		public void copyFrom(ExpressionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BinaryExprContext extends ExpressionContext {
		public ExpressionContext left;
		public Token op;
		public ExpressionContext right;
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode MULT() { return getToken(VLangParseParser.MULT, 0); }
		public TerminalNode DIV() { return getToken(VLangParseParser.DIV, 0); }
		public TerminalNode MOD() { return getToken(VLangParseParser.MOD, 0); }
		public TerminalNode PLUS() { return getToken(VLangParseParser.PLUS, 0); }
		public TerminalNode MINUS() { return getToken(VLangParseParser.MINUS, 0); }
		public TerminalNode LE() { return getToken(VLangParseParser.LE, 0); }
		public TerminalNode LT() { return getToken(VLangParseParser.LT, 0); }
		public TerminalNode GE() { return getToken(VLangParseParser.GE, 0); }
		public TerminalNode GT() { return getToken(VLangParseParser.GT, 0); }
		public TerminalNode EQ() { return getToken(VLangParseParser.EQ, 0); }
		public TerminalNode NE() { return getToken(VLangParseParser.NE, 0); }
		public TerminalNode AND() { return getToken(VLangParseParser.AND, 0); }
		public TerminalNode OR() { return getToken(VLangParseParser.OR, 0); }
		public BinaryExprContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrimaryExprContext extends ExpressionContext {
		public Primary_expressionContext primary_expression() {
			return getRuleContext(Primary_expressionContext.class,0);
		}
		public PrimaryExprContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParenthesizedExprContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode LPAREN() { return getToken(VLangParseParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(VLangParseParser.RPAREN, 0); }
		public ParenthesizedExprContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UnaryExprContext extends ExpressionContext {
		public Token op;
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode NOT() { return getToken(VLangParseParser.NOT, 0); }
		public TerminalNode MINUS() { return getToken(VLangParseParser.MINUS, 0); }
		public UnaryExprContext(ExpressionContext ctx) { copyFrom(ctx); }
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 14;
		enterRecursionRule(_localctx, 14, RULE_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(71);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PRINT:
			case PRINTLN:
			case NIL:
			case TRUE:
			case FALSE:
			case LPAREN:
			case INT_LITERAL:
			case FLOAT_LITERAL:
			case STRING_LITERAL:
			case ID:
				{
				_localctx = new PrimaryExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(68);
				primary_expression();
				}
				break;
			case MINUS:
			case NOT:
				{
				_localctx = new UnaryExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(69);
				((UnaryExprContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==MINUS || _la==NOT) ) {
					((UnaryExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(70);
				expression(7);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(98);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(96);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
					case 1:
						{
						_localctx = new BinaryExprContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(73);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(74);
						((BinaryExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 229376L) != 0)) ) {
							((BinaryExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(75);
						((BinaryExprContext)_localctx).right = expression(7);
						}
						break;
					case 2:
						{
						_localctx = new BinaryExprContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(76);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(77);
						((BinaryExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==PLUS || _la==MINUS) ) {
							((BinaryExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(78);
						((BinaryExprContext)_localctx).right = expression(6);
						}
						break;
					case 3:
						{
						_localctx = new BinaryExprContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(79);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(80);
						((BinaryExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 125829120L) != 0)) ) {
							((BinaryExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(81);
						((BinaryExprContext)_localctx).right = expression(5);
						}
						break;
					case 4:
						{
						_localctx = new BinaryExprContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(82);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(83);
						((BinaryExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==EQ || _la==NE) ) {
							((BinaryExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(84);
						((BinaryExprContext)_localctx).right = expression(4);
						}
						break;
					case 5:
						{
						_localctx = new BinaryExprContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(85);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(86);
						((BinaryExprContext)_localctx).op = match(AND);
						setState(87);
						((BinaryExprContext)_localctx).right = expression(3);
						}
						break;
					case 6:
						{
						_localctx = new BinaryExprContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(88);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(89);
						((BinaryExprContext)_localctx).op = match(OR);
						setState(90);
						((BinaryExprContext)_localctx).right = expression(2);
						}
						break;
					case 7:
						{
						_localctx = new ParenthesizedExprContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(91);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(92);
						match(LPAREN);
						setState(93);
						expression(0);
						setState(94);
						match(RPAREN);
						}
						break;
					}
					} 
				}
				setState(100);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Primary_expressionContext extends ParserRuleContext {
		public Primary_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primary_expression; }
	 
		public Primary_expressionContext() { }
		public void copyFrom(Primary_expressionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StringLiteralContext extends Primary_expressionContext {
		public TerminalNode STRING_LITERAL() { return getToken(VLangParseParser.STRING_LITERAL, 0); }
		public StringLiteralContext(Primary_expressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TrueLiteralContext extends Primary_expressionContext {
		public TerminalNode TRUE() { return getToken(VLangParseParser.TRUE, 0); }
		public TrueLiteralContext(Primary_expressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BuiltinCallContext extends Primary_expressionContext {
		public Builtin_function_callContext builtin_function_call() {
			return getRuleContext(Builtin_function_callContext.class,0);
		}
		public BuiltinCallContext(Primary_expressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdentifierExprContext extends Primary_expressionContext {
		public TerminalNode ID() { return getToken(VLangParseParser.ID, 0); }
		public IdentifierExprContext(Primary_expressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FloatLiteralContext extends Primary_expressionContext {
		public TerminalNode FLOAT_LITERAL() { return getToken(VLangParseParser.FLOAT_LITERAL, 0); }
		public FloatLiteralContext(Primary_expressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NilLiteralContext extends Primary_expressionContext {
		public TerminalNode NIL() { return getToken(VLangParseParser.NIL, 0); }
		public NilLiteralContext(Primary_expressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IntLiteralContext extends Primary_expressionContext {
		public TerminalNode INT_LITERAL() { return getToken(VLangParseParser.INT_LITERAL, 0); }
		public IntLiteralContext(Primary_expressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParenExprContext extends Primary_expressionContext {
		public TerminalNode LPAREN() { return getToken(VLangParseParser.LPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(VLangParseParser.RPAREN, 0); }
		public ParenExprContext(Primary_expressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FalseLiteralContext extends Primary_expressionContext {
		public TerminalNode FALSE() { return getToken(VLangParseParser.FALSE, 0); }
		public FalseLiteralContext(Primary_expressionContext ctx) { copyFrom(ctx); }
	}

	public final Primary_expressionContext primary_expression() throws RecognitionException {
		Primary_expressionContext _localctx = new Primary_expressionContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_primary_expression);
		try {
			setState(113);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ID:
				_localctx = new IdentifierExprContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(101);
				match(ID);
				}
				break;
			case INT_LITERAL:
				_localctx = new IntLiteralContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(102);
				match(INT_LITERAL);
				}
				break;
			case FLOAT_LITERAL:
				_localctx = new FloatLiteralContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(103);
				match(FLOAT_LITERAL);
				}
				break;
			case STRING_LITERAL:
				_localctx = new StringLiteralContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(104);
				match(STRING_LITERAL);
				}
				break;
			case TRUE:
				_localctx = new TrueLiteralContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(105);
				match(TRUE);
				}
				break;
			case FALSE:
				_localctx = new FalseLiteralContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(106);
				match(FALSE);
				}
				break;
			case NIL:
				_localctx = new NilLiteralContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(107);
				match(NIL);
				}
				break;
			case PRINT:
			case PRINTLN:
				_localctx = new BuiltinCallContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(108);
				builtin_function_call();
				}
				break;
			case LPAREN:
				_localctx = new ParenExprContext(_localctx);
				enterOuterAlt(_localctx, 9);
				{
				setState(109);
				match(LPAREN);
				setState(110);
				expression(0);
				setState(111);
				match(RPAREN);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Builtin_function_callContext extends ParserRuleContext {
		public Builtin_function_callContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_builtin_function_call; }
	 
		public Builtin_function_callContext() { }
		public void copyFrom(Builtin_function_callContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrintCallContext extends Builtin_function_callContext {
		public TerminalNode PRINT() { return getToken(VLangParseParser.PRINT, 0); }
		public TerminalNode LPAREN() { return getToken(VLangParseParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(VLangParseParser.RPAREN, 0); }
		public Argument_listContext argument_list() {
			return getRuleContext(Argument_listContext.class,0);
		}
		public PrintCallContext(Builtin_function_callContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrintlnCallContext extends Builtin_function_callContext {
		public TerminalNode PRINTLN() { return getToken(VLangParseParser.PRINTLN, 0); }
		public TerminalNode LPAREN() { return getToken(VLangParseParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(VLangParseParser.RPAREN, 0); }
		public Argument_listContext argument_list() {
			return getRuleContext(Argument_listContext.class,0);
		}
		public PrintlnCallContext(Builtin_function_callContext ctx) { copyFrom(ctx); }
	}

	public final Builtin_function_callContext builtin_function_call() throws RecognitionException {
		Builtin_function_callContext _localctx = new Builtin_function_callContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_builtin_function_call);
		int _la;
		try {
			setState(127);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PRINT:
				_localctx = new PrintCallContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(115);
				match(PRINT);
				setState(116);
				match(LPAREN);
				setState(118);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 130459654200L) != 0)) {
					{
					setState(117);
					argument_list();
					}
				}

				setState(120);
				match(RPAREN);
				}
				break;
			case PRINTLN:
				_localctx = new PrintlnCallContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(121);
				match(PRINTLN);
				setState(122);
				match(LPAREN);
				setState(124);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 130459654200L) != 0)) {
					{
					setState(123);
					argument_list();
					}
				}

				setState(126);
				match(RPAREN);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Argument_listContext extends ParserRuleContext {
		public Argument_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_argument_list; }
	 
		public Argument_listContext() { }
		public void copyFrom(Argument_listContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ArgListContext extends Argument_listContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(VLangParseParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(VLangParseParser.COMMA, i);
		}
		public ArgListContext(Argument_listContext ctx) { copyFrom(ctx); }
	}

	public final Argument_listContext argument_list() throws RecognitionException {
		Argument_listContext _localctx = new Argument_listContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_argument_list);
		int _la;
		try {
			_localctx = new ArgListContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(129);
			expression(0);
			setState(134);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(130);
				match(COMMA);
				setState(131);
				expression(0);
				}
				}
				setState(136);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 7:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 6);
		case 1:
			return precpred(_ctx, 5);
		case 2:
			return precpred(_ctx, 4);
		case 3:
			return precpred(_ctx, 3);
		case 4:
			return precpred(_ctx, 2);
		case 5:
			return precpred(_ctx, 1);
		case 6:
			return precpred(_ctx, 8);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001\'\u008a\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0001\u0000\u0005\u0000\u0018"+
		"\b\u0000\n\u0000\f\u0000\u001b\t\u0000\u0001\u0000\u0003\u0000\u001e\b"+
		"\u0000\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0003\u0001\u0003\u0003\u0003\'\b\u0003\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0003\u0004.\b\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u00045\b\u0004\u0001"+
		"\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006B\b"+
		"\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0003\u0007H\b"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0005"+
		"\u0007a\b\u0007\n\u0007\f\u0007d\t\u0007\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0003"+
		"\br\b\b\u0001\t\u0001\t\u0001\t\u0003\tw\b\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0003\t}\b\t\u0001\t\u0003\t\u0080\b\t\u0001\n\u0001\n\u0001\n\u0005"+
		"\n\u0085\b\n\n\n\f\n\u0088\t\n\u0001\n\u0000\u0001\u000e\u000b\u0000\u0002"+
		"\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0000\u0006\u0002\u0000\u0006"+
		"\n$$\u0002\u0000\u000e\u000e\u001d\u001d\u0001\u0000\u000f\u0011\u0001"+
		"\u0000\r\u000e\u0001\u0000\u0017\u001a\u0001\u0000\u0015\u0016\u0099\u0000"+
		"\u0019\u0001\u0000\u0000\u0000\u0002\u001f\u0001\u0000\u0000\u0000\u0004"+
		"!\u0001\u0000\u0000\u0000\u0006&\u0001\u0000\u0000\u0000\b4\u0001\u0000"+
		"\u0000\u0000\n6\u0001\u0000\u0000\u0000\fA\u0001\u0000\u0000\u0000\u000e"+
		"G\u0001\u0000\u0000\u0000\u0010q\u0001\u0000\u0000\u0000\u0012\u007f\u0001"+
		"\u0000\u0000\u0000\u0014\u0081\u0001\u0000\u0000\u0000\u0016\u0018\u0003"+
		"\u0004\u0002\u0000\u0017\u0016\u0001\u0000\u0000\u0000\u0018\u001b\u0001"+
		"\u0000\u0000\u0000\u0019\u0017\u0001\u0000\u0000\u0000\u0019\u001a\u0001"+
		"\u0000\u0000\u0000\u001a\u001d\u0001\u0000\u0000\u0000\u001b\u0019\u0001"+
		"\u0000\u0000\u0000\u001c\u001e\u0005\u0000\u0000\u0001\u001d\u001c\u0001"+
		"\u0000\u0000\u0000\u001d\u001e\u0001\u0000\u0000\u0000\u001e\u0001\u0001"+
		"\u0000\u0000\u0000\u001f \u0005\u0001\u0000\u0000 \u0003\u0001\u0000\u0000"+
		"\u0000!\"\u0003\u0006\u0003\u0000\"#\u0003\u0002\u0001\u0000#\u0005\u0001"+
		"\u0000\u0000\u0000$\'\u0003\b\u0004\u0000%\'\u0003\f\u0006\u0000&$\u0001"+
		"\u0000\u0000\u0000&%\u0001\u0000\u0000\u0000\'\u0007\u0001\u0000\u0000"+
		"\u0000()\u0005\u0002\u0000\u0000)*\u0005$\u0000\u0000*-\u0003\n\u0005"+
		"\u0000+,\u0005\u0012\u0000\u0000,.\u0003\u000e\u0007\u0000-+\u0001\u0000"+
		"\u0000\u0000-.\u0001\u0000\u0000\u0000.5\u0001\u0000\u0000\u0000/0\u0005"+
		"$\u0000\u000001\u0003\n\u0005\u000012\u0005\u0012\u0000\u000023\u0003"+
		"\u000e\u0007\u000035\u0001\u0000\u0000\u00004(\u0001\u0000\u0000\u0000"+
		"4/\u0001\u0000\u0000\u00005\t\u0001\u0000\u0000\u000067\u0007\u0000\u0000"+
		"\u00007\u000b\u0001\u0000\u0000\u000089\u0005$\u0000\u00009:\u0005\u0012"+
		"\u0000\u0000:B\u0003\u000e\u0007\u0000;<\u0005$\u0000\u0000<=\u0005\u0013"+
		"\u0000\u0000=B\u0003\u000e\u0007\u0000>?\u0005$\u0000\u0000?@\u0005\u0014"+
		"\u0000\u0000@B\u0003\u000e\u0007\u0000A8\u0001\u0000\u0000\u0000A;\u0001"+
		"\u0000\u0000\u0000A>\u0001\u0000\u0000\u0000B\r\u0001\u0000\u0000\u0000"+
		"CD\u0006\u0007\uffff\uffff\u0000DH\u0003\u0010\b\u0000EF\u0007\u0001\u0000"+
		"\u0000FH\u0003\u000e\u0007\u0007GC\u0001\u0000\u0000\u0000GE\u0001\u0000"+
		"\u0000\u0000Hb\u0001\u0000\u0000\u0000IJ\n\u0006\u0000\u0000JK\u0007\u0002"+
		"\u0000\u0000Ka\u0003\u000e\u0007\u0007LM\n\u0005\u0000\u0000MN\u0007\u0003"+
		"\u0000\u0000Na\u0003\u000e\u0007\u0006OP\n\u0004\u0000\u0000PQ\u0007\u0004"+
		"\u0000\u0000Qa\u0003\u000e\u0007\u0005RS\n\u0003\u0000\u0000ST\u0007\u0005"+
		"\u0000\u0000Ta\u0003\u000e\u0007\u0004UV\n\u0002\u0000\u0000VW\u0005\u001b"+
		"\u0000\u0000Wa\u0003\u000e\u0007\u0003XY\n\u0001\u0000\u0000YZ\u0005\u001c"+
		"\u0000\u0000Za\u0003\u000e\u0007\u0002[\\\n\b\u0000\u0000\\]\u0005\u001e"+
		"\u0000\u0000]^\u0003\u000e\u0007\u0000^_\u0005\u001f\u0000\u0000_a\u0001"+
		"\u0000\u0000\u0000`I\u0001\u0000\u0000\u0000`L\u0001\u0000\u0000\u0000"+
		"`O\u0001\u0000\u0000\u0000`R\u0001\u0000\u0000\u0000`U\u0001\u0000\u0000"+
		"\u0000`X\u0001\u0000\u0000\u0000`[\u0001\u0000\u0000\u0000ad\u0001\u0000"+
		"\u0000\u0000b`\u0001\u0000\u0000\u0000bc\u0001\u0000\u0000\u0000c\u000f"+
		"\u0001\u0000\u0000\u0000db\u0001\u0000\u0000\u0000er\u0005$\u0000\u0000"+
		"fr\u0005!\u0000\u0000gr\u0005\"\u0000\u0000hr\u0005#\u0000\u0000ir\u0005"+
		"\u000b\u0000\u0000jr\u0005\f\u0000\u0000kr\u0005\u0005\u0000\u0000lr\u0003"+
		"\u0012\t\u0000mn\u0005\u001e\u0000\u0000no\u0003\u000e\u0007\u0000op\u0005"+
		"\u001f\u0000\u0000pr\u0001\u0000\u0000\u0000qe\u0001\u0000\u0000\u0000"+
		"qf\u0001\u0000\u0000\u0000qg\u0001\u0000\u0000\u0000qh\u0001\u0000\u0000"+
		"\u0000qi\u0001\u0000\u0000\u0000qj\u0001\u0000\u0000\u0000qk\u0001\u0000"+
		"\u0000\u0000ql\u0001\u0000\u0000\u0000qm\u0001\u0000\u0000\u0000r\u0011"+
		"\u0001\u0000\u0000\u0000st\u0005\u0003\u0000\u0000tv\u0005\u001e\u0000"+
		"\u0000uw\u0003\u0014\n\u0000vu\u0001\u0000\u0000\u0000vw\u0001\u0000\u0000"+
		"\u0000wx\u0001\u0000\u0000\u0000x\u0080\u0005\u001f\u0000\u0000yz\u0005"+
		"\u0004\u0000\u0000z|\u0005\u001e\u0000\u0000{}\u0003\u0014\n\u0000|{\u0001"+
		"\u0000\u0000\u0000|}\u0001\u0000\u0000\u0000}~\u0001\u0000\u0000\u0000"+
		"~\u0080\u0005\u001f\u0000\u0000\u007fs\u0001\u0000\u0000\u0000\u007fy"+
		"\u0001\u0000\u0000\u0000\u0080\u0013\u0001\u0000\u0000\u0000\u0081\u0086"+
		"\u0003\u000e\u0007\u0000\u0082\u0083\u0005 \u0000\u0000\u0083\u0085\u0003"+
		"\u000e\u0007\u0000\u0084\u0082\u0001\u0000\u0000\u0000\u0085\u0088\u0001"+
		"\u0000\u0000\u0000\u0086\u0084\u0001\u0000\u0000\u0000\u0086\u0087\u0001"+
		"\u0000\u0000\u0000\u0087\u0015\u0001\u0000\u0000\u0000\u0088\u0086\u0001"+
		"\u0000\u0000\u0000\u000e\u0019\u001d&-4AG`bqv|\u007f\u0086";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}