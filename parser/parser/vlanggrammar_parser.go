// Code generated from C:/Users/72358/Documents/Compi_2/Semana_01/awesomeProject/grammar/VLangGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // VLangGrammar
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

type VLangGrammarParser struct {
	*antlr.BaseParser
}

var VLangGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func vlanggrammarParserInit() {
	staticData := &VLangGrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "'mut'", "'func'", "'struct'", "'if'", "'else'", "'for'", "'switch'",
		"'case'", "'default'", "'break'", "'continue'", "'return'", "'in'",
		"'nil'", "'int'", "'float64'", "'string'", "'bool'", "'true'", "'false'",
		"'+'", "'-'", "'*'", "'/'", "'%'", "'='", "'+='", "'-='", "'=='", "'!='",
		"'<'", "'<='", "'>'", "'>='", "'&&'", "'||'", "'!'", "'('", "')'", "'{'",
		"'}'", "'['", "']'", "';'", "','", "'.'", "':'", "'&'",
	}
	staticData.SymbolicNames = []string{
		"", "MUT", "FUNC", "STRUCT", "IF", "ELSE", "FOR", "SWITCH", "CASE",
		"DEFAULT", "BREAK", "CONTINUE", "RETURN", "IN", "NIL", "INT_TYPE", "FLOAT_TYPE",
		"STRING_TYPE", "BOOL_TYPE", "TRUE", "FALSE", "PLUS", "MINUS", "MULT",
		"DIV", "MOD", "ASSIGN", "PLUS_ASSIGN", "MINUS_ASSIGN", "EQ", "NE", "LT",
		"LE", "GT", "GE", "AND", "OR", "NOT", "LPAREN", "RPAREN", "LBRACE",
		"RBRACE", "LBRACKET", "RBRACKET", "SEMICOLON", "COMMA", "DOT", "COLON",
		"AMPERSAND", "INT", "FLOAT", "STRING", "RUNE", "ID", "LINE_COMMENT",
		"BLOCK_COMMENT", "WS", "IDENTIFIER", "DECLARE_ASSIGN", "RUNE_TYPE",
		"MULTIPLY", "DIVIDE", "MODULO", "LESS_THAN", "LESS_EQUAL", "GREATER_THAN",
		"GREATER_EQUAL", "EQUAL", "NOT_EQUAL", "INT_LITERAL", "FLOAT_LITERAL",
		"STRING_LITERAL", "RUNE_LITERAL", "PRINT", "ATOI", "PARSE_FLOAT", "TYPE_OF",
		"INDEX_OF", "JOIN", "LEN", "APPEND",
	}
	staticData.RuleNames = []string{
		"program", "declaration", "variable_declaration", "type_annotation",
		"slice_type", "function_declaration", "receiver", "parameter_list",
		"parameter", "struct_declaration", "struct_field", "statement", "assignment_statement",
		"if_statement", "for_statement", "switch_statement", "case_clause",
		"default_clause", "break_statement", "continue_statement", "return_statement",
		"expression_statement", "block_statement", "expression", "primary_expression",
		"slice_literal", "struct_literal", "field_assignment_list", "field_assignment",
		"builtin_function_call", "argument_list", "expression_list",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 80, 449, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 1, 0, 1, 0, 1, 0, 5, 0, 68, 8, 0, 10, 0, 12, 0, 71, 9, 0, 1, 0, 5,
		0, 74, 8, 0, 10, 0, 12, 0, 77, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 83,
		8, 1, 1, 2, 1, 2, 1, 2, 3, 2, 88, 8, 2, 1, 2, 1, 2, 3, 2, 92, 8, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 102, 8, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 111, 8, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 121, 8, 5, 1, 5, 1, 5, 3, 5, 125, 8, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 135, 8, 5, 1, 5,
		1, 5, 3, 5, 139, 8, 5, 1, 5, 1, 5, 3, 5, 143, 8, 5, 1, 6, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 7, 5, 7, 151, 8, 7, 10, 7, 12, 7, 154, 9, 7, 1, 8, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 9, 1, 9, 4, 9, 163, 8, 9, 11, 9, 12, 9, 164, 1, 9, 1,
		9, 1, 10, 1, 10, 1, 10, 3, 10, 172, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 184, 8, 11, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 207,
		8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 216, 8,
		13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 236, 8,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 242, 8, 14, 1, 15, 1, 15, 1, 15,
		1, 15, 5, 15, 248, 8, 15, 10, 15, 12, 15, 251, 9, 15, 1, 15, 3, 15, 254,
		8, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 262, 8, 16, 10,
		16, 12, 16, 265, 9, 16, 1, 17, 1, 17, 1, 17, 5, 17, 270, 8, 17, 10, 17,
		12, 17, 273, 9, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 3, 20, 281,
		8, 20, 1, 21, 1, 21, 1, 22, 1, 22, 5, 22, 287, 8, 22, 10, 22, 12, 22, 290,
		9, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 300,
		8, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 331,
		8, 23, 1, 23, 5, 23, 334, 8, 23, 10, 23, 12, 23, 337, 9, 23, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 3, 24, 354, 8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 3, 25, 361, 8, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 3, 26, 368, 8,
		26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 5, 27, 375, 8, 27, 10, 27, 12, 27,
		378, 9, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 3, 29, 387,
		8, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 29, 3, 29, 431, 8, 29, 1, 30, 1, 30, 1, 30, 5, 30, 436, 8, 30, 10, 30,
		12, 30, 439, 9, 30, 1, 31, 1, 31, 1, 31, 5, 31, 444, 8, 31, 10, 31, 12,
		31, 447, 9, 31, 1, 31, 0, 1, 46, 32, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54,
		56, 58, 60, 62, 0, 4, 1, 0, 60, 62, 1, 0, 21, 22, 1, 0, 63, 66, 1, 0, 67,
		68, 499, 0, 69, 1, 0, 0, 0, 2, 82, 1, 0, 0, 0, 4, 101, 1, 0, 0, 0, 6, 110,
		1, 0, 0, 0, 8, 112, 1, 0, 0, 0, 10, 142, 1, 0, 0, 0, 12, 144, 1, 0, 0,
		0, 14, 147, 1, 0, 0, 0, 16, 155, 1, 0, 0, 0, 18, 158, 1, 0, 0, 0, 20, 168,
		1, 0, 0, 0, 22, 183, 1, 0, 0, 0, 24, 206, 1, 0, 0, 0, 26, 208, 1, 0, 0,
		0, 28, 241, 1, 0, 0, 0, 30, 243, 1, 0, 0, 0, 32, 257, 1, 0, 0, 0, 34, 266,
		1, 0, 0, 0, 36, 274, 1, 0, 0, 0, 38, 276, 1, 0, 0, 0, 40, 278, 1, 0, 0,
		0, 42, 282, 1, 0, 0, 0, 44, 284, 1, 0, 0, 0, 46, 299, 1, 0, 0, 0, 48, 353,
		1, 0, 0, 0, 50, 355, 1, 0, 0, 0, 52, 364, 1, 0, 0, 0, 54, 371, 1, 0, 0,
		0, 56, 379, 1, 0, 0, 0, 58, 430, 1, 0, 0, 0, 60, 432, 1, 0, 0, 0, 62, 440,
		1, 0, 0, 0, 64, 68, 3, 2, 1, 0, 65, 68, 3, 10, 5, 0, 66, 68, 3, 18, 9,
		0, 67, 64, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 66, 1, 0, 0, 0, 68, 71,
		1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 75, 1, 0, 0, 0,
		71, 69, 1, 0, 0, 0, 72, 74, 3, 10, 5, 0, 73, 72, 1, 0, 0, 0, 74, 77, 1,
		0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 78, 1, 0, 0, 0, 77,
		75, 1, 0, 0, 0, 78, 79, 5, 0, 0, 1, 79, 1, 1, 0, 0, 0, 80, 83, 3, 4, 2,
		0, 81, 83, 3, 24, 12, 0, 82, 80, 1, 0, 0, 0, 82, 81, 1, 0, 0, 0, 83, 3,
		1, 0, 0, 0, 84, 85, 5, 1, 0, 0, 85, 87, 5, 57, 0, 0, 86, 88, 3, 6, 3, 0,
		87, 86, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 90, 5,
		26, 0, 0, 90, 92, 3, 46, 23, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0,
		92, 102, 1, 0, 0, 0, 93, 94, 5, 57, 0, 0, 94, 95, 3, 6, 3, 0, 95, 96, 5,
		26, 0, 0, 96, 97, 3, 46, 23, 0, 97, 102, 1, 0, 0, 0, 98, 99, 5, 57, 0,
		0, 99, 100, 5, 58, 0, 0, 100, 102, 3, 46, 23, 0, 101, 84, 1, 0, 0, 0, 101,
		93, 1, 0, 0, 0, 101, 98, 1, 0, 0, 0, 102, 5, 1, 0, 0, 0, 103, 111, 5, 15,
		0, 0, 104, 111, 5, 16, 0, 0, 105, 111, 5, 17, 0, 0, 106, 111, 5, 18, 0,
		0, 107, 111, 5, 59, 0, 0, 108, 111, 3, 8, 4, 0, 109, 111, 5, 57, 0, 0,
		110, 103, 1, 0, 0, 0, 110, 104, 1, 0, 0, 0, 110, 105, 1, 0, 0, 0, 110,
		106, 1, 0, 0, 0, 110, 107, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 110, 109,
		1, 0, 0, 0, 111, 7, 1, 0, 0, 0, 112, 113, 5, 42, 0, 0, 113, 114, 5, 43,
		0, 0, 114, 115, 3, 6, 3, 0, 115, 9, 1, 0, 0, 0, 116, 117, 5, 2, 0, 0, 117,
		118, 5, 57, 0, 0, 118, 120, 5, 38, 0, 0, 119, 121, 3, 14, 7, 0, 120, 119,
		1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 124, 5, 39,
		0, 0, 123, 125, 3, 6, 3, 0, 124, 123, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0,
		125, 126, 1, 0, 0, 0, 126, 143, 3, 44, 22, 0, 127, 128, 5, 2, 0, 0, 128,
		129, 5, 38, 0, 0, 129, 130, 3, 12, 6, 0, 130, 131, 5, 39, 0, 0, 131, 132,
		5, 57, 0, 0, 132, 134, 5, 38, 0, 0, 133, 135, 3, 14, 7, 0, 134, 133, 1,
		0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 138, 5, 39, 0,
		0, 137, 139, 3, 6, 3, 0, 138, 137, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139,
		140, 1, 0, 0, 0, 140, 141, 3, 44, 22, 0, 141, 143, 1, 0, 0, 0, 142, 116,
		1, 0, 0, 0, 142, 127, 1, 0, 0, 0, 143, 11, 1, 0, 0, 0, 144, 145, 5, 57,
		0, 0, 145, 146, 3, 6, 3, 0, 146, 13, 1, 0, 0, 0, 147, 152, 3, 16, 8, 0,
		148, 149, 5, 45, 0, 0, 149, 151, 3, 16, 8, 0, 150, 148, 1, 0, 0, 0, 151,
		154, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 15, 1,
		0, 0, 0, 154, 152, 1, 0, 0, 0, 155, 156, 5, 57, 0, 0, 156, 157, 3, 6, 3,
		0, 157, 17, 1, 0, 0, 0, 158, 159, 5, 3, 0, 0, 159, 160, 5, 57, 0, 0, 160,
		162, 5, 40, 0, 0, 161, 163, 3, 20, 10, 0, 162, 161, 1, 0, 0, 0, 163, 164,
		1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 166, 1, 0,
		0, 0, 166, 167, 5, 41, 0, 0, 167, 19, 1, 0, 0, 0, 168, 169, 3, 6, 3, 0,
		169, 171, 5, 57, 0, 0, 170, 172, 5, 44, 0, 0, 171, 170, 1, 0, 0, 0, 171,
		172, 1, 0, 0, 0, 172, 21, 1, 0, 0, 0, 173, 184, 3, 2, 1, 0, 174, 184, 3,
		24, 12, 0, 175, 184, 3, 26, 13, 0, 176, 184, 3, 28, 14, 0, 177, 184, 3,
		30, 15, 0, 178, 184, 3, 36, 18, 0, 179, 184, 3, 38, 19, 0, 180, 184, 3,
		40, 20, 0, 181, 184, 3, 42, 21, 0, 182, 184, 3, 44, 22, 0, 183, 173, 1,
		0, 0, 0, 183, 174, 1, 0, 0, 0, 183, 175, 1, 0, 0, 0, 183, 176, 1, 0, 0,
		0, 183, 177, 1, 0, 0, 0, 183, 178, 1, 0, 0, 0, 183, 179, 1, 0, 0, 0, 183,
		180, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 182, 1, 0, 0, 0, 184, 23, 1,
		0, 0, 0, 185, 186, 5, 57, 0, 0, 186, 187, 5, 26, 0, 0, 187, 207, 3, 46,
		23, 0, 188, 189, 5, 57, 0, 0, 189, 190, 5, 27, 0, 0, 190, 207, 3, 46, 23,
		0, 191, 192, 5, 57, 0, 0, 192, 193, 5, 28, 0, 0, 193, 207, 3, 46, 23, 0,
		194, 195, 5, 57, 0, 0, 195, 196, 5, 42, 0, 0, 196, 197, 3, 46, 23, 0, 197,
		198, 5, 43, 0, 0, 198, 199, 5, 26, 0, 0, 199, 200, 3, 46, 23, 0, 200, 207,
		1, 0, 0, 0, 201, 202, 5, 57, 0, 0, 202, 203, 5, 46, 0, 0, 203, 204, 5,
		57, 0, 0, 204, 205, 5, 26, 0, 0, 205, 207, 3, 46, 23, 0, 206, 185, 1, 0,
		0, 0, 206, 188, 1, 0, 0, 0, 206, 191, 1, 0, 0, 0, 206, 194, 1, 0, 0, 0,
		206, 201, 1, 0, 0, 0, 207, 25, 1, 0, 0, 0, 208, 209, 5, 4, 0, 0, 209, 210,
		3, 46, 23, 0, 210, 215, 3, 44, 22, 0, 211, 212, 5, 5, 0, 0, 212, 216, 3,
		26, 13, 0, 213, 214, 5, 5, 0, 0, 214, 216, 3, 44, 22, 0, 215, 211, 1, 0,
		0, 0, 215, 213, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 27, 1, 0, 0, 0,
		217, 218, 5, 6, 0, 0, 218, 219, 3, 46, 23, 0, 219, 220, 3, 44, 22, 0, 220,
		242, 1, 0, 0, 0, 221, 222, 5, 6, 0, 0, 222, 223, 5, 57, 0, 0, 223, 224,
		5, 58, 0, 0, 224, 225, 3, 46, 23, 0, 225, 226, 5, 44, 0, 0, 226, 227, 3,
		46, 23, 0, 227, 228, 5, 44, 0, 0, 228, 229, 3, 24, 12, 0, 229, 230, 3,
		44, 22, 0, 230, 242, 1, 0, 0, 0, 231, 232, 5, 6, 0, 0, 232, 235, 5, 57,
		0, 0, 233, 234, 5, 45, 0, 0, 234, 236, 5, 57, 0, 0, 235, 233, 1, 0, 0,
		0, 235, 236, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237, 238, 5, 13, 0, 0, 238,
		239, 3, 46, 23, 0, 239, 240, 3, 44, 22, 0, 240, 242, 1, 0, 0, 0, 241, 217,
		1, 0, 0, 0, 241, 221, 1, 0, 0, 0, 241, 231, 1, 0, 0, 0, 242, 29, 1, 0,
		0, 0, 243, 244, 5, 7, 0, 0, 244, 245, 3, 46, 23, 0, 245, 249, 5, 40, 0,
		0, 246, 248, 3, 32, 16, 0, 247, 246, 1, 0, 0, 0, 248, 251, 1, 0, 0, 0,
		249, 247, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 253, 1, 0, 0, 0, 251,
		249, 1, 0, 0, 0, 252, 254, 3, 34, 17, 0, 253, 252, 1, 0, 0, 0, 253, 254,
		1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 256, 5, 41, 0, 0, 256, 31, 1, 0,
		0, 0, 257, 258, 5, 8, 0, 0, 258, 259, 3, 46, 23, 0, 259, 263, 5, 47, 0,
		0, 260, 262, 3, 22, 11, 0, 261, 260, 1, 0, 0, 0, 262, 265, 1, 0, 0, 0,
		263, 261, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 33, 1, 0, 0, 0, 265, 263,
		1, 0, 0, 0, 266, 267, 5, 9, 0, 0, 267, 271, 5, 47, 0, 0, 268, 270, 3, 22,
		11, 0, 269, 268, 1, 0, 0, 0, 270, 273, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0,
		271, 272, 1, 0, 0, 0, 272, 35, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 274, 275,
		5, 10, 0, 0, 275, 37, 1, 0, 0, 0, 276, 277, 5, 11, 0, 0, 277, 39, 1, 0,
		0, 0, 278, 280, 5, 12, 0, 0, 279, 281, 3, 46, 23, 0, 280, 279, 1, 0, 0,
		0, 280, 281, 1, 0, 0, 0, 281, 41, 1, 0, 0, 0, 282, 283, 3, 46, 23, 0, 283,
		43, 1, 0, 0, 0, 284, 288, 5, 40, 0, 0, 285, 287, 3, 22, 11, 0, 286, 285,
		1, 0, 0, 0, 287, 290, 1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 288, 289, 1, 0,
		0, 0, 289, 291, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 291, 292, 5, 41, 0, 0,
		292, 45, 1, 0, 0, 0, 293, 294, 6, 23, -1, 0, 294, 300, 3, 48, 24, 0, 295,
		296, 5, 22, 0, 0, 296, 300, 3, 46, 23, 8, 297, 298, 5, 37, 0, 0, 298, 300,
		3, 46, 23, 7, 299, 293, 1, 0, 0, 0, 299, 295, 1, 0, 0, 0, 299, 297, 1,
		0, 0, 0, 300, 335, 1, 0, 0, 0, 301, 302, 10, 6, 0, 0, 302, 303, 7, 0, 0,
		0, 303, 334, 3, 46, 23, 7, 304, 305, 10, 5, 0, 0, 305, 306, 7, 1, 0, 0,
		306, 334, 3, 46, 23, 6, 307, 308, 10, 4, 0, 0, 308, 309, 7, 2, 0, 0, 309,
		334, 3, 46, 23, 5, 310, 311, 10, 3, 0, 0, 311, 312, 7, 3, 0, 0, 312, 334,
		3, 46, 23, 4, 313, 314, 10, 2, 0, 0, 314, 315, 5, 35, 0, 0, 315, 334, 3,
		46, 23, 3, 316, 317, 10, 1, 0, 0, 317, 318, 5, 36, 0, 0, 318, 334, 3, 46,
		23, 2, 319, 320, 10, 11, 0, 0, 320, 321, 5, 46, 0, 0, 321, 334, 5, 57,
		0, 0, 322, 323, 10, 10, 0, 0, 323, 324, 5, 42, 0, 0, 324, 325, 3, 46, 23,
		0, 325, 326, 5, 43, 0, 0, 326, 334, 1, 0, 0, 0, 327, 328, 10, 9, 0, 0,
		328, 330, 5, 38, 0, 0, 329, 331, 3, 60, 30, 0, 330, 329, 1, 0, 0, 0, 330,
		331, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332, 334, 5, 39, 0, 0, 333, 301,
		1, 0, 0, 0, 333, 304, 1, 0, 0, 0, 333, 307, 1, 0, 0, 0, 333, 310, 1, 0,
		0, 0, 333, 313, 1, 0, 0, 0, 333, 316, 1, 0, 0, 0, 333, 319, 1, 0, 0, 0,
		333, 322, 1, 0, 0, 0, 333, 327, 1, 0, 0, 0, 334, 337, 1, 0, 0, 0, 335,
		333, 1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336, 47, 1, 0, 0, 0, 337, 335, 1,
		0, 0, 0, 338, 354, 5, 57, 0, 0, 339, 354, 5, 69, 0, 0, 340, 354, 5, 70,
		0, 0, 341, 354, 5, 71, 0, 0, 342, 354, 5, 72, 0, 0, 343, 354, 5, 19, 0,
		0, 344, 354, 5, 20, 0, 0, 345, 354, 5, 14, 0, 0, 346, 354, 3, 50, 25, 0,
		347, 354, 3, 52, 26, 0, 348, 354, 3, 58, 29, 0, 349, 350, 5, 38, 0, 0,
		350, 351, 3, 46, 23, 0, 351, 352, 5, 39, 0, 0, 352, 354, 1, 0, 0, 0, 353,
		338, 1, 0, 0, 0, 353, 339, 1, 0, 0, 0, 353, 340, 1, 0, 0, 0, 353, 341,
		1, 0, 0, 0, 353, 342, 1, 0, 0, 0, 353, 343, 1, 0, 0, 0, 353, 344, 1, 0,
		0, 0, 353, 345, 1, 0, 0, 0, 353, 346, 1, 0, 0, 0, 353, 347, 1, 0, 0, 0,
		353, 348, 1, 0, 0, 0, 353, 349, 1, 0, 0, 0, 354, 49, 1, 0, 0, 0, 355, 356,
		5, 42, 0, 0, 356, 357, 5, 43, 0, 0, 357, 358, 3, 6, 3, 0, 358, 360, 5,
		40, 0, 0, 359, 361, 3, 62, 31, 0, 360, 359, 1, 0, 0, 0, 360, 361, 1, 0,
		0, 0, 361, 362, 1, 0, 0, 0, 362, 363, 5, 41, 0, 0, 363, 51, 1, 0, 0, 0,
		364, 365, 5, 57, 0, 0, 365, 367, 5, 40, 0, 0, 366, 368, 3, 54, 27, 0, 367,
		366, 1, 0, 0, 0, 367, 368, 1, 0, 0, 0, 368, 369, 1, 0, 0, 0, 369, 370,
		5, 41, 0, 0, 370, 53, 1, 0, 0, 0, 371, 376, 3, 56, 28, 0, 372, 373, 5,
		45, 0, 0, 373, 375, 3, 56, 28, 0, 374, 372, 1, 0, 0, 0, 375, 378, 1, 0,
		0, 0, 376, 374, 1, 0, 0, 0, 376, 377, 1, 0, 0, 0, 377, 55, 1, 0, 0, 0,
		378, 376, 1, 0, 0, 0, 379, 380, 5, 57, 0, 0, 380, 381, 5, 47, 0, 0, 381,
		382, 3, 46, 23, 0, 382, 57, 1, 0, 0, 0, 383, 384, 5, 73, 0, 0, 384, 386,
		5, 38, 0, 0, 385, 387, 3, 60, 30, 0, 386, 385, 1, 0, 0, 0, 386, 387, 1,
		0, 0, 0, 387, 388, 1, 0, 0, 0, 388, 431, 5, 39, 0, 0, 389, 390, 5, 74,
		0, 0, 390, 391, 5, 38, 0, 0, 391, 392, 3, 46, 23, 0, 392, 393, 5, 39, 0,
		0, 393, 431, 1, 0, 0, 0, 394, 395, 5, 75, 0, 0, 395, 396, 5, 38, 0, 0,
		396, 397, 3, 46, 23, 0, 397, 398, 5, 39, 0, 0, 398, 431, 1, 0, 0, 0, 399,
		400, 5, 76, 0, 0, 400, 401, 5, 38, 0, 0, 401, 402, 3, 46, 23, 0, 402, 403,
		5, 39, 0, 0, 403, 431, 1, 0, 0, 0, 404, 405, 5, 77, 0, 0, 405, 406, 5,
		38, 0, 0, 406, 407, 3, 46, 23, 0, 407, 408, 5, 45, 0, 0, 408, 409, 3, 46,
		23, 0, 409, 410, 5, 39, 0, 0, 410, 431, 1, 0, 0, 0, 411, 412, 5, 78, 0,
		0, 412, 413, 5, 38, 0, 0, 413, 414, 3, 46, 23, 0, 414, 415, 5, 45, 0, 0,
		415, 416, 3, 46, 23, 0, 416, 417, 5, 39, 0, 0, 417, 431, 1, 0, 0, 0, 418,
		419, 5, 79, 0, 0, 419, 420, 5, 38, 0, 0, 420, 421, 3, 46, 23, 0, 421, 422,
		5, 39, 0, 0, 422, 431, 1, 0, 0, 0, 423, 424, 5, 80, 0, 0, 424, 425, 5,
		38, 0, 0, 425, 426, 3, 46, 23, 0, 426, 427, 5, 45, 0, 0, 427, 428, 3, 46,
		23, 0, 428, 429, 5, 39, 0, 0, 429, 431, 1, 0, 0, 0, 430, 383, 1, 0, 0,
		0, 430, 389, 1, 0, 0, 0, 430, 394, 1, 0, 0, 0, 430, 399, 1, 0, 0, 0, 430,
		404, 1, 0, 0, 0, 430, 411, 1, 0, 0, 0, 430, 418, 1, 0, 0, 0, 430, 423,
		1, 0, 0, 0, 431, 59, 1, 0, 0, 0, 432, 437, 3, 46, 23, 0, 433, 434, 5, 45,
		0, 0, 434, 436, 3, 46, 23, 0, 435, 433, 1, 0, 0, 0, 436, 439, 1, 0, 0,
		0, 437, 435, 1, 0, 0, 0, 437, 438, 1, 0, 0, 0, 438, 61, 1, 0, 0, 0, 439,
		437, 1, 0, 0, 0, 440, 445, 3, 46, 23, 0, 441, 442, 5, 45, 0, 0, 442, 444,
		3, 46, 23, 0, 443, 441, 1, 0, 0, 0, 444, 447, 1, 0, 0, 0, 445, 443, 1,
		0, 0, 0, 445, 446, 1, 0, 0, 0, 446, 63, 1, 0, 0, 0, 447, 445, 1, 0, 0,
		0, 39, 67, 69, 75, 82, 87, 91, 101, 110, 120, 124, 134, 138, 142, 152,
		164, 171, 183, 206, 215, 235, 241, 249, 253, 263, 271, 280, 288, 299, 330,
		333, 335, 353, 360, 367, 376, 386, 430, 437, 445,
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

// VLangGrammarParserInit initializes any static state used to implement VLangGrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVLangGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func VLangGrammarParserInit() {
	staticData := &VLangGrammarParserStaticData
	staticData.once.Do(vlanggrammarParserInit)
}

// NewVLangGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewVLangGrammarParser(input antlr.TokenStream) *VLangGrammarParser {
	VLangGrammarParserInit()
	this := new(VLangGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &VLangGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "VLangGrammar.g4"

	return this
}

// VLangGrammarParser tokens.
const (
	VLangGrammarParserEOF            = antlr.TokenEOF
	VLangGrammarParserMUT            = 1
	VLangGrammarParserFUNC           = 2
	VLangGrammarParserSTRUCT         = 3
	VLangGrammarParserIF             = 4
	VLangGrammarParserELSE           = 5
	VLangGrammarParserFOR            = 6
	VLangGrammarParserSWITCH         = 7
	VLangGrammarParserCASE           = 8
	VLangGrammarParserDEFAULT        = 9
	VLangGrammarParserBREAK          = 10
	VLangGrammarParserCONTINUE       = 11
	VLangGrammarParserRETURN         = 12
	VLangGrammarParserIN             = 13
	VLangGrammarParserNIL            = 14
	VLangGrammarParserINT_TYPE       = 15
	VLangGrammarParserFLOAT_TYPE     = 16
	VLangGrammarParserSTRING_TYPE    = 17
	VLangGrammarParserBOOL_TYPE      = 18
	VLangGrammarParserTRUE           = 19
	VLangGrammarParserFALSE          = 20
	VLangGrammarParserPLUS           = 21
	VLangGrammarParserMINUS          = 22
	VLangGrammarParserMULT           = 23
	VLangGrammarParserDIV            = 24
	VLangGrammarParserMOD            = 25
	VLangGrammarParserASSIGN         = 26
	VLangGrammarParserPLUS_ASSIGN    = 27
	VLangGrammarParserMINUS_ASSIGN   = 28
	VLangGrammarParserEQ             = 29
	VLangGrammarParserNE             = 30
	VLangGrammarParserLT             = 31
	VLangGrammarParserLE             = 32
	VLangGrammarParserGT             = 33
	VLangGrammarParserGE             = 34
	VLangGrammarParserAND            = 35
	VLangGrammarParserOR             = 36
	VLangGrammarParserNOT            = 37
	VLangGrammarParserLPAREN         = 38
	VLangGrammarParserRPAREN         = 39
	VLangGrammarParserLBRACE         = 40
	VLangGrammarParserRBRACE         = 41
	VLangGrammarParserLBRACKET       = 42
	VLangGrammarParserRBRACKET       = 43
	VLangGrammarParserSEMICOLON      = 44
	VLangGrammarParserCOMMA          = 45
	VLangGrammarParserDOT            = 46
	VLangGrammarParserCOLON          = 47
	VLangGrammarParserAMPERSAND      = 48
	VLangGrammarParserINT            = 49
	VLangGrammarParserFLOAT          = 50
	VLangGrammarParserSTRING         = 51
	VLangGrammarParserRUNE           = 52
	VLangGrammarParserID             = 53
	VLangGrammarParserLINE_COMMENT   = 54
	VLangGrammarParserBLOCK_COMMENT  = 55
	VLangGrammarParserWS             = 56
	VLangGrammarParserIDENTIFIER     = 57
	VLangGrammarParserDECLARE_ASSIGN = 58
	VLangGrammarParserRUNE_TYPE      = 59
	VLangGrammarParserMULTIPLY       = 60
	VLangGrammarParserDIVIDE         = 61
	VLangGrammarParserMODULO         = 62
	VLangGrammarParserLESS_THAN      = 63
	VLangGrammarParserLESS_EQUAL     = 64
	VLangGrammarParserGREATER_THAN   = 65
	VLangGrammarParserGREATER_EQUAL  = 66
	VLangGrammarParserEQUAL          = 67
	VLangGrammarParserNOT_EQUAL      = 68
	VLangGrammarParserINT_LITERAL    = 69
	VLangGrammarParserFLOAT_LITERAL  = 70
	VLangGrammarParserSTRING_LITERAL = 71
	VLangGrammarParserRUNE_LITERAL   = 72
	VLangGrammarParserPRINT          = 73
	VLangGrammarParserATOI           = 74
	VLangGrammarParserPARSE_FLOAT    = 75
	VLangGrammarParserTYPE_OF        = 76
	VLangGrammarParserINDEX_OF       = 77
	VLangGrammarParserJOIN           = 78
	VLangGrammarParserLEN            = 79
	VLangGrammarParserAPPEND         = 80
)

// VLangGrammarParser rules.
const (
	VLangGrammarParserRULE_program               = 0
	VLangGrammarParserRULE_declaration           = 1
	VLangGrammarParserRULE_variable_declaration  = 2
	VLangGrammarParserRULE_type_annotation       = 3
	VLangGrammarParserRULE_slice_type            = 4
	VLangGrammarParserRULE_function_declaration  = 5
	VLangGrammarParserRULE_receiver              = 6
	VLangGrammarParserRULE_parameter_list        = 7
	VLangGrammarParserRULE_parameter             = 8
	VLangGrammarParserRULE_struct_declaration    = 9
	VLangGrammarParserRULE_struct_field          = 10
	VLangGrammarParserRULE_statement             = 11
	VLangGrammarParserRULE_assignment_statement  = 12
	VLangGrammarParserRULE_if_statement          = 13
	VLangGrammarParserRULE_for_statement         = 14
	VLangGrammarParserRULE_switch_statement      = 15
	VLangGrammarParserRULE_case_clause           = 16
	VLangGrammarParserRULE_default_clause        = 17
	VLangGrammarParserRULE_break_statement       = 18
	VLangGrammarParserRULE_continue_statement    = 19
	VLangGrammarParserRULE_return_statement      = 20
	VLangGrammarParserRULE_expression_statement  = 21
	VLangGrammarParserRULE_block_statement       = 22
	VLangGrammarParserRULE_expression            = 23
	VLangGrammarParserRULE_primary_expression    = 24
	VLangGrammarParserRULE_slice_literal         = 25
	VLangGrammarParserRULE_struct_literal        = 26
	VLangGrammarParserRULE_field_assignment_list = 27
	VLangGrammarParserRULE_field_assignment      = 28
	VLangGrammarParserRULE_builtin_function_call = 29
	VLangGrammarParserRULE_argument_list         = 30
	VLangGrammarParserRULE_expression_list       = 31
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllDeclaration() []IDeclarationContext
	Declaration(i int) IDeclarationContext
	AllFunction_declaration() []IFunction_declarationContext
	Function_declaration(i int) IFunction_declarationContext
	AllStruct_declaration() []IStruct_declarationContext
	Struct_declaration(i int) IStruct_declarationContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserEOF, 0)
}

func (s *ProgramContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
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

	return t.(IDeclarationContext)
}

func (s *ProgramContext) AllFunction_declaration() []IFunction_declarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunction_declarationContext); ok {
			len++
		}
	}

	tst := make([]IFunction_declarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunction_declarationContext); ok {
			tst[i] = t.(IFunction_declarationContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Function_declaration(i int) IFunction_declarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_declarationContext); ok {
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

	return t.(IFunction_declarationContext)
}

func (s *ProgramContext) AllStruct_declaration() []IStruct_declarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStruct_declarationContext); ok {
			len++
		}
	}

	tst := make([]IStruct_declarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStruct_declarationContext); ok {
			tst[i] = t.(IStruct_declarationContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Struct_declaration(i int) IStruct_declarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_declarationContext); ok {
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

	return t.(IStruct_declarationContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VLangGrammarParserRULE_program)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(67)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case VLangGrammarParserMUT, VLangGrammarParserIDENTIFIER:
				{
					p.SetState(64)
					p.Declaration()
				}

			case VLangGrammarParserFUNC:
				{
					p.SetState(65)
					p.Function_declaration()
				}

			case VLangGrammarParserSTRUCT:
				{
					p.SetState(66)
					p.Struct_declaration()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		}
		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VLangGrammarParserFUNC {
		{
			p.SetState(72)
			p.Function_declaration()
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(78)
		p.Match(VLangGrammarParserEOF)
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

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Variable_declaration() IVariable_declarationContext
	Assignment_statement() IAssignment_statementContext

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
	p.RuleIndex = VLangGrammarParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_declaration

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

func (s *DeclarationContext) Assignment_statement() IAssignment_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignment_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignment_statementContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VLangGrammarParserRULE_declaration)
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.Variable_declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.Assignment_statement()
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
	p.RuleIndex = VLangGrammarParserRULE_variable_declaration
	return p
}

func InitEmptyVariable_declarationContext(p *Variable_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_variable_declaration
}

func (*Variable_declarationContext) IsVariable_declarationContext() {}

func NewVariable_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Variable_declarationContext {
	var p = new(Variable_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_variable_declaration

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

type InferredVarDeclContext struct {
	Variable_declarationContext
}

func NewInferredVarDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InferredVarDeclContext {
	var p = new(InferredVarDeclContext)

	InitEmptyVariable_declarationContext(&p.Variable_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Variable_declarationContext))

	return p
}

func (s *InferredVarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InferredVarDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserIDENTIFIER, 0)
}

func (s *InferredVarDeclContext) DECLARE_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserDECLARE_ASSIGN, 0)
}

func (s *InferredVarDeclContext) Expression() IExpressionContext {
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

func (s *InferredVarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterInferredVarDecl(s)
	}
}

func (s *InferredVarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitInferredVarDecl(s)
	}
}

func (s *InferredVarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitInferredVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
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

func (s *TypedVarDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserIDENTIFIER, 0)
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
	return s.GetToken(VLangGrammarParserASSIGN, 0)
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
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterTypedVarDecl(s)
	}
}

func (s *TypedVarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitTypedVarDecl(s)
	}
}

func (s *TypedVarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
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
	return s.GetToken(VLangGrammarParserMUT, 0)
}

func (s *MutableVarDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserIDENTIFIER, 0)
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
	return s.GetToken(VLangGrammarParserASSIGN, 0)
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
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterMutableVarDecl(s)
	}
}

func (s *MutableVarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitMutableVarDecl(s)
	}
}

func (s *MutableVarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitMutableVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Variable_declaration() (localctx IVariable_declarationContext) {
	localctx = NewVariable_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, VLangGrammarParserRULE_variable_declaration)
	var _la int

	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMutableVarDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.Match(VLangGrammarParserMUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.Match(VLangGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(86)
				p.Type_annotation()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == VLangGrammarParserASSIGN {
			{
				p.SetState(89)
				p.Match(VLangGrammarParserASSIGN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(90)
				p.expression(0)
			}

		}

	case 2:
		localctx = NewTypedVarDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.Match(VLangGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.Type_annotation()
		}
		{
			p.SetState(95)
			p.Match(VLangGrammarParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)
			p.expression(0)
		}

	case 3:
		localctx = NewInferredVarDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(98)
			p.Match(VLangGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.Match(VLangGrammarParserDECLARE_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
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
	RUNE_TYPE() antlr.TerminalNode
	Slice_type() ISlice_typeContext
	IDENTIFIER() antlr.TerminalNode

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
	p.RuleIndex = VLangGrammarParserRULE_type_annotation
	return p
}

func InitEmptyType_annotationContext(p *Type_annotationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_type_annotation
}

func (*Type_annotationContext) IsType_annotationContext() {}

func NewType_annotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_annotationContext {
	var p = new(Type_annotationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_type_annotation

	return p
}

func (s *Type_annotationContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_annotationContext) INT_TYPE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserINT_TYPE, 0)
}

func (s *Type_annotationContext) FLOAT_TYPE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserFLOAT_TYPE, 0)
}

func (s *Type_annotationContext) STRING_TYPE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserSTRING_TYPE, 0)
}

func (s *Type_annotationContext) BOOL_TYPE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserBOOL_TYPE, 0)
}

func (s *Type_annotationContext) RUNE_TYPE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserRUNE_TYPE, 0)
}

func (s *Type_annotationContext) Slice_type() ISlice_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISlice_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISlice_typeContext)
}

func (s *Type_annotationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserIDENTIFIER, 0)
}

func (s *Type_annotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_annotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_annotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterType_annotation(s)
	}
}

func (s *Type_annotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitType_annotation(s)
	}
}

func (s *Type_annotationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitType_annotation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Type_annotation() (localctx IType_annotationContext) {
	localctx = NewType_annotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, VLangGrammarParserRULE_type_annotation)
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VLangGrammarParserINT_TYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.Match(VLangGrammarParserINT_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangGrammarParserFLOAT_TYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.Match(VLangGrammarParserFLOAT_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangGrammarParserSTRING_TYPE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(105)
			p.Match(VLangGrammarParserSTRING_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangGrammarParserBOOL_TYPE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(106)
			p.Match(VLangGrammarParserBOOL_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangGrammarParserRUNE_TYPE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(107)
			p.Match(VLangGrammarParserRUNE_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangGrammarParserLBRACKET:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(108)
			p.Slice_type()
		}

	case VLangGrammarParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(109)
			p.Match(VLangGrammarParserIDENTIFIER)
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

// ISlice_typeContext is an interface to support dynamic dispatch.
type ISlice_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	Type_annotation() IType_annotationContext

	// IsSlice_typeContext differentiates from other interfaces.
	IsSlice_typeContext()
}

type Slice_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySlice_typeContext() *Slice_typeContext {
	var p = new(Slice_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_slice_type
	return p
}

func InitEmptySlice_typeContext(p *Slice_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_slice_type
}

func (*Slice_typeContext) IsSlice_typeContext() {}

func NewSlice_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Slice_typeContext {
	var p = new(Slice_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_slice_type

	return p
}

func (s *Slice_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Slice_typeContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserLBRACKET, 0)
}

func (s *Slice_typeContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserRBRACKET, 0)
}

func (s *Slice_typeContext) Type_annotation() IType_annotationContext {
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

func (s *Slice_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Slice_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Slice_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterSlice_type(s)
	}
}

func (s *Slice_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitSlice_type(s)
	}
}

func (s *Slice_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitSlice_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Slice_type() (localctx ISlice_typeContext) {
	localctx = NewSlice_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, VLangGrammarParserRULE_slice_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(VLangGrammarParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.Match(VLangGrammarParserRBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)
		p.Type_annotation()
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

// IFunction_declarationContext is an interface to support dynamic dispatch.
type IFunction_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	AllLPAREN() []antlr.TerminalNode
	LPAREN(i int) antlr.TerminalNode
	AllRPAREN() []antlr.TerminalNode
	RPAREN(i int) antlr.TerminalNode
	Block_statement() IBlock_statementContext
	Parameter_list() IParameter_listContext
	Type_annotation() IType_annotationContext
	Receiver() IReceiverContext

	// IsFunction_declarationContext differentiates from other interfaces.
	IsFunction_declarationContext()
}

type Function_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_declarationContext() *Function_declarationContext {
	var p = new(Function_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_function_declaration
	return p
}

func InitEmptyFunction_declarationContext(p *Function_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_function_declaration
}

func (*Function_declarationContext) IsFunction_declarationContext() {}

func NewFunction_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_declarationContext {
	var p = new(Function_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_function_declaration

	return p
}

func (s *Function_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_declarationContext) FUNC() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserFUNC, 0)
}

func (s *Function_declarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserIDENTIFIER, 0)
}

func (s *Function_declarationContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(VLangGrammarParserLPAREN)
}

func (s *Function_declarationContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserLPAREN, i)
}

func (s *Function_declarationContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(VLangGrammarParserRPAREN)
}

func (s *Function_declarationContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserRPAREN, i)
}

func (s *Function_declarationContext) Block_statement() IBlock_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_statementContext)
}

func (s *Function_declarationContext) Parameter_list() IParameter_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameter_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameter_listContext)
}

func (s *Function_declarationContext) Type_annotation() IType_annotationContext {
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

func (s *Function_declarationContext) Receiver() IReceiverContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReceiverContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReceiverContext)
}

func (s *Function_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterFunction_declaration(s)
	}
}

func (s *Function_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitFunction_declaration(s)
	}
}

func (s *Function_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitFunction_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Function_declaration() (localctx IFunction_declarationContext) {
	localctx = NewFunction_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, VLangGrammarParserRULE_function_declaration)
	var _la int

	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.Match(VLangGrammarParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Match(VLangGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.Match(VLangGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == VLangGrammarParserIDENTIFIER {
			{
				p.SetState(119)
				p.Parameter_list()
			}

		}
		{
			p.SetState(122)
			p.Match(VLangGrammarParserRPAREN)
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

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&720580338426281984) != 0 {
			{
				p.SetState(123)
				p.Type_annotation()
			}

		}
		{
			p.SetState(126)
			p.Block_statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(127)
			p.Match(VLangGrammarParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.Match(VLangGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)
			p.Receiver()
		}
		{
			p.SetState(130)
			p.Match(VLangGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.Match(VLangGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Match(VLangGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == VLangGrammarParserIDENTIFIER {
			{
				p.SetState(133)
				p.Parameter_list()
			}

		}
		{
			p.SetState(136)
			p.Match(VLangGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&720580338426281984) != 0 {
			{
				p.SetState(137)
				p.Type_annotation()
			}

		}
		{
			p.SetState(140)
			p.Block_statement()
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

// IReceiverContext is an interface to support dynamic dispatch.
type IReceiverContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Type_annotation() IType_annotationContext

	// IsReceiverContext differentiates from other interfaces.
	IsReceiverContext()
}

type ReceiverContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReceiverContext() *ReceiverContext {
	var p = new(ReceiverContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_receiver
	return p
}

func InitEmptyReceiverContext(p *ReceiverContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_receiver
}

func (*ReceiverContext) IsReceiverContext() {}

func NewReceiverContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReceiverContext {
	var p = new(ReceiverContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_receiver

	return p
}

func (s *ReceiverContext) GetParser() antlr.Parser { return s.parser }

func (s *ReceiverContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserIDENTIFIER, 0)
}

func (s *ReceiverContext) Type_annotation() IType_annotationContext {
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

func (s *ReceiverContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReceiverContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReceiverContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterReceiver(s)
	}
}

func (s *ReceiverContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitReceiver(s)
	}
}

func (s *ReceiverContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitReceiver(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Receiver() (localctx IReceiverContext) {
	localctx = NewReceiverContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, VLangGrammarParserRULE_receiver)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(VLangGrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Type_annotation()
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

// IParameter_listContext is an interface to support dynamic dispatch.
type IParameter_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParameter_listContext differentiates from other interfaces.
	IsParameter_listContext()
}

type Parameter_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameter_listContext() *Parameter_listContext {
	var p = new(Parameter_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_parameter_list
	return p
}

func InitEmptyParameter_listContext(p *Parameter_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_parameter_list
}

func (*Parameter_listContext) IsParameter_listContext() {}

func NewParameter_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parameter_listContext {
	var p = new(Parameter_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_parameter_list

	return p
}

func (s *Parameter_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Parameter_listContext) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *Parameter_listContext) Parameter(i int) IParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
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

	return t.(IParameterContext)
}

func (s *Parameter_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VLangGrammarParserCOMMA)
}

func (s *Parameter_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserCOMMA, i)
}

func (s *Parameter_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parameter_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Parameter_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterParameter_list(s)
	}
}

func (s *Parameter_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitParameter_list(s)
	}
}

func (s *Parameter_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitParameter_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Parameter_list() (localctx IParameter_listContext) {
	localctx = NewParameter_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, VLangGrammarParserRULE_parameter_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Parameter()
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VLangGrammarParserCOMMA {
		{
			p.SetState(148)
			p.Match(VLangGrammarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.Parameter()
		}

		p.SetState(154)
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

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Type_annotation() IType_annotationContext

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_parameter
	return p
}

func InitEmptyParameterContext(p *ParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_parameter
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserIDENTIFIER, 0)
}

func (s *ParameterContext) Type_annotation() IType_annotationContext {
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

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, VLangGrammarParserRULE_parameter)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(VLangGrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Type_annotation()
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

// IStruct_declarationContext is an interface to support dynamic dispatch.
type IStruct_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRUCT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStruct_field() []IStruct_fieldContext
	Struct_field(i int) IStruct_fieldContext

	// IsStruct_declarationContext differentiates from other interfaces.
	IsStruct_declarationContext()
}

type Struct_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_declarationContext() *Struct_declarationContext {
	var p = new(Struct_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_struct_declaration
	return p
}

func InitEmptyStruct_declarationContext(p *Struct_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_struct_declaration
}

func (*Struct_declarationContext) IsStruct_declarationContext() {}

func NewStruct_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_declarationContext {
	var p = new(Struct_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_struct_declaration

	return p
}

func (s *Struct_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_declarationContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserSTRUCT, 0)
}

func (s *Struct_declarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserIDENTIFIER, 0)
}

func (s *Struct_declarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserLBRACE, 0)
}

func (s *Struct_declarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserRBRACE, 0)
}

func (s *Struct_declarationContext) AllStruct_field() []IStruct_fieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStruct_fieldContext); ok {
			len++
		}
	}

	tst := make([]IStruct_fieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStruct_fieldContext); ok {
			tst[i] = t.(IStruct_fieldContext)
			i++
		}
	}

	return tst
}

func (s *Struct_declarationContext) Struct_field(i int) IStruct_fieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_fieldContext); ok {
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

	return t.(IStruct_fieldContext)
}

func (s *Struct_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Struct_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterStruct_declaration(s)
	}
}

func (s *Struct_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitStruct_declaration(s)
	}
}

func (s *Struct_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitStruct_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Struct_declaration() (localctx IStruct_declarationContext) {
	localctx = NewStruct_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, VLangGrammarParserRULE_struct_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Match(VLangGrammarParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.Match(VLangGrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)
		p.Match(VLangGrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&720580338426281984) != 0) {
		{
			p.SetState(161)
			p.Struct_field()
		}

		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(166)
		p.Match(VLangGrammarParserRBRACE)
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

// IStruct_fieldContext is an interface to support dynamic dispatch.
type IStruct_fieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_annotation() IType_annotationContext
	IDENTIFIER() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode

	// IsStruct_fieldContext differentiates from other interfaces.
	IsStruct_fieldContext()
}

type Struct_fieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_fieldContext() *Struct_fieldContext {
	var p = new(Struct_fieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_struct_field
	return p
}

func InitEmptyStruct_fieldContext(p *Struct_fieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_struct_field
}

func (*Struct_fieldContext) IsStruct_fieldContext() {}

func NewStruct_fieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_fieldContext {
	var p = new(Struct_fieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_struct_field

	return p
}

func (s *Struct_fieldContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_fieldContext) Type_annotation() IType_annotationContext {
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

func (s *Struct_fieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserIDENTIFIER, 0)
}

func (s *Struct_fieldContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserSEMICOLON, 0)
}

func (s *Struct_fieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_fieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Struct_fieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterStruct_field(s)
	}
}

func (s *Struct_fieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitStruct_field(s)
	}
}

func (s *Struct_fieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitStruct_field(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Struct_field() (localctx IStruct_fieldContext) {
	localctx = NewStruct_fieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, VLangGrammarParserRULE_struct_field)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Type_annotation()
	}
	{
		p.SetState(169)
		p.Match(VLangGrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VLangGrammarParserSEMICOLON {
		{
			p.SetState(170)
			p.Match(VLangGrammarParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declaration() IDeclarationContext
	Assignment_statement() IAssignment_statementContext
	If_statement() IIf_statementContext
	For_statement() IFor_statementContext
	Switch_statement() ISwitch_statementContext
	Break_statement() IBreak_statementContext
	Continue_statement() IContinue_statementContext
	Return_statement() IReturn_statementContext
	Expression_statement() IExpression_statementContext
	Block_statement() IBlock_statementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Declaration() IDeclarationContext {
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

func (s *StatementContext) Assignment_statement() IAssignment_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignment_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignment_statementContext)
}

func (s *StatementContext) If_statement() IIf_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_statementContext)
}

func (s *StatementContext) For_statement() IFor_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_statementContext)
}

func (s *StatementContext) Switch_statement() ISwitch_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_statementContext)
}

func (s *StatementContext) Break_statement() IBreak_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreak_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreak_statementContext)
}

func (s *StatementContext) Continue_statement() IContinue_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinue_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinue_statementContext)
}

func (s *StatementContext) Return_statement() IReturn_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturn_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturn_statementContext)
}

func (s *StatementContext) Expression_statement() IExpression_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpression_statementContext)
}

func (s *StatementContext) Block_statement() IBlock_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_statementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, VLangGrammarParserRULE_statement)
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(173)
			p.Declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(174)
			p.Assignment_statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(175)
			p.If_statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(176)
			p.For_statement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(177)
			p.Switch_statement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(178)
			p.Break_statement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(179)
			p.Continue_statement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(180)
			p.Return_statement()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(181)
			p.Expression_statement()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(182)
			p.Block_statement()
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

// IAssignment_statementContext is an interface to support dynamic dispatch.
type IAssignment_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssignment_statementContext differentiates from other interfaces.
	IsAssignment_statementContext()
}

type Assignment_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignment_statementContext() *Assignment_statementContext {
	var p = new(Assignment_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_assignment_statement
	return p
}

func InitEmptyAssignment_statementContext(p *Assignment_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_assignment_statement
}

func (*Assignment_statementContext) IsAssignment_statementContext() {}

func NewAssignment_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assignment_statementContext {
	var p = new(Assignment_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_assignment_statement

	return p
}

func (s *Assignment_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Assignment_statementContext) CopyAll(ctx *Assignment_statementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Assignment_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assignment_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MinusAssignmentContext struct {
	Assignment_statementContext
}

func NewMinusAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MinusAssignmentContext {
	var p = new(MinusAssignmentContext)

	InitEmptyAssignment_statementContext(&p.Assignment_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assignment_statementContext))

	return p
}

func (s *MinusAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinusAssignmentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserIDENTIFIER, 0)
}

func (s *MinusAssignmentContext) MINUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserMINUS_ASSIGN, 0)
}

func (s *MinusAssignmentContext) Expression() IExpressionContext {
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

func (s *MinusAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterMinusAssignment(s)
	}
}

func (s *MinusAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitMinusAssignment(s)
	}
}

func (s *MinusAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitMinusAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type PlusAssignmentContext struct {
	Assignment_statementContext
}

func NewPlusAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlusAssignmentContext {
	var p = new(PlusAssignmentContext)

	InitEmptyAssignment_statementContext(&p.Assignment_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assignment_statementContext))

	return p
}

func (s *PlusAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlusAssignmentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserIDENTIFIER, 0)
}

func (s *PlusAssignmentContext) PLUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserPLUS_ASSIGN, 0)
}

func (s *PlusAssignmentContext) Expression() IExpressionContext {
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

func (s *PlusAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterPlusAssignment(s)
	}
}

func (s *PlusAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitPlusAssignment(s)
	}
}

func (s *PlusAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitPlusAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type FieldAssignmentContext struct {
	Assignment_statementContext
}

func NewFieldAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldAssignmentContext {
	var p = new(FieldAssignmentContext)

	InitEmptyAssignment_statementContext(&p.Assignment_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assignment_statementContext))

	return p
}

func (s *FieldAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldAssignmentContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(VLangGrammarParserIDENTIFIER)
}

func (s *FieldAssignmentContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserIDENTIFIER, i)
}

func (s *FieldAssignmentContext) DOT() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserDOT, 0)
}

func (s *FieldAssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserASSIGN, 0)
}

func (s *FieldAssignmentContext) Expression() IExpressionContext {
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

func (s *FieldAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterFieldAssignment(s)
	}
}

func (s *FieldAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitFieldAssignment(s)
	}
}

func (s *FieldAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitFieldAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type SimpleAssignmentContext struct {
	Assignment_statementContext
}

func NewSimpleAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleAssignmentContext {
	var p = new(SimpleAssignmentContext)

	InitEmptyAssignment_statementContext(&p.Assignment_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assignment_statementContext))

	return p
}

func (s *SimpleAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleAssignmentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserIDENTIFIER, 0)
}

func (s *SimpleAssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserASSIGN, 0)
}

func (s *SimpleAssignmentContext) Expression() IExpressionContext {
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

func (s *SimpleAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterSimpleAssignment(s)
	}
}

func (s *SimpleAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitSimpleAssignment(s)
	}
}

func (s *SimpleAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitSimpleAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type IndexAssignmentContext struct {
	Assignment_statementContext
}

func NewIndexAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexAssignmentContext {
	var p = new(IndexAssignmentContext)

	InitEmptyAssignment_statementContext(&p.Assignment_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assignment_statementContext))

	return p
}

func (s *IndexAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexAssignmentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserIDENTIFIER, 0)
}

func (s *IndexAssignmentContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserLBRACKET, 0)
}

func (s *IndexAssignmentContext) AllExpression() []IExpressionContext {
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

func (s *IndexAssignmentContext) Expression(i int) IExpressionContext {
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

func (s *IndexAssignmentContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserRBRACKET, 0)
}

func (s *IndexAssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserASSIGN, 0)
}

func (s *IndexAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterIndexAssignment(s)
	}
}

func (s *IndexAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitIndexAssignment(s)
	}
}

func (s *IndexAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitIndexAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Assignment_statement() (localctx IAssignment_statementContext) {
	localctx = NewAssignment_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, VLangGrammarParserRULE_assignment_statement)
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSimpleAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)
			p.Match(VLangGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.Match(VLangGrammarParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.expression(0)
		}

	case 2:
		localctx = NewPlusAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(188)
			p.Match(VLangGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.Match(VLangGrammarParserPLUS_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.expression(0)
		}

	case 3:
		localctx = NewMinusAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(191)
			p.Match(VLangGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(192)
			p.Match(VLangGrammarParserMINUS_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.expression(0)
		}

	case 4:
		localctx = NewIndexAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(194)
			p.Match(VLangGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(195)
			p.Match(VLangGrammarParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)
			p.expression(0)
		}
		{
			p.SetState(197)
			p.Match(VLangGrammarParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.Match(VLangGrammarParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
			p.expression(0)
		}

	case 5:
		localctx = NewFieldAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(201)
			p.Match(VLangGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(202)
			p.Match(VLangGrammarParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)
			p.Match(VLangGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)
			p.Match(VLangGrammarParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(205)
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

// IIf_statementContext is an interface to support dynamic dispatch.
type IIf_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Expression() IExpressionContext
	AllBlock_statement() []IBlock_statementContext
	Block_statement(i int) IBlock_statementContext
	ELSE() antlr.TerminalNode
	If_statement() IIf_statementContext

	// IsIf_statementContext differentiates from other interfaces.
	IsIf_statementContext()
}

type If_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_statementContext() *If_statementContext {
	var p = new(If_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_if_statement
	return p
}

func InitEmptyIf_statementContext(p *If_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_if_statement
}

func (*If_statementContext) IsIf_statementContext() {}

func NewIf_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_statementContext {
	var p = new(If_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_if_statement

	return p
}

func (s *If_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *If_statementContext) IF() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserIF, 0)
}

func (s *If_statementContext) Expression() IExpressionContext {
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

func (s *If_statementContext) AllBlock_statement() []IBlock_statementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlock_statementContext); ok {
			len++
		}
	}

	tst := make([]IBlock_statementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlock_statementContext); ok {
			tst[i] = t.(IBlock_statementContext)
			i++
		}
	}

	return tst
}

func (s *If_statementContext) Block_statement(i int) IBlock_statementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_statementContext); ok {
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

	return t.(IBlock_statementContext)
}

func (s *If_statementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserELSE, 0)
}

func (s *If_statementContext) If_statement() IIf_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_statementContext)
}

func (s *If_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterIf_statement(s)
	}
}

func (s *If_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitIf_statement(s)
	}
}

func (s *If_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitIf_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) If_statement() (localctx IIf_statementContext) {
	localctx = NewIf_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, VLangGrammarParserRULE_if_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(VLangGrammarParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.expression(0)
	}
	{
		p.SetState(210)
		p.Block_statement()
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(211)
			p.Match(VLangGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(212)
			p.If_statement()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(213)
			p.Match(VLangGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(214)
			p.Block_statement()
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

// IFor_statementContext is an interface to support dynamic dispatch.
type IFor_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFor_statementContext differentiates from other interfaces.
	IsFor_statementContext()
}

type For_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_statementContext() *For_statementContext {
	var p = new(For_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_for_statement
	return p
}

func InitEmptyFor_statementContext(p *For_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_for_statement
}

func (*For_statementContext) IsFor_statementContext() {}

func NewFor_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_statementContext {
	var p = new(For_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_for_statement

	return p
}

func (s *For_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *For_statementContext) CopyAll(ctx *For_statementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *For_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForInContext struct {
	For_statementContext
}

func NewForInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForInContext {
	var p = new(ForInContext)

	InitEmptyFor_statementContext(&p.For_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*For_statementContext))

	return p
}

func (s *ForInContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInContext) FOR() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserFOR, 0)
}

func (s *ForInContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(VLangGrammarParserIDENTIFIER)
}

func (s *ForInContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserIDENTIFIER, i)
}

func (s *ForInContext) IN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserIN, 0)
}

func (s *ForInContext) Expression() IExpressionContext {
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

func (s *ForInContext) Block_statement() IBlock_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_statementContext)
}

func (s *ForInContext) COMMA() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserCOMMA, 0)
}

func (s *ForInContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterForIn(s)
	}
}

func (s *ForInContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitForIn(s)
	}
}

func (s *ForInContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitForIn(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForLoopContext struct {
	For_statementContext
}

func NewForLoopContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForLoopContext {
	var p = new(ForLoopContext)

	InitEmptyFor_statementContext(&p.For_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*For_statementContext))

	return p
}

func (s *ForLoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForLoopContext) FOR() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserFOR, 0)
}

func (s *ForLoopContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserIDENTIFIER, 0)
}

func (s *ForLoopContext) DECLARE_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserDECLARE_ASSIGN, 0)
}

func (s *ForLoopContext) AllExpression() []IExpressionContext {
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

func (s *ForLoopContext) Expression(i int) IExpressionContext {
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

func (s *ForLoopContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(VLangGrammarParserSEMICOLON)
}

func (s *ForLoopContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserSEMICOLON, i)
}

func (s *ForLoopContext) Assignment_statement() IAssignment_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignment_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignment_statementContext)
}

func (s *ForLoopContext) Block_statement() IBlock_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_statementContext)
}

func (s *ForLoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterForLoop(s)
	}
}

func (s *ForLoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitForLoop(s)
	}
}

func (s *ForLoopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitForLoop(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForConditionContext struct {
	For_statementContext
}

func NewForConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForConditionContext {
	var p = new(ForConditionContext)

	InitEmptyFor_statementContext(&p.For_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*For_statementContext))

	return p
}

func (s *ForConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForConditionContext) FOR() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserFOR, 0)
}

func (s *ForConditionContext) Expression() IExpressionContext {
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

func (s *ForConditionContext) Block_statement() IBlock_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_statementContext)
}

func (s *ForConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterForCondition(s)
	}
}

func (s *ForConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitForCondition(s)
	}
}

func (s *ForConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitForCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) For_statement() (localctx IFor_statementContext) {
	localctx = NewFor_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, VLangGrammarParserRULE_for_statement)
	var _la int

	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(217)
			p.Match(VLangGrammarParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)
			p.expression(0)
		}
		{
			p.SetState(219)
			p.Block_statement()
		}

	case 2:
		localctx = NewForLoopContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(221)
			p.Match(VLangGrammarParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(222)
			p.Match(VLangGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
			p.Match(VLangGrammarParserDECLARE_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)
			p.expression(0)
		}
		{
			p.SetState(225)
			p.Match(VLangGrammarParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(226)
			p.expression(0)
		}
		{
			p.SetState(227)
			p.Match(VLangGrammarParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(228)
			p.Assignment_statement()
		}
		{
			p.SetState(229)
			p.Block_statement()
		}

	case 3:
		localctx = NewForInContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(231)
			p.Match(VLangGrammarParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(232)
			p.Match(VLangGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == VLangGrammarParserCOMMA {
			{
				p.SetState(233)
				p.Match(VLangGrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(234)
				p.Match(VLangGrammarParserIDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(237)
			p.Match(VLangGrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.expression(0)
		}
		{
			p.SetState(239)
			p.Block_statement()
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

// ISwitch_statementContext is an interface to support dynamic dispatch.
type ISwitch_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SWITCH() antlr.TerminalNode
	Expression() IExpressionContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllCase_clause() []ICase_clauseContext
	Case_clause(i int) ICase_clauseContext
	Default_clause() IDefault_clauseContext

	// IsSwitch_statementContext differentiates from other interfaces.
	IsSwitch_statementContext()
}

type Switch_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_statementContext() *Switch_statementContext {
	var p = new(Switch_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_switch_statement
	return p
}

func InitEmptySwitch_statementContext(p *Switch_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_switch_statement
}

func (*Switch_statementContext) IsSwitch_statementContext() {}

func NewSwitch_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_statementContext {
	var p = new(Switch_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_switch_statement

	return p
}

func (s *Switch_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_statementContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserSWITCH, 0)
}

func (s *Switch_statementContext) Expression() IExpressionContext {
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

func (s *Switch_statementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserLBRACE, 0)
}

func (s *Switch_statementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserRBRACE, 0)
}

func (s *Switch_statementContext) AllCase_clause() []ICase_clauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICase_clauseContext); ok {
			len++
		}
	}

	tst := make([]ICase_clauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICase_clauseContext); ok {
			tst[i] = t.(ICase_clauseContext)
			i++
		}
	}

	return tst
}

func (s *Switch_statementContext) Case_clause(i int) ICase_clauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICase_clauseContext); ok {
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

	return t.(ICase_clauseContext)
}

func (s *Switch_statementContext) Default_clause() IDefault_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefault_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefault_clauseContext)
}

func (s *Switch_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Switch_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterSwitch_statement(s)
	}
}

func (s *Switch_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitSwitch_statement(s)
	}
}

func (s *Switch_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitSwitch_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Switch_statement() (localctx ISwitch_statementContext) {
	localctx = NewSwitch_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, VLangGrammarParserRULE_switch_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(VLangGrammarParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.expression(0)
	}
	{
		p.SetState(245)
		p.Match(VLangGrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VLangGrammarParserCASE {
		{
			p.SetState(246)
			p.Case_clause()
		}

		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VLangGrammarParserDEFAULT {
		{
			p.SetState(252)
			p.Default_clause()
		}

	}
	{
		p.SetState(255)
		p.Match(VLangGrammarParserRBRACE)
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

// ICase_clauseContext is an interface to support dynamic dispatch.
type ICase_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	Expression() IExpressionContext
	COLON() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsCase_clauseContext differentiates from other interfaces.
	IsCase_clauseContext()
}

type Case_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCase_clauseContext() *Case_clauseContext {
	var p = new(Case_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_case_clause
	return p
}

func InitEmptyCase_clauseContext(p *Case_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_case_clause
}

func (*Case_clauseContext) IsCase_clauseContext() {}

func NewCase_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Case_clauseContext {
	var p = new(Case_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_case_clause

	return p
}

func (s *Case_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Case_clauseContext) CASE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserCASE, 0)
}

func (s *Case_clauseContext) Expression() IExpressionContext {
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

func (s *Case_clauseContext) COLON() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserCOLON, 0)
}

func (s *Case_clauseContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *Case_clauseContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *Case_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Case_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Case_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterCase_clause(s)
	}
}

func (s *Case_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitCase_clause(s)
	}
}

func (s *Case_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitCase_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Case_clause() (localctx ICase_clauseContext) {
	localctx = NewCase_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, VLangGrammarParserRULE_case_clause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		p.Match(VLangGrammarParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
		p.expression(0)
	}
	{
		p.SetState(259)
		p.Match(VLangGrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144121097956646098) != 0) || ((int64((_la-69)) & ^0x3f) == 0 && ((int64(1)<<(_la-69))&4095) != 0) {
		{
			p.SetState(260)
			p.Statement()
		}

		p.SetState(265)
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

// IDefault_clauseContext is an interface to support dynamic dispatch.
type IDefault_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFAULT() antlr.TerminalNode
	COLON() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsDefault_clauseContext differentiates from other interfaces.
	IsDefault_clauseContext()
}

type Default_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefault_clauseContext() *Default_clauseContext {
	var p = new(Default_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_default_clause
	return p
}

func InitEmptyDefault_clauseContext(p *Default_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_default_clause
}

func (*Default_clauseContext) IsDefault_clauseContext() {}

func NewDefault_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Default_clauseContext {
	var p = new(Default_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_default_clause

	return p
}

func (s *Default_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Default_clauseContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserDEFAULT, 0)
}

func (s *Default_clauseContext) COLON() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserCOLON, 0)
}

func (s *Default_clauseContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *Default_clauseContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *Default_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Default_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Default_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterDefault_clause(s)
	}
}

func (s *Default_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitDefault_clause(s)
	}
}

func (s *Default_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitDefault_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Default_clause() (localctx IDefault_clauseContext) {
	localctx = NewDefault_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, VLangGrammarParserRULE_default_clause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Match(VLangGrammarParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(267)
		p.Match(VLangGrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144121097956646098) != 0) || ((int64((_la-69)) & ^0x3f) == 0 && ((int64(1)<<(_la-69))&4095) != 0) {
		{
			p.SetState(268)
			p.Statement()
		}

		p.SetState(273)
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

// IBreak_statementContext is an interface to support dynamic dispatch.
type IBreak_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BREAK() antlr.TerminalNode

	// IsBreak_statementContext differentiates from other interfaces.
	IsBreak_statementContext()
}

type Break_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreak_statementContext() *Break_statementContext {
	var p = new(Break_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_break_statement
	return p
}

func InitEmptyBreak_statementContext(p *Break_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_break_statement
}

func (*Break_statementContext) IsBreak_statementContext() {}

func NewBreak_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Break_statementContext {
	var p = new(Break_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_break_statement

	return p
}

func (s *Break_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Break_statementContext) BREAK() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserBREAK, 0)
}

func (s *Break_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Break_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Break_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterBreak_statement(s)
	}
}

func (s *Break_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitBreak_statement(s)
	}
}

func (s *Break_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitBreak_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Break_statement() (localctx IBreak_statementContext) {
	localctx = NewBreak_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, VLangGrammarParserRULE_break_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(274)
		p.Match(VLangGrammarParserBREAK)
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

// IContinue_statementContext is an interface to support dynamic dispatch.
type IContinue_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUE() antlr.TerminalNode

	// IsContinue_statementContext differentiates from other interfaces.
	IsContinue_statementContext()
}

type Continue_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinue_statementContext() *Continue_statementContext {
	var p = new(Continue_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_continue_statement
	return p
}

func InitEmptyContinue_statementContext(p *Continue_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_continue_statement
}

func (*Continue_statementContext) IsContinue_statementContext() {}

func NewContinue_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Continue_statementContext {
	var p = new(Continue_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_continue_statement

	return p
}

func (s *Continue_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Continue_statementContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserCONTINUE, 0)
}

func (s *Continue_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Continue_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Continue_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterContinue_statement(s)
	}
}

func (s *Continue_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitContinue_statement(s)
	}
}

func (s *Continue_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitContinue_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Continue_statement() (localctx IContinue_statementContext) {
	localctx = NewContinue_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, VLangGrammarParserRULE_continue_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.Match(VLangGrammarParserCONTINUE)
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

// IReturn_statementContext is an interface to support dynamic dispatch.
type IReturn_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsReturn_statementContext differentiates from other interfaces.
	IsReturn_statementContext()
}

type Return_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturn_statementContext() *Return_statementContext {
	var p = new(Return_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_return_statement
	return p
}

func InitEmptyReturn_statementContext(p *Return_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_return_statement
}

func (*Return_statementContext) IsReturn_statementContext() {}

func NewReturn_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_statementContext {
	var p = new(Return_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_return_statement

	return p
}

func (s *Return_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_statementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserRETURN, 0)
}

func (s *Return_statementContext) Expression() IExpressionContext {
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

func (s *Return_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterReturn_statement(s)
	}
}

func (s *Return_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitReturn_statement(s)
	}
}

func (s *Return_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitReturn_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Return_statement() (localctx IReturn_statementContext) {
	localctx = NewReturn_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, VLangGrammarParserRULE_return_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(278)
		p.Match(VLangGrammarParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(279)
			p.expression(0)
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

// IExpression_statementContext is an interface to support dynamic dispatch.
type IExpression_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsExpression_statementContext differentiates from other interfaces.
	IsExpression_statementContext()
}

type Expression_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_statementContext() *Expression_statementContext {
	var p = new(Expression_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_expression_statement
	return p
}

func InitEmptyExpression_statementContext(p *Expression_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_expression_statement
}

func (*Expression_statementContext) IsExpression_statementContext() {}

func NewExpression_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_statementContext {
	var p = new(Expression_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_expression_statement

	return p
}

func (s *Expression_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_statementContext) Expression() IExpressionContext {
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

func (s *Expression_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterExpression_statement(s)
	}
}

func (s *Expression_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitExpression_statement(s)
	}
}

func (s *Expression_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitExpression_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Expression_statement() (localctx IExpression_statementContext) {
	localctx = NewExpression_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, VLangGrammarParserRULE_expression_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.expression(0)
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

// IBlock_statementContext is an interface to support dynamic dispatch.
type IBlock_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlock_statementContext differentiates from other interfaces.
	IsBlock_statementContext()
}

type Block_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlock_statementContext() *Block_statementContext {
	var p = new(Block_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_block_statement
	return p
}

func InitEmptyBlock_statementContext(p *Block_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_block_statement
}

func (*Block_statementContext) IsBlock_statementContext() {}

func NewBlock_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_statementContext {
	var p = new(Block_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_block_statement

	return p
}

func (s *Block_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_statementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserLBRACE, 0)
}

func (s *Block_statementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserRBRACE, 0)
}

func (s *Block_statementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *Block_statementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *Block_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterBlock_statement(s)
	}
}

func (s *Block_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitBlock_statement(s)
	}
}

func (s *Block_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitBlock_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Block_statement() (localctx IBlock_statementContext) {
	localctx = NewBlock_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, VLangGrammarParserRULE_block_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.Match(VLangGrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144121097956646098) != 0) || ((int64((_la-69)) & ^0x3f) == 0 && ((int64(1)<<(_la-69))&4095) != 0) {
		{
			p.SetState(285)
			p.Statement()
		}

		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(291)
		p.Match(VLangGrammarParserRBRACE)
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
	p.RuleIndex = VLangGrammarParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_expression

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

type UnaryNotContext struct {
	ExpressionContext
}

func NewUnaryNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryNotContext {
	var p = new(UnaryNotContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryNotContext) NOT() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserNOT, 0)
}

func (s *UnaryNotContext) Expression() IExpressionContext {
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

func (s *UnaryNotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterUnaryNot(s)
	}
}

func (s *UnaryNotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitUnaryNot(s)
	}
}

func (s *UnaryNotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitUnaryNot(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulDivModContext struct {
	ExpressionContext
	op antlr.Token
}

func NewMulDivModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivModContext {
	var p = new(MulDivModContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MulDivModContext) GetOp() antlr.Token { return s.op }

func (s *MulDivModContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivModContext) AllExpression() []IExpressionContext {
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

func (s *MulDivModContext) Expression(i int) IExpressionContext {
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

func (s *MulDivModContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserMULTIPLY, 0)
}

func (s *MulDivModContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserDIVIDE, 0)
}

func (s *MulDivModContext) MODULO() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserMODULO, 0)
}

func (s *MulDivModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterMulDivMod(s)
	}
}

func (s *MulDivModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitMulDivMod(s)
	}
}

func (s *MulDivModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitMulDivMod(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubContext struct {
	ExpressionContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpression() []IExpressionContext {
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

func (s *AddSubContext) Expression(i int) IExpressionContext {
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

func (s *AddSubContext) PLUS() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserPLUS, 0)
}

func (s *AddSubContext) MINUS() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserMINUS, 0)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (s *AddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type IndexAccessContext struct {
	ExpressionContext
}

func NewIndexAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexAccessContext {
	var p = new(IndexAccessContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IndexAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexAccessContext) AllExpression() []IExpressionContext {
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

func (s *IndexAccessContext) Expression(i int) IExpressionContext {
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

func (s *IndexAccessContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserLBRACKET, 0)
}

func (s *IndexAccessContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserRBRACKET, 0)
}

func (s *IndexAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterIndexAccess(s)
	}
}

func (s *IndexAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitIndexAccess(s)
	}
}

func (s *IndexAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitIndexAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryContext struct {
	ExpressionContext
}

func NewPrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryContext {
	var p = new(PrimaryContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) Primary_expression() IPrimary_expressionContext {
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

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicalAndContext struct {
	ExpressionContext
}

func NewLogicalAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalAndContext {
	var p = new(LogicalAndContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LogicalAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndContext) AllExpression() []IExpressionContext {
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

func (s *LogicalAndContext) Expression(i int) IExpressionContext {
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

func (s *LogicalAndContext) AND() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserAND, 0)
}

func (s *LogicalAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterLogicalAnd(s)
	}
}

func (s *LogicalAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitLogicalAnd(s)
	}
}

func (s *LogicalAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitLogicalAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type RelationalContext struct {
	ExpressionContext
	op antlr.Token
}

func NewRelationalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationalContext {
	var p = new(RelationalContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *RelationalContext) GetOp() antlr.Token { return s.op }

func (s *RelationalContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelationalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalContext) AllExpression() []IExpressionContext {
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

func (s *RelationalContext) Expression(i int) IExpressionContext {
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

func (s *RelationalContext) LESS_THAN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserLESS_THAN, 0)
}

func (s *RelationalContext) LESS_EQUAL() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserLESS_EQUAL, 0)
}

func (s *RelationalContext) GREATER_THAN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserGREATER_THAN, 0)
}

func (s *RelationalContext) GREATER_EQUAL() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserGREATER_EQUAL, 0)
}

func (s *RelationalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterRelational(s)
	}
}

func (s *RelationalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitRelational(s)
	}
}

func (s *RelationalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitRelational(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryMinusContext struct {
	ExpressionContext
}

func NewUnaryMinusContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryMinusContext {
	var p = new(UnaryMinusContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryMinusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMinusContext) MINUS() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserMINUS, 0)
}

func (s *UnaryMinusContext) Expression() IExpressionContext {
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

func (s *UnaryMinusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterUnaryMinus(s)
	}
}

func (s *UnaryMinusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitUnaryMinus(s)
	}
}

func (s *UnaryMinusContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitUnaryMinus(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualityContext struct {
	ExpressionContext
	op antlr.Token
}

func NewEqualityContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityContext {
	var p = new(EqualityContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *EqualityContext) GetOp() antlr.Token { return s.op }

func (s *EqualityContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqualityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityContext) AllExpression() []IExpressionContext {
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

func (s *EqualityContext) Expression(i int) IExpressionContext {
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

func (s *EqualityContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserEQUAL, 0)
}

func (s *EqualityContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserNOT_EQUAL, 0)
}

func (s *EqualityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterEquality(s)
	}
}

func (s *EqualityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitEquality(s)
	}
}

func (s *EqualityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitEquality(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallContext struct {
	ExpressionContext
}

func NewFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallContext {
	var p = new(FunctionCallContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) Expression() IExpressionContext {
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

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserRPAREN, 0)
}

func (s *FunctionCallContext) Argument_list() IArgument_listContext {
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

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type FieldAccessContext struct {
	ExpressionContext
}

func NewFieldAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldAccessContext {
	var p = new(FieldAccessContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FieldAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldAccessContext) Expression() IExpressionContext {
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

func (s *FieldAccessContext) DOT() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserDOT, 0)
}

func (s *FieldAccessContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserIDENTIFIER, 0)
}

func (s *FieldAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterFieldAccess(s)
	}
}

func (s *FieldAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitFieldAccess(s)
	}
}

func (s *FieldAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitFieldAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicalOrContext struct {
	ExpressionContext
}

func NewLogicalOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalOrContext {
	var p = new(LogicalOrContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LogicalOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrContext) AllExpression() []IExpressionContext {
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

func (s *LogicalOrContext) Expression(i int) IExpressionContext {
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

func (s *LogicalOrContext) OR() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserOR, 0)
}

func (s *LogicalOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterLogicalOr(s)
	}
}

func (s *LogicalOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitLogicalOr(s)
	}
}

func (s *LogicalOrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitLogicalOr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *VLangGrammarParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 46
	p.EnterRecursionRule(localctx, 46, VLangGrammarParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VLangGrammarParserNIL, VLangGrammarParserTRUE, VLangGrammarParserFALSE, VLangGrammarParserLPAREN, VLangGrammarParserLBRACKET, VLangGrammarParserIDENTIFIER, VLangGrammarParserINT_LITERAL, VLangGrammarParserFLOAT_LITERAL, VLangGrammarParserSTRING_LITERAL, VLangGrammarParserRUNE_LITERAL, VLangGrammarParserPRINT, VLangGrammarParserATOI, VLangGrammarParserPARSE_FLOAT, VLangGrammarParserTYPE_OF, VLangGrammarParserINDEX_OF, VLangGrammarParserJOIN, VLangGrammarParserLEN, VLangGrammarParserAPPEND:
		localctx = NewPrimaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(294)
			p.Primary_expression()
		}

	case VLangGrammarParserMINUS:
		localctx = NewUnaryMinusContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(295)
			p.Match(VLangGrammarParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(296)
			p.expression(8)
		}

	case VLangGrammarParserNOT:
		localctx = NewUnaryNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(297)
			p.Match(VLangGrammarParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(298)
			p.expression(7)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(333)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivModContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VLangGrammarParserRULE_expression)
				p.SetState(301)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(302)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivModContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8070450532247928832) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivModContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(303)
					p.expression(7)
				}

			case 2:
				localctx = NewAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VLangGrammarParserRULE_expression)
				p.SetState(304)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(305)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VLangGrammarParserPLUS || _la == VLangGrammarParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(306)
					p.expression(6)
				}

			case 3:
				localctx = NewRelationalContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VLangGrammarParserRULE_expression)
				p.SetState(307)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(308)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelationalContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-63)) & ^0x3f) == 0 && ((int64(1)<<(_la-63))&15) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelationalContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(309)
					p.expression(5)
				}

			case 4:
				localctx = NewEqualityContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VLangGrammarParserRULE_expression)
				p.SetState(310)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(311)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqualityContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VLangGrammarParserEQUAL || _la == VLangGrammarParserNOT_EQUAL) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqualityContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(312)
					p.expression(4)
				}

			case 5:
				localctx = NewLogicalAndContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VLangGrammarParserRULE_expression)
				p.SetState(313)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(314)
					p.Match(VLangGrammarParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(315)
					p.expression(3)
				}

			case 6:
				localctx = NewLogicalOrContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VLangGrammarParserRULE_expression)
				p.SetState(316)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(317)
					p.Match(VLangGrammarParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(318)
					p.expression(2)
				}

			case 7:
				localctx = NewFieldAccessContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VLangGrammarParserRULE_expression)
				p.SetState(319)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(320)
					p.Match(VLangGrammarParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(321)
					p.Match(VLangGrammarParserIDENTIFIER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 8:
				localctx = NewIndexAccessContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VLangGrammarParserRULE_expression)
				p.SetState(322)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(323)
					p.Match(VLangGrammarParserLBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(324)
					p.expression(0)
				}
				{
					p.SetState(325)
					p.Match(VLangGrammarParserRBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 9:
				localctx = NewFunctionCallContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VLangGrammarParserRULE_expression)
				p.SetState(327)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(328)
					p.Match(VLangGrammarParserLPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(330)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144119998445010944) != 0) || ((int64((_la-69)) & ^0x3f) == 0 && ((int64(1)<<(_la-69))&4095) != 0) {
					{
						p.SetState(329)
						p.Argument_list()
					}

				}
				{
					p.SetState(332)
					p.Match(VLangGrammarParserRPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(337)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
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
	p.RuleIndex = VLangGrammarParserRULE_primary_expression
	return p
}

func InitEmptyPrimary_expressionContext(p *Primary_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_primary_expression
}

func (*Primary_expressionContext) IsPrimary_expressionContext() {}

func NewPrimary_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Primary_expressionContext {
	var p = new(Primary_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_primary_expression

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

type SliceLiteralExprContext struct {
	Primary_expressionContext
}

func NewSliceLiteralExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceLiteralExprContext {
	var p = new(SliceLiteralExprContext)

	InitEmptyPrimary_expressionContext(&p.Primary_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Primary_expressionContext))

	return p
}

func (s *SliceLiteralExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceLiteralExprContext) Slice_literal() ISlice_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISlice_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISlice_literalContext)
}

func (s *SliceLiteralExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterSliceLiteralExpr(s)
	}
}

func (s *SliceLiteralExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitSliceLiteralExpr(s)
	}
}

func (s *SliceLiteralExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitSliceLiteralExpr(s)

	default:
		return t.VisitChildren(s)
	}
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
	return s.GetToken(VLangGrammarParserSTRING_LITERAL, 0)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
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
	return s.GetToken(VLangGrammarParserTRUE, 0)
}

func (s *TrueLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterTrueLiteral(s)
	}
}

func (s *TrueLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitTrueLiteral(s)
	}
}

func (s *TrueLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
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
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterBuiltinCall(s)
	}
}

func (s *BuiltinCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitBuiltinCall(s)
	}
}

func (s *BuiltinCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
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

func (s *IdentifierExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserIDENTIFIER, 0)
}

func (s *IdentifierExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterIdentifierExpr(s)
	}
}

func (s *IdentifierExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitIdentifierExpr(s)
	}
}

func (s *IdentifierExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
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
	return s.GetToken(VLangGrammarParserFLOAT_LITERAL, 0)
}

func (s *FloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
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
	return s.GetToken(VLangGrammarParserNIL, 0)
}

func (s *NilLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterNilLiteral(s)
	}
}

func (s *NilLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitNilLiteral(s)
	}
}

func (s *NilLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
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
	return s.GetToken(VLangGrammarParserINT_LITERAL, 0)
}

func (s *IntLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterIntLiteral(s)
	}
}

func (s *IntLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitIntLiteral(s)
	}
}

func (s *IntLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitIntLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructLiteralExprContext struct {
	Primary_expressionContext
}

func NewStructLiteralExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructLiteralExprContext {
	var p = new(StructLiteralExprContext)

	InitEmptyPrimary_expressionContext(&p.Primary_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Primary_expressionContext))

	return p
}

func (s *StructLiteralExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructLiteralExprContext) Struct_literal() IStruct_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_literalContext)
}

func (s *StructLiteralExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterStructLiteralExpr(s)
	}
}

func (s *StructLiteralExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitStructLiteralExpr(s)
	}
}

func (s *StructLiteralExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitStructLiteralExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type RuneLiteralContext struct {
	Primary_expressionContext
}

func NewRuneLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RuneLiteralContext {
	var p = new(RuneLiteralContext)

	InitEmptyPrimary_expressionContext(&p.Primary_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Primary_expressionContext))

	return p
}

func (s *RuneLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuneLiteralContext) RUNE_LITERAL() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserRUNE_LITERAL, 0)
}

func (s *RuneLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterRuneLiteral(s)
	}
}

func (s *RuneLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitRuneLiteral(s)
	}
}

func (s *RuneLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitRuneLiteral(s)

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
	return s.GetToken(VLangGrammarParserLPAREN, 0)
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
	return s.GetToken(VLangGrammarParserRPAREN, 0)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
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
	return s.GetToken(VLangGrammarParserFALSE, 0)
}

func (s *FalseLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterFalseLiteral(s)
	}
}

func (s *FalseLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitFalseLiteral(s)
	}
}

func (s *FalseLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitFalseLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Primary_expression() (localctx IPrimary_expressionContext) {
	localctx = NewPrimary_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, VLangGrammarParserRULE_primary_expression)
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIdentifierExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(338)
			p.Match(VLangGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewIntLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(339)
			p.Match(VLangGrammarParserINT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewFloatLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(340)
			p.Match(VLangGrammarParserFLOAT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(341)
			p.Match(VLangGrammarParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewRuneLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(342)
			p.Match(VLangGrammarParserRUNE_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewTrueLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(343)
			p.Match(VLangGrammarParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewFalseLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(344)
			p.Match(VLangGrammarParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewNilLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(345)
			p.Match(VLangGrammarParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewSliceLiteralExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(346)
			p.Slice_literal()
		}

	case 10:
		localctx = NewStructLiteralExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(347)
			p.Struct_literal()
		}

	case 11:
		localctx = NewBuiltinCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(348)
			p.Builtin_function_call()
		}

	case 12:
		localctx = NewParenExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(349)
			p.Match(VLangGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(350)
			p.expression(0)
		}
		{
			p.SetState(351)
			p.Match(VLangGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// ISlice_literalContext is an interface to support dynamic dispatch.
type ISlice_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	Type_annotation() IType_annotationContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	Expression_list() IExpression_listContext

	// IsSlice_literalContext differentiates from other interfaces.
	IsSlice_literalContext()
}

type Slice_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySlice_literalContext() *Slice_literalContext {
	var p = new(Slice_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_slice_literal
	return p
}

func InitEmptySlice_literalContext(p *Slice_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_slice_literal
}

func (*Slice_literalContext) IsSlice_literalContext() {}

func NewSlice_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Slice_literalContext {
	var p = new(Slice_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_slice_literal

	return p
}

func (s *Slice_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Slice_literalContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserLBRACKET, 0)
}

func (s *Slice_literalContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserRBRACKET, 0)
}

func (s *Slice_literalContext) Type_annotation() IType_annotationContext {
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

func (s *Slice_literalContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserLBRACE, 0)
}

func (s *Slice_literalContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserRBRACE, 0)
}

func (s *Slice_literalContext) Expression_list() IExpression_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpression_listContext)
}

func (s *Slice_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Slice_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Slice_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterSlice_literal(s)
	}
}

func (s *Slice_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitSlice_literal(s)
	}
}

func (s *Slice_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitSlice_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Slice_literal() (localctx ISlice_literalContext) {
	localctx = NewSlice_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, VLangGrammarParserRULE_slice_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.Match(VLangGrammarParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(356)
		p.Match(VLangGrammarParserRBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(357)
		p.Type_annotation()
	}
	{
		p.SetState(358)
		p.Match(VLangGrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144119998445010944) != 0) || ((int64((_la-69)) & ^0x3f) == 0 && ((int64(1)<<(_la-69))&4095) != 0) {
		{
			p.SetState(359)
			p.Expression_list()
		}

	}
	{
		p.SetState(362)
		p.Match(VLangGrammarParserRBRACE)
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

// IStruct_literalContext is an interface to support dynamic dispatch.
type IStruct_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	Field_assignment_list() IField_assignment_listContext

	// IsStruct_literalContext differentiates from other interfaces.
	IsStruct_literalContext()
}

type Struct_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_literalContext() *Struct_literalContext {
	var p = new(Struct_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_struct_literal
	return p
}

func InitEmptyStruct_literalContext(p *Struct_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_struct_literal
}

func (*Struct_literalContext) IsStruct_literalContext() {}

func NewStruct_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_literalContext {
	var p = new(Struct_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_struct_literal

	return p
}

func (s *Struct_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_literalContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserIDENTIFIER, 0)
}

func (s *Struct_literalContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserLBRACE, 0)
}

func (s *Struct_literalContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserRBRACE, 0)
}

func (s *Struct_literalContext) Field_assignment_list() IField_assignment_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_assignment_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField_assignment_listContext)
}

func (s *Struct_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Struct_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterStruct_literal(s)
	}
}

func (s *Struct_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitStruct_literal(s)
	}
}

func (s *Struct_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitStruct_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Struct_literal() (localctx IStruct_literalContext) {
	localctx = NewStruct_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, VLangGrammarParserRULE_struct_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(364)
		p.Match(VLangGrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(365)
		p.Match(VLangGrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VLangGrammarParserIDENTIFIER {
		{
			p.SetState(366)
			p.Field_assignment_list()
		}

	}
	{
		p.SetState(369)
		p.Match(VLangGrammarParserRBRACE)
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

// IField_assignment_listContext is an interface to support dynamic dispatch.
type IField_assignment_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllField_assignment() []IField_assignmentContext
	Field_assignment(i int) IField_assignmentContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsField_assignment_listContext differentiates from other interfaces.
	IsField_assignment_listContext()
}

type Field_assignment_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_assignment_listContext() *Field_assignment_listContext {
	var p = new(Field_assignment_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_field_assignment_list
	return p
}

func InitEmptyField_assignment_listContext(p *Field_assignment_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_field_assignment_list
}

func (*Field_assignment_listContext) IsField_assignment_listContext() {}

func NewField_assignment_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_assignment_listContext {
	var p = new(Field_assignment_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_field_assignment_list

	return p
}

func (s *Field_assignment_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_assignment_listContext) AllField_assignment() []IField_assignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IField_assignmentContext); ok {
			len++
		}
	}

	tst := make([]IField_assignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IField_assignmentContext); ok {
			tst[i] = t.(IField_assignmentContext)
			i++
		}
	}

	return tst
}

func (s *Field_assignment_listContext) Field_assignment(i int) IField_assignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_assignmentContext); ok {
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

	return t.(IField_assignmentContext)
}

func (s *Field_assignment_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VLangGrammarParserCOMMA)
}

func (s *Field_assignment_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserCOMMA, i)
}

func (s *Field_assignment_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_assignment_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_assignment_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterField_assignment_list(s)
	}
}

func (s *Field_assignment_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitField_assignment_list(s)
	}
}

func (s *Field_assignment_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitField_assignment_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Field_assignment_list() (localctx IField_assignment_listContext) {
	localctx = NewField_assignment_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, VLangGrammarParserRULE_field_assignment_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(371)
		p.Field_assignment()
	}
	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VLangGrammarParserCOMMA {
		{
			p.SetState(372)
			p.Match(VLangGrammarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)
			p.Field_assignment()
		}

		p.SetState(378)
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

// IField_assignmentContext is an interface to support dynamic dispatch.
type IField_assignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Expression() IExpressionContext

	// IsField_assignmentContext differentiates from other interfaces.
	IsField_assignmentContext()
}

type Field_assignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_assignmentContext() *Field_assignmentContext {
	var p = new(Field_assignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_field_assignment
	return p
}

func InitEmptyField_assignmentContext(p *Field_assignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_field_assignment
}

func (*Field_assignmentContext) IsField_assignmentContext() {}

func NewField_assignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_assignmentContext {
	var p = new(Field_assignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_field_assignment

	return p
}

func (s *Field_assignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_assignmentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserIDENTIFIER, 0)
}

func (s *Field_assignmentContext) COLON() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserCOLON, 0)
}

func (s *Field_assignmentContext) Expression() IExpressionContext {
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

func (s *Field_assignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_assignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_assignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterField_assignment(s)
	}
}

func (s *Field_assignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitField_assignment(s)
	}
}

func (s *Field_assignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitField_assignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Field_assignment() (localctx IField_assignmentContext) {
	localctx = NewField_assignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, VLangGrammarParserRULE_field_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(379)
		p.Match(VLangGrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(380)
		p.Match(VLangGrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(381)
		p.expression(0)
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
	p.RuleIndex = VLangGrammarParserRULE_builtin_function_call
	return p
}

func InitEmptyBuiltin_function_callContext(p *Builtin_function_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_builtin_function_call
}

func (*Builtin_function_callContext) IsBuiltin_function_callContext() {}

func NewBuiltin_function_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Builtin_function_callContext {
	var p = new(Builtin_function_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_builtin_function_call

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

type ParseFloatCallContext struct {
	Builtin_function_callContext
}

func NewParseFloatCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParseFloatCallContext {
	var p = new(ParseFloatCallContext)

	InitEmptyBuiltin_function_callContext(&p.Builtin_function_callContext)
	p.parser = parser
	p.CopyAll(ctx.(*Builtin_function_callContext))

	return p
}

func (s *ParseFloatCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseFloatCallContext) PARSE_FLOAT() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserPARSE_FLOAT, 0)
}

func (s *ParseFloatCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserLPAREN, 0)
}

func (s *ParseFloatCallContext) Expression() IExpressionContext {
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

func (s *ParseFloatCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserRPAREN, 0)
}

func (s *ParseFloatCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterParseFloatCall(s)
	}
}

func (s *ParseFloatCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitParseFloatCall(s)
	}
}

func (s *ParseFloatCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitParseFloatCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type LenCallContext struct {
	Builtin_function_callContext
}

func NewLenCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LenCallContext {
	var p = new(LenCallContext)

	InitEmptyBuiltin_function_callContext(&p.Builtin_function_callContext)
	p.parser = parser
	p.CopyAll(ctx.(*Builtin_function_callContext))

	return p
}

func (s *LenCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LenCallContext) LEN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserLEN, 0)
}

func (s *LenCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserLPAREN, 0)
}

func (s *LenCallContext) Expression() IExpressionContext {
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

func (s *LenCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserRPAREN, 0)
}

func (s *LenCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterLenCall(s)
	}
}

func (s *LenCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitLenCall(s)
	}
}

func (s *LenCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitLenCall(s)

	default:
		return t.VisitChildren(s)
	}
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
	return s.GetToken(VLangGrammarParserPRINT, 0)
}

func (s *PrintCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserLPAREN, 0)
}

func (s *PrintCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserRPAREN, 0)
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
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterPrintCall(s)
	}
}

func (s *PrintCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitPrintCall(s)
	}
}

func (s *PrintCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitPrintCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type IndexOfCallContext struct {
	Builtin_function_callContext
}

func NewIndexOfCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexOfCallContext {
	var p = new(IndexOfCallContext)

	InitEmptyBuiltin_function_callContext(&p.Builtin_function_callContext)
	p.parser = parser
	p.CopyAll(ctx.(*Builtin_function_callContext))

	return p
}

func (s *IndexOfCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexOfCallContext) INDEX_OF() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserINDEX_OF, 0)
}

func (s *IndexOfCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserLPAREN, 0)
}

func (s *IndexOfCallContext) AllExpression() []IExpressionContext {
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

func (s *IndexOfCallContext) Expression(i int) IExpressionContext {
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

func (s *IndexOfCallContext) COMMA() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserCOMMA, 0)
}

func (s *IndexOfCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserRPAREN, 0)
}

func (s *IndexOfCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterIndexOfCall(s)
	}
}

func (s *IndexOfCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitIndexOfCall(s)
	}
}

func (s *IndexOfCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitIndexOfCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeOfCallContext struct {
	Builtin_function_callContext
}

func NewTypeOfCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeOfCallContext {
	var p = new(TypeOfCallContext)

	InitEmptyBuiltin_function_callContext(&p.Builtin_function_callContext)
	p.parser = parser
	p.CopyAll(ctx.(*Builtin_function_callContext))

	return p
}

func (s *TypeOfCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeOfCallContext) TYPE_OF() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserTYPE_OF, 0)
}

func (s *TypeOfCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserLPAREN, 0)
}

func (s *TypeOfCallContext) Expression() IExpressionContext {
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

func (s *TypeOfCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserRPAREN, 0)
}

func (s *TypeOfCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterTypeOfCall(s)
	}
}

func (s *TypeOfCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitTypeOfCall(s)
	}
}

func (s *TypeOfCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitTypeOfCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type AtoiCallContext struct {
	Builtin_function_callContext
}

func NewAtoiCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtoiCallContext {
	var p = new(AtoiCallContext)

	InitEmptyBuiltin_function_callContext(&p.Builtin_function_callContext)
	p.parser = parser
	p.CopyAll(ctx.(*Builtin_function_callContext))

	return p
}

func (s *AtoiCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtoiCallContext) ATOI() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserATOI, 0)
}

func (s *AtoiCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserLPAREN, 0)
}

func (s *AtoiCallContext) Expression() IExpressionContext {
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

func (s *AtoiCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserRPAREN, 0)
}

func (s *AtoiCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterAtoiCall(s)
	}
}

func (s *AtoiCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitAtoiCall(s)
	}
}

func (s *AtoiCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitAtoiCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type JoinCallContext struct {
	Builtin_function_callContext
}

func NewJoinCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JoinCallContext {
	var p = new(JoinCallContext)

	InitEmptyBuiltin_function_callContext(&p.Builtin_function_callContext)
	p.parser = parser
	p.CopyAll(ctx.(*Builtin_function_callContext))

	return p
}

func (s *JoinCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinCallContext) JOIN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserJOIN, 0)
}

func (s *JoinCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserLPAREN, 0)
}

func (s *JoinCallContext) AllExpression() []IExpressionContext {
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

func (s *JoinCallContext) Expression(i int) IExpressionContext {
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

func (s *JoinCallContext) COMMA() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserCOMMA, 0)
}

func (s *JoinCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserRPAREN, 0)
}

func (s *JoinCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterJoinCall(s)
	}
}

func (s *JoinCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitJoinCall(s)
	}
}

func (s *JoinCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitJoinCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type AppendCallContext struct {
	Builtin_function_callContext
}

func NewAppendCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AppendCallContext {
	var p = new(AppendCallContext)

	InitEmptyBuiltin_function_callContext(&p.Builtin_function_callContext)
	p.parser = parser
	p.CopyAll(ctx.(*Builtin_function_callContext))

	return p
}

func (s *AppendCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendCallContext) APPEND() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserAPPEND, 0)
}

func (s *AppendCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserLPAREN, 0)
}

func (s *AppendCallContext) AllExpression() []IExpressionContext {
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

func (s *AppendCallContext) Expression(i int) IExpressionContext {
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

func (s *AppendCallContext) COMMA() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserCOMMA, 0)
}

func (s *AppendCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserRPAREN, 0)
}

func (s *AppendCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterAppendCall(s)
	}
}

func (s *AppendCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitAppendCall(s)
	}
}

func (s *AppendCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitAppendCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Builtin_function_call() (localctx IBuiltin_function_callContext) {
	localctx = NewBuiltin_function_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, VLangGrammarParserRULE_builtin_function_call)
	var _la int

	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VLangGrammarParserPRINT:
		localctx = NewPrintCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(383)
			p.Match(VLangGrammarParserPRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(384)
			p.Match(VLangGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144119998445010944) != 0) || ((int64((_la-69)) & ^0x3f) == 0 && ((int64(1)<<(_la-69))&4095) != 0) {
			{
				p.SetState(385)
				p.Argument_list()
			}

		}
		{
			p.SetState(388)
			p.Match(VLangGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangGrammarParserATOI:
		localctx = NewAtoiCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(389)
			p.Match(VLangGrammarParserATOI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(390)
			p.Match(VLangGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(391)
			p.expression(0)
		}
		{
			p.SetState(392)
			p.Match(VLangGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangGrammarParserPARSE_FLOAT:
		localctx = NewParseFloatCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(394)
			p.Match(VLangGrammarParserPARSE_FLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(395)
			p.Match(VLangGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(396)
			p.expression(0)
		}
		{
			p.SetState(397)
			p.Match(VLangGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangGrammarParserTYPE_OF:
		localctx = NewTypeOfCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(399)
			p.Match(VLangGrammarParserTYPE_OF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(400)
			p.Match(VLangGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(401)
			p.expression(0)
		}
		{
			p.SetState(402)
			p.Match(VLangGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangGrammarParserINDEX_OF:
		localctx = NewIndexOfCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(404)
			p.Match(VLangGrammarParserINDEX_OF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)
			p.Match(VLangGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(406)
			p.expression(0)
		}
		{
			p.SetState(407)
			p.Match(VLangGrammarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(408)
			p.expression(0)
		}
		{
			p.SetState(409)
			p.Match(VLangGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangGrammarParserJOIN:
		localctx = NewJoinCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(411)
			p.Match(VLangGrammarParserJOIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(412)
			p.Match(VLangGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(413)
			p.expression(0)
		}
		{
			p.SetState(414)
			p.Match(VLangGrammarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(415)
			p.expression(0)
		}
		{
			p.SetState(416)
			p.Match(VLangGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangGrammarParserLEN:
		localctx = NewLenCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(418)
			p.Match(VLangGrammarParserLEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(419)
			p.Match(VLangGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(420)
			p.expression(0)
		}
		{
			p.SetState(421)
			p.Match(VLangGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangGrammarParserAPPEND:
		localctx = NewAppendCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(423)
			p.Match(VLangGrammarParserAPPEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(424)
			p.Match(VLangGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(425)
			p.expression(0)
		}
		{
			p.SetState(426)
			p.Match(VLangGrammarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(427)
			p.expression(0)
		}
		{
			p.SetState(428)
			p.Match(VLangGrammarParserRPAREN)
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

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

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
	p.RuleIndex = VLangGrammarParserRULE_argument_list
	return p
}

func InitEmptyArgument_listContext(p *Argument_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_argument_list
}

func (*Argument_listContext) IsArgument_listContext() {}

func NewArgument_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Argument_listContext {
	var p = new(Argument_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_argument_list

	return p
}

func (s *Argument_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Argument_listContext) AllExpression() []IExpressionContext {
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

func (s *Argument_listContext) Expression(i int) IExpressionContext {
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

func (s *Argument_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VLangGrammarParserCOMMA)
}

func (s *Argument_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserCOMMA, i)
}

func (s *Argument_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Argument_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Argument_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterArgument_list(s)
	}
}

func (s *Argument_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitArgument_list(s)
	}
}

func (s *Argument_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitArgument_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Argument_list() (localctx IArgument_listContext) {
	localctx = NewArgument_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, VLangGrammarParserRULE_argument_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(432)
		p.expression(0)
	}
	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VLangGrammarParserCOMMA {
		{
			p.SetState(433)
			p.Match(VLangGrammarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(434)
			p.expression(0)
		}

		p.SetState(439)
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

// IExpression_listContext is an interface to support dynamic dispatch.
type IExpression_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsExpression_listContext differentiates from other interfaces.
	IsExpression_listContext()
}

type Expression_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_listContext() *Expression_listContext {
	var p = new(Expression_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_expression_list
	return p
}

func InitEmptyExpression_listContext(p *Expression_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarParserRULE_expression_list
}

func (*Expression_listContext) IsExpression_listContext() {}

func NewExpression_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_listContext {
	var p = new(Expression_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarParserRULE_expression_list

	return p
}

func (s *Expression_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_listContext) AllExpression() []IExpressionContext {
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

func (s *Expression_listContext) Expression(i int) IExpressionContext {
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

func (s *Expression_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VLangGrammarParserCOMMA)
}

func (s *Expression_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VLangGrammarParserCOMMA, i)
}

func (s *Expression_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterExpression_list(s)
	}
}

func (s *Expression_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitExpression_list(s)
	}
}

func (s *Expression_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitExpression_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammarParser) Expression_list() (localctx IExpression_listContext) {
	localctx = NewExpression_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, VLangGrammarParserRULE_expression_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(440)
		p.expression(0)
	}
	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VLangGrammarParserCOMMA {
		{
			p.SetState(441)
			p.Match(VLangGrammarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(442)
			p.expression(0)
		}

		p.SetState(447)
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

func (p *VLangGrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 23:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *VLangGrammarParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
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
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 9)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
