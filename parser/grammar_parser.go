// Generated from Grammar.g4 by ANTLR 4.7.

package parser // Grammar

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 19, 45, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 3, 5, 3, 19, 10, 3, 3, 3, 3, 3, 3, 3, 5, 3, 24, 10, 3,
	3, 4, 3, 4, 3, 5, 5, 5, 29, 10, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 35, 10,
	5, 12, 5, 14, 5, 38, 11, 5, 5, 5, 40, 10, 5, 3, 5, 5, 5, 43, 10, 5, 3,
	5, 2, 2, 6, 2, 4, 6, 8, 2, 3, 3, 2, 3, 6, 2, 47, 2, 10, 3, 2, 2, 2, 4,
	23, 3, 2, 2, 2, 6, 25, 3, 2, 2, 2, 8, 28, 3, 2, 2, 2, 10, 11, 7, 8, 2,
	2, 11, 12, 7, 8, 2, 2, 12, 13, 5, 4, 3, 2, 13, 14, 7, 9, 2, 2, 14, 15,
	7, 9, 2, 2, 15, 16, 7, 2, 2, 3, 16, 3, 3, 2, 2, 2, 17, 19, 5, 6, 4, 2,
	18, 17, 3, 2, 2, 2, 18, 19, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 24, 5,
	8, 5, 2, 21, 24, 7, 13, 2, 2, 22, 24, 7, 14, 2, 2, 23, 18, 3, 2, 2, 2,
	23, 21, 3, 2, 2, 2, 23, 22, 3, 2, 2, 2, 24, 5, 3, 2, 2, 2, 25, 26, 9, 2,
	2, 2, 26, 7, 3, 2, 2, 2, 27, 29, 7, 7, 2, 2, 28, 27, 3, 2, 2, 2, 28, 29,
	3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 39, 7, 12, 2, 2, 31, 36, 7, 16, 2,
	2, 32, 33, 7, 12, 2, 2, 33, 35, 7, 16, 2, 2, 34, 32, 3, 2, 2, 2, 35, 38,
	3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 40, 3, 2, 2, 2,
	38, 36, 3, 2, 2, 2, 39, 31, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 42, 3,
	2, 2, 2, 41, 43, 7, 11, 2, 2, 42, 41, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43,
	9, 3, 2, 2, 2, 8, 18, 23, 28, 36, 39, 42,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'range'", "'if'", "'with'", "'block'", "'&'", "'{'", "'}'", "'|'",
	"'>>>'", "'.'", "'end'", "'stop'", "'$'",
}
var symbolicNames = []string{
	"", "", "", "", "", "R_AND", "R_BRACE_L", "R_BRACE_R", "R_PIPE", "R_SPREAD",
	"R_DOT", "R_END", "R_STOP", "R_DOLAR", "ID", "R_LINE_COMMENT", "R_WS",
	"R_COMMENT",
}

var ruleNames = []string{
	"expression", "actionTypes", "command", "path",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type GrammarParser struct {
	*antlr.BaseParser
}

func NewGrammarParser(input antlr.TokenStream) *GrammarParser {
	this := new(GrammarParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Grammar.g4"

	return this
}

// GrammarParser tokens.
const (
	GrammarParserEOF            = antlr.TokenEOF
	GrammarParserT__0           = 1
	GrammarParserT__1           = 2
	GrammarParserT__2           = 3
	GrammarParserT__3           = 4
	GrammarParserR_AND          = 5
	GrammarParserR_BRACE_L      = 6
	GrammarParserR_BRACE_R      = 7
	GrammarParserR_PIPE         = 8
	GrammarParserR_SPREAD       = 9
	GrammarParserR_DOT          = 10
	GrammarParserR_END          = 11
	GrammarParserR_STOP         = 12
	GrammarParserR_DOLAR        = 13
	GrammarParserID             = 14
	GrammarParserR_LINE_COMMENT = 15
	GrammarParserR_WS           = 16
	GrammarParserR_COMMENT      = 17
)

// GrammarParser rules.
const (
	GrammarParserRULE_expression  = 0
	GrammarParserRULE_actionTypes = 1
	GrammarParserRULE_command     = 2
	GrammarParserRULE_path        = 3
)

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllR_BRACE_L() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserR_BRACE_L)
}

func (s *ExpressionContext) R_BRACE_L(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserR_BRACE_L, i)
}

func (s *ExpressionContext) ActionTypes() IActionTypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IActionTypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IActionTypesContext)
}

func (s *ExpressionContext) AllR_BRACE_R() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserR_BRACE_R)
}

func (s *ExpressionContext) R_BRACE_R(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserR_BRACE_R, i)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(GrammarParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *GrammarParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GrammarParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(8)
		p.Match(GrammarParserR_BRACE_L)
	}
	{
		p.SetState(9)
		p.Match(GrammarParserR_BRACE_L)
	}
	{
		p.SetState(10)
		p.ActionTypes()
	}
	{
		p.SetState(11)
		p.Match(GrammarParserR_BRACE_R)
	}
	{
		p.SetState(12)
		p.Match(GrammarParserR_BRACE_R)
	}
	{
		p.SetState(13)
		p.Match(GrammarParserEOF)
	}

	return localctx
}

// IActionTypesContext is an interface to support dynamic dispatch.
type IActionTypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActionTypesContext differentiates from other interfaces.
	IsActionTypesContext()
}

type ActionTypesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionTypesContext() *ActionTypesContext {
	var p = new(ActionTypesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_actionTypes
	return p
}

func (*ActionTypesContext) IsActionTypesContext() {}

func NewActionTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionTypesContext {
	var p = new(ActionTypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_actionTypes

	return p
}

func (s *ActionTypesContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionTypesContext) Path() IPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *ActionTypesContext) Command() ICommandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *ActionTypesContext) R_END() antlr.TerminalNode {
	return s.GetToken(GrammarParserR_END, 0)
}

func (s *ActionTypesContext) R_STOP() antlr.TerminalNode {
	return s.GetToken(GrammarParserR_STOP, 0)
}

func (s *ActionTypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionTypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionTypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterActionTypes(s)
	}
}

func (s *ActionTypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitActionTypes(s)
	}
}

func (p *GrammarParser) ActionTypes() (localctx IActionTypesContext) {
	localctx = NewActionTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GrammarParserRULE_actionTypes)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(21)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserT__0, GrammarParserT__1, GrammarParserT__2, GrammarParserT__3, GrammarParserR_AND, GrammarParserR_DOT:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(16)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GrammarParserT__0)|(1<<GrammarParserT__1)|(1<<GrammarParserT__2)|(1<<GrammarParserT__3))) != 0 {
			{
				p.SetState(15)
				p.Command()
			}

		}
		{
			p.SetState(18)
			p.Path()
		}

	case GrammarParserR_END:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(19)
			p.Match(GrammarParserR_END)
		}

	case GrammarParserR_STOP:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(20)
			p.Match(GrammarParserR_STOP)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_command
	return p
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }
func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterCommand(s)
	}
}

func (s *CommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitCommand(s)
	}
}

func (p *GrammarParser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GrammarParserRULE_command)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(23)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GrammarParserT__0)|(1<<GrammarParserT__1)|(1<<GrammarParserT__2)|(1<<GrammarParserT__3))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IPathContext is an interface to support dynamic dispatch.
type IPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPathContext differentiates from other interfaces.
	IsPathContext()
}

type PathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext {
	var p = new(PathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_path
	return p
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_path

	return p
}

func (s *PathContext) GetParser() antlr.Parser { return s.parser }

func (s *PathContext) AllR_DOT() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserR_DOT)
}

func (s *PathContext) R_DOT(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserR_DOT, i)
}

func (s *PathContext) R_AND() antlr.TerminalNode {
	return s.GetToken(GrammarParserR_AND, 0)
}

func (s *PathContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserID)
}

func (s *PathContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserID, i)
}

func (s *PathContext) R_SPREAD() antlr.TerminalNode {
	return s.GetToken(GrammarParserR_SPREAD, 0)
}

func (s *PathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterPath(s)
	}
}

func (s *PathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitPath(s)
	}
}

func (p *GrammarParser) Path() (localctx IPathContext) {
	localctx = NewPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GrammarParserRULE_path)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserR_AND {
		{
			p.SetState(25)
			p.Match(GrammarParserR_AND)
		}

	}
	{
		p.SetState(28)
		p.Match(GrammarParserR_DOT)
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserID {
		{
			p.SetState(29)
			p.Match(GrammarParserID)
		}
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GrammarParserR_DOT {
			{
				p.SetState(30)
				p.Match(GrammarParserR_DOT)
			}
			{
				p.SetState(31)
				p.Match(GrammarParserID)
			}

			p.SetState(36)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserR_SPREAD {
		{
			p.SetState(39)
			p.Match(GrammarParserR_SPREAD)
		}

	}

	return localctx
}
