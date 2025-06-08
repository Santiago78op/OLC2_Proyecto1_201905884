// Code generated from grammar/VLangParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // VLangParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type VLangParserParser struct {
	*antlr.BaseParser
}

var VLangParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func vlangparserParserInit() {
	staticData := &VLangParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'\\n'", "'mut'", "'print'", "'println'", "'nil'", "'int'", "'float64'",
		"'string'", "'bool'", "'slice'", "'true'", "'false'", "'+'", "'-'",
		"'*'", "'/'", "'%'", "'='", "'+='", "'-='", "'=='", "'!='", "'<'", "'<='",
		"'>'", "'>='", "'&&'", "'||'", "'!'", "'('", "')'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "", "MUT", "PRINT", "PRINTLN", "NIL", "INT_TYPE", "FLOAT_TYPE",
		"STRING_TYPE", "BOOL_TYPE", "SLICE_TYPE", "TRUE", "FALSE", "PLUS", "MINUS",
		"MULT", "DIV", "MOD", "ASSIGN", "PLUS_ASSIGN", "MINUS_ASSIGN", "EQ",
		"NE", "LT", "LE", "GT", "GE", "AND", "OR", "NOT", "LPAREN", "RPAREN",
		"COMMA", "INT_LITERAL", "FLOAT_LITERAL", "STRING_LITERAL", "ID", "LINE_COMMENT",
		"BLOCK_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"prog", "delim", "stmt", "declaration", "variable_declaration", "type_annotation",
		"assignment_declaration", "expression", "primary_expression", "builtin_function_call",
		"argument_list",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 39, 138, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 5, 0, 24, 8, 0, 10, 0, 12, 0, 27, 9, 0, 1, 0, 3, 0, 30, 8, 0,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 39, 8, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 3, 4, 46, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 53,
		8, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		3, 6, 66, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 72, 8, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 97, 8, 7, 10,
		7, 12, 7, 100, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 3, 8, 114, 8, 8, 1, 9, 1, 9, 1, 9, 3, 9, 119, 8, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 125, 8, 9, 1, 9, 3, 9, 128, 8, 9, 1, 10,
		1, 10, 1, 10, 5, 10, 133, 8, 10, 10, 10, 12, 10, 136, 9, 10, 1, 10, 0,
		1, 14, 11, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 0, 6, 2, 0, 6, 10, 36,
		36, 2, 0, 14, 14, 29, 29, 1, 0, 15, 17, 1, 0, 13, 14, 1, 0, 23, 26, 1,
		0, 21, 22, 153, 0, 25, 1, 0, 0, 0, 2, 31, 1, 0, 0, 0, 4, 33, 1, 0, 0, 0,
		6, 38, 1, 0, 0, 0, 8, 52, 1, 0, 0, 0, 10, 54, 1, 0, 0, 0, 12, 65, 1, 0,
		0, 0, 14, 71, 1, 0, 0, 0, 16, 113, 1, 0, 0, 0, 18, 127, 1, 0, 0, 0, 20,
		129, 1, 0, 0, 0, 22, 24, 3, 4, 2, 0, 23, 22, 1, 0, 0, 0, 24, 27, 1, 0,
		0, 0, 25, 23, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 29, 1, 0, 0, 0, 27, 25,
		1, 0, 0, 0, 28, 30, 5, 0, 0, 1, 29, 28, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0,
		30, 1, 1, 0, 0, 0, 31, 32, 5, 1, 0, 0, 32, 3, 1, 0, 0, 0, 33, 34, 3, 6,
		3, 0, 34, 35, 3, 2, 1, 0, 35, 5, 1, 0, 0, 0, 36, 39, 3, 8, 4, 0, 37, 39,
		3, 12, 6, 0, 38, 36, 1, 0, 0, 0, 38, 37, 1, 0, 0, 0, 39, 7, 1, 0, 0, 0,
		40, 41, 5, 2, 0, 0, 41, 42, 5, 36, 0, 0, 42, 45, 3, 10, 5, 0, 43, 44, 5,
		18, 0, 0, 44, 46, 3, 14, 7, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0,
		46, 53, 1, 0, 0, 0, 47, 48, 5, 36, 0, 0, 48, 49, 3, 10, 5, 0, 49, 50, 5,
		18, 0, 0, 50, 51, 3, 14, 7, 0, 51, 53, 1, 0, 0, 0, 52, 40, 1, 0, 0, 0,
		52, 47, 1, 0, 0, 0, 53, 9, 1, 0, 0, 0, 54, 55, 7, 0, 0, 0, 55, 11, 1, 0,
		0, 0, 56, 57, 5, 36, 0, 0, 57, 58, 5, 18, 0, 0, 58, 66, 3, 14, 7, 0, 59,
		60, 5, 36, 0, 0, 60, 61, 5, 19, 0, 0, 61, 66, 3, 14, 7, 0, 62, 63, 5, 36,
		0, 0, 63, 64, 5, 20, 0, 0, 64, 66, 3, 14, 7, 0, 65, 56, 1, 0, 0, 0, 65,
		59, 1, 0, 0, 0, 65, 62, 1, 0, 0, 0, 66, 13, 1, 0, 0, 0, 67, 68, 6, 7, -1,
		0, 68, 72, 3, 16, 8, 0, 69, 70, 7, 1, 0, 0, 70, 72, 3, 14, 7, 7, 71, 67,
		1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 72, 98, 1, 0, 0, 0, 73, 74, 10, 6, 0, 0,
		74, 75, 7, 2, 0, 0, 75, 97, 3, 14, 7, 7, 76, 77, 10, 5, 0, 0, 77, 78, 7,
		3, 0, 0, 78, 97, 3, 14, 7, 6, 79, 80, 10, 4, 0, 0, 80, 81, 7, 4, 0, 0,
		81, 97, 3, 14, 7, 5, 82, 83, 10, 3, 0, 0, 83, 84, 7, 5, 0, 0, 84, 97, 3,
		14, 7, 4, 85, 86, 10, 2, 0, 0, 86, 87, 5, 27, 0, 0, 87, 97, 3, 14, 7, 3,
		88, 89, 10, 1, 0, 0, 89, 90, 5, 28, 0, 0, 90, 97, 3, 14, 7, 2, 91, 92,
		10, 8, 0, 0, 92, 93, 5, 30, 0, 0, 93, 94, 3, 14, 7, 0, 94, 95, 5, 31, 0,
		0, 95, 97, 1, 0, 0, 0, 96, 73, 1, 0, 0, 0, 96, 76, 1, 0, 0, 0, 96, 79,
		1, 0, 0, 0, 96, 82, 1, 0, 0, 0, 96, 85, 1, 0, 0, 0, 96, 88, 1, 0, 0, 0,
		96, 91, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1,
		0, 0, 0, 99, 15, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101, 114, 5, 36, 0, 0,
		102, 114, 5, 33, 0, 0, 103, 114, 5, 34, 0, 0, 104, 114, 5, 35, 0, 0, 105,
		114, 5, 11, 0, 0, 106, 114, 5, 12, 0, 0, 107, 114, 5, 5, 0, 0, 108, 114,
		3, 18, 9, 0, 109, 110, 5, 30, 0, 0, 110, 111, 3, 14, 7, 0, 111, 112, 5,
		31, 0, 0, 112, 114, 1, 0, 0, 0, 113, 101, 1, 0, 0, 0, 113, 102, 1, 0, 0,
		0, 113, 103, 1, 0, 0, 0, 113, 104, 1, 0, 0, 0, 113, 105, 1, 0, 0, 0, 113,
		106, 1, 0, 0, 0, 113, 107, 1, 0, 0, 0, 113, 108, 1, 0, 0, 0, 113, 109,
		1, 0, 0, 0, 114, 17, 1, 0, 0, 0, 115, 116, 5, 3, 0, 0, 116, 118, 5, 30,
		0, 0, 117, 119, 3, 20, 10, 0, 118, 117, 1, 0, 0, 0, 118, 119, 1, 0, 0,
		0, 119, 120, 1, 0, 0, 0, 120, 128, 5, 31, 0, 0, 121, 122, 5, 4, 0, 0, 122,
		124, 5, 30, 0, 0, 123, 125, 3, 20, 10, 0, 124, 123, 1, 0, 0, 0, 124, 125,
		1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 128, 5, 31, 0, 0, 127, 115, 1, 0,
		0, 0, 127, 121, 1, 0, 0, 0, 128, 19, 1, 0, 0, 0, 129, 134, 3, 14, 7, 0,
		130, 131, 5, 32, 0, 0, 131, 133, 3, 14, 7, 0, 132, 130, 1, 0, 0, 0, 133,
		136, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 21, 1,
		0, 0, 0, 136, 134, 1, 0, 0, 0, 14, 25, 29, 38, 45, 52, 65, 71, 96, 98,
		113, 118, 124, 127, 134,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// VLangParserParserInit initializes any static state used to implement VLangParserParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVLangParserParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func VLangParserParserInit() {
	staticData := &VLangParserParserStaticData
	staticData.once.Do(vlangparserParserInit)
}

// NewVLangParserParser produces a new parser instance for the optional input antlr.TokenStream.
func NewVLangParserParser(input antlr.TokenStream) *VLangParserParser {
	VLangParserParserInit()
	this := new(VLangParserParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &VLangParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "VLangParser.g4"

	return this
}

// VLangParserParser tokens.
const (
	VLangParserParserEOF            = antlr.TokenEOF
	VLangParserParserT__0           = 1
	VLangParserParserMUT            = 2
	VLangParserParserPRINT          = 3
	VLangParserParserPRINTLN        = 4
	VLangParserParserNIL            = 5
	VLangParserParserINT_TYPE       = 6
	VLangParserParserFLOAT_TYPE     = 7
	VLangParserParserSTRING_TYPE    = 8
	VLangParserParserBOOL_TYPE      = 9
	VLangParserParserSLICE_TYPE     = 10
	VLangParserParserTRUE           = 11
	VLangParserParserFALSE          = 12
	VLangParserParserPLUS           = 13
	VLangParserParserMINUS          = 14
	VLangParserParserMULT           = 15
	VLangParserParserDIV            = 16
	VLangParserParserMOD            = 17
	VLangParserParserASSIGN         = 18
	VLangParserParserPLUS_ASSIGN    = 19
	VLangParserParserMINUS_ASSIGN   = 20
	VLangParserParserEQ             = 21
	VLangParserParserNE             = 22
	VLangParserParserLT             = 23
	VLangParserParserLE             = 24
	VLangParserParserGT             = 25
	VLangParserParserGE             = 26
	VLangParserParserAND            = 27
	VLangParserParserOR             = 28
	VLangParserParserNOT            = 29
	VLangParserParserLPAREN         = 30
	VLangParserParserRPAREN         = 31
	VLangParserParserCOMMA          = 32
	VLangParserParserINT_LITERAL    = 33
	VLangParserParserFLOAT_LITERAL  = 34
	VLangParserParserSTRING_LITERAL = 35
	VLangParserParserID             = 36
	VLangParserParserLINE_COMMENT   = 37
	VLangParserParserBLOCK_COMMENT  = 38
	VLangParserParserWS             = 39
)

// VLangParserParser rules.
const (
	VLangParserParserRULE_prog                   = 0
	VLangParserParserRULE_delim                  = 1
	VLangParserParserRULE_stmt                   = 2
	VLangParserParserRULE_declaration            = 3
	VLangParserParserRULE_variable_declaration   = 4
	VLangParserParserRULE_type_annotation        = 5
	VLangParserParserRULE_assignment_declaration = 6
	VLangParserParserRULE_expression             = 7
	VLangParserParserRULE_primary_expression     = 8
	VLangParserParserRULE_builtin_function_call  = 9
	VLangParserParserRULE_argument_list          = 10
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext
	EOF() antlr.TerminalNode

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangParserParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangParserParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangParserParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(VLangParserParserEOF, 0)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangParserParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VLangParserParserRULE_prog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VLangParserParserMUT || _la == VLangParserParserID {
		{
			p.SetState(22)
			p.Stmt()
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(28)
			p.Match(VLangParserParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDelimContext is an interface to support dynamic dispatch.
type IDelimContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDelimContext differentiates from other interfaces.
	IsDelimContext()
}

type DelimContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelimContext() *DelimContext {
	var p = new(DelimContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangParserParserRULE_delim
	return p
}

func InitEmptyDelimContext(p *DelimContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangParserParserRULE_delim
}

func (*DelimContext) IsDelimContext() {}

func NewDelimContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DelimContext {
	var p = new(DelimContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangParserParserRULE_delim

	return p
}

func (s *DelimContext) GetParser() antlr.Parser { return s.parser }
func (s *DelimContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DelimContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DelimContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterDelim(s)
	}
}

func (s *DelimContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitDelim(s)
	}
}

func (s *DelimContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitDelim(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangParserParser) Delim() (localctx IDelimContext) {
	localctx = NewDelimContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VLangParserParserRULE_delim)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Match(VLangParserParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declaration() IDeclarationContext
	Delim() IDelimContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangParserParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangParserParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangParserParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Declaration() IDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StmtContext) Delim() IDelimContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDelimContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDelimContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangParserParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, VLangParserParserRULE_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)
		p.Declaration()
	}
	{
		p.SetState(34)
		p.Delim()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Variable_declaration() IVariable_declarationContext
	Assignment_declaration() IAssignment_declarationContext

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangParserParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangParserParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangParserParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) Variable_declaration() IVariable_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariable_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariable_declarationContext)
}

func (s *DeclarationContext) Assignment_declaration() IAssignment_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignment_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignment_declarationContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangParserParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, VLangParserParserRULE_declaration)
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(36)
			p.Variable_declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.Assignment_declaration()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariable_declarationContext is an interface to support dynamic dispatch.
type IVariable_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVariable_declarationContext differentiates from other interfaces.
	IsVariable_declarationContext()
}

type Variable_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariable_declarationContext() *Variable_declarationContext {
	var p = new(Variable_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangParserParserRULE_variable_declaration
	return p
}

func InitEmptyVariable_declarationContext(p *Variable_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangParserParserRULE_variable_declaration
}

func (*Variable_declarationContext) IsVariable_declarationContext() {}

func NewVariable_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Variable_declarationContext {
	var p = new(Variable_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangParserParserRULE_variable_declaration

	return p
}

func (s *Variable_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Variable_declarationContext) CopyAll(ctx *Variable_declarationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Variable_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Variable_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypedVarDeclContext struct {
	Variable_declarationContext
}

func NewTypedVarDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypedVarDeclContext {
	var p = new(TypedVarDeclContext)

	InitEmptyVariable_declarationContext(&p.Variable_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Variable_declarationContext))

	return p
}

func (s *TypedVarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedVarDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(VLangParserParserID, 0)
}

func (s *TypedVarDeclContext) Type_annotation() IType_annotationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_annotationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_annotationContext)
}

func (s *TypedVarDeclContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VLangParserParserASSIGN, 0)
}

func (s *TypedVarDeclContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TypedVarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterTypedVarDecl(s)
	}
}

func (s *TypedVarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitTypedVarDecl(s)
	}
}

func (s *TypedVarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitTypedVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type MutableVarDeclContext struct {
	Variable_declarationContext
}

func NewMutableVarDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MutableVarDeclContext {
	var p = new(MutableVarDeclContext)

	InitEmptyVariable_declarationContext(&p.Variable_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Variable_declarationContext))

	return p
}

func (s *MutableVarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutableVarDeclContext) MUT() antlr.TerminalNode {
	return s.GetToken(VLangParserParserMUT, 0)
}

func (s *MutableVarDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(VLangParserParserID, 0)
}

func (s *MutableVarDeclContext) Type_annotation() IType_annotationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_annotationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_annotationContext)
}

func (s *MutableVarDeclContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VLangParserParserASSIGN, 0)
}

func (s *MutableVarDeclContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MutableVarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterMutableVarDecl(s)
	}
}

func (s *MutableVarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitMutableVarDecl(s)
	}
}

func (s *MutableVarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitMutableVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangParserParser) Variable_declaration() (localctx IVariable_declarationContext) {
	localctx = NewVariable_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, VLangParserParserRULE_variable_declaration)
	var _la int

	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VLangParserParserMUT:
		localctx = NewMutableVarDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(40)
			p.Match(VLangParserParserMUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(41)
			p.Match(VLangParserParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(42)
			p.Type_annotation()
		}
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == VLangParserParserASSIGN {
			{
				p.SetState(43)
				p.Match(VLangParserParserASSIGN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(44)
				p.expression(0)
			}

		}

	case VLangParserParserID:
		localctx = NewTypedVarDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)
			p.Match(VLangParserParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.Type_annotation()
		}
		{
			p.SetState(49)
			p.Match(VLangParserParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.expression(0)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IType_annotationContext is an interface to support dynamic dispatch.
type IType_annotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_TYPE() antlr.TerminalNode
	FLOAT_TYPE() antlr.TerminalNode
	STRING_TYPE() antlr.TerminalNode
	BOOL_TYPE() antlr.TerminalNode
	SLICE_TYPE() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsType_annotationContext differentiates from other interfaces.
	IsType_annotationContext()
}

type Type_annotationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_annotationContext() *Type_annotationContext {
	var p = new(Type_annotationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangParserParserRULE_type_annotation
	return p
}

func InitEmptyType_annotationContext(p *Type_annotationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangParserParserRULE_type_annotation
}

func (*Type_annotationContext) IsType_annotationContext() {}

func NewType_annotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_annotationContext {
	var p = new(Type_annotationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangParserParserRULE_type_annotation

	return p
}

func (s *Type_annotationContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_annotationContext) INT_TYPE() antlr.TerminalNode {
	return s.GetToken(VLangParserParserINT_TYPE, 0)
}

func (s *Type_annotationContext) FLOAT_TYPE() antlr.TerminalNode {
	return s.GetToken(VLangParserParserFLOAT_TYPE, 0)
}

func (s *Type_annotationContext) STRING_TYPE() antlr.TerminalNode {
	return s.GetToken(VLangParserParserSTRING_TYPE, 0)
}

func (s *Type_annotationContext) BOOL_TYPE() antlr.TerminalNode {
	return s.GetToken(VLangParserParserBOOL_TYPE, 0)
}

func (s *Type_annotationContext) SLICE_TYPE() antlr.TerminalNode {
	return s.GetToken(VLangParserParserSLICE_TYPE, 0)
}

func (s *Type_annotationContext) ID() antlr.TerminalNode {
	return s.GetToken(VLangParserParserID, 0)
}

func (s *Type_annotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_annotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_annotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterType_annotation(s)
	}
}

func (s *Type_annotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitType_annotation(s)
	}
}

func (s *Type_annotationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitType_annotation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangParserParser) Type_annotation() (localctx IType_annotationContext) {
	localctx = NewType_annotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, VLangParserParserRULE_type_annotation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68719478720) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignment_declarationContext is an interface to support dynamic dispatch.
type IAssignment_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssignment_declarationContext differentiates from other interfaces.
	IsAssignment_declarationContext()
}

type Assignment_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignment_declarationContext() *Assignment_declarationContext {
	var p = new(Assignment_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangParserParserRULE_assignment_declaration
	return p
}

func InitEmptyAssignment_declarationContext(p *Assignment_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangParserParserRULE_assignment_declaration
}

func (*Assignment_declarationContext) IsAssignment_declarationContext() {}

func NewAssignment_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assignment_declarationContext {
	var p = new(Assignment_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangParserParserRULE_assignment_declaration

	return p
}

func (s *Assignment_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Assignment_declarationContext) CopyAll(ctx *Assignment_declarationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Assignment_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assignment_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MinusAssignmentDeclContext struct {
	Assignment_declarationContext
}

func NewMinusAssignmentDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MinusAssignmentDeclContext {
	var p = new(MinusAssignmentDeclContext)

	InitEmptyAssignment_declarationContext(&p.Assignment_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assignment_declarationContext))

	return p
}

func (s *MinusAssignmentDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinusAssignmentDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(VLangParserParserID, 0)
}

func (s *MinusAssignmentDeclContext) MINUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VLangParserParserMINUS_ASSIGN, 0)
}

func (s *MinusAssignmentDeclContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MinusAssignmentDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterMinusAssignmentDecl(s)
	}
}

func (s *MinusAssignmentDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitMinusAssignmentDecl(s)
	}
}

func (s *MinusAssignmentDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitMinusAssignmentDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type PlusAssignmentDeclContext struct {
	Assignment_declarationContext
}

func NewPlusAssignmentDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlusAssignmentDeclContext {
	var p = new(PlusAssignmentDeclContext)

	InitEmptyAssignment_declarationContext(&p.Assignment_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assignment_declarationContext))

	return p
}

func (s *PlusAssignmentDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlusAssignmentDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(VLangParserParserID, 0)
}

func (s *PlusAssignmentDeclContext) PLUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VLangParserParserPLUS_ASSIGN, 0)
}

func (s *PlusAssignmentDeclContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PlusAssignmentDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterPlusAssignmentDecl(s)
	}
}

func (s *PlusAssignmentDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitPlusAssignmentDecl(s)
	}
}

func (s *PlusAssignmentDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitPlusAssignmentDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentDeclContext struct {
	Assignment_declarationContext
}

func NewAssignmentDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentDeclContext {
	var p = new(AssignmentDeclContext)

	InitEmptyAssignment_declarationContext(&p.Assignment_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assignment_declarationContext))

	return p
}

func (s *AssignmentDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(VLangParserParserID, 0)
}

func (s *AssignmentDeclContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VLangParserParserASSIGN, 0)
}

func (s *AssignmentDeclContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterAssignmentDecl(s)
	}
}

func (s *AssignmentDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitAssignmentDecl(s)
	}
}

func (s *AssignmentDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitAssignmentDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangParserParser) Assignment_declaration() (localctx IAssignment_declarationContext) {
	localctx = NewAssignment_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, VLangParserParserRULE_assignment_declaration)
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignmentDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.Match(VLangParserParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.Match(VLangParserParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)
			p.expression(0)
		}

	case 2:
		localctx = NewPlusAssignmentDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(59)
			p.Match(VLangParserParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(60)
			p.Match(VLangParserParserPLUS_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)
			p.expression(0)
		}

	case 3:
		localctx = NewMinusAssignmentDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(62)
			p.Match(VLangParserParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.Match(VLangParserParserMINUS_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangParserParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangParserParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangParserParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BinaryExprContext struct {
	ExpressionContext
	left  IExpressionContext
	op    antlr.Token
	right IExpressionContext
}

func NewBinaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExprContext {
	var p = new(BinaryExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryExprContext) GetOp() antlr.Token { return s.op }

func (s *BinaryExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinaryExprContext) GetLeft() IExpressionContext { return s.left }

func (s *BinaryExprContext) GetRight() IExpressionContext { return s.right }

func (s *BinaryExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *BinaryExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *BinaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinaryExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinaryExprContext) MULT() antlr.TerminalNode {
	return s.GetToken(VLangParserParserMULT, 0)
}

func (s *BinaryExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(VLangParserParserDIV, 0)
}

func (s *BinaryExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(VLangParserParserMOD, 0)
}

func (s *BinaryExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(VLangParserParserPLUS, 0)
}

func (s *BinaryExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(VLangParserParserMINUS, 0)
}

func (s *BinaryExprContext) LE() antlr.TerminalNode {
	return s.GetToken(VLangParserParserLE, 0)
}

func (s *BinaryExprContext) LT() antlr.TerminalNode {
	return s.GetToken(VLangParserParserLT, 0)
}

func (s *BinaryExprContext) GE() antlr.TerminalNode {
	return s.GetToken(VLangParserParserGE, 0)
}

func (s *BinaryExprContext) GT() antlr.TerminalNode {
	return s.GetToken(VLangParserParserGT, 0)
}

func (s *BinaryExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(VLangParserParserEQ, 0)
}

func (s *BinaryExprContext) NE() antlr.TerminalNode {
	return s.GetToken(VLangParserParserNE, 0)
}

func (s *BinaryExprContext) AND() antlr.TerminalNode {
	return s.GetToken(VLangParserParserAND, 0)
}

func (s *BinaryExprContext) OR() antlr.TerminalNode {
	return s.GetToken(VLangParserParserOR, 0)
}

func (s *BinaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterBinaryExpr(s)
	}
}

func (s *BinaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitBinaryExpr(s)
	}
}

func (s *BinaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitBinaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryExprContext struct {
	ExpressionContext
}

func NewPrimaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) Primary_expression() IPrimary_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimary_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimary_expressionContext)
}

func (s *PrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitPrimaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenthesizedExprContext struct {
	ExpressionContext
}

func NewParenthesizedExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesizedExprContext {
	var p = new(ParenthesizedExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParenthesizedExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesizedExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ParenthesizedExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenthesizedExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VLangParserParserLPAREN, 0)
}

func (s *ParenthesizedExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VLangParserParserRPAREN, 0)
}

func (s *ParenthesizedExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterParenthesizedExpr(s)
	}
}

func (s *ParenthesizedExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitParenthesizedExpr(s)
	}
}

func (s *ParenthesizedExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitParenthesizedExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryExprContext struct {
	ExpressionContext
	op antlr.Token
}

func NewUnaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExprContext {
	var p = new(UnaryExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryExprContext) GetOp() antlr.Token { return s.op }

func (s *UnaryExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(VLangParserParserNOT, 0)
}

func (s *UnaryExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(VLangParserParserMINUS, 0)
}

func (s *UnaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterUnaryExpr(s)
	}
}

func (s *UnaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitUnaryExpr(s)
	}
}

func (s *UnaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitUnaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangParserParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *VLangParserParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, VLangParserParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VLangParserParserPRINT, VLangParserParserPRINTLN, VLangParserParserNIL, VLangParserParserTRUE, VLangParserParserFALSE, VLangParserParserLPAREN, VLangParserParserINT_LITERAL, VLangParserParserFLOAT_LITERAL, VLangParserParserSTRING_LITERAL, VLangParserParserID:
		localctx = NewPrimaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(68)
			p.Primary_expression()
		}

	case VLangParserParserMINUS, VLangParserParserNOT:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(69)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == VLangParserParserMINUS || _la == VLangParserParserNOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(70)
			p.expression(7)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VLangParserParserRULE_expression)
				p.SetState(73)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(74)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&229376) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(75)

					var _x = p.expression(7)

					localctx.(*BinaryExprContext).right = _x
				}

			case 2:
				localctx = NewBinaryExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VLangParserParserRULE_expression)
				p.SetState(76)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(77)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VLangParserParserPLUS || _la == VLangParserParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(78)

					var _x = p.expression(6)

					localctx.(*BinaryExprContext).right = _x
				}

			case 3:
				localctx = NewBinaryExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VLangParserParserRULE_expression)
				p.SetState(79)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(80)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&125829120) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(81)

					var _x = p.expression(5)

					localctx.(*BinaryExprContext).right = _x
				}

			case 4:
				localctx = NewBinaryExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VLangParserParserRULE_expression)
				p.SetState(82)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(83)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VLangParserParserEQ || _la == VLangParserParserNE) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(84)

					var _x = p.expression(4)

					localctx.(*BinaryExprContext).right = _x
				}

			case 5:
				localctx = NewBinaryExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VLangParserParserRULE_expression)
				p.SetState(85)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(86)

					var _m = p.Match(VLangParserParserAND)

					localctx.(*BinaryExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(87)

					var _x = p.expression(3)

					localctx.(*BinaryExprContext).right = _x
				}

			case 6:
				localctx = NewBinaryExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VLangParserParserRULE_expression)
				p.SetState(88)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(89)

					var _m = p.Match(VLangParserParserOR)

					localctx.(*BinaryExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(90)

					var _x = p.expression(2)

					localctx.(*BinaryExprContext).right = _x
				}

			case 7:
				localctx = NewParenthesizedExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VLangParserParserRULE_expression)
				p.SetState(91)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(92)
					p.Match(VLangParserParserLPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(93)
					p.expression(0)
				}
				{
					p.SetState(94)
					p.Match(VLangParserParserRPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimary_expressionContext is an interface to support dynamic dispatch.
type IPrimary_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrimary_expressionContext differentiates from other interfaces.
	IsPrimary_expressionContext()
}

type Primary_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimary_expressionContext() *Primary_expressionContext {
	var p = new(Primary_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangParserParserRULE_primary_expression
	return p
}

func InitEmptyPrimary_expressionContext(p *Primary_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangParserParserRULE_primary_expression
}

func (*Primary_expressionContext) IsPrimary_expressionContext() {}

func NewPrimary_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Primary_expressionContext {
	var p = new(Primary_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangParserParserRULE_primary_expression

	return p
}

func (s *Primary_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Primary_expressionContext) CopyAll(ctx *Primary_expressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Primary_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primary_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringLiteralContext struct {
	Primary_expressionContext
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext {
	var p = new(StringLiteralContext)

	InitEmptyPrimary_expressionContext(&p.Primary_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Primary_expressionContext))

	return p
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(VLangParserParserSTRING_LITERAL, 0)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type TrueLiteralContext struct {
	Primary_expressionContext
}

func NewTrueLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueLiteralContext {
	var p = new(TrueLiteralContext)

	InitEmptyPrimary_expressionContext(&p.Primary_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Primary_expressionContext))

	return p
}

func (s *TrueLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(VLangParserParserTRUE, 0)
}

func (s *TrueLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterTrueLiteral(s)
	}
}

func (s *TrueLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitTrueLiteral(s)
	}
}

func (s *TrueLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitTrueLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type BuiltinCallContext struct {
	Primary_expressionContext
}

func NewBuiltinCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BuiltinCallContext {
	var p = new(BuiltinCallContext)

	InitEmptyPrimary_expressionContext(&p.Primary_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Primary_expressionContext))

	return p
}

func (s *BuiltinCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuiltinCallContext) Builtin_function_call() IBuiltin_function_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBuiltin_function_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBuiltin_function_callContext)
}

func (s *BuiltinCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterBuiltinCall(s)
	}
}

func (s *BuiltinCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitBuiltinCall(s)
	}
}

func (s *BuiltinCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitBuiltinCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierExprContext struct {
	Primary_expressionContext
}

func NewIdentifierExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExprContext {
	var p = new(IdentifierExprContext)

	InitEmptyPrimary_expressionContext(&p.Primary_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Primary_expressionContext))

	return p
}

func (s *IdentifierExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExprContext) ID() antlr.TerminalNode {
	return s.GetToken(VLangParserParserID, 0)
}

func (s *IdentifierExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterIdentifierExpr(s)
	}
}

func (s *IdentifierExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitIdentifierExpr(s)
	}
}

func (s *IdentifierExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitIdentifierExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatLiteralContext struct {
	Primary_expressionContext
}

func NewFloatLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	InitEmptyPrimary_expressionContext(&p.Primary_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Primary_expressionContext))

	return p
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(VLangParserParserFLOAT_LITERAL, 0)
}

func (s *FloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitFloatLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type NilLiteralContext struct {
	Primary_expressionContext
}

func NewNilLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilLiteralContext {
	var p = new(NilLiteralContext)

	InitEmptyPrimary_expressionContext(&p.Primary_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Primary_expressionContext))

	return p
}

func (s *NilLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilLiteralContext) NIL() antlr.TerminalNode {
	return s.GetToken(VLangParserParserNIL, 0)
}

func (s *NilLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterNilLiteral(s)
	}
}

func (s *NilLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitNilLiteral(s)
	}
}

func (s *NilLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitNilLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntLiteralContext struct {
	Primary_expressionContext
}

func NewIntLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntLiteralContext {
	var p = new(IntLiteralContext)

	InitEmptyPrimary_expressionContext(&p.Primary_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Primary_expressionContext))

	return p
}

func (s *IntLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLiteralContext) INT_LITERAL() antlr.TerminalNode {
	return s.GetToken(VLangParserParserINT_LITERAL, 0)
}

func (s *IntLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterIntLiteral(s)
	}
}

func (s *IntLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitIntLiteral(s)
	}
}

func (s *IntLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitIntLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExprContext struct {
	Primary_expressionContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	InitEmptyPrimary_expressionContext(&p.Primary_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Primary_expressionContext))

	return p
}

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VLangParserParserLPAREN, 0)
}

func (s *ParenExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VLangParserParserRPAREN, 0)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitParenExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type FalseLiteralContext struct {
	Primary_expressionContext
}

func NewFalseLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalseLiteralContext {
	var p = new(FalseLiteralContext)

	InitEmptyPrimary_expressionContext(&p.Primary_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Primary_expressionContext))

	return p
}

func (s *FalseLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(VLangParserParserFALSE, 0)
}

func (s *FalseLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterFalseLiteral(s)
	}
}

func (s *FalseLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitFalseLiteral(s)
	}
}

func (s *FalseLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitFalseLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangParserParser) Primary_expression() (localctx IPrimary_expressionContext) {
	localctx = NewPrimary_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, VLangParserParserRULE_primary_expression)
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VLangParserParserID:
		localctx = NewIdentifierExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)
			p.Match(VLangParserParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangParserParserINT_LITERAL:
		localctx = NewIntLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)
			p.Match(VLangParserParserINT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangParserParserFLOAT_LITERAL:
		localctx = NewFloatLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(103)
			p.Match(VLangParserParserFLOAT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangParserParserSTRING_LITERAL:
		localctx = NewStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(104)
			p.Match(VLangParserParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangParserParserTRUE:
		localctx = NewTrueLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(105)
			p.Match(VLangParserParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangParserParserFALSE:
		localctx = NewFalseLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(106)
			p.Match(VLangParserParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangParserParserNIL:
		localctx = NewNilLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(107)
			p.Match(VLangParserParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangParserParserPRINT, VLangParserParserPRINTLN:
		localctx = NewBuiltinCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(108)
			p.Builtin_function_call()
		}

	case VLangParserParserLPAREN:
		localctx = NewParenExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(109)
			p.Match(VLangParserParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)
			p.expression(0)
		}
		{
			p.SetState(111)
			p.Match(VLangParserParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBuiltin_function_callContext is an interface to support dynamic dispatch.
type IBuiltin_function_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBuiltin_function_callContext differentiates from other interfaces.
	IsBuiltin_function_callContext()
}

type Builtin_function_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuiltin_function_callContext() *Builtin_function_callContext {
	var p = new(Builtin_function_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangParserParserRULE_builtin_function_call
	return p
}

func InitEmptyBuiltin_function_callContext(p *Builtin_function_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangParserParserRULE_builtin_function_call
}

func (*Builtin_function_callContext) IsBuiltin_function_callContext() {}

func NewBuiltin_function_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Builtin_function_callContext {
	var p = new(Builtin_function_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangParserParserRULE_builtin_function_call

	return p
}

func (s *Builtin_function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Builtin_function_callContext) CopyAll(ctx *Builtin_function_callContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Builtin_function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Builtin_function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrintCallContext struct {
	Builtin_function_callContext
}

func NewPrintCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintCallContext {
	var p = new(PrintCallContext)

	InitEmptyBuiltin_function_callContext(&p.Builtin_function_callContext)
	p.parser = parser
	p.CopyAll(ctx.(*Builtin_function_callContext))

	return p
}

func (s *PrintCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintCallContext) PRINT() antlr.TerminalNode {
	return s.GetToken(VLangParserParserPRINT, 0)
}

func (s *PrintCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VLangParserParserLPAREN, 0)
}

func (s *PrintCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VLangParserParserRPAREN, 0)
}

func (s *PrintCallContext) Argument_list() IArgument_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgument_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgument_listContext)
}

func (s *PrintCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterPrintCall(s)
	}
}

func (s *PrintCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitPrintCall(s)
	}
}

func (s *PrintCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitPrintCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintlnCallContext struct {
	Builtin_function_callContext
}

func NewPrintlnCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintlnCallContext {
	var p = new(PrintlnCallContext)

	InitEmptyBuiltin_function_callContext(&p.Builtin_function_callContext)
	p.parser = parser
	p.CopyAll(ctx.(*Builtin_function_callContext))

	return p
}

func (s *PrintlnCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintlnCallContext) PRINTLN() antlr.TerminalNode {
	return s.GetToken(VLangParserParserPRINTLN, 0)
}

func (s *PrintlnCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VLangParserParserLPAREN, 0)
}

func (s *PrintlnCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VLangParserParserRPAREN, 0)
}

func (s *PrintlnCallContext) Argument_list() IArgument_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgument_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgument_listContext)
}

func (s *PrintlnCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterPrintlnCall(s)
	}
}

func (s *PrintlnCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitPrintlnCall(s)
	}
}

func (s *PrintlnCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitPrintlnCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangParserParser) Builtin_function_call() (localctx IBuiltin_function_callContext) {
	localctx = NewBuiltin_function_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, VLangParserParserRULE_builtin_function_call)
	var _la int

	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VLangParserParserPRINT:
		localctx = NewPrintCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(115)
			p.Match(VLangParserParserPRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.Match(VLangParserParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&130459654200) != 0 {
			{
				p.SetState(117)
				p.Argument_list()
			}

		}
		{
			p.SetState(120)
			p.Match(VLangParserParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangParserParserPRINTLN:
		localctx = NewPrintlnCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(121)
			p.Match(VLangParserParserPRINTLN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.Match(VLangParserParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&130459654200) != 0 {
			{
				p.SetState(123)
				p.Argument_list()
			}

		}
		{
			p.SetState(126)
			p.Match(VLangParserParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgument_listContext is an interface to support dynamic dispatch.
type IArgument_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArgument_listContext differentiates from other interfaces.
	IsArgument_listContext()
}

type Argument_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgument_listContext() *Argument_listContext {
	var p = new(Argument_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangParserParserRULE_argument_list
	return p
}

func InitEmptyArgument_listContext(p *Argument_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangParserParserRULE_argument_list
}

func (*Argument_listContext) IsArgument_listContext() {}

func NewArgument_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Argument_listContext {
	var p = new(Argument_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangParserParserRULE_argument_list

	return p
}

func (s *Argument_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Argument_listContext) CopyAll(ctx *Argument_listContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Argument_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Argument_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArgListContext struct {
	Argument_listContext
}

func NewArgListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArgListContext {
	var p = new(ArgListContext)

	InitEmptyArgument_listContext(&p.Argument_listContext)
	p.parser = parser
	p.CopyAll(ctx.(*Argument_listContext))

	return p
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VLangParserParserCOMMA)
}

func (s *ArgListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VLangParserParserCOMMA, i)
}

func (s *ArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.EnterArgList(s)
	}
}

func (s *ArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangParserListener); ok {
		listenerT.ExitArgList(s)
	}
}

func (s *ArgListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangParserVisitor:
		return t.VisitArgList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangParserParser) Argument_list() (localctx IArgument_listContext) {
	localctx = NewArgument_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, VLangParserParserRULE_argument_list)
	var _la int

	localctx = NewArgListContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.expression(0)
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VLangParserParserCOMMA {
		{
			p.SetState(130)
			p.Match(VLangParserParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.expression(0)
		}

		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *VLangParserParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 7:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *VLangParserParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
