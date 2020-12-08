// Generated from Grammar.g4 by ANTLR 4.7.

package parser // Grammar

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGrammarListener is a complete listener for a parse tree produced by GrammarParser.
type BaseGrammarListener struct{}

var _ GrammarListener = &BaseGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGrammarListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGrammarListener) ExitExpression(ctx *ExpressionContext) {}

// EnterActionTypes is called when production actionTypes is entered.
func (s *BaseGrammarListener) EnterActionTypes(ctx *ActionTypesContext) {}

// ExitActionTypes is called when production actionTypes is exited.
func (s *BaseGrammarListener) ExitActionTypes(ctx *ActionTypesContext) {}

// EnterCommand is called when production command is entered.
func (s *BaseGrammarListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BaseGrammarListener) ExitCommand(ctx *CommandContext) {}

// EnterPath is called when production path is entered.
func (s *BaseGrammarListener) EnterPath(ctx *PathContext) {}

// ExitPath is called when production path is exited.
func (s *BaseGrammarListener) ExitPath(ctx *PathContext) {}
