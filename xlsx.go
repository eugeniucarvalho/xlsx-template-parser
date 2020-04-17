package xlsx

import (
	//"fmt"

	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"git.nanocomp.dcc.ufmg.br/mggrafeno/platform.v2/utils/xlsx/parser"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/davecgh/go-spew/spew"

	"github.com/jeremywohl/flatten"
)

// Xlsx Represents template struct
type Xlsx struct {
    file *excelize.File
    data  map[string]interface{}
    Error chan error
    Progress chan int
    Done chan bool
}

// Options for render has only one property WrapTextInAllCells for wrapping text
type Options struct {
    WrapTextInAllCells bool
}

var (
    rgx             = regexp.MustCompile(`\{\{\s*(\.\w+(\.\w+)*)\}\}`)
    eachRegex       = regexp.MustCompile(`\{\{\#each\s+(\w+)\s*\}\}`)
    endRegex        = regexp.MustCompile(`\{\{\s*\/end\s*\}\}`)
    stopRegex        = regexp.MustCompile(`\{\{\s*\/\$stop\s*\}\}`)
    filtro          = regexp.MustCompile(`\{\{\s*filtro\\|([\w:]*)\s*\}\}.*`)                      
    removeMustache  = regexp.MustCompile(`[{}]*`)
    hasMustache     = regexp.MustCompile(`{{[\w\.]*}}`)
)

// ReadTemplate reads template from disk and stores it in a struct
func FromTemplate(path string) (xlsx *Xlsx,err error) {
    
    xlsx = &Xlsx{
        Error: make(chan error),
        Progress: make(chan int),
        Done: make(chan bool),
        data: map[string]interface{}{},
    }
    
    xlsx.file, err = excelize.OpenFile(path)
    
    return 
}

func (xlsx *Xlsx) normalizeData(in interface{}) (data map[string]interface{}, err error) {
    var (
		variablesJson string
		part string
		base string
		out           []byte
	)
    
    data = map[string]interface{}{}

    if out, err = json.Marshal(in); err != nil {
		return
	}

	if variablesJson, err = flatten.FlattenString(string(out), ".", flatten.DotStyle); err != nil {
		return
	}

    if err = json.Unmarshal([]byte(variablesJson), &data); err != nil {
        return
    }
    
    for path := range data {
        base = ""
        parts := strings.Split(path,".")
        parts = parts[1:]
        for {
            if len(parts) == 0 {
                break
            }
            part = parts[0]
            parts = parts[1:]
            
            base = fmt.Sprintf("%s.%s", base, part)
            
            if _, err := strconv.ParseInt(part,10,64); err == nil {
                if _, found := data[base]; !found {
                    data[base] = true
                }
            }
        }
    } 
    
    return
}

// Render renders report and stores it in a struct
func (xlsx *Xlsx) Render(in interface{}) (err error) {
    return xlsx.RenderWithOptions(in, new(Options))
}

// RenderWithOptions renders report with options provided and stores it in a struct
func (xlsx *Xlsx) RenderWithOptions(data interface{}, options *Options) (err error) {
    
    var (
        file = xlsx.file
        wg = sync.WaitGroup{}
    )

    if xlsx.data, err  = xlsx.normalizeData(data); err != nil {
        return
    }

    spew.Dump(xlsx.data)
    
    for _, name := range file.GetSheetMap() {
        wg.Add(1)
        go func ()  {
            defer func(){
                wg.Done()
            }()
            if err = xlsx.renderSheet(name); err != nil {
                xlsx.Error <-err
            }
        }()
    }
    wg.Wait()
  
    // xlsx.Done <- true
    return nil
}

func (xlsx *Xlsx) Save(path string) (err error) {
    return xlsx.file.SaveAs(path)
}

func (xlsx *Xlsx) renderSheet(name string) (err error) {
    
    var (
        rows *excelize.Rows
        row []string
        index = 1
    )
    
    if rows, err = xlsx.file.Rows(name); err != nil {
        println(err.Error())
        return
    }

    for rows.Next() {
        if row, err = rows.Columns(); err != nil {
            println(err.Error())
            return
        }
        
        if xlsx.hasStop(row) {
            return
        }
    
        xlsx.evalRow(name, row, index)
        index++
    }
    
    return 
}

func (xlsx *Xlsx) hasStop(row []string) bool {
    return row == nil ||  stopRegex.MatchString(row[0])
}

func (xlsx *Xlsx) evalRow(sheet string, row []string, rowIndex int) (err error) {
    // var (
    //     file = xlsx.file
    // )
    
    for index, val := range row {
        xlsx.evalCell(sheet, row, rowIndex, index, val)
    }
    return
}

func (xlsx *Xlsx) evalCell(sheet string, row []string, rowIndex, index int, value string) (err error) {

    value = strings.Trim(value," ")
    
    if value == ""  || !strings.Contains(value,"{{") {
        return
    }

    listener := &ExpressionListener{
        xlsx: xlsx,
        Source: value,
        Sheet: sheet,
        Row: row,
        RowIndex: rowIndex,
        Index: index,
    }

	is := antlr.NewInputStream(value)
	lexer := parser.NewGrammarLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewGrammarParser(stream)


	// Finally parse the expression
    antlr.ParseTreeWalkerDefault.Walk(listener, p.Expression())
    return
}

var alphabet = strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ","")

func colToAxis(col, row int) string{
    return fmt.Sprintf("%s%d", alphabet[col], row)
}
  // for si, _   := range xlsx.Sheets {
    //     ctrl    := 0
    //     ctx     := getCtx(in, si)
    //     sheet   := xlsx.Sheets[si]
    //     //Verifica se a planilha possui linhas para processar
    //     lenRows := len(sheet.Rows)

    //     if lenRows == 0 { continue }

    //     for {
    //         //fmt.Println("Read row", ctrl, len(sheet.Rows))
    //         if ctrl == lenRows { break }

    //         row := sheet.Rows[ctrl] 
    //         //if(!ok) {continue}

    //         if isTerminador(row) {
    //             sheet.Rows = sheet.Rows[:ctrl]
    //             break
    //         }

    //         prop := isList(row)
    //         format := map[string]interface{}{}
    //         types, ok := ctx[prop + "Types"];
    //         if ok {
    //             format = reflect.ValueOf(types).Interface().(map[string]interface{})
    //         }
    //         if prop != "" && isArray(ctx, prop) {
    //             row := sheet.Rows[ctrl + 1] 
               
    //             lista := reflect.ValueOf(ctx[prop])
    //             for i := 0; i < lista.Len(); i++ {
    //                 rowcp := copyRow(row)
    //                 RenderRow(sheet, rowcp, lista.Index(i).Interface(), format)
    //                 //adiciona a linha nova
    //                 sheet.Rows = append(sheet.Rows[:ctrl], append([]*Row{rowcp} , sheet.Rows[ctrl:]...)...)
    //                 ctrl++
    //             }
    //             //Remove os elementos de controle do template
    //             sheet.Rows = append(sheet.Rows[:ctrl], sheet.Rows[ctrl+3:]...)
    //         } else {
    //             RenderRow(sheet, sheet.Rows[ctrl], ctx, format);
    //             ctrl++ 
    //         }
    //     }
    // }
// func copyRow(from *Row) (to *Row) {
// 	to = &Row{Sheet: from.Sheet}
//         to.Height = from.Height
// 	for _, cell := range from.Cells {
//             ncell := to.AddCell()
//            // ncell.SetFormula(cell.Formula())
//             ncell.SetStyle(cell.GetStyle())
//             ncell.Value = cell.Value
//             ncell.HMerge = cell.HMerge
//             ncell.VMerge = cell.VMerge
//             ncell.Hidden = cell.Hidden
//             ncell.NumFmt = cell.NumFmt
// 	}
//     return
// }

// func checkAutoFilter(row *Row) (*XlsxAutoFilter, bool) {
//     cell := row.Cells[0]
//     if match := filtro.FindAllStringSubmatch(cell.Value, -1); len(match) > 0{
//         //fmt.Println("Match auto filter", cell.Value)
//         //spew.Dump(match)

//         autoFilter := &XlsxAutoFilter{}
//         //autoFilter.Ref = fmt.Sprintf("A1:%s1", numericToLetters(maxCell))
//         autoFilter.Ref = match[0][1]
//         cell.Value = filtro.ReplaceAllString(cell.Value, "")
//         return autoFilter, true
//     }
//     return nil, false
// }

// func isList(in *Row) string {
//     for _, cell := range in.Cells {
//         if cell.Value == "" { continue }
//         if match := eachRegex.FindAllStringSubmatch(cell.Value, -1); match != nil {
//             return match[0][1]
//         }
//     }
//     return ""
// }

// func isTerminador(in *Row) bool {
//     for _, cell := range in.Cells {
//         if match := endRegex.FindAllStringSubmatch(cell.Value, -1); match != nil {
//             return true;
//         }
//     }
//     return false
// }

// func RenderRow(s *Sheet, in *Row, ctx interface{}, format map[string]interface{}){
//    //fmt.Println("UTC: ", timeLocationUTC)
//     //spew.Dump(ctx)
//     for _, cell := range in.Cells {
//         if !hasMustache.MatchString(cell.Value) {continue}

//         prop := removeMustache.ReplaceAllString(cell.Value, "")
//         v, _ := ctx.(map[string]interface{})[prop]
        
//         if  v == nil {
//             cell.SetString("");
//             continue;
//         }

//         codf, _ := format[prop]
//         //fmt.Println("Cell prop", "-"+cell.Value+"-",prop, codf , v, "\n")
//         var typec int
        
//         if codf != nil {
//             typec = int(codf.(float64))
//         }else {
//             typec = 2
//         }
//         switch typec {
            
//             case 3 :// bool
//                 vs := "Não"
//                 if v.(bool) {vs = "Sim"}
//                 cell.SetString(vs)
//             case 4 :// date
//                 //fmt.Println("set Date", v.(float64), int64(v.(float64)))
//                 cell.SetDate(F64toTime(v.(float64)))
//             case 5 :// datetime
//                 cell.SetDateTime(F64toTime(v.(float64)))
//                 //cell.SetDateTimeWithFormat(v.(float64), "")
//             case 6 :// func
//                 //fmt.Println("Formular", v.(string))
//                 cell.SetFormula(v.(string))
//             case 7 :// money
//                 //cell.SetFloatWithFormat(v.(float64), `0.00,_-R$ * #.##0,00_-;-R$ * #.##0,00_-;_-R$ * "-"_-;_-@_-`)
//                 cell.SetFloatWithFormat(v.(float64), "R$ * #.#,00 ;[red]-R$ * #.#,00;R$ * 0,00") 
//             case 8 :
//                 cell.SetHyperLink(v.(string))
//             case 9 :
//                 //  xtime := F64toTime(v.(float64))
//                 //fmt.Println("Time: ", xtime)
//                 //cell.SetDateTimeWithFormat(timeToExcelTime(F64toTime(v.(float64))), "h:mm") 
//                 cell.SetTime(F64toTime(v.(float64))) 
//             default:
//                 cell.Value = raymond.MustRender(cell.Value, ctx)
//         }
//     }
// }

// func F64toTime(v float64) time.Time {
//     //date := time.Unix(int64(v), 0)
//     return time.Unix(int64(v), 0)//date.In(timeLocationUTC)
// }

// func getCtx(in interface{}, i int) map[string]interface{} {
//     if ctx, ok := in.(map[string]interface{}); ok {
//         return ctx
//     }
//     if ctxSlice, ok := in.([]interface{}); ok {
//         if len(ctxSlice) > i {
//             _ctx := ctxSlice[i]
//             if ctx, ok := _ctx.(map[string]interface{}); ok {
//                 return ctx
//             }
//         }
//         return nil
//     }
//     return nil
// }

// func isArray(in map[string]interface{}, prop string) bool {
//     val, ok := in[prop]
//     if !ok || val == nil { return false }
    
//     switch reflect.TypeOf(val).Kind() {
//         case reflect.Array, reflect.Slice:
//             return true
//     }
//     return false
// }

/*


func init(){
    if err := SetTimeZone("America/Fortaleza"); err != nil {
        fmt.Println("XLSX.Error: ", err.Error())
    }
}

func SetTimeZone(locale string) (err error) {
    fmt.Println("XLSX.TimeZone: ", locale)
    timeLocationUTC , err = time.LoadLocation(locale)
    return
}

*/
