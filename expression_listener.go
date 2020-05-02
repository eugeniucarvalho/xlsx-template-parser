package xlsx

import (
	"fmt"
	"strings"

	"git.nanocomp.dcc.ufmg.br/mggrafeno/platform.v2/utils/xlsx/parser"
)

type CommandHandler = func(expression *ExpressionListener, ctx *parser.ActionTypesContext)

type ExpressionListener struct {
    *parser.BaseGrammarListener
    xlsx *Xlsx
    Source string
    Sheet string
    Expression string
    Row []string
    RowIndex int
    Index int
    Skip int
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
            file = xlsx.file
            rowIndex = expression.RowIndex + 1
            sheet = expression.Sheet
            data = xlsx.data
            index = 0
            basePath string
            rowBase = rowIndex
            rowAdded = 0
        )

        for {
            value, _ := file.GetCellValue(
                sheet,
                CoordinatesToCellName(expression.Index, rowIndex),
            )
            if value == "{{stop}}" ||  value == "{{end}}"{
                break
            }
            rowIndex++
        }

        terminationIndex := rowIndex

        rows := xlsx.Rows[expression.RowIndex: terminationIndex - 1]
        
        templateRowsCount := len(rows)
        expression.Skip = rowIndex - expression.RowIndex 

        rowBase = expression.RowIndex

        for {
            arrayIndexKey := fmt.Sprintf("%s.%d.$idx", expression.BasePath, index)
            
            if _, found := data[arrayIndexKey]; !found {
                break
            }
            
            file.DuplicateRowTo(sheet,rowBase, rowBase + templateRowsCount)

            
            basePath = arrayIndexKey[: len(arrayIndexKey) - 5 ]
            
            rowAdded += templateRowsCount

            for templateRowIndex, row := range rows {
                context := *expression
                context.BasePath = basePath
                context.Row = row
                context.RowIndex = rowBase + templateRowIndex
                xlsx.EvalRow(&context)
            }

            rowBase += templateRowsCount

            index++
        }
        
        for i := 0; i <= (templateRowsCount + 1) ; i++ {
            file.RemoveRow(sheet, rowBase)
        }
        expression.RowIndex = expression.RowIndex + rowAdded - 1
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

func (expression *ExpressionListener) SpreadHandler(ctx *parser.ActionTypesContext) {
    var (
        xlsx = expression.xlsx
        file = xlsx.file
        sheet = expression.Sheet
        data = xlsx.data
        colIndex = expression.Index
        index = 0
    )

    for {
        valueIndexKey := fmt.Sprintf("%s.%d", expression.BasePath, index)
        arrayIndexKey := fmt.Sprintf("%s.$idx", valueIndexKey)
        
        if _, found := data[arrayIndexKey]; !found {
            break
        }
        
        file.SetCellValue(
            sheet,
            CoordinatesToCellName(colIndex, expression.RowIndex),
            data[valueIndexKey],
        )

        colIndex++
        index++
    }

}
// EnterActionTypes is called when production actionTypes is entered.
func (expression *ExpressionListener) EnterActionTypes(ctx *parser.ActionTypesContext) {
    var (
        xlsx = expression.xlsx
        data = xlsx.data
        value interface{}
        found bool
        path string
        pathContext = ctx.Path()
        spreadContext interface{}
    )
    
    defer func() {
        expression.replace = value
    }()
    
    value = ""

    if pathContext != nil {
        path = pathContext.GetText()

        if spreadContext = pathContext.(*parser.PathContext).R_SPREAD(); spreadContext != nil {
            path = path[:len(path)-3]
        }
        
        if path[:1] == "&" {
            path = strings.Replace(path,"&",expression.BasePath,1)
        }
    }
    if ctx.Command() != nil {
        expression.BasePath = path
        expression.EvalCommand(ctx)
        
        value = nil
        
    }  else if ctx.R_END() != nil {

        value = nil
    
    } else if spreadContext != nil {
        expression.BasePath = path
        expression.SpreadHandler(ctx)

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
        
        value := strings.ReplaceAll(
            expression.Source,
            expression.Expression,
            fmt.Sprintf("%v", expression.replace),
        )

        expression.xlsx.file.SetCellValue(
            expression.Sheet,
            CoordinatesToCellName(expression.Index, expression.RowIndex),
            value,
        )
    }
}
