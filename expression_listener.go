package xlsx

import (
	//"fmt"

	// "encoding/json"
	// "fmt"
	// "regexp"
	// "strings"
	// "sync"

	"fmt"

	"git.nanocomp.dcc.ufmg.br/mggrafeno/platform.v2/utils/xlsx/parser"
	// "github.com/360EntSecGroup-Skylar/excelize/v2"
	// "github.com/jeremywohl/flatten"
)

type CommandHandler = func(expression *ExpressionListener, ctx *parser.ActionTypesContext)

type ExpressionListener struct {
    *parser.BaseGrammarListener
    xlsx *Xlsx
    Source string
    Sheet string
    Row []string
    RowIndex int
    Index int
    BasePath string
    replace interface{}
}

var (
    commands = map[string]CommandHandler{}
)

func init(){
    commands["range"] = func(expression *ExpressionListener, ctx *parser.ActionTypesContext) {
        var (
            xlsx = expression.xlsx
            data = xlsx.data
            file = xlsx.file
            rowIndex = expression.RowIndex + 1
            sheet = expression.Sheet
            index = 0
            basePath string
        )

        for {
            value, _ := file.GetCellValue(sheet, colToAxis(expression.Index, rowIndex))
            if value  == "{{end}}" {
                break
            }
            rowIndex++
        }

        template, _ := file.GetRows(sheet)
        rows := template[expression.RowIndex: rowIndex - 1]
        
        for {
            basePath = fmt.Sprintf("%s.%d", expression.BasePath, index)
            if _, found := data[basePath]; !found {
                break
            }
         
            fmt.Println("tinha o index ", basePath)

            for _, row := range rows {
                context := *expression
                context.BasePath = basePath
                xlsx.EvalRow(context)
            }

            index++
        }
    }

    commands["if"] = func(expression *ExpressionListener, ctx *parser.ActionTypesContext) {

    }

    commands["with"] = func(expression *ExpressionListener,ctx *parser.ActionTypesContext) {

    }

    commands["block"] = func(expression *ExpressionListener,ctx *parser.ActionTypesContext) {

    }
}


func (expression *ExpressionListener) EvalCommand(ctx *parser.ActionTypesContext) {
    var (
        found bool
        handler CommandHandler
        id = ctx.Command().GetText()
    )

    if handler, found = commands[id]; !found {
        fmt.Printf("ERROR: '%s' command handler not found \n", id)
        return
    }
    
    handler(expression, ctx)
}
// EnterActionTypes is called when production actionTypes is entered.
func (expression *ExpressionListener) EnterActionTypes(ctx *parser.ActionTypesContext) {
    var (
        xlsx = expression.xlsx
        data = xlsx.data
        value interface{}
        found bool
        path string
    )
    
    defer func() {
        expression.replace = value
    }()
    
    value = ""

    if ctx.Path() != nil {
        path = ctx.Path().GetText()
    }
    fmt.Println("evaluate",
    expression.Source,
    ctx.Command() != nil,
    ctx.Path() != nil,
    ctx.R_END() != nil,
)

    if ctx.Command() != nil {
        expression.BasePath = path
        // xlsx.file.RemoveRow(expression.Sheet, expression.RowIndex)
        
        expression.EvalCommand(ctx)
        
    }  else if ctx.R_END() != nil {
        
        // xlsx.file.RemoveRow(expression.Sheet, expression.RowIndex)
        value = nil
    
    } else if value, found = data[path]; !found {
        value = ""
    }
    
}

// ExitActionTypes is called when production actionTypes is exited.
func (expression *ExpressionListener) ExitActionTypes(ctx *parser.ActionTypesContext) {
    
}
// ExitExpression is called when production expression is exited.
func (expression *ExpressionListener) ExitExpression(ctx *parser.ExpressionContext) {

    if expression.replace != nil {
        expression.xlsx.file.SetCellValue(
            expression.Sheet,
            colToAxis(expression.Index, expression.RowIndex),
            expression.replace,
        )
    }
}
