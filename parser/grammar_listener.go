// Generated from Grammar.g4 by ANTLR 4.7.

package parser // Grammar

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GrammarListener is a complete listener for a parse tree produced by GrammarParser.
type GrammarListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterActionTypes is called when entering the actionTypes production.
	EnterActionTypes(c *ActionTypesContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterPath is called when entering the path production.
	EnterPath(c *PathContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitActionTypes is called when exiting the actionTypes production.
	ExitActionTypes(c *ActionTypesContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitPath is called when exiting the path production.
	ExitPath(c *PathContext)
}
