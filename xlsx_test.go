package xlsx

import (
    "time"
    "strings"
    "testing"
)

type Sample struct {
	Name string `json:"name"`
	Characterizations []string `json:"characterizations"`
}

func Test(t *testing.T) {
    
    var (
        file *Xlsx
        err error
	)

	if file, err = FromTemplate("./input.xlsx"); err != nil {
        t.Error(err)
        return
    }
    
    now := time.Now().Format("01/02/2006")
    characterizations := strings.Split(",,,,X,,,,,X,,,,,,X,,,,",",")
    
    data := map[string]interface{}{
        "code": "023/2020",
        "objectives": " Verificar as propriedades dos materiais sintetizados para verificar se a funcionalização do grafeno com enxofre ocorreu conforme o esperado.",
        "samples":[]Sample{
			{ Name: "20180709Pa_s03_fg_snp_Sco_t1", Characterizations: characterizations},
			{ Name: "20180709Pa_s03_fg_snp_Sco_t2", Characterizations: characterizations},
			{ Name: "20180709Pa_s03_fg_snp_Sco_t3", Characterizations: characterizations},
		},
        "timestamp": map[string]interface{}{
			"created": now,
			"approved": now,
			"archived": now,
		},
    }


    file.Render(data)
	
	file.Save("./input.saved.xlsx")
}