// Code generated from grammar/Calculator.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CalculatorLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var CalculatorLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func calculatorlexerLexerInit() {
	staticData := &CalculatorLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'='", "'*'", "'/'", "'+'", "'-'", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "ID", "INT", "NEWLINE", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "ID", "INT",
		"NEWLINE", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 11, 59, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 4, 7, 39, 8, 7, 11, 7, 12, 7, 40, 1, 8, 4,
		8, 44, 8, 8, 11, 8, 12, 8, 45, 1, 9, 4, 9, 49, 8, 9, 11, 9, 12, 9, 50,
		1, 10, 4, 10, 54, 8, 10, 11, 10, 12, 10, 55, 1, 10, 1, 10, 0, 0, 11, 1,
		1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		1, 0, 4, 2, 0, 65, 90, 97, 122, 1, 0, 48, 57, 2, 0, 10, 10, 13, 13, 2,
		0, 9, 9, 32, 32, 62, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 1, 23, 1, 0, 0, 0, 3, 25, 1, 0, 0, 0, 5, 27, 1, 0, 0, 0, 7, 29,
		1, 0, 0, 0, 9, 31, 1, 0, 0, 0, 11, 33, 1, 0, 0, 0, 13, 35, 1, 0, 0, 0,
		15, 38, 1, 0, 0, 0, 17, 43, 1, 0, 0, 0, 19, 48, 1, 0, 0, 0, 21, 53, 1,
		0, 0, 0, 23, 24, 5, 61, 0, 0, 24, 2, 1, 0, 0, 0, 25, 26, 5, 42, 0, 0, 26,
		4, 1, 0, 0, 0, 27, 28, 5, 47, 0, 0, 28, 6, 1, 0, 0, 0, 29, 30, 5, 43, 0,
		0, 30, 8, 1, 0, 0, 0, 31, 32, 5, 45, 0, 0, 32, 10, 1, 0, 0, 0, 33, 34,
		5, 40, 0, 0, 34, 12, 1, 0, 0, 0, 35, 36, 5, 41, 0, 0, 36, 14, 1, 0, 0,
		0, 37, 39, 7, 0, 0, 0, 38, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 38,
		1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 16, 1, 0, 0, 0, 42, 44, 7, 1, 0, 0,
		43, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1,
		0, 0, 0, 46, 18, 1, 0, 0, 0, 47, 49, 7, 2, 0, 0, 48, 47, 1, 0, 0, 0, 49,
		50, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 20, 1, 0, 0,
		0, 52, 54, 7, 3, 0, 0, 53, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 53,
		1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 58, 6, 10, 0, 0,
		58, 22, 1, 0, 0, 0, 5, 0, 40, 45, 50, 55, 1, 6, 0, 0,
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

// CalculatorLexerInit initializes any static state used to implement CalculatorLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCalculatorLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CalculatorLexerInit() {
	staticData := &CalculatorLexerLexerStaticData
	staticData.once.Do(calculatorlexerLexerInit)
}

// NewCalculatorLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCalculatorLexer(input antlr.CharStream) *CalculatorLexer {
	CalculatorLexerInit()
	l := new(CalculatorLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &CalculatorLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Calculator.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalculatorLexer tokens.
const (
	CalculatorLexerT__0    = 1
	CalculatorLexerT__1    = 2
	CalculatorLexerT__2    = 3
	CalculatorLexerT__3    = 4
	CalculatorLexerT__4    = 5
	CalculatorLexerT__5    = 6
	CalculatorLexerT__6    = 7
	CalculatorLexerID      = 8
	CalculatorLexerINT     = 9
	CalculatorLexerNEWLINE = 10
	CalculatorLexerWS      = 11
)
