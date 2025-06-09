// Code generated from grammar/VLangGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // VLangGrammar
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

type VLangGrammar struct {
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
		"", "'mut'", "'func'", "'struct'", "'slice'", "'if'", "'else'", "'switch'",
		"'case'", "'default'", "'for'", "'while'", "'in'", "'break'", "'continue'",
		"'return'", "'int'", "'float64'", "'string'", "'bool'", "'rune'", "'+'",
		"'-'", "'*'", "'/'", "'%'", "'='", "'+='", "'-='", "'=='", "'!='", "'<'",
		"'<='", "'>'", "'>='", "'&&'", "'||'", "'!'", "'('", "')'", "'{'", "'}'",
		"'['", "']'", "';'", "':'", "'.'", "','", "'\\n'", "", "", "", "", "",
		"'nil'",
	}
	staticData.SymbolicNames = []string{
		"", "MUT", "FUNC", "STR", "SLICE_TYPE", "IF_KW", "ELSE_KW", "SWITCH_KW",
		"CASE_KW", "DEFAULT_KW", "FOR_KW", "WHILE_KW", "IN_KW", "BREAK_KW",
		"CONTINUE_KW", "RETURN_KW", "INT_TYPE", "FLOAT_TYPE", "STRING_TYPE",
		"BOOL_TYPE", "RUNE_TYPE", "PLUS", "MINUS", "MULT", "DIV", "MOD", "ASSIGN",
		"PLUS_ASSIGN", "MINUS_ASSIGN", "EQ", "NE", "LT", "LE", "GT", "GE", "AND",
		"OR", "NOT", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK",
		"SEMI", "COLON", "DOT", "COMMA", "NEWLINE", "INT_LITERAL", "FLOAT_LITERAL",
		"CHAR_LITERAL", "STRING_LITERAL", "BOOL_LITERAL", "NIL_LITERAL", "ID",
		"LINE_COMMENT", "BLOCK_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"prog", "delim", "stmt", "decl_stmt", "type_annotation", "var_type",
		"assign_stmt", "id_pattern", "literal", "expression", "vect_expr", "vect_item",
		"vect_prop", "vect_func", "repeating", "type", "vector_type", "matrix_type",
		"aux_matrix_type", "if_stmt", "if_chain", "else_stmt", "switch_stmt",
		"switch_case", "default_case", "while_stmt", "for_stmt", "range", "transfer_stmt",
		"func_call", "arg_list", "func_arg", "func_dcl", "param_list", "func_param",
		"strct_dcl", "struct_prop", "struct_vector",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 58, 471, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 1, 0, 5, 0, 78, 8, 0, 10, 0, 12, 0, 81, 9, 0, 1, 0, 3, 0,
		84, 8, 0, 1, 1, 3, 1, 87, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 3, 2, 109, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 116, 8, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 123, 8, 3, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 137, 8, 6, 1, 7,
		1, 7, 1, 7, 5, 7, 142, 8, 7, 10, 7, 12, 7, 145, 9, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 3, 8, 152, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 169, 8, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 189, 8, 9, 10, 9, 12, 9, 192, 9, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 198, 8, 10, 10, 10, 12, 10, 201, 9,
		10, 3, 10, 203, 8, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		4, 11, 212, 8, 11, 11, 11, 12, 11, 213, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 3, 14, 226, 8, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 3,
		15, 240, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 3, 17, 252, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1,
		19, 1, 19, 5, 19, 261, 8, 19, 10, 19, 12, 19, 264, 9, 19, 1, 19, 3, 19,
		267, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 273, 8, 20, 10, 20, 12,
		20, 276, 9, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 5, 21, 283, 8, 21, 10,
		21, 12, 21, 286, 9, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22,
		294, 8, 22, 10, 22, 12, 22, 297, 9, 22, 1, 22, 3, 22, 300, 8, 22, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 308, 8, 23, 10, 23, 12, 23, 311,
		9, 23, 1, 24, 1, 24, 1, 24, 5, 24, 316, 8, 24, 10, 24, 12, 24, 319, 9,
		24, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25, 325, 8, 25, 10, 25, 12, 25, 328,
		9, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 336, 8, 26, 10,
		26, 12, 26, 339, 9, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 5, 26, 351, 8, 26, 10, 26, 12, 26, 354, 9, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 365,
		8, 26, 1, 26, 1, 26, 5, 26, 369, 8, 26, 10, 26, 12, 26, 372, 9, 26, 1,
		26, 1, 26, 3, 26, 376, 8, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 28, 1, 28, 3, 28, 386, 8, 28, 1, 28, 1, 28, 3, 28, 390, 8, 28, 1, 29,
		1, 29, 1, 29, 3, 29, 395, 8, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 5,
		30, 402, 8, 30, 10, 30, 12, 30, 405, 9, 30, 1, 31, 3, 31, 408, 8, 31, 1,
		31, 1, 31, 3, 31, 412, 8, 31, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 418, 8,
		32, 1, 32, 1, 32, 3, 32, 422, 8, 32, 1, 32, 1, 32, 5, 32, 426, 8, 32, 10,
		32, 12, 32, 429, 9, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 5, 33, 436,
		8, 33, 10, 33, 12, 33, 439, 9, 33, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1,
		35, 1, 35, 5, 35, 448, 8, 35, 10, 35, 12, 35, 451, 9, 35, 1, 35, 1, 35,
		1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 459, 8, 36, 1, 36, 1, 36, 3, 36, 463,
		8, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 0, 1, 18, 38, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
		40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74,
		0, 7, 3, 0, 4, 4, 16, 19, 55, 55, 1, 0, 27, 28, 2, 0, 22, 22, 37, 37, 1,
		0, 23, 25, 1, 0, 21, 22, 1, 0, 31, 34, 1, 0, 29, 30, 502, 0, 79, 1, 0,
		0, 0, 2, 86, 1, 0, 0, 0, 4, 108, 1, 0, 0, 0, 6, 122, 1, 0, 0, 0, 8, 124,
		1, 0, 0, 0, 10, 126, 1, 0, 0, 0, 12, 136, 1, 0, 0, 0, 14, 138, 1, 0, 0,
		0, 16, 151, 1, 0, 0, 0, 18, 168, 1, 0, 0, 0, 20, 193, 1, 0, 0, 0, 22, 206,
		1, 0, 0, 0, 24, 215, 1, 0, 0, 0, 26, 219, 1, 0, 0, 0, 28, 225, 1, 0, 0,
		0, 30, 239, 1, 0, 0, 0, 32, 241, 1, 0, 0, 0, 34, 251, 1, 0, 0, 0, 36, 253,
		1, 0, 0, 0, 38, 257, 1, 0, 0, 0, 40, 268, 1, 0, 0, 0, 42, 279, 1, 0, 0,
		0, 44, 289, 1, 0, 0, 0, 46, 303, 1, 0, 0, 0, 48, 312, 1, 0, 0, 0, 50, 320,
		1, 0, 0, 0, 52, 375, 1, 0, 0, 0, 54, 377, 1, 0, 0, 0, 56, 389, 1, 0, 0,
		0, 58, 391, 1, 0, 0, 0, 60, 398, 1, 0, 0, 0, 62, 407, 1, 0, 0, 0, 64, 413,
		1, 0, 0, 0, 66, 432, 1, 0, 0, 0, 68, 440, 1, 0, 0, 0, 70, 443, 1, 0, 0,
		0, 72, 454, 1, 0, 0, 0, 74, 464, 1, 0, 0, 0, 76, 78, 3, 4, 2, 0, 77, 76,
		1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0,
		80, 83, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 82, 84, 5, 0, 0, 1, 83, 82, 1,
		0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 1, 1, 0, 0, 0, 85, 87, 5, 48, 0, 0, 86,
		85, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 3, 1, 0, 0, 0, 88, 89, 3, 6, 3,
		0, 89, 90, 3, 2, 1, 0, 90, 109, 1, 0, 0, 0, 91, 92, 3, 12, 6, 0, 92, 93,
		3, 2, 1, 0, 93, 109, 1, 0, 0, 0, 94, 95, 3, 56, 28, 0, 95, 96, 3, 2, 1,
		0, 96, 109, 1, 0, 0, 0, 97, 109, 3, 38, 19, 0, 98, 109, 3, 44, 22, 0, 99,
		109, 3, 52, 26, 0, 100, 101, 3, 58, 29, 0, 101, 102, 3, 2, 1, 0, 102, 109,
		1, 0, 0, 0, 103, 104, 3, 26, 13, 0, 104, 105, 3, 2, 1, 0, 105, 109, 1,
		0, 0, 0, 106, 109, 3, 64, 32, 0, 107, 109, 3, 70, 35, 0, 108, 88, 1, 0,
		0, 0, 108, 91, 1, 0, 0, 0, 108, 94, 1, 0, 0, 0, 108, 97, 1, 0, 0, 0, 108,
		98, 1, 0, 0, 0, 108, 99, 1, 0, 0, 0, 108, 100, 1, 0, 0, 0, 108, 103, 1,
		0, 0, 0, 108, 106, 1, 0, 0, 0, 108, 107, 1, 0, 0, 0, 109, 5, 1, 0, 0, 0,
		110, 111, 3, 10, 5, 0, 111, 112, 5, 55, 0, 0, 112, 115, 3, 8, 4, 0, 113,
		114, 5, 26, 0, 0, 114, 116, 3, 18, 9, 0, 115, 113, 1, 0, 0, 0, 115, 116,
		1, 0, 0, 0, 116, 123, 1, 0, 0, 0, 117, 118, 5, 55, 0, 0, 118, 119, 3, 8,
		4, 0, 119, 120, 5, 26, 0, 0, 120, 121, 3, 18, 9, 0, 121, 123, 1, 0, 0,
		0, 122, 110, 1, 0, 0, 0, 122, 117, 1, 0, 0, 0, 123, 7, 1, 0, 0, 0, 124,
		125, 7, 0, 0, 0, 125, 9, 1, 0, 0, 0, 126, 127, 5, 1, 0, 0, 127, 11, 1,
		0, 0, 0, 128, 129, 3, 14, 7, 0, 129, 130, 5, 26, 0, 0, 130, 131, 3, 18,
		9, 0, 131, 137, 1, 0, 0, 0, 132, 133, 3, 14, 7, 0, 133, 134, 7, 1, 0, 0,
		134, 135, 3, 18, 9, 0, 135, 137, 1, 0, 0, 0, 136, 128, 1, 0, 0, 0, 136,
		132, 1, 0, 0, 0, 137, 13, 1, 0, 0, 0, 138, 143, 5, 55, 0, 0, 139, 140,
		5, 46, 0, 0, 140, 142, 5, 55, 0, 0, 141, 139, 1, 0, 0, 0, 142, 145, 1,
		0, 0, 0, 143, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 15, 1, 0, 0,
		0, 145, 143, 1, 0, 0, 0, 146, 152, 5, 49, 0, 0, 147, 152, 5, 50, 0, 0,
		148, 152, 5, 52, 0, 0, 149, 152, 5, 53, 0, 0, 150, 152, 5, 54, 0, 0, 151,
		146, 1, 0, 0, 0, 151, 147, 1, 0, 0, 0, 151, 148, 1, 0, 0, 0, 151, 149,
		1, 0, 0, 0, 151, 150, 1, 0, 0, 0, 152, 17, 1, 0, 0, 0, 153, 154, 6, 9,
		-1, 0, 154, 155, 5, 38, 0, 0, 155, 156, 3, 18, 9, 0, 156, 157, 5, 39, 0,
		0, 157, 169, 1, 0, 0, 0, 158, 169, 3, 58, 29, 0, 159, 169, 3, 14, 7, 0,
		160, 169, 3, 22, 11, 0, 161, 169, 3, 24, 12, 0, 162, 169, 3, 26, 13, 0,
		163, 169, 3, 16, 8, 0, 164, 169, 3, 20, 10, 0, 165, 169, 3, 28, 14, 0,
		166, 167, 7, 2, 0, 0, 167, 169, 3, 18, 9, 7, 168, 153, 1, 0, 0, 0, 168,
		158, 1, 0, 0, 0, 168, 159, 1, 0, 0, 0, 168, 160, 1, 0, 0, 0, 168, 161,
		1, 0, 0, 0, 168, 162, 1, 0, 0, 0, 168, 163, 1, 0, 0, 0, 168, 164, 1, 0,
		0, 0, 168, 165, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 169, 190, 1, 0, 0, 0,
		170, 171, 10, 6, 0, 0, 171, 172, 7, 3, 0, 0, 172, 189, 3, 18, 9, 7, 173,
		174, 10, 5, 0, 0, 174, 175, 7, 4, 0, 0, 175, 189, 3, 18, 9, 6, 176, 177,
		10, 4, 0, 0, 177, 178, 7, 5, 0, 0, 178, 189, 3, 18, 9, 5, 179, 180, 10,
		3, 0, 0, 180, 181, 7, 6, 0, 0, 181, 189, 3, 18, 9, 4, 182, 183, 10, 2,
		0, 0, 183, 184, 5, 35, 0, 0, 184, 189, 3, 18, 9, 3, 185, 186, 10, 1, 0,
		0, 186, 187, 5, 36, 0, 0, 187, 189, 3, 18, 9, 2, 188, 170, 1, 0, 0, 0,
		188, 173, 1, 0, 0, 0, 188, 176, 1, 0, 0, 0, 188, 179, 1, 0, 0, 0, 188,
		182, 1, 0, 0, 0, 188, 185, 1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190, 188,
		1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 19, 1, 0, 0, 0, 192, 190, 1, 0,
		0, 0, 193, 202, 5, 42, 0, 0, 194, 199, 3, 18, 9, 0, 195, 196, 5, 47, 0,
		0, 196, 198, 3, 18, 9, 0, 197, 195, 1, 0, 0, 0, 198, 201, 1, 0, 0, 0, 199,
		197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 203, 1, 0, 0, 0, 201, 199,
		1, 0, 0, 0, 202, 194, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 204, 1, 0,
		0, 0, 204, 205, 5, 43, 0, 0, 205, 21, 1, 0, 0, 0, 206, 211, 3, 14, 7, 0,
		207, 208, 5, 40, 0, 0, 208, 209, 3, 18, 9, 0, 209, 210, 5, 41, 0, 0, 210,
		212, 1, 0, 0, 0, 211, 207, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 211,
		1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 23, 1, 0, 0, 0, 215, 216, 3, 22,
		11, 0, 216, 217, 5, 46, 0, 0, 217, 218, 3, 14, 7, 0, 218, 25, 1, 0, 0,
		0, 219, 220, 3, 22, 11, 0, 220, 221, 5, 46, 0, 0, 221, 222, 3, 58, 29,
		0, 222, 27, 1, 0, 0, 0, 223, 226, 3, 10, 5, 0, 224, 226, 3, 34, 17, 0,
		225, 223, 1, 0, 0, 0, 225, 224, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227,
		228, 5, 38, 0, 0, 228, 229, 5, 55, 0, 0, 229, 230, 5, 45, 0, 0, 230, 231,
		3, 18, 9, 0, 231, 232, 5, 47, 0, 0, 232, 233, 5, 55, 0, 0, 233, 234, 5,
		45, 0, 0, 234, 235, 3, 18, 9, 0, 235, 236, 5, 39, 0, 0, 236, 29, 1, 0,
		0, 0, 237, 240, 3, 32, 16, 0, 238, 240, 3, 34, 17, 0, 239, 237, 1, 0, 0,
		0, 239, 238, 1, 0, 0, 0, 240, 31, 1, 0, 0, 0, 241, 242, 5, 40, 0, 0, 242,
		243, 5, 55, 0, 0, 243, 244, 5, 43, 0, 0, 244, 33, 1, 0, 0, 0, 245, 252,
		3, 36, 18, 0, 246, 247, 5, 42, 0, 0, 247, 248, 5, 42, 0, 0, 248, 249, 5,
		55, 0, 0, 249, 250, 5, 43, 0, 0, 250, 252, 5, 43, 0, 0, 251, 245, 1, 0,
		0, 0, 251, 246, 1, 0, 0, 0, 252, 35, 1, 0, 0, 0, 253, 254, 5, 42, 0, 0,
		254, 255, 3, 34, 17, 0, 255, 256, 5, 43, 0, 0, 256, 37, 1, 0, 0, 0, 257,
		262, 3, 40, 20, 0, 258, 259, 5, 6, 0, 0, 259, 261, 3, 40, 20, 0, 260, 258,
		1, 0, 0, 0, 261, 264, 1, 0, 0, 0, 262, 260, 1, 0, 0, 0, 262, 263, 1, 0,
		0, 0, 263, 266, 1, 0, 0, 0, 264, 262, 1, 0, 0, 0, 265, 267, 3, 42, 21,
		0, 266, 265, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 39, 1, 0, 0, 0, 268,
		269, 5, 5, 0, 0, 269, 270, 3, 18, 9, 0, 270, 274, 5, 40, 0, 0, 271, 273,
		3, 4, 2, 0, 272, 271, 1, 0, 0, 0, 273, 276, 1, 0, 0, 0, 274, 272, 1, 0,
		0, 0, 274, 275, 1, 0, 0, 0, 275, 277, 1, 0, 0, 0, 276, 274, 1, 0, 0, 0,
		277, 278, 5, 41, 0, 0, 278, 41, 1, 0, 0, 0, 279, 280, 5, 6, 0, 0, 280,
		284, 5, 40, 0, 0, 281, 283, 3, 4, 2, 0, 282, 281, 1, 0, 0, 0, 283, 286,
		1, 0, 0, 0, 284, 282, 1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 287, 1, 0,
		0, 0, 286, 284, 1, 0, 0, 0, 287, 288, 5, 41, 0, 0, 288, 43, 1, 0, 0, 0,
		289, 290, 5, 7, 0, 0, 290, 291, 3, 18, 9, 0, 291, 295, 5, 40, 0, 0, 292,
		294, 3, 46, 23, 0, 293, 292, 1, 0, 0, 0, 294, 297, 1, 0, 0, 0, 295, 293,
		1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 299, 1, 0, 0, 0, 297, 295, 1, 0,
		0, 0, 298, 300, 3, 48, 24, 0, 299, 298, 1, 0, 0, 0, 299, 300, 1, 0, 0,
		0, 300, 301, 1, 0, 0, 0, 301, 302, 5, 41, 0, 0, 302, 45, 1, 0, 0, 0, 303,
		304, 5, 8, 0, 0, 304, 305, 3, 18, 9, 0, 305, 309, 5, 45, 0, 0, 306, 308,
		3, 4, 2, 0, 307, 306, 1, 0, 0, 0, 308, 311, 1, 0, 0, 0, 309, 307, 1, 0,
		0, 0, 309, 310, 1, 0, 0, 0, 310, 47, 1, 0, 0, 0, 311, 309, 1, 0, 0, 0,
		312, 313, 5, 9, 0, 0, 313, 317, 5, 45, 0, 0, 314, 316, 3, 4, 2, 0, 315,
		314, 1, 0, 0, 0, 316, 319, 1, 0, 0, 0, 317, 315, 1, 0, 0, 0, 317, 318,
		1, 0, 0, 0, 318, 49, 1, 0, 0, 0, 319, 317, 1, 0, 0, 0, 320, 321, 5, 11,
		0, 0, 321, 322, 3, 18, 9, 0, 322, 326, 5, 40, 0, 0, 323, 325, 3, 4, 2,
		0, 324, 323, 1, 0, 0, 0, 325, 328, 1, 0, 0, 0, 326, 324, 1, 0, 0, 0, 326,
		327, 1, 0, 0, 0, 327, 329, 1, 0, 0, 0, 328, 326, 1, 0, 0, 0, 329, 330,
		5, 41, 0, 0, 330, 51, 1, 0, 0, 0, 331, 332, 5, 10, 0, 0, 332, 333, 3, 18,
		9, 0, 333, 337, 5, 40, 0, 0, 334, 336, 3, 4, 2, 0, 335, 334, 1, 0, 0, 0,
		336, 339, 1, 0, 0, 0, 337, 335, 1, 0, 0, 0, 337, 338, 1, 0, 0, 0, 338,
		340, 1, 0, 0, 0, 339, 337, 1, 0, 0, 0, 340, 341, 5, 41, 0, 0, 341, 376,
		1, 0, 0, 0, 342, 343, 5, 10, 0, 0, 343, 344, 3, 12, 6, 0, 344, 345, 5,
		44, 0, 0, 345, 346, 3, 18, 9, 0, 346, 347, 5, 44, 0, 0, 347, 348, 3, 12,
		6, 0, 348, 352, 5, 40, 0, 0, 349, 351, 3, 4, 2, 0, 350, 349, 1, 0, 0, 0,
		351, 354, 1, 0, 0, 0, 352, 350, 1, 0, 0, 0, 352, 353, 1, 0, 0, 0, 353,
		355, 1, 0, 0, 0, 354, 352, 1, 0, 0, 0, 355, 356, 5, 41, 0, 0, 356, 376,
		1, 0, 0, 0, 357, 358, 5, 10, 0, 0, 358, 359, 5, 55, 0, 0, 359, 360, 5,
		47, 0, 0, 360, 361, 3, 18, 9, 0, 361, 364, 5, 12, 0, 0, 362, 365, 3, 18,
		9, 0, 363, 365, 3, 54, 27, 0, 364, 362, 1, 0, 0, 0, 364, 363, 1, 0, 0,
		0, 365, 366, 1, 0, 0, 0, 366, 370, 5, 40, 0, 0, 367, 369, 3, 4, 2, 0, 368,
		367, 1, 0, 0, 0, 369, 372, 1, 0, 0, 0, 370, 368, 1, 0, 0, 0, 370, 371,
		1, 0, 0, 0, 371, 373, 1, 0, 0, 0, 372, 370, 1, 0, 0, 0, 373, 374, 5, 41,
		0, 0, 374, 376, 1, 0, 0, 0, 375, 331, 1, 0, 0, 0, 375, 342, 1, 0, 0, 0,
		375, 357, 1, 0, 0, 0, 376, 53, 1, 0, 0, 0, 377, 378, 3, 18, 9, 0, 378,
		379, 5, 46, 0, 0, 379, 380, 5, 46, 0, 0, 380, 381, 5, 46, 0, 0, 381, 382,
		3, 18, 9, 0, 382, 55, 1, 0, 0, 0, 383, 385, 5, 15, 0, 0, 384, 386, 3, 18,
		9, 0, 385, 384, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 390, 1, 0, 0, 0,
		387, 390, 5, 13, 0, 0, 388, 390, 5, 14, 0, 0, 389, 383, 1, 0, 0, 0, 389,
		387, 1, 0, 0, 0, 389, 388, 1, 0, 0, 0, 390, 57, 1, 0, 0, 0, 391, 392, 3,
		14, 7, 0, 392, 394, 5, 38, 0, 0, 393, 395, 3, 60, 30, 0, 394, 393, 1, 0,
		0, 0, 394, 395, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396, 397, 5, 39, 0, 0,
		397, 59, 1, 0, 0, 0, 398, 403, 3, 62, 31, 0, 399, 400, 5, 47, 0, 0, 400,
		402, 3, 62, 31, 0, 401, 399, 1, 0, 0, 0, 402, 405, 1, 0, 0, 0, 403, 401,
		1, 0, 0, 0, 403, 404, 1, 0, 0, 0, 404, 61, 1, 0, 0, 0, 405, 403, 1, 0,
		0, 0, 406, 408, 5, 55, 0, 0, 407, 406, 1, 0, 0, 0, 407, 408, 1, 0, 0, 0,
		408, 411, 1, 0, 0, 0, 409, 412, 3, 14, 7, 0, 410, 412, 3, 18, 9, 0, 411,
		409, 1, 0, 0, 0, 411, 410, 1, 0, 0, 0, 412, 63, 1, 0, 0, 0, 413, 414, 5,
		2, 0, 0, 414, 415, 5, 55, 0, 0, 415, 417, 5, 38, 0, 0, 416, 418, 3, 66,
		33, 0, 417, 416, 1, 0, 0, 0, 417, 418, 1, 0, 0, 0, 418, 419, 1, 0, 0, 0,
		419, 421, 5, 39, 0, 0, 420, 422, 3, 30, 15, 0, 421, 420, 1, 0, 0, 0, 421,
		422, 1, 0, 0, 0, 422, 423, 1, 0, 0, 0, 423, 427, 5, 40, 0, 0, 424, 426,
		3, 4, 2, 0, 425, 424, 1, 0, 0, 0, 426, 429, 1, 0, 0, 0, 427, 425, 1, 0,
		0, 0, 427, 428, 1, 0, 0, 0, 428, 430, 1, 0, 0, 0, 429, 427, 1, 0, 0, 0,
		430, 431, 5, 41, 0, 0, 431, 65, 1, 0, 0, 0, 432, 437, 3, 68, 34, 0, 433,
		434, 5, 47, 0, 0, 434, 436, 3, 68, 34, 0, 435, 433, 1, 0, 0, 0, 436, 439,
		1, 0, 0, 0, 437, 435, 1, 0, 0, 0, 437, 438, 1, 0, 0, 0, 438, 67, 1, 0,
		0, 0, 439, 437, 1, 0, 0, 0, 440, 441, 5, 55, 0, 0, 441, 442, 3, 30, 15,
		0, 442, 69, 1, 0, 0, 0, 443, 444, 5, 3, 0, 0, 444, 445, 5, 55, 0, 0, 445,
		449, 5, 40, 0, 0, 446, 448, 3, 72, 36, 0, 447, 446, 1, 0, 0, 0, 448, 451,
		1, 0, 0, 0, 449, 447, 1, 0, 0, 0, 449, 450, 1, 0, 0, 0, 450, 452, 1, 0,
		0, 0, 451, 449, 1, 0, 0, 0, 452, 453, 5, 41, 0, 0, 453, 71, 1, 0, 0, 0,
		454, 455, 3, 10, 5, 0, 455, 458, 5, 55, 0, 0, 456, 457, 5, 45, 0, 0, 457,
		459, 3, 30, 15, 0, 458, 456, 1, 0, 0, 0, 458, 459, 1, 0, 0, 0, 459, 462,
		1, 0, 0, 0, 460, 461, 5, 26, 0, 0, 461, 463, 3, 18, 9, 0, 462, 460, 1,
		0, 0, 0, 462, 463, 1, 0, 0, 0, 463, 73, 1, 0, 0, 0, 464, 465, 5, 42, 0,
		0, 465, 466, 5, 55, 0, 0, 466, 467, 5, 43, 0, 0, 467, 468, 5, 38, 0, 0,
		468, 469, 5, 39, 0, 0, 469, 75, 1, 0, 0, 0, 45, 79, 83, 86, 108, 115, 122,
		136, 143, 151, 168, 188, 190, 199, 202, 213, 225, 239, 251, 262, 266, 274,
		284, 295, 299, 309, 317, 326, 337, 352, 364, 370, 375, 385, 389, 394, 403,
		407, 411, 417, 421, 427, 437, 449, 458, 462,
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

// VLangGrammarInit initializes any static state used to implement VLangGrammar. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVLangGrammar(). You can call this function if you wish to initialize the static state ahead
// of time.
func VLangGrammarInit() {
	staticData := &VLangGrammarParserStaticData
	staticData.once.Do(vlanggrammarParserInit)
}

// NewVLangGrammar produces a new parser instance for the optional input antlr.TokenStream.
func NewVLangGrammar(input antlr.TokenStream) *VLangGrammar {
	VLangGrammarInit()
	this := new(VLangGrammar)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &VLangGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "VLangGrammar.g4"

	return this
}

// VLangGrammar tokens.
const (
	VLangGrammarEOF            = antlr.TokenEOF
	VLangGrammarMUT            = 1
	VLangGrammarFUNC           = 2
	VLangGrammarSTR            = 3
	VLangGrammarSLICE_TYPE     = 4
	VLangGrammarIF_KW          = 5
	VLangGrammarELSE_KW        = 6
	VLangGrammarSWITCH_KW      = 7
	VLangGrammarCASE_KW        = 8
	VLangGrammarDEFAULT_KW     = 9
	VLangGrammarFOR_KW         = 10
	VLangGrammarWHILE_KW       = 11
	VLangGrammarIN_KW          = 12
	VLangGrammarBREAK_KW       = 13
	VLangGrammarCONTINUE_KW    = 14
	VLangGrammarRETURN_KW      = 15
	VLangGrammarINT_TYPE       = 16
	VLangGrammarFLOAT_TYPE     = 17
	VLangGrammarSTRING_TYPE    = 18
	VLangGrammarBOOL_TYPE      = 19
	VLangGrammarRUNE_TYPE      = 20
	VLangGrammarPLUS           = 21
	VLangGrammarMINUS          = 22
	VLangGrammarMULT           = 23
	VLangGrammarDIV            = 24
	VLangGrammarMOD            = 25
	VLangGrammarASSIGN         = 26
	VLangGrammarPLUS_ASSIGN    = 27
	VLangGrammarMINUS_ASSIGN   = 28
	VLangGrammarEQ             = 29
	VLangGrammarNE             = 30
	VLangGrammarLT             = 31
	VLangGrammarLE             = 32
	VLangGrammarGT             = 33
	VLangGrammarGE             = 34
	VLangGrammarAND            = 35
	VLangGrammarOR             = 36
	VLangGrammarNOT            = 37
	VLangGrammarLPAREN         = 38
	VLangGrammarRPAREN         = 39
	VLangGrammarLBRACE         = 40
	VLangGrammarRBRACE         = 41
	VLangGrammarLBRACK         = 42
	VLangGrammarRBRACK         = 43
	VLangGrammarSEMI           = 44
	VLangGrammarCOLON          = 45
	VLangGrammarDOT            = 46
	VLangGrammarCOMMA          = 47
	VLangGrammarNEWLINE        = 48
	VLangGrammarINT_LITERAL    = 49
	VLangGrammarFLOAT_LITERAL  = 50
	VLangGrammarCHAR_LITERAL   = 51
	VLangGrammarSTRING_LITERAL = 52
	VLangGrammarBOOL_LITERAL   = 53
	VLangGrammarNIL_LITERAL    = 54
	VLangGrammarID             = 55
	VLangGrammarLINE_COMMENT   = 56
	VLangGrammarBLOCK_COMMENT  = 57
	VLangGrammarWS             = 58
)

// VLangGrammar rules.
const (
	VLangGrammarRULE_prog            = 0
	VLangGrammarRULE_delim           = 1
	VLangGrammarRULE_stmt            = 2
	VLangGrammarRULE_decl_stmt       = 3
	VLangGrammarRULE_type_annotation = 4
	VLangGrammarRULE_var_type        = 5
	VLangGrammarRULE_assign_stmt     = 6
	VLangGrammarRULE_id_pattern      = 7
	VLangGrammarRULE_literal         = 8
	VLangGrammarRULE_expression      = 9
	VLangGrammarRULE_vect_expr       = 10
	VLangGrammarRULE_vect_item       = 11
	VLangGrammarRULE_vect_prop       = 12
	VLangGrammarRULE_vect_func       = 13
	VLangGrammarRULE_repeating       = 14
	VLangGrammarRULE_type            = 15
	VLangGrammarRULE_vector_type     = 16
	VLangGrammarRULE_matrix_type     = 17
	VLangGrammarRULE_aux_matrix_type = 18
	VLangGrammarRULE_if_stmt         = 19
	VLangGrammarRULE_if_chain        = 20
	VLangGrammarRULE_else_stmt       = 21
	VLangGrammarRULE_switch_stmt     = 22
	VLangGrammarRULE_switch_case     = 23
	VLangGrammarRULE_default_case    = 24
	VLangGrammarRULE_while_stmt      = 25
	VLangGrammarRULE_for_stmt        = 26
	VLangGrammarRULE_range           = 27
	VLangGrammarRULE_transfer_stmt   = 28
	VLangGrammarRULE_func_call       = 29
	VLangGrammarRULE_arg_list        = 30
	VLangGrammarRULE_func_arg        = 31
	VLangGrammarRULE_func_dcl        = 32
	VLangGrammarRULE_param_list      = 33
	VLangGrammarRULE_func_param      = 34
	VLangGrammarRULE_strct_dcl       = 35
	VLangGrammarRULE_struct_prop     = 36
	VLangGrammarRULE_struct_vector   = 37
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
	p.RuleIndex = VLangGrammarRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_prog

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
	return s.GetToken(VLangGrammarEOF, 0)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VLangGrammarRULE_prog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36028797019022510) != 0 {
		{
			p.SetState(76)
			p.Stmt()
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(82)
			p.Match(VLangGrammarEOF)
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

	// Getter signatures
	NEWLINE() antlr.TerminalNode

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
	p.RuleIndex = VLangGrammarRULE_delim
	return p
}

func InitEmptyDelimContext(p *DelimContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_delim
}

func (*DelimContext) IsDelimContext() {}

func NewDelimContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DelimContext {
	var p = new(DelimContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_delim

	return p
}

func (s *DelimContext) GetParser() antlr.Parser { return s.parser }

func (s *DelimContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarNEWLINE, 0)
}

func (s *DelimContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DelimContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DelimContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterDelim(s)
	}
}

func (s *DelimContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitDelim(s)
	}
}

func (s *DelimContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitDelim(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Delim() (localctx IDelimContext) {
	localctx = NewDelimContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VLangGrammarRULE_delim)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VLangGrammarNEWLINE {
		{
			p.SetState(85)
			p.Match(VLangGrammarNEWLINE)
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

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Decl_stmt() IDecl_stmtContext
	Delim() IDelimContext
	Assign_stmt() IAssign_stmtContext
	Transfer_stmt() ITransfer_stmtContext
	If_stmt() IIf_stmtContext
	Switch_stmt() ISwitch_stmtContext
	For_stmt() IFor_stmtContext
	Func_call() IFunc_callContext
	Vect_func() IVect_funcContext
	Func_dcl() IFunc_dclContext
	Strct_dcl() IStrct_dclContext

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
	p.RuleIndex = VLangGrammarRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Decl_stmt() IDecl_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecl_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecl_stmtContext)
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

func (s *StmtContext) Assign_stmt() IAssign_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_stmtContext)
}

func (s *StmtContext) Transfer_stmt() ITransfer_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransfer_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITransfer_stmtContext)
}

func (s *StmtContext) If_stmt() IIf_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_stmtContext)
}

func (s *StmtContext) Switch_stmt() ISwitch_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_stmtContext)
}

func (s *StmtContext) For_stmt() IFor_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_stmtContext)
}

func (s *StmtContext) Func_call() IFunc_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_callContext)
}

func (s *StmtContext) Vect_func() IVect_funcContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVect_funcContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVect_funcContext)
}

func (s *StmtContext) Func_dcl() IFunc_dclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_dclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_dclContext)
}

func (s *StmtContext) Strct_dcl() IStrct_dclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrct_dclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrct_dclContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, VLangGrammarRULE_stmt)
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Decl_stmt()
		}
		{
			p.SetState(89)
			p.Delim()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)
			p.Assign_stmt()
		}
		{
			p.SetState(92)
			p.Delim()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(94)
			p.Transfer_stmt()
		}
		{
			p.SetState(95)
			p.Delim()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(97)
			p.If_stmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(98)
			p.Switch_stmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(99)
			p.For_stmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(100)
			p.Func_call()
		}
		{
			p.SetState(101)
			p.Delim()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(103)
			p.Vect_func()
		}
		{
			p.SetState(104)
			p.Delim()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(106)
			p.Func_dcl()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(107)
			p.Strct_dcl()
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

// IDecl_stmtContext is an interface to support dynamic dispatch.
type IDecl_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDecl_stmtContext differentiates from other interfaces.
	IsDecl_stmtContext()
}

type Decl_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecl_stmtContext() *Decl_stmtContext {
	var p = new(Decl_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_decl_stmt
	return p
}

func InitEmptyDecl_stmtContext(p *Decl_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_decl_stmt
}

func (*Decl_stmtContext) IsDecl_stmtContext() {}

func NewDecl_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Decl_stmtContext {
	var p = new(Decl_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_decl_stmt

	return p
}

func (s *Decl_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Decl_stmtContext) CopyAll(ctx *Decl_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Decl_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Decl_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VarAssDeclContext struct {
	Decl_stmtContext
}

func NewVarAssDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarAssDeclContext {
	var p = new(VarAssDeclContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *VarAssDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarAssDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(VLangGrammarID, 0)
}

func (s *VarAssDeclContext) Type_annotation() IType_annotationContext {
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

func (s *VarAssDeclContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarASSIGN, 0)
}

func (s *VarAssDeclContext) Expression() IExpressionContext {
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

func (s *VarAssDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterVarAssDecl(s)
	}
}

func (s *VarAssDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitVarAssDecl(s)
	}
}

func (s *VarAssDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitVarAssDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type MutVarDeclContext struct {
	Decl_stmtContext
}

func NewMutVarDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MutVarDeclContext {
	var p = new(MutVarDeclContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *MutVarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutVarDeclContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *MutVarDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(VLangGrammarID, 0)
}

func (s *MutVarDeclContext) Type_annotation() IType_annotationContext {
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

func (s *MutVarDeclContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarASSIGN, 0)
}

func (s *MutVarDeclContext) Expression() IExpressionContext {
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

func (s *MutVarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterMutVarDecl(s)
	}
}

func (s *MutVarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitMutVarDecl(s)
	}
}

func (s *MutVarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitMutVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Decl_stmt() (localctx IDecl_stmtContext) {
	localctx = NewDecl_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, VLangGrammarRULE_decl_stmt)
	var _la int

	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VLangGrammarMUT:
		localctx = NewMutVarDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.Var_type()
		}
		{
			p.SetState(111)
			p.Match(VLangGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.Type_annotation()
		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == VLangGrammarASSIGN {
			{
				p.SetState(113)
				p.Match(VLangGrammarASSIGN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(114)
				p.expression(0)
			}

		}

	case VLangGrammarID:
		localctx = NewVarAssDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.Match(VLangGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.Type_annotation()
		}
		{
			p.SetState(119)
			p.Match(VLangGrammarASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
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
	p.RuleIndex = VLangGrammarRULE_type_annotation
	return p
}

func InitEmptyType_annotationContext(p *Type_annotationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_type_annotation
}

func (*Type_annotationContext) IsType_annotationContext() {}

func NewType_annotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_annotationContext {
	var p = new(Type_annotationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_type_annotation

	return p
}

func (s *Type_annotationContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_annotationContext) INT_TYPE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarINT_TYPE, 0)
}

func (s *Type_annotationContext) FLOAT_TYPE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarFLOAT_TYPE, 0)
}

func (s *Type_annotationContext) STRING_TYPE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarSTRING_TYPE, 0)
}

func (s *Type_annotationContext) BOOL_TYPE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarBOOL_TYPE, 0)
}

func (s *Type_annotationContext) SLICE_TYPE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarSLICE_TYPE, 0)
}

func (s *Type_annotationContext) ID() antlr.TerminalNode {
	return s.GetToken(VLangGrammarID, 0)
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

func (p *VLangGrammar) Type_annotation() (localctx IType_annotationContext) {
	localctx = NewType_annotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, VLangGrammarRULE_type_annotation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36028797019947024) != 0) {
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

// IVar_typeContext is an interface to support dynamic dispatch.
type IVar_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MUT() antlr.TerminalNode

	// IsVar_typeContext differentiates from other interfaces.
	IsVar_typeContext()
}

type Var_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_typeContext() *Var_typeContext {
	var p = new(Var_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_var_type
	return p
}

func InitEmptyVar_typeContext(p *Var_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_var_type
}

func (*Var_typeContext) IsVar_typeContext() {}

func NewVar_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_typeContext {
	var p = new(Var_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_var_type

	return p
}

func (s *Var_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_typeContext) MUT() antlr.TerminalNode {
	return s.GetToken(VLangGrammarMUT, 0)
}

func (s *Var_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterVar_type(s)
	}
}

func (s *Var_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitVar_type(s)
	}
}

func (s *Var_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitVar_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Var_type() (localctx IVar_typeContext) {
	localctx = NewVar_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, VLangGrammarRULE_var_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(VLangGrammarMUT)
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

// IAssign_stmtContext is an interface to support dynamic dispatch.
type IAssign_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssign_stmtContext differentiates from other interfaces.
	IsAssign_stmtContext()
}

type Assign_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_stmtContext() *Assign_stmtContext {
	var p = new(Assign_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_assign_stmt
	return p
}

func InitEmptyAssign_stmtContext(p *Assign_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_assign_stmt
}

func (*Assign_stmtContext) IsAssign_stmtContext() {}

func NewAssign_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_stmtContext {
	var p = new(Assign_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_assign_stmt

	return p
}

func (s *Assign_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_stmtContext) CopyAll(ctx *Assign_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Assign_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AugmentedAssignmentDeclContext struct {
	Assign_stmtContext
	op antlr.Token
}

func NewAugmentedAssignmentDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AugmentedAssignmentDeclContext {
	var p = new(AugmentedAssignmentDeclContext)

	InitEmptyAssign_stmtContext(&p.Assign_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_stmtContext))

	return p
}

func (s *AugmentedAssignmentDeclContext) GetOp() antlr.Token { return s.op }

func (s *AugmentedAssignmentDeclContext) SetOp(v antlr.Token) { s.op = v }

func (s *AugmentedAssignmentDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AugmentedAssignmentDeclContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *AugmentedAssignmentDeclContext) Expression() IExpressionContext {
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

func (s *AugmentedAssignmentDeclContext) PLUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarPLUS_ASSIGN, 0)
}

func (s *AugmentedAssignmentDeclContext) MINUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarMINUS_ASSIGN, 0)
}

func (s *AugmentedAssignmentDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterAugmentedAssignmentDecl(s)
	}
}

func (s *AugmentedAssignmentDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitAugmentedAssignmentDecl(s)
	}
}

func (s *AugmentedAssignmentDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitAugmentedAssignmentDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentDeclContext struct {
	Assign_stmtContext
}

func NewAssignmentDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentDeclContext {
	var p = new(AssignmentDeclContext)

	InitEmptyAssign_stmtContext(&p.Assign_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_stmtContext))

	return p
}

func (s *AssignmentDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentDeclContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *AssignmentDeclContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarASSIGN, 0)
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
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterAssignmentDecl(s)
	}
}

func (s *AssignmentDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitAssignmentDecl(s)
	}
}

func (s *AssignmentDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitAssignmentDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Assign_stmt() (localctx IAssign_stmtContext) {
	localctx = NewAssign_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, VLangGrammarRULE_assign_stmt)
	var _la int

	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignmentDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.Id_pattern()
		}
		{
			p.SetState(129)
			p.Match(VLangGrammarASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.expression(0)
		}

	case 2:
		localctx = NewAugmentedAssignmentDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)
			p.Id_pattern()
		}
		{
			p.SetState(133)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AugmentedAssignmentDeclContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == VLangGrammarPLUS_ASSIGN || _la == VLangGrammarMINUS_ASSIGN) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AugmentedAssignmentDeclContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(134)
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

// IId_patternContext is an interface to support dynamic dispatch.
type IId_patternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsId_patternContext differentiates from other interfaces.
	IsId_patternContext()
}

type Id_patternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyId_patternContext() *Id_patternContext {
	var p = new(Id_patternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_id_pattern
	return p
}

func InitEmptyId_patternContext(p *Id_patternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_id_pattern
}

func (*Id_patternContext) IsId_patternContext() {}

func NewId_patternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Id_patternContext {
	var p = new(Id_patternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_id_pattern

	return p
}

func (s *Id_patternContext) GetParser() antlr.Parser { return s.parser }

func (s *Id_patternContext) CopyAll(ctx *Id_patternContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Id_patternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Id_patternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdPatternContext struct {
	Id_patternContext
}

func NewIdPatternContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdPatternContext {
	var p = new(IdPatternContext)

	InitEmptyId_patternContext(&p.Id_patternContext)
	p.parser = parser
	p.CopyAll(ctx.(*Id_patternContext))

	return p
}

func (s *IdPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdPatternContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VLangGrammarID)
}

func (s *IdPatternContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VLangGrammarID, i)
}

func (s *IdPatternContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(VLangGrammarDOT)
}

func (s *IdPatternContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(VLangGrammarDOT, i)
}

func (s *IdPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterIdPattern(s)
	}
}

func (s *IdPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitIdPattern(s)
	}
}

func (s *IdPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitIdPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Id_pattern() (localctx IId_patternContext) {
	localctx = NewId_patternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, VLangGrammarRULE_id_pattern)
	var _alt int

	localctx = NewIdPatternContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(VLangGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(139)
				p.Match(VLangGrammarDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(140)
				p.Match(VLangGrammarID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyAll(ctx *LiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringLiteralContext struct {
	LiteralContext
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext {
	var p = new(StringLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(VLangGrammarSTRING_LITERAL, 0)
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

type BoolLiteralContext struct {
	LiteralContext
}

func NewBoolLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolLiteralContext {
	var p = new(BoolLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *BoolLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLiteralContext) BOOL_LITERAL() antlr.TerminalNode {
	return s.GetToken(VLangGrammarBOOL_LITERAL, 0)
}

func (s *BoolLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterBoolLiteral(s)
	}
}

func (s *BoolLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitBoolLiteral(s)
	}
}

func (s *BoolLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitBoolLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatLiteralContext struct {
	LiteralContext
}

func NewFloatLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(VLangGrammarFLOAT_LITERAL, 0)
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
	LiteralContext
}

func NewNilLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilLiteralContext {
	var p = new(NilLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *NilLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilLiteralContext) NIL_LITERAL() antlr.TerminalNode {
	return s.GetToken(VLangGrammarNIL_LITERAL, 0)
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
	LiteralContext
}

func NewIntLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntLiteralContext {
	var p = new(IntLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *IntLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLiteralContext) INT_LITERAL() antlr.TerminalNode {
	return s.GetToken(VLangGrammarINT_LITERAL, 0)
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

func (p *VLangGrammar) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, VLangGrammarRULE_literal)
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VLangGrammarINT_LITERAL:
		localctx = NewIntLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(146)
			p.Match(VLangGrammarINT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangGrammarFLOAT_LITERAL:
		localctx = NewFloatLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(147)
			p.Match(VLangGrammarFLOAT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangGrammarSTRING_LITERAL:
		localctx = NewStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(148)
			p.Match(VLangGrammarSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangGrammarBOOL_LITERAL:
		localctx = NewBoolLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(149)
			p.Match(VLangGrammarBOOL_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangGrammarNIL_LITERAL:
		localctx = NewNilLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(150)
			p.Match(VLangGrammarNIL_LITERAL)
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
	p.RuleIndex = VLangGrammarRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_expression

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

type RepeatingExprContext struct {
	ExpressionContext
}

func NewRepeatingExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RepeatingExprContext {
	var p = new(RepeatingExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *RepeatingExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatingExprContext) Repeating() IRepeatingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRepeatingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRepeatingContext)
}

func (s *RepeatingExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterRepeatingExpr(s)
	}
}

func (s *RepeatingExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitRepeatingExpr(s)
	}
}

func (s *RepeatingExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitRepeatingExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorPropertyExprContext struct {
	ExpressionContext
}

func NewVectorPropertyExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorPropertyExprContext {
	var p = new(VectorPropertyExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *VectorPropertyExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorPropertyExprContext) Vect_prop() IVect_propContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVect_propContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVect_propContext)
}

func (s *VectorPropertyExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterVectorPropertyExpr(s)
	}
}

func (s *VectorPropertyExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitVectorPropertyExpr(s)
	}
}

func (s *VectorPropertyExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitVectorPropertyExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorItemExprContext struct {
	ExpressionContext
}

func NewVectorItemExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorItemExprContext {
	var p = new(VectorItemExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *VectorItemExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorItemExprContext) Vect_item() IVect_itemContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVect_itemContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVect_itemContext)
}

func (s *VectorItemExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterVectorItemExpr(s)
	}
}

func (s *VectorItemExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitVectorItemExpr(s)
	}
}

func (s *VectorItemExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitVectorItemExpr(s)

	default:
		return t.VisitChildren(s)
	}
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
	return s.GetToken(VLangGrammarMULT, 0)
}

func (s *BinaryExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(VLangGrammarDIV, 0)
}

func (s *BinaryExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(VLangGrammarMOD, 0)
}

func (s *BinaryExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(VLangGrammarPLUS, 0)
}

func (s *BinaryExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(VLangGrammarMINUS, 0)
}

func (s *BinaryExprContext) LE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarLE, 0)
}

func (s *BinaryExprContext) LT() antlr.TerminalNode {
	return s.GetToken(VLangGrammarLT, 0)
}

func (s *BinaryExprContext) GE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarGE, 0)
}

func (s *BinaryExprContext) GT() antlr.TerminalNode {
	return s.GetToken(VLangGrammarGT, 0)
}

func (s *BinaryExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(VLangGrammarEQ, 0)
}

func (s *BinaryExprContext) NE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarNE, 0)
}

func (s *BinaryExprContext) AND() antlr.TerminalNode {
	return s.GetToken(VLangGrammarAND, 0)
}

func (s *BinaryExprContext) OR() antlr.TerminalNode {
	return s.GetToken(VLangGrammarOR, 0)
}

func (s *BinaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterBinaryExpr(s)
	}
}

func (s *BinaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitBinaryExpr(s)
	}
}

func (s *BinaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitBinaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParensExprContext struct {
	ExpressionContext
}

func NewParensExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensExprContext {
	var p = new(ParensExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParensExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarLPAREN, 0)
}

func (s *ParensExprContext) Expression() IExpressionContext {
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

func (s *ParensExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarRPAREN, 0)
}

func (s *ParensExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterParensExpr(s)
	}
}

func (s *ParensExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitParensExpr(s)
	}
}

func (s *ParensExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitParensExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralExprContext struct {
	ExpressionContext
}

func NewLiteralExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExprContext {
	var p = new(LiteralExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LiteralExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExprContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterLiteralExpr(s)
	}
}

func (s *LiteralExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitLiteralExpr(s)
	}
}

func (s *LiteralExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitLiteralExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorFuncCallExprContext struct {
	ExpressionContext
}

func NewVectorFuncCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorFuncCallExprContext {
	var p = new(VectorFuncCallExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *VectorFuncCallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorFuncCallExprContext) Vect_func() IVect_funcContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVect_funcContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVect_funcContext)
}

func (s *VectorFuncCallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterVectorFuncCallExpr(s)
	}
}

func (s *VectorFuncCallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitVectorFuncCallExpr(s)
	}
}

func (s *VectorFuncCallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitVectorFuncCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorExprContext struct {
	ExpressionContext
}

func NewVectorExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorExprContext {
	var p = new(VectorExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *VectorExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorExprContext) Vect_expr() IVect_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVect_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVect_exprContext)
}

func (s *VectorExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterVectorExpr(s)
	}
}

func (s *VectorExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitVectorExpr(s)
	}
}

func (s *VectorExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitVectorExpr(s)

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
	return s.GetToken(VLangGrammarNOT, 0)
}

func (s *UnaryExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(VLangGrammarMINUS, 0)
}

func (s *UnaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterUnaryExpr(s)
	}
}

func (s *UnaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitUnaryExpr(s)
	}
}

func (s *UnaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitUnaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdPatternExprContext struct {
	ExpressionContext
}

func NewIdPatternExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdPatternExprContext {
	var p = new(IdPatternExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IdPatternExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdPatternExprContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *IdPatternExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterIdPatternExpr(s)
	}
}

func (s *IdPatternExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitIdPatternExpr(s)
	}
}

func (s *IdPatternExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitIdPatternExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type FuncCallExprContext struct {
	ExpressionContext
}

func NewFuncCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncCallExprContext {
	var p = new(FuncCallExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FuncCallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallExprContext) Func_call() IFunc_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_callContext)
}

func (s *FuncCallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterFuncCallExpr(s)
	}
}

func (s *FuncCallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitFuncCallExpr(s)
	}
}

func (s *FuncCallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitFuncCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *VLangGrammar) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, VLangGrammarRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParensExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(154)
			p.Match(VLangGrammarLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(155)
			p.expression(0)
		}
		{
			p.SetState(156)
			p.Match(VLangGrammarRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewFuncCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(158)
			p.Func_call()
		}

	case 3:
		localctx = NewIdPatternExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(159)
			p.Id_pattern()
		}

	case 4:
		localctx = NewVectorItemExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(160)
			p.Vect_item()
		}

	case 5:
		localctx = NewVectorPropertyExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(161)
			p.Vect_prop()
		}

	case 6:
		localctx = NewVectorFuncCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(162)
			p.Vect_func()
		}

	case 7:
		localctx = NewLiteralExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(163)
			p.Literal()
		}

	case 8:
		localctx = NewVectorExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(164)
			p.Vect_expr()
		}

	case 9:
		localctx = NewRepeatingExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(165)
			p.Repeating()
		}

	case 10:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(166)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == VLangGrammarMINUS || _la == VLangGrammarNOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(167)
			p.expression(7)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(188)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VLangGrammarRULE_expression)
				p.SetState(170)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(171)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&58720256) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(172)

					var _x = p.expression(7)

					localctx.(*BinaryExprContext).right = _x
				}

			case 2:
				localctx = NewBinaryExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VLangGrammarRULE_expression)
				p.SetState(173)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(174)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VLangGrammarPLUS || _la == VLangGrammarMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(175)

					var _x = p.expression(6)

					localctx.(*BinaryExprContext).right = _x
				}

			case 3:
				localctx = NewBinaryExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VLangGrammarRULE_expression)
				p.SetState(176)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(177)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32212254720) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(178)

					var _x = p.expression(5)

					localctx.(*BinaryExprContext).right = _x
				}

			case 4:
				localctx = NewBinaryExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VLangGrammarRULE_expression)
				p.SetState(179)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(180)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VLangGrammarEQ || _la == VLangGrammarNE) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(181)

					var _x = p.expression(4)

					localctx.(*BinaryExprContext).right = _x
				}

			case 5:
				localctx = NewBinaryExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VLangGrammarRULE_expression)
				p.SetState(182)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(183)

					var _m = p.Match(VLangGrammarAND)

					localctx.(*BinaryExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(184)

					var _x = p.expression(3)

					localctx.(*BinaryExprContext).right = _x
				}

			case 6:
				localctx = NewBinaryExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VLangGrammarRULE_expression)
				p.SetState(185)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(186)

					var _m = p.Match(VLangGrammarOR)

					localctx.(*BinaryExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(187)

					var _x = p.expression(2)

					localctx.(*BinaryExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
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

// IVect_exprContext is an interface to support dynamic dispatch.
type IVect_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVect_exprContext differentiates from other interfaces.
	IsVect_exprContext()
}

type Vect_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVect_exprContext() *Vect_exprContext {
	var p = new(Vect_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_vect_expr
	return p
}

func InitEmptyVect_exprContext(p *Vect_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_vect_expr
}

func (*Vect_exprContext) IsVect_exprContext() {}

func NewVect_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vect_exprContext {
	var p = new(Vect_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_vect_expr

	return p
}

func (s *Vect_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Vect_exprContext) CopyAll(ctx *Vect_exprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Vect_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vect_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VectorItemLisContext struct {
	Vect_exprContext
}

func NewVectorItemLisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorItemLisContext {
	var p = new(VectorItemLisContext)

	InitEmptyVect_exprContext(&p.Vect_exprContext)
	p.parser = parser
	p.CopyAll(ctx.(*Vect_exprContext))

	return p
}

func (s *VectorItemLisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorItemLisContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(VLangGrammarLBRACK, 0)
}

func (s *VectorItemLisContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(VLangGrammarRBRACK, 0)
}

func (s *VectorItemLisContext) AllExpression() []IExpressionContext {
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

func (s *VectorItemLisContext) Expression(i int) IExpressionContext {
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

func (s *VectorItemLisContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VLangGrammarCOMMA)
}

func (s *VectorItemLisContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VLangGrammarCOMMA, i)
}

func (s *VectorItemLisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterVectorItemLis(s)
	}
}

func (s *VectorItemLisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitVectorItemLis(s)
	}
}

func (s *VectorItemLisContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitVectorItemLis(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Vect_expr() (localctx IVect_exprContext) {
	localctx = NewVect_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, VLangGrammarRULE_vect_expr)
	var _la int

	localctx = NewVectorItemLisContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.Match(VLangGrammarLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&69247654638387202) != 0 {
		{
			p.SetState(194)
			p.expression(0)
		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VLangGrammarCOMMA {
			{
				p.SetState(195)
				p.Match(VLangGrammarCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(196)
				p.expression(0)
			}

			p.SetState(201)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(204)
		p.Match(VLangGrammarRBRACK)
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

// IVect_itemContext is an interface to support dynamic dispatch.
type IVect_itemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVect_itemContext differentiates from other interfaces.
	IsVect_itemContext()
}

type Vect_itemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVect_itemContext() *Vect_itemContext {
	var p = new(Vect_itemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_vect_item
	return p
}

func InitEmptyVect_itemContext(p *Vect_itemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_vect_item
}

func (*Vect_itemContext) IsVect_itemContext() {}

func NewVect_itemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vect_itemContext {
	var p = new(Vect_itemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_vect_item

	return p
}

func (s *Vect_itemContext) GetParser() antlr.Parser { return s.parser }

func (s *Vect_itemContext) CopyAll(ctx *Vect_itemContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Vect_itemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vect_itemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VectorItemContext struct {
	Vect_itemContext
}

func NewVectorItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorItemContext {
	var p = new(VectorItemContext)

	InitEmptyVect_itemContext(&p.Vect_itemContext)
	p.parser = parser
	p.CopyAll(ctx.(*Vect_itemContext))

	return p
}

func (s *VectorItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorItemContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *VectorItemContext) AllLBRACE() []antlr.TerminalNode {
	return s.GetTokens(VLangGrammarLBRACE)
}

func (s *VectorItemContext) LBRACE(i int) antlr.TerminalNode {
	return s.GetToken(VLangGrammarLBRACE, i)
}

func (s *VectorItemContext) AllExpression() []IExpressionContext {
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

func (s *VectorItemContext) Expression(i int) IExpressionContext {
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

func (s *VectorItemContext) AllRBRACE() []antlr.TerminalNode {
	return s.GetTokens(VLangGrammarRBRACE)
}

func (s *VectorItemContext) RBRACE(i int) antlr.TerminalNode {
	return s.GetToken(VLangGrammarRBRACE, i)
}

func (s *VectorItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterVectorItem(s)
	}
}

func (s *VectorItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitVectorItem(s)
	}
}

func (s *VectorItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitVectorItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Vect_item() (localctx IVect_itemContext) {
	localctx = NewVect_itemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, VLangGrammarRULE_vect_item)
	var _alt int

	localctx = NewVectorItemContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Id_pattern()
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(207)
				p.Match(VLangGrammarLBRACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(208)
				p.expression(0)
			}
			{
				p.SetState(209)
				p.Match(VLangGrammarRBRACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVect_propContext is an interface to support dynamic dispatch.
type IVect_propContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVect_propContext differentiates from other interfaces.
	IsVect_propContext()
}

type Vect_propContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVect_propContext() *Vect_propContext {
	var p = new(Vect_propContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_vect_prop
	return p
}

func InitEmptyVect_propContext(p *Vect_propContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_vect_prop
}

func (*Vect_propContext) IsVect_propContext() {}

func NewVect_propContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vect_propContext {
	var p = new(Vect_propContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_vect_prop

	return p
}

func (s *Vect_propContext) GetParser() antlr.Parser { return s.parser }

func (s *Vect_propContext) CopyAll(ctx *Vect_propContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Vect_propContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vect_propContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VectorPropertyContext struct {
	Vect_propContext
}

func NewVectorPropertyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorPropertyContext {
	var p = new(VectorPropertyContext)

	InitEmptyVect_propContext(&p.Vect_propContext)
	p.parser = parser
	p.CopyAll(ctx.(*Vect_propContext))

	return p
}

func (s *VectorPropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorPropertyContext) Vect_item() IVect_itemContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVect_itemContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVect_itemContext)
}

func (s *VectorPropertyContext) DOT() antlr.TerminalNode {
	return s.GetToken(VLangGrammarDOT, 0)
}

func (s *VectorPropertyContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *VectorPropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterVectorProperty(s)
	}
}

func (s *VectorPropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitVectorProperty(s)
	}
}

func (s *VectorPropertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitVectorProperty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Vect_prop() (localctx IVect_propContext) {
	localctx = NewVect_propContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, VLangGrammarRULE_vect_prop)
	localctx = NewVectorPropertyContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Vect_item()
	}
	{
		p.SetState(216)
		p.Match(VLangGrammarDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)
		p.Id_pattern()
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

// IVect_funcContext is an interface to support dynamic dispatch.
type IVect_funcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVect_funcContext differentiates from other interfaces.
	IsVect_funcContext()
}

type Vect_funcContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVect_funcContext() *Vect_funcContext {
	var p = new(Vect_funcContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_vect_func
	return p
}

func InitEmptyVect_funcContext(p *Vect_funcContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_vect_func
}

func (*Vect_funcContext) IsVect_funcContext() {}

func NewVect_funcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vect_funcContext {
	var p = new(Vect_funcContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_vect_func

	return p
}

func (s *Vect_funcContext) GetParser() antlr.Parser { return s.parser }

func (s *Vect_funcContext) CopyAll(ctx *Vect_funcContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Vect_funcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vect_funcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VectorFuncCallContext struct {
	Vect_funcContext
}

func NewVectorFuncCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorFuncCallContext {
	var p = new(VectorFuncCallContext)

	InitEmptyVect_funcContext(&p.Vect_funcContext)
	p.parser = parser
	p.CopyAll(ctx.(*Vect_funcContext))

	return p
}

func (s *VectorFuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorFuncCallContext) Vect_item() IVect_itemContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVect_itemContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVect_itemContext)
}

func (s *VectorFuncCallContext) DOT() antlr.TerminalNode {
	return s.GetToken(VLangGrammarDOT, 0)
}

func (s *VectorFuncCallContext) Func_call() IFunc_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_callContext)
}

func (s *VectorFuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterVectorFuncCall(s)
	}
}

func (s *VectorFuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitVectorFuncCall(s)
	}
}

func (s *VectorFuncCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitVectorFuncCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Vect_func() (localctx IVect_funcContext) {
	localctx = NewVect_funcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, VLangGrammarRULE_vect_func)
	localctx = NewVectorFuncCallContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(219)
		p.Vect_item()
	}
	{
		p.SetState(220)
		p.Match(VLangGrammarDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(221)
		p.Func_call()
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

// IRepeatingContext is an interface to support dynamic dispatch.
type IRepeatingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRepeatingContext differentiates from other interfaces.
	IsRepeatingContext()
}

type RepeatingContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRepeatingContext() *RepeatingContext {
	var p = new(RepeatingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_repeating
	return p
}

func InitEmptyRepeatingContext(p *RepeatingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_repeating
}

func (*RepeatingContext) IsRepeatingContext() {}

func NewRepeatingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepeatingContext {
	var p = new(RepeatingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_repeating

	return p
}

func (s *RepeatingContext) GetParser() antlr.Parser { return s.parser }

func (s *RepeatingContext) CopyAll(ctx *RepeatingContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RepeatingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RepeatingDeclContext struct {
	RepeatingContext
}

func NewRepeatingDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RepeatingDeclContext {
	var p = new(RepeatingDeclContext)

	InitEmptyRepeatingContext(&p.RepeatingContext)
	p.parser = parser
	p.CopyAll(ctx.(*RepeatingContext))

	return p
}

func (s *RepeatingDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatingDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarLPAREN, 0)
}

func (s *RepeatingDeclContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VLangGrammarID)
}

func (s *RepeatingDeclContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VLangGrammarID, i)
}

func (s *RepeatingDeclContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(VLangGrammarCOLON)
}

func (s *RepeatingDeclContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(VLangGrammarCOLON, i)
}

func (s *RepeatingDeclContext) AllExpression() []IExpressionContext {
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

func (s *RepeatingDeclContext) Expression(i int) IExpressionContext {
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

func (s *RepeatingDeclContext) COMMA() antlr.TerminalNode {
	return s.GetToken(VLangGrammarCOMMA, 0)
}

func (s *RepeatingDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarRPAREN, 0)
}

func (s *RepeatingDeclContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *RepeatingDeclContext) Matrix_type() IMatrix_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrix_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrix_typeContext)
}

func (s *RepeatingDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterRepeatingDecl(s)
	}
}

func (s *RepeatingDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitRepeatingDecl(s)
	}
}

func (s *RepeatingDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitRepeatingDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Repeating() (localctx IRepeatingContext) {
	localctx = NewRepeatingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, VLangGrammarRULE_repeating)
	localctx = NewRepeatingDeclContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VLangGrammarMUT:
		{
			p.SetState(223)
			p.Var_type()
		}

	case VLangGrammarLBRACK:
		{
			p.SetState(224)
			p.Matrix_type()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(227)
		p.Match(VLangGrammarLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(228)
		p.Match(VLangGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.Match(VLangGrammarCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.expression(0)
	}
	{
		p.SetState(231)
		p.Match(VLangGrammarCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)
		p.Match(VLangGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(233)
		p.Match(VLangGrammarCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(234)
		p.expression(0)
	}
	{
		p.SetState(235)
		p.Match(VLangGrammarRPAREN)
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Vector_type() IVector_typeContext
	Matrix_type() IMatrix_typeContext

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) Vector_type() IVector_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVector_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVector_typeContext)
}

func (s *TypeContext) Matrix_type() IMatrix_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrix_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrix_typeContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, VLangGrammarRULE_type)
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VLangGrammarLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(237)
			p.Vector_type()
		}

	case VLangGrammarLBRACK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(238)
			p.Matrix_type()
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

// IVector_typeContext is an interface to support dynamic dispatch.
type IVector_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	ID() antlr.TerminalNode
	RBRACK() antlr.TerminalNode

	// IsVector_typeContext differentiates from other interfaces.
	IsVector_typeContext()
}

type Vector_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVector_typeContext() *Vector_typeContext {
	var p = new(Vector_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_vector_type
	return p
}

func InitEmptyVector_typeContext(p *Vector_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_vector_type
}

func (*Vector_typeContext) IsVector_typeContext() {}

func NewVector_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vector_typeContext {
	var p = new(Vector_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_vector_type

	return p
}

func (s *Vector_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Vector_typeContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarLBRACE, 0)
}

func (s *Vector_typeContext) ID() antlr.TerminalNode {
	return s.GetToken(VLangGrammarID, 0)
}

func (s *Vector_typeContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(VLangGrammarRBRACK, 0)
}

func (s *Vector_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vector_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Vector_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterVector_type(s)
	}
}

func (s *Vector_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitVector_type(s)
	}
}

func (s *Vector_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitVector_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Vector_type() (localctx IVector_typeContext) {
	localctx = NewVector_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, VLangGrammarRULE_vector_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Match(VLangGrammarLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(242)
		p.Match(VLangGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.Match(VLangGrammarRBRACK)
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

// IMatrix_typeContext is an interface to support dynamic dispatch.
type IMatrix_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Aux_matrix_type() IAux_matrix_typeContext
	AllLBRACK() []antlr.TerminalNode
	LBRACK(i int) antlr.TerminalNode
	ID() antlr.TerminalNode
	AllRBRACK() []antlr.TerminalNode
	RBRACK(i int) antlr.TerminalNode

	// IsMatrix_typeContext differentiates from other interfaces.
	IsMatrix_typeContext()
}

type Matrix_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrix_typeContext() *Matrix_typeContext {
	var p = new(Matrix_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_matrix_type
	return p
}

func InitEmptyMatrix_typeContext(p *Matrix_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_matrix_type
}

func (*Matrix_typeContext) IsMatrix_typeContext() {}

func NewMatrix_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Matrix_typeContext {
	var p = new(Matrix_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_matrix_type

	return p
}

func (s *Matrix_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Matrix_typeContext) Aux_matrix_type() IAux_matrix_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAux_matrix_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAux_matrix_typeContext)
}

func (s *Matrix_typeContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(VLangGrammarLBRACK)
}

func (s *Matrix_typeContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(VLangGrammarLBRACK, i)
}

func (s *Matrix_typeContext) ID() antlr.TerminalNode {
	return s.GetToken(VLangGrammarID, 0)
}

func (s *Matrix_typeContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(VLangGrammarRBRACK)
}

func (s *Matrix_typeContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(VLangGrammarRBRACK, i)
}

func (s *Matrix_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Matrix_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Matrix_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterMatrix_type(s)
	}
}

func (s *Matrix_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitMatrix_type(s)
	}
}

func (s *Matrix_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitMatrix_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Matrix_type() (localctx IMatrix_typeContext) {
	localctx = NewMatrix_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, VLangGrammarRULE_matrix_type)
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(245)
			p.Aux_matrix_type()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(246)
			p.Match(VLangGrammarLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(247)
			p.Match(VLangGrammarLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(248)
			p.Match(VLangGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(249)
			p.Match(VLangGrammarRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(250)
			p.Match(VLangGrammarRBRACK)
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

// IAux_matrix_typeContext is an interface to support dynamic dispatch.
type IAux_matrix_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	Matrix_type() IMatrix_typeContext
	RBRACK() antlr.TerminalNode

	// IsAux_matrix_typeContext differentiates from other interfaces.
	IsAux_matrix_typeContext()
}

type Aux_matrix_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAux_matrix_typeContext() *Aux_matrix_typeContext {
	var p = new(Aux_matrix_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_aux_matrix_type
	return p
}

func InitEmptyAux_matrix_typeContext(p *Aux_matrix_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_aux_matrix_type
}

func (*Aux_matrix_typeContext) IsAux_matrix_typeContext() {}

func NewAux_matrix_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Aux_matrix_typeContext {
	var p = new(Aux_matrix_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_aux_matrix_type

	return p
}

func (s *Aux_matrix_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Aux_matrix_typeContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(VLangGrammarLBRACK, 0)
}

func (s *Aux_matrix_typeContext) Matrix_type() IMatrix_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrix_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrix_typeContext)
}

func (s *Aux_matrix_typeContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(VLangGrammarRBRACK, 0)
}

func (s *Aux_matrix_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Aux_matrix_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Aux_matrix_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterAux_matrix_type(s)
	}
}

func (s *Aux_matrix_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitAux_matrix_type(s)
	}
}

func (s *Aux_matrix_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitAux_matrix_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Aux_matrix_type() (localctx IAux_matrix_typeContext) {
	localctx = NewAux_matrix_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, VLangGrammarRULE_aux_matrix_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		p.Match(VLangGrammarLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(254)
		p.Matrix_type()
	}
	{
		p.SetState(255)
		p.Match(VLangGrammarRBRACK)
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

// IIf_stmtContext is an interface to support dynamic dispatch.
type IIf_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIf_stmtContext differentiates from other interfaces.
	IsIf_stmtContext()
}

type If_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_stmtContext() *If_stmtContext {
	var p = new(If_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_if_stmt
	return p
}

func InitEmptyIf_stmtContext(p *If_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_if_stmt
}

func (*If_stmtContext) IsIf_stmtContext() {}

func NewIf_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_stmtContext {
	var p = new(If_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_if_stmt

	return p
}

func (s *If_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *If_stmtContext) CopyAll(ctx *If_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *If_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfStmtContext struct {
	If_stmtContext
}

func NewIfStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStmtContext {
	var p = new(IfStmtContext)

	InitEmptyIf_stmtContext(&p.If_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_stmtContext))

	return p
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) AllIf_chain() []IIf_chainContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIf_chainContext); ok {
			len++
		}
	}

	tst := make([]IIf_chainContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIf_chainContext); ok {
			tst[i] = t.(IIf_chainContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) If_chain(i int) IIf_chainContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_chainContext); ok {
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

	return t.(IIf_chainContext)
}

func (s *IfStmtContext) AllELSE_KW() []antlr.TerminalNode {
	return s.GetTokens(VLangGrammarELSE_KW)
}

func (s *IfStmtContext) ELSE_KW(i int) antlr.TerminalNode {
	return s.GetToken(VLangGrammarELSE_KW, i)
}

func (s *IfStmtContext) Else_stmt() IElse_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElse_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElse_stmtContext)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) If_stmt() (localctx IIf_stmtContext) {
	localctx = NewIf_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, VLangGrammarRULE_if_stmt)
	var _la int

	var _alt int

	localctx = NewIfStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		p.If_chain()
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(258)
				p.Match(VLangGrammarELSE_KW)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(259)
				p.If_chain()
			}

		}
		p.SetState(264)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VLangGrammarELSE_KW {
		{
			p.SetState(265)
			p.Else_stmt()
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

// IIf_chainContext is an interface to support dynamic dispatch.
type IIf_chainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIf_chainContext differentiates from other interfaces.
	IsIf_chainContext()
}

type If_chainContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_chainContext() *If_chainContext {
	var p = new(If_chainContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_if_chain
	return p
}

func InitEmptyIf_chainContext(p *If_chainContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_if_chain
}

func (*If_chainContext) IsIf_chainContext() {}

func NewIf_chainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_chainContext {
	var p = new(If_chainContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_if_chain

	return p
}

func (s *If_chainContext) GetParser() antlr.Parser { return s.parser }

func (s *If_chainContext) CopyAll(ctx *If_chainContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *If_chainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_chainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfChainContext struct {
	If_chainContext
}

func NewIfChainContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfChainContext {
	var p = new(IfChainContext)

	InitEmptyIf_chainContext(&p.If_chainContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_chainContext))

	return p
}

func (s *IfChainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfChainContext) IF_KW() antlr.TerminalNode {
	return s.GetToken(VLangGrammarIF_KW, 0)
}

func (s *IfChainContext) Expression() IExpressionContext {
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

func (s *IfChainContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarLBRACE, 0)
}

func (s *IfChainContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarRBRACE, 0)
}

func (s *IfChainContext) AllStmt() []IStmtContext {
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

func (s *IfChainContext) Stmt(i int) IStmtContext {
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

func (s *IfChainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterIfChain(s)
	}
}

func (s *IfChainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitIfChain(s)
	}
}

func (s *IfChainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitIfChain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) If_chain() (localctx IIf_chainContext) {
	localctx = NewIf_chainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, VLangGrammarRULE_if_chain)
	var _la int

	localctx = NewIfChainContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Match(VLangGrammarIF_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)
		p.expression(0)
	}
	{
		p.SetState(270)
		p.Match(VLangGrammarLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36028797019022510) != 0 {
		{
			p.SetState(271)
			p.Stmt()
		}

		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(277)
		p.Match(VLangGrammarRBRACE)
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

// IElse_stmtContext is an interface to support dynamic dispatch.
type IElse_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsElse_stmtContext differentiates from other interfaces.
	IsElse_stmtContext()
}

type Else_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElse_stmtContext() *Else_stmtContext {
	var p = new(Else_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_else_stmt
	return p
}

func InitEmptyElse_stmtContext(p *Else_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_else_stmt
}

func (*Else_stmtContext) IsElse_stmtContext() {}

func NewElse_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_stmtContext {
	var p = new(Else_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_else_stmt

	return p
}

func (s *Else_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_stmtContext) CopyAll(ctx *Else_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Else_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ElseStmtContext struct {
	Else_stmtContext
}

func NewElseStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElseStmtContext {
	var p = new(ElseStmtContext)

	InitEmptyElse_stmtContext(&p.Else_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Else_stmtContext))

	return p
}

func (s *ElseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStmtContext) ELSE_KW() antlr.TerminalNode {
	return s.GetToken(VLangGrammarELSE_KW, 0)
}

func (s *ElseStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarLBRACE, 0)
}

func (s *ElseStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarRBRACE, 0)
}

func (s *ElseStmtContext) AllStmt() []IStmtContext {
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

func (s *ElseStmtContext) Stmt(i int) IStmtContext {
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

func (s *ElseStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterElseStmt(s)
	}
}

func (s *ElseStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitElseStmt(s)
	}
}

func (s *ElseStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitElseStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Else_stmt() (localctx IElse_stmtContext) {
	localctx = NewElse_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, VLangGrammarRULE_else_stmt)
	var _la int

	localctx = NewElseStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.Match(VLangGrammarELSE_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(280)
		p.Match(VLangGrammarLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36028797019022510) != 0 {
		{
			p.SetState(281)
			p.Stmt()
		}

		p.SetState(286)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(287)
		p.Match(VLangGrammarRBRACE)
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

// ISwitch_stmtContext is an interface to support dynamic dispatch.
type ISwitch_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSwitch_stmtContext differentiates from other interfaces.
	IsSwitch_stmtContext()
}

type Switch_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_stmtContext() *Switch_stmtContext {
	var p = new(Switch_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_switch_stmt
	return p
}

func InitEmptySwitch_stmtContext(p *Switch_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_switch_stmt
}

func (*Switch_stmtContext) IsSwitch_stmtContext() {}

func NewSwitch_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_stmtContext {
	var p = new(Switch_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_switch_stmt

	return p
}

func (s *Switch_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_stmtContext) CopyAll(ctx *Switch_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Switch_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SwitchStmtContext struct {
	Switch_stmtContext
}

func NewSwitchStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchStmtContext {
	var p = new(SwitchStmtContext)

	InitEmptySwitch_stmtContext(&p.Switch_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Switch_stmtContext))

	return p
}

func (s *SwitchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStmtContext) SWITCH_KW() antlr.TerminalNode {
	return s.GetToken(VLangGrammarSWITCH_KW, 0)
}

func (s *SwitchStmtContext) Expression() IExpressionContext {
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

func (s *SwitchStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarLBRACE, 0)
}

func (s *SwitchStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarRBRACE, 0)
}

func (s *SwitchStmtContext) AllSwitch_case() []ISwitch_caseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISwitch_caseContext); ok {
			len++
		}
	}

	tst := make([]ISwitch_caseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISwitch_caseContext); ok {
			tst[i] = t.(ISwitch_caseContext)
			i++
		}
	}

	return tst
}

func (s *SwitchStmtContext) Switch_case(i int) ISwitch_caseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_caseContext); ok {
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

	return t.(ISwitch_caseContext)
}

func (s *SwitchStmtContext) Default_case() IDefault_caseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefault_caseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefault_caseContext)
}

func (s *SwitchStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitSwitchStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Switch_stmt() (localctx ISwitch_stmtContext) {
	localctx = NewSwitch_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, VLangGrammarRULE_switch_stmt)
	var _la int

	localctx = NewSwitchStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		p.Match(VLangGrammarSWITCH_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(290)
		p.expression(0)
	}
	{
		p.SetState(291)
		p.Match(VLangGrammarLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VLangGrammarCASE_KW {
		{
			p.SetState(292)
			p.Switch_case()
		}

		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VLangGrammarDEFAULT_KW {
		{
			p.SetState(298)
			p.Default_case()
		}

	}
	{
		p.SetState(301)
		p.Match(VLangGrammarRBRACE)
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

// ISwitch_caseContext is an interface to support dynamic dispatch.
type ISwitch_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSwitch_caseContext differentiates from other interfaces.
	IsSwitch_caseContext()
}

type Switch_caseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_caseContext() *Switch_caseContext {
	var p = new(Switch_caseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_switch_case
	return p
}

func InitEmptySwitch_caseContext(p *Switch_caseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_switch_case
}

func (*Switch_caseContext) IsSwitch_caseContext() {}

func NewSwitch_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_caseContext {
	var p = new(Switch_caseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_switch_case

	return p
}

func (s *Switch_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_caseContext) CopyAll(ctx *Switch_caseContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Switch_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SwitchCaseContext struct {
	Switch_caseContext
}

func NewSwitchCaseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchCaseContext {
	var p = new(SwitchCaseContext)

	InitEmptySwitch_caseContext(&p.Switch_caseContext)
	p.parser = parser
	p.CopyAll(ctx.(*Switch_caseContext))

	return p
}

func (s *SwitchCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchCaseContext) CASE_KW() antlr.TerminalNode {
	return s.GetToken(VLangGrammarCASE_KW, 0)
}

func (s *SwitchCaseContext) Expression() IExpressionContext {
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

func (s *SwitchCaseContext) COLON() antlr.TerminalNode {
	return s.GetToken(VLangGrammarCOLON, 0)
}

func (s *SwitchCaseContext) AllStmt() []IStmtContext {
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

func (s *SwitchCaseContext) Stmt(i int) IStmtContext {
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

func (s *SwitchCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterSwitchCase(s)
	}
}

func (s *SwitchCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitSwitchCase(s)
	}
}

func (s *SwitchCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitSwitchCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Switch_case() (localctx ISwitch_caseContext) {
	localctx = NewSwitch_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, VLangGrammarRULE_switch_case)
	var _la int

	localctx = NewSwitchCaseContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
		p.Match(VLangGrammarCASE_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(304)
		p.expression(0)
	}
	{
		p.SetState(305)
		p.Match(VLangGrammarCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36028797019022510) != 0 {
		{
			p.SetState(306)
			p.Stmt()
		}

		p.SetState(311)
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

// IDefault_caseContext is an interface to support dynamic dispatch.
type IDefault_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDefault_caseContext differentiates from other interfaces.
	IsDefault_caseContext()
}

type Default_caseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefault_caseContext() *Default_caseContext {
	var p = new(Default_caseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_default_case
	return p
}

func InitEmptyDefault_caseContext(p *Default_caseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_default_case
}

func (*Default_caseContext) IsDefault_caseContext() {}

func NewDefault_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Default_caseContext {
	var p = new(Default_caseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_default_case

	return p
}

func (s *Default_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Default_caseContext) CopyAll(ctx *Default_caseContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Default_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Default_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DefaultCaseContext struct {
	Default_caseContext
}

func NewDefaultCaseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefaultCaseContext {
	var p = new(DefaultCaseContext)

	InitEmptyDefault_caseContext(&p.Default_caseContext)
	p.parser = parser
	p.CopyAll(ctx.(*Default_caseContext))

	return p
}

func (s *DefaultCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultCaseContext) DEFAULT_KW() antlr.TerminalNode {
	return s.GetToken(VLangGrammarDEFAULT_KW, 0)
}

func (s *DefaultCaseContext) COLON() antlr.TerminalNode {
	return s.GetToken(VLangGrammarCOLON, 0)
}

func (s *DefaultCaseContext) AllStmt() []IStmtContext {
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

func (s *DefaultCaseContext) Stmt(i int) IStmtContext {
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

func (s *DefaultCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterDefaultCase(s)
	}
}

func (s *DefaultCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitDefaultCase(s)
	}
}

func (s *DefaultCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitDefaultCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Default_case() (localctx IDefault_caseContext) {
	localctx = NewDefault_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, VLangGrammarRULE_default_case)
	var _la int

	localctx = NewDefaultCaseContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)
		p.Match(VLangGrammarDEFAULT_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(313)
		p.Match(VLangGrammarCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36028797019022510) != 0 {
		{
			p.SetState(314)
			p.Stmt()
		}

		p.SetState(319)
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

// IWhile_stmtContext is an interface to support dynamic dispatch.
type IWhile_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsWhile_stmtContext differentiates from other interfaces.
	IsWhile_stmtContext()
}

type While_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_stmtContext() *While_stmtContext {
	var p = new(While_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_while_stmt
	return p
}

func InitEmptyWhile_stmtContext(p *While_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_while_stmt
}

func (*While_stmtContext) IsWhile_stmtContext() {}

func NewWhile_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_stmtContext {
	var p = new(While_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_while_stmt

	return p
}

func (s *While_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *While_stmtContext) CopyAll(ctx *While_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *While_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type WhileStmtContext struct {
	While_stmtContext
}

func NewWhileStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileStmtContext {
	var p = new(WhileStmtContext)

	InitEmptyWhile_stmtContext(&p.While_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*While_stmtContext))

	return p
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) WHILE_KW() antlr.TerminalNode {
	return s.GetToken(VLangGrammarWHILE_KW, 0)
}

func (s *WhileStmtContext) Expression() IExpressionContext {
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

func (s *WhileStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarLBRACE, 0)
}

func (s *WhileStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarRBRACE, 0)
}

func (s *WhileStmtContext) AllStmt() []IStmtContext {
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

func (s *WhileStmtContext) Stmt(i int) IStmtContext {
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

func (s *WhileStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterWhileStmt(s)
	}
}

func (s *WhileStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitWhileStmt(s)
	}
}

func (s *WhileStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitWhileStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) While_stmt() (localctx IWhile_stmtContext) {
	localctx = NewWhile_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, VLangGrammarRULE_while_stmt)
	var _la int

	localctx = NewWhileStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		p.Match(VLangGrammarWHILE_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(321)
		p.expression(0)
	}
	{
		p.SetState(322)
		p.Match(VLangGrammarLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36028797019022510) != 0 {
		{
			p.SetState(323)
			p.Stmt()
		}

		p.SetState(328)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(329)
		p.Match(VLangGrammarRBRACE)
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

// IFor_stmtContext is an interface to support dynamic dispatch.
type IFor_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFor_stmtContext differentiates from other interfaces.
	IsFor_stmtContext()
}

type For_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_stmtContext() *For_stmtContext {
	var p = new(For_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_for_stmt
	return p
}

func InitEmptyFor_stmtContext(p *For_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_for_stmt
}

func (*For_stmtContext) IsFor_stmtContext() {}

func NewFor_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_stmtContext {
	var p = new(For_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_for_stmt

	return p
}

func (s *For_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *For_stmtContext) CopyAll(ctx *For_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *For_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForStmtCondContext struct {
	For_stmtContext
}

func NewForStmtCondContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForStmtCondContext {
	var p = new(ForStmtCondContext)

	InitEmptyFor_stmtContext(&p.For_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*For_stmtContext))

	return p
}

func (s *ForStmtCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtCondContext) FOR_KW() antlr.TerminalNode {
	return s.GetToken(VLangGrammarFOR_KW, 0)
}

func (s *ForStmtCondContext) Expression() IExpressionContext {
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

func (s *ForStmtCondContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarLBRACE, 0)
}

func (s *ForStmtCondContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarRBRACE, 0)
}

func (s *ForStmtCondContext) AllStmt() []IStmtContext {
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

func (s *ForStmtCondContext) Stmt(i int) IStmtContext {
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

func (s *ForStmtCondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterForStmtCond(s)
	}
}

func (s *ForStmtCondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitForStmtCond(s)
	}
}

func (s *ForStmtCondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitForStmtCond(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForAssCondContext struct {
	For_stmtContext
}

func NewForAssCondContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForAssCondContext {
	var p = new(ForAssCondContext)

	InitEmptyFor_stmtContext(&p.For_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*For_stmtContext))

	return p
}

func (s *ForAssCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForAssCondContext) FOR_KW() antlr.TerminalNode {
	return s.GetToken(VLangGrammarFOR_KW, 0)
}

func (s *ForAssCondContext) AllAssign_stmt() []IAssign_stmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssign_stmtContext); ok {
			len++
		}
	}

	tst := make([]IAssign_stmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssign_stmtContext); ok {
			tst[i] = t.(IAssign_stmtContext)
			i++
		}
	}

	return tst
}

func (s *ForAssCondContext) Assign_stmt(i int) IAssign_stmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_stmtContext); ok {
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

	return t.(IAssign_stmtContext)
}

func (s *ForAssCondContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(VLangGrammarSEMI)
}

func (s *ForAssCondContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(VLangGrammarSEMI, i)
}

func (s *ForAssCondContext) Expression() IExpressionContext {
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

func (s *ForAssCondContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarLBRACE, 0)
}

func (s *ForAssCondContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarRBRACE, 0)
}

func (s *ForAssCondContext) AllStmt() []IStmtContext {
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

func (s *ForAssCondContext) Stmt(i int) IStmtContext {
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

func (s *ForAssCondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterForAssCond(s)
	}
}

func (s *ForAssCondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitForAssCond(s)
	}
}

func (s *ForAssCondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitForAssCond(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForStmtContext struct {
	For_stmtContext
}

func NewForStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForStmtContext {
	var p = new(ForStmtContext)

	InitEmptyFor_stmtContext(&p.For_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*For_stmtContext))

	return p
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) FOR_KW() antlr.TerminalNode {
	return s.GetToken(VLangGrammarFOR_KW, 0)
}

func (s *ForStmtContext) ID() antlr.TerminalNode {
	return s.GetToken(VLangGrammarID, 0)
}

func (s *ForStmtContext) COMMA() antlr.TerminalNode {
	return s.GetToken(VLangGrammarCOMMA, 0)
}

func (s *ForStmtContext) AllExpression() []IExpressionContext {
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

func (s *ForStmtContext) Expression(i int) IExpressionContext {
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

func (s *ForStmtContext) IN_KW() antlr.TerminalNode {
	return s.GetToken(VLangGrammarIN_KW, 0)
}

func (s *ForStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarLBRACE, 0)
}

func (s *ForStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarRBRACE, 0)
}

func (s *ForStmtContext) Range_() IRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeContext)
}

func (s *ForStmtContext) AllStmt() []IStmtContext {
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

func (s *ForStmtContext) Stmt(i int) IStmtContext {
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

func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterForStmt(s)
	}
}

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitForStmt(s)
	}
}

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitForStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) For_stmt() (localctx IFor_stmtContext) {
	localctx = NewFor_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, VLangGrammarRULE_for_stmt)
	var _la int

	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForStmtCondContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(331)
			p.Match(VLangGrammarFOR_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.expression(0)
		}
		{
			p.SetState(333)
			p.Match(VLangGrammarLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(337)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36028797019022510) != 0 {
			{
				p.SetState(334)
				p.Stmt()
			}

			p.SetState(339)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(340)
			p.Match(VLangGrammarRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewForAssCondContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(342)
			p.Match(VLangGrammarFOR_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			p.Assign_stmt()
		}
		{
			p.SetState(344)
			p.Match(VLangGrammarSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(345)
			p.expression(0)
		}
		{
			p.SetState(346)
			p.Match(VLangGrammarSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(347)
			p.Assign_stmt()
		}
		{
			p.SetState(348)
			p.Match(VLangGrammarLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(352)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36028797019022510) != 0 {
			{
				p.SetState(349)
				p.Stmt()
			}

			p.SetState(354)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(355)
			p.Match(VLangGrammarRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewForStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(357)
			p.Match(VLangGrammarFOR_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(358)
			p.Match(VLangGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(359)
			p.Match(VLangGrammarCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(360)
			p.expression(0)
		}
		{
			p.SetState(361)
			p.Match(VLangGrammarIN_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(364)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(362)
				p.expression(0)
			}

		case 2:
			{
				p.SetState(363)
				p.Range_()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		{
			p.SetState(366)
			p.Match(VLangGrammarLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(370)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36028797019022510) != 0 {
			{
				p.SetState(367)
				p.Stmt()
			}

			p.SetState(372)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(373)
			p.Match(VLangGrammarRBRACE)
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

// IRangeContext is an interface to support dynamic dispatch.
type IRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRangeContext differentiates from other interfaces.
	IsRangeContext()
}

type RangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeContext() *RangeContext {
	var p = new(RangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_range
	return p
}

func InitEmptyRangeContext(p *RangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_range
}

func (*RangeContext) IsRangeContext() {}

func NewRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeContext {
	var p = new(RangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_range

	return p
}

func (s *RangeContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeContext) CopyAll(ctx *RangeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NumericRangeContext struct {
	RangeContext
}

func NewNumericRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumericRangeContext {
	var p = new(NumericRangeContext)

	InitEmptyRangeContext(&p.RangeContext)
	p.parser = parser
	p.CopyAll(ctx.(*RangeContext))

	return p
}

func (s *NumericRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericRangeContext) AllExpression() []IExpressionContext {
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

func (s *NumericRangeContext) Expression(i int) IExpressionContext {
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

func (s *NumericRangeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(VLangGrammarDOT)
}

func (s *NumericRangeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(VLangGrammarDOT, i)
}

func (s *NumericRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterNumericRange(s)
	}
}

func (s *NumericRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitNumericRange(s)
	}
}

func (s *NumericRangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitNumericRange(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Range_() (localctx IRangeContext) {
	localctx = NewRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, VLangGrammarRULE_range)
	localctx = NewNumericRangeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(377)
		p.expression(0)
	}
	{
		p.SetState(378)
		p.Match(VLangGrammarDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(379)
		p.Match(VLangGrammarDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(380)
		p.Match(VLangGrammarDOT)
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

// ITransfer_stmtContext is an interface to support dynamic dispatch.
type ITransfer_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTransfer_stmtContext differentiates from other interfaces.
	IsTransfer_stmtContext()
}

type Transfer_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransfer_stmtContext() *Transfer_stmtContext {
	var p = new(Transfer_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_transfer_stmt
	return p
}

func InitEmptyTransfer_stmtContext(p *Transfer_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_transfer_stmt
}

func (*Transfer_stmtContext) IsTransfer_stmtContext() {}

func NewTransfer_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Transfer_stmtContext {
	var p = new(Transfer_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_transfer_stmt

	return p
}

func (s *Transfer_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Transfer_stmtContext) CopyAll(ctx *Transfer_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Transfer_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Transfer_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ContinueStmtContext struct {
	Transfer_stmtContext
}

func NewContinueStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	InitEmptyTransfer_stmtContext(&p.Transfer_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Transfer_stmtContext))

	return p
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) CONTINUE_KW() antlr.TerminalNode {
	return s.GetToken(VLangGrammarCONTINUE_KW, 0)
}

func (s *ContinueStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterContinueStmt(s)
	}
}

func (s *ContinueStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitContinueStmt(s)
	}
}

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type BreakStmtContext struct {
	Transfer_stmtContext
}

func NewBreakStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStmtContext {
	var p = new(BreakStmtContext)

	InitEmptyTransfer_stmtContext(&p.Transfer_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Transfer_stmtContext))

	return p
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) BREAK_KW() antlr.TerminalNode {
	return s.GetToken(VLangGrammarBREAK_KW, 0)
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStmtContext struct {
	Transfer_stmtContext
}

func NewReturnStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	InitEmptyTransfer_stmtContext(&p.Transfer_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Transfer_stmtContext))

	return p
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) RETURN_KW() antlr.TerminalNode {
	return s.GetToken(VLangGrammarRETURN_KW, 0)
}

func (s *ReturnStmtContext) Expression() IExpressionContext {
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

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Transfer_stmt() (localctx ITransfer_stmtContext) {
	localctx = NewTransfer_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, VLangGrammarRULE_transfer_stmt)
	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VLangGrammarRETURN_KW:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(383)
			p.Match(VLangGrammarRETURN_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(385)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(384)
				p.expression(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case VLangGrammarBREAK_KW:
		localctx = NewBreakStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(387)
			p.Match(VLangGrammarBREAK_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VLangGrammarCONTINUE_KW:
		localctx = NewContinueStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(388)
			p.Match(VLangGrammarCONTINUE_KW)
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

// IFunc_callContext is an interface to support dynamic dispatch.
type IFunc_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunc_callContext differentiates from other interfaces.
	IsFunc_callContext()
}

type Func_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_callContext() *Func_callContext {
	var p = new(Func_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_func_call
	return p
}

func InitEmptyFunc_callContext(p *Func_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_func_call
}

func (*Func_callContext) IsFunc_callContext() {}

func NewFunc_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_callContext {
	var p = new(Func_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_func_call

	return p
}

func (s *Func_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_callContext) CopyAll(ctx *Func_callContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Func_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncCallContext struct {
	Func_callContext
}

func NewFuncCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncCallContext {
	var p = new(FuncCallContext)

	InitEmptyFunc_callContext(&p.Func_callContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_callContext))

	return p
}

func (s *FuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *FuncCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarLPAREN, 0)
}

func (s *FuncCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarRPAREN, 0)
}

func (s *FuncCallContext) Arg_list() IArg_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArg_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArg_listContext)
}

func (s *FuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterFuncCall(s)
	}
}

func (s *FuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitFuncCall(s)
	}
}

func (s *FuncCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitFuncCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Func_call() (localctx IFunc_callContext) {
	localctx = NewFunc_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, VLangGrammarRULE_func_call)
	var _la int

	localctx = NewFuncCallContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(391)
		p.Id_pattern()
	}
	{
		p.SetState(392)
		p.Match(VLangGrammarLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&69247654638387202) != 0 {
		{
			p.SetState(393)
			p.Arg_list()
		}

	}
	{
		p.SetState(396)
		p.Match(VLangGrammarRPAREN)
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

// IArg_listContext is an interface to support dynamic dispatch.
type IArg_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArg_listContext differentiates from other interfaces.
	IsArg_listContext()
}

type Arg_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArg_listContext() *Arg_listContext {
	var p = new(Arg_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_arg_list
	return p
}

func InitEmptyArg_listContext(p *Arg_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_arg_list
}

func (*Arg_listContext) IsArg_listContext() {}

func NewArg_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arg_listContext {
	var p = new(Arg_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_arg_list

	return p
}

func (s *Arg_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Arg_listContext) CopyAll(ctx *Arg_listContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Arg_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arg_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArgListContext struct {
	Arg_listContext
}

func NewArgListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArgListContext {
	var p = new(ArgListContext)

	InitEmptyArg_listContext(&p.Arg_listContext)
	p.parser = parser
	p.CopyAll(ctx.(*Arg_listContext))

	return p
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) AllFunc_arg() []IFunc_argContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunc_argContext); ok {
			len++
		}
	}

	tst := make([]IFunc_argContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunc_argContext); ok {
			tst[i] = t.(IFunc_argContext)
			i++
		}
	}

	return tst
}

func (s *ArgListContext) Func_arg(i int) IFunc_argContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_argContext); ok {
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

	return t.(IFunc_argContext)
}

func (s *ArgListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VLangGrammarCOMMA)
}

func (s *ArgListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VLangGrammarCOMMA, i)
}

func (s *ArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterArgList(s)
	}
}

func (s *ArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitArgList(s)
	}
}

func (s *ArgListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitArgList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Arg_list() (localctx IArg_listContext) {
	localctx = NewArg_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, VLangGrammarRULE_arg_list)
	var _la int

	localctx = NewArgListContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(398)
		p.Func_arg()
	}
	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VLangGrammarCOMMA {
		{
			p.SetState(399)
			p.Match(VLangGrammarCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(400)
			p.Func_arg()
		}

		p.SetState(405)
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

// IFunc_argContext is an interface to support dynamic dispatch.
type IFunc_argContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunc_argContext differentiates from other interfaces.
	IsFunc_argContext()
}

type Func_argContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_argContext() *Func_argContext {
	var p = new(Func_argContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_func_arg
	return p
}

func InitEmptyFunc_argContext(p *Func_argContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_func_arg
}

func (*Func_argContext) IsFunc_argContext() {}

func NewFunc_argContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_argContext {
	var p = new(Func_argContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_func_arg

	return p
}

func (s *Func_argContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_argContext) CopyAll(ctx *Func_argContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Func_argContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_argContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncArgContext struct {
	Func_argContext
}

func NewFuncArgContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncArgContext {
	var p = new(FuncArgContext)

	InitEmptyFunc_argContext(&p.Func_argContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_argContext))

	return p
}

func (s *FuncArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *FuncArgContext) Expression() IExpressionContext {
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

func (s *FuncArgContext) ID() antlr.TerminalNode {
	return s.GetToken(VLangGrammarID, 0)
}

func (s *FuncArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterFuncArg(s)
	}
}

func (s *FuncArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitFuncArg(s)
	}
}

func (s *FuncArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitFuncArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Func_arg() (localctx IFunc_argContext) {
	localctx = NewFunc_argContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, VLangGrammarRULE_func_arg)
	localctx = NewFuncArgContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(407)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(406)
			p.Match(VLangGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(411)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(409)
			p.Id_pattern()
		}

	case 2:
		{
			p.SetState(410)
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

// IFunc_dclContext is an interface to support dynamic dispatch.
type IFunc_dclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunc_dclContext differentiates from other interfaces.
	IsFunc_dclContext()
}

type Func_dclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_dclContext() *Func_dclContext {
	var p = new(Func_dclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_func_dcl
	return p
}

func InitEmptyFunc_dclContext(p *Func_dclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_func_dcl
}

func (*Func_dclContext) IsFunc_dclContext() {}

func NewFunc_dclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_dclContext {
	var p = new(Func_dclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_func_dcl

	return p
}

func (s *Func_dclContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_dclContext) CopyAll(ctx *Func_dclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Func_dclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_dclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncDeclContext struct {
	Func_dclContext
}

func NewFuncDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncDeclContext {
	var p = new(FuncDeclContext)

	InitEmptyFunc_dclContext(&p.Func_dclContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_dclContext))

	return p
}

func (s *FuncDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclContext) FUNC() antlr.TerminalNode {
	return s.GetToken(VLangGrammarFUNC, 0)
}

func (s *FuncDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(VLangGrammarID, 0)
}

func (s *FuncDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarLPAREN, 0)
}

func (s *FuncDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarRPAREN, 0)
}

func (s *FuncDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarLBRACE, 0)
}

func (s *FuncDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarRBRACE, 0)
}

func (s *FuncDeclContext) Param_list() IParam_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParam_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParam_listContext)
}

func (s *FuncDeclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FuncDeclContext) AllStmt() []IStmtContext {
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

func (s *FuncDeclContext) Stmt(i int) IStmtContext {
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

func (s *FuncDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterFuncDecl(s)
	}
}

func (s *FuncDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitFuncDecl(s)
	}
}

func (s *FuncDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitFuncDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Func_dcl() (localctx IFunc_dclContext) {
	localctx = NewFunc_dclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, VLangGrammarRULE_func_dcl)
	var _la int

	localctx = NewFuncDeclContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(413)
		p.Match(VLangGrammarFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(414)
		p.Match(VLangGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(415)
		p.Match(VLangGrammarLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(417)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VLangGrammarID {
		{
			p.SetState(416)
			p.Param_list()
		}

	}
	{
		p.SetState(419)
		p.Match(VLangGrammarRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(421)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(420)
			p.Type_()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(423)
		p.Match(VLangGrammarLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36028797019022510) != 0 {
		{
			p.SetState(424)
			p.Stmt()
		}

		p.SetState(429)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(430)
		p.Match(VLangGrammarRBRACE)
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

// IParam_listContext is an interface to support dynamic dispatch.
type IParam_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsParam_listContext differentiates from other interfaces.
	IsParam_listContext()
}

type Param_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParam_listContext() *Param_listContext {
	var p = new(Param_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_param_list
	return p
}

func InitEmptyParam_listContext(p *Param_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_param_list
}

func (*Param_listContext) IsParam_listContext() {}

func NewParam_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Param_listContext {
	var p = new(Param_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_param_list

	return p
}

func (s *Param_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Param_listContext) CopyAll(ctx *Param_listContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Param_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Param_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParamListContext struct {
	Param_listContext
}

func NewParamListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParamListContext {
	var p = new(ParamListContext)

	InitEmptyParam_listContext(&p.Param_listContext)
	p.parser = parser
	p.CopyAll(ctx.(*Param_listContext))

	return p
}

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) AllFunc_param() []IFunc_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunc_paramContext); ok {
			len++
		}
	}

	tst := make([]IFunc_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunc_paramContext); ok {
			tst[i] = t.(IFunc_paramContext)
			i++
		}
	}

	return tst
}

func (s *ParamListContext) Func_param(i int) IFunc_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_paramContext); ok {
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

	return t.(IFunc_paramContext)
}

func (s *ParamListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VLangGrammarCOMMA)
}

func (s *ParamListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VLangGrammarCOMMA, i)
}

func (s *ParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterParamList(s)
	}
}

func (s *ParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitParamList(s)
	}
}

func (s *ParamListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitParamList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Param_list() (localctx IParam_listContext) {
	localctx = NewParam_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, VLangGrammarRULE_param_list)
	var _la int

	localctx = NewParamListContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(432)
		p.Func_param()
	}
	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VLangGrammarCOMMA {
		{
			p.SetState(433)
			p.Match(VLangGrammarCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(434)
			p.Func_param()
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

// IFunc_paramContext is an interface to support dynamic dispatch.
type IFunc_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunc_paramContext differentiates from other interfaces.
	IsFunc_paramContext()
}

type Func_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_paramContext() *Func_paramContext {
	var p = new(Func_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_func_param
	return p
}

func InitEmptyFunc_paramContext(p *Func_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_func_param
}

func (*Func_paramContext) IsFunc_paramContext() {}

func NewFunc_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_paramContext {
	var p = new(Func_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_func_param

	return p
}

func (s *Func_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_paramContext) CopyAll(ctx *Func_paramContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Func_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncParamContext struct {
	Func_paramContext
}

func NewFuncParamContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncParamContext {
	var p = new(FuncParamContext)

	InitEmptyFunc_paramContext(&p.Func_paramContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_paramContext))

	return p
}

func (s *FuncParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamContext) ID() antlr.TerminalNode {
	return s.GetToken(VLangGrammarID, 0)
}

func (s *FuncParamContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FuncParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterFuncParam(s)
	}
}

func (s *FuncParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitFuncParam(s)
	}
}

func (s *FuncParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitFuncParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Func_param() (localctx IFunc_paramContext) {
	localctx = NewFunc_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, VLangGrammarRULE_func_param)
	localctx = NewFuncParamContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(440)
		p.Match(VLangGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(441)
		p.Type_()
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

// IStrct_dclContext is an interface to support dynamic dispatch.
type IStrct_dclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStrct_dclContext differentiates from other interfaces.
	IsStrct_dclContext()
}

type Strct_dclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrct_dclContext() *Strct_dclContext {
	var p = new(Strct_dclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_strct_dcl
	return p
}

func InitEmptyStrct_dclContext(p *Strct_dclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_strct_dcl
}

func (*Strct_dclContext) IsStrct_dclContext() {}

func NewStrct_dclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Strct_dclContext {
	var p = new(Strct_dclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_strct_dcl

	return p
}

func (s *Strct_dclContext) GetParser() antlr.Parser { return s.parser }

func (s *Strct_dclContext) CopyAll(ctx *Strct_dclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Strct_dclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Strct_dclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructDeclContext struct {
	Strct_dclContext
}

func NewStructDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructDeclContext {
	var p = new(StructDeclContext)

	InitEmptyStrct_dclContext(&p.Strct_dclContext)
	p.parser = parser
	p.CopyAll(ctx.(*Strct_dclContext))

	return p
}

func (s *StructDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclContext) STR() antlr.TerminalNode {
	return s.GetToken(VLangGrammarSTR, 0)
}

func (s *StructDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(VLangGrammarID, 0)
}

func (s *StructDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarLBRACE, 0)
}

func (s *StructDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VLangGrammarRBRACE, 0)
}

func (s *StructDeclContext) AllStruct_prop() []IStruct_propContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStruct_propContext); ok {
			len++
		}
	}

	tst := make([]IStruct_propContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStruct_propContext); ok {
			tst[i] = t.(IStruct_propContext)
			i++
		}
	}

	return tst
}

func (s *StructDeclContext) Struct_prop(i int) IStruct_propContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_propContext); ok {
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

	return t.(IStruct_propContext)
}

func (s *StructDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterStructDecl(s)
	}
}

func (s *StructDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitStructDecl(s)
	}
}

func (s *StructDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitStructDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Strct_dcl() (localctx IStrct_dclContext) {
	localctx = NewStrct_dclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, VLangGrammarRULE_strct_dcl)
	var _la int

	localctx = NewStructDeclContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(443)
		p.Match(VLangGrammarSTR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(444)
		p.Match(VLangGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(445)
		p.Match(VLangGrammarLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(449)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VLangGrammarMUT {
		{
			p.SetState(446)
			p.Struct_prop()
		}

		p.SetState(451)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(452)
		p.Match(VLangGrammarRBRACE)
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

// IStruct_propContext is an interface to support dynamic dispatch.
type IStruct_propContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStruct_propContext differentiates from other interfaces.
	IsStruct_propContext()
}

type Struct_propContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_propContext() *Struct_propContext {
	var p = new(Struct_propContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_struct_prop
	return p
}

func InitEmptyStruct_propContext(p *Struct_propContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_struct_prop
}

func (*Struct_propContext) IsStruct_propContext() {}

func NewStruct_propContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_propContext {
	var p = new(Struct_propContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_struct_prop

	return p
}

func (s *Struct_propContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_propContext) CopyAll(ctx *Struct_propContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Struct_propContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_propContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructAttrContext struct {
	Struct_propContext
}

func NewStructAttrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructAttrContext {
	var p = new(StructAttrContext)

	InitEmptyStruct_propContext(&p.Struct_propContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_propContext))

	return p
}

func (s *StructAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructAttrContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *StructAttrContext) ID() antlr.TerminalNode {
	return s.GetToken(VLangGrammarID, 0)
}

func (s *StructAttrContext) COLON() antlr.TerminalNode {
	return s.GetToken(VLangGrammarCOLON, 0)
}

func (s *StructAttrContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *StructAttrContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarASSIGN, 0)
}

func (s *StructAttrContext) Expression() IExpressionContext {
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

func (s *StructAttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterStructAttr(s)
	}
}

func (s *StructAttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitStructAttr(s)
	}
}

func (s *StructAttrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitStructAttr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Struct_prop() (localctx IStruct_propContext) {
	localctx = NewStruct_propContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, VLangGrammarRULE_struct_prop)
	var _la int

	localctx = NewStructAttrContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(454)
		p.Var_type()
	}
	{
		p.SetState(455)
		p.Match(VLangGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VLangGrammarCOLON {
		{
			p.SetState(456)
			p.Match(VLangGrammarCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(457)
			p.Type_()
		}

	}
	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VLangGrammarASSIGN {
		{
			p.SetState(460)
			p.Match(VLangGrammarASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(461)
			p.expression(0)
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

// IStruct_vectorContext is an interface to support dynamic dispatch.
type IStruct_vectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStruct_vectorContext differentiates from other interfaces.
	IsStruct_vectorContext()
}

type Struct_vectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_vectorContext() *Struct_vectorContext {
	var p = new(Struct_vectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_struct_vector
	return p
}

func InitEmptyStruct_vectorContext(p *Struct_vectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VLangGrammarRULE_struct_vector
}

func (*Struct_vectorContext) IsStruct_vectorContext() {}

func NewStruct_vectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_vectorContext {
	var p = new(Struct_vectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VLangGrammarRULE_struct_vector

	return p
}

func (s *Struct_vectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_vectorContext) CopyAll(ctx *Struct_vectorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Struct_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_vectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructVectorContext struct {
	Struct_vectorContext
}

func NewStructVectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructVectorContext {
	var p = new(StructVectorContext)

	InitEmptyStruct_vectorContext(&p.Struct_vectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_vectorContext))

	return p
}

func (s *StructVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructVectorContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(VLangGrammarLBRACK, 0)
}

func (s *StructVectorContext) ID() antlr.TerminalNode {
	return s.GetToken(VLangGrammarID, 0)
}

func (s *StructVectorContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(VLangGrammarRBRACK, 0)
}

func (s *StructVectorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarLPAREN, 0)
}

func (s *StructVectorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VLangGrammarRPAREN, 0)
}

func (s *StructVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.EnterStructVector(s)
	}
}

func (s *StructVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VLangGrammarListener); ok {
		listenerT.ExitStructVector(s)
	}
}

func (s *StructVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VLangGrammarVisitor:
		return t.VisitStructVector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VLangGrammar) Struct_vector() (localctx IStruct_vectorContext) {
	localctx = NewStruct_vectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, VLangGrammarRULE_struct_vector)
	localctx = NewStructVectorContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(464)
		p.Match(VLangGrammarLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(465)
		p.Match(VLangGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(466)
		p.Match(VLangGrammarRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(467)
		p.Match(VLangGrammarLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(468)
		p.Match(VLangGrammarRPAREN)
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

func (p *VLangGrammar) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *VLangGrammar) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
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

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
