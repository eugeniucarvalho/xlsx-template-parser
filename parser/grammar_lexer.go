// Generated from Grammar.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 18, 136,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 7, 16, 90, 10,
	16, 12, 16, 14, 16, 93, 11, 16, 3, 17, 6, 17, 96, 10, 17, 13, 17, 14, 17,
	97, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 104, 10, 18, 12, 18, 14, 18, 107,
	11, 18, 3, 18, 5, 18, 110, 10, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 6,
	19, 117, 10, 19, 13, 19, 14, 19, 118, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20,
	3, 20, 7, 20, 127, 10, 20, 12, 20, 14, 20, 130, 11, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 4, 105, 128, 2, 21, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 2, 27, 2, 29, 2, 31, 14,
	33, 15, 35, 16, 37, 17, 39, 18, 3, 2, 6, 3, 2, 50, 59, 5, 2, 67, 92, 97,
	97, 99, 124, 3, 2, 67, 92, 5, 2, 11, 12, 15, 15, 34, 34, 2, 139, 2, 3,
	3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11,
	3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2,
	19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2,
	2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2,
	2, 3, 41, 3, 2, 2, 2, 5, 47, 3, 2, 2, 2, 7, 50, 3, 2, 2, 2, 9, 55, 3, 2,
	2, 2, 11, 61, 3, 2, 2, 2, 13, 63, 3, 2, 2, 2, 15, 65, 3, 2, 2, 2, 17, 67,
	3, 2, 2, 2, 19, 69, 3, 2, 2, 2, 21, 71, 3, 2, 2, 2, 23, 75, 3, 2, 2, 2,
	25, 80, 3, 2, 2, 2, 27, 82, 3, 2, 2, 2, 29, 84, 3, 2, 2, 2, 31, 86, 3,
	2, 2, 2, 33, 95, 3, 2, 2, 2, 35, 99, 3, 2, 2, 2, 37, 116, 3, 2, 2, 2, 39,
	122, 3, 2, 2, 2, 41, 42, 7, 116, 2, 2, 42, 43, 7, 99, 2, 2, 43, 44, 7,
	112, 2, 2, 44, 45, 7, 105, 2, 2, 45, 46, 7, 103, 2, 2, 46, 4, 3, 2, 2,
	2, 47, 48, 7, 107, 2, 2, 48, 49, 7, 104, 2, 2, 49, 6, 3, 2, 2, 2, 50, 51,
	7, 121, 2, 2, 51, 52, 7, 107, 2, 2, 52, 53, 7, 118, 2, 2, 53, 54, 7, 106,
	2, 2, 54, 8, 3, 2, 2, 2, 55, 56, 7, 100, 2, 2, 56, 57, 7, 110, 2, 2, 57,
	58, 7, 113, 2, 2, 58, 59, 7, 101, 2, 2, 59, 60, 7, 109, 2, 2, 60, 10, 3,
	2, 2, 2, 61, 62, 7, 40, 2, 2, 62, 12, 3, 2, 2, 2, 63, 64, 7, 125, 2, 2,
	64, 14, 3, 2, 2, 2, 65, 66, 7, 127, 2, 2, 66, 16, 3, 2, 2, 2, 67, 68, 7,
	126, 2, 2, 68, 18, 3, 2, 2, 2, 69, 70, 7, 48, 2, 2, 70, 20, 3, 2, 2, 2,
	71, 72, 7, 103, 2, 2, 72, 73, 7, 112, 2, 2, 73, 74, 7, 102, 2, 2, 74, 22,
	3, 2, 2, 2, 75, 76, 7, 117, 2, 2, 76, 77, 7, 118, 2, 2, 77, 78, 7, 113,
	2, 2, 78, 79, 7, 114, 2, 2, 79, 24, 3, 2, 2, 2, 80, 81, 9, 2, 2, 2, 81,
	26, 3, 2, 2, 2, 82, 83, 9, 3, 2, 2, 83, 28, 3, 2, 2, 2, 84, 85, 9, 4, 2,
	2, 85, 30, 3, 2, 2, 2, 86, 91, 5, 27, 14, 2, 87, 90, 5, 27, 14, 2, 88,
	90, 5, 25, 13, 2, 89, 87, 3, 2, 2, 2, 89, 88, 3, 2, 2, 2, 90, 93, 3, 2,
	2, 2, 91, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 32, 3, 2, 2, 2, 93, 91,
	3, 2, 2, 2, 94, 96, 5, 25, 13, 2, 95, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2,
	2, 97, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 34, 3, 2, 2, 2, 99, 100,
	7, 49, 2, 2, 100, 101, 7, 49, 2, 2, 101, 105, 3, 2, 2, 2, 102, 104, 11,
	2, 2, 2, 103, 102, 3, 2, 2, 2, 104, 107, 3, 2, 2, 2, 105, 106, 3, 2, 2,
	2, 105, 103, 3, 2, 2, 2, 106, 109, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 108,
	110, 7, 15, 2, 2, 109, 108, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 111,
	3, 2, 2, 2, 111, 112, 7, 12, 2, 2, 112, 113, 3, 2, 2, 2, 113, 114, 8, 18,
	2, 2, 114, 36, 3, 2, 2, 2, 115, 117, 9, 5, 2, 2, 116, 115, 3, 2, 2, 2,
	117, 118, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119,
	120, 3, 2, 2, 2, 120, 121, 8, 19, 2, 2, 121, 38, 3, 2, 2, 2, 122, 123,
	7, 49, 2, 2, 123, 124, 7, 44, 2, 2, 124, 128, 3, 2, 2, 2, 125, 127, 11,
	2, 2, 2, 126, 125, 3, 2, 2, 2, 127, 130, 3, 2, 2, 2, 128, 129, 3, 2, 2,
	2, 128, 126, 3, 2, 2, 2, 129, 131, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 131,
	132, 7, 44, 2, 2, 132, 133, 7, 49, 2, 2, 133, 134, 3, 2, 2, 2, 134, 135,
	8, 20, 2, 2, 135, 40, 3, 2, 2, 2, 10, 2, 89, 91, 97, 105, 109, 118, 128,
	3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'range'", "'if'", "'with'", "'block'", "'&'", "'{'", "'}'", "'|'",
	"'.'", "'end'", "'stop'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "R_AND", "R_BRACE_L", "R_BRACE_R", "R_PIPE", "R_DOT",
	"R_END", "R_STOP", "ID", "T_INTEIRO", "R_LINE_COMMENT", "R_WS", "R_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "R_AND", "R_BRACE_L", "R_BRACE_R", "R_PIPE",
	"R_DOT", "R_END", "R_STOP", "DIGITO", "CHARACTER", "CHARACTER_UP", "ID",
	"T_INTEIRO", "R_LINE_COMMENT", "R_WS", "R_COMMENT",
}

type GrammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewGrammarLexer(input antlr.CharStream) *GrammarLexer {

	l := new(GrammarLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Grammar.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GrammarLexer tokens.
const (
	GrammarLexerT__0           = 1
	GrammarLexerT__1           = 2
	GrammarLexerT__2           = 3
	GrammarLexerT__3           = 4
	GrammarLexerR_AND          = 5
	GrammarLexerR_BRACE_L      = 6
	GrammarLexerR_BRACE_R      = 7
	GrammarLexerR_PIPE         = 8
	GrammarLexerR_DOT          = 9
	GrammarLexerR_END          = 10
	GrammarLexerR_STOP         = 11
	GrammarLexerID             = 12
	GrammarLexerT_INTEIRO      = 13
	GrammarLexerR_LINE_COMMENT = 14
	GrammarLexerR_WS           = 15
	GrammarLexerR_COMMENT      = 16
)
