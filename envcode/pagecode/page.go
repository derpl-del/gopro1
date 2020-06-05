package pagecode

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strconv"

	"github.com/derpl-del/gopro1/envcode/dbcode"
	"github.com/derpl-del/gopro1/envcode/structcode"
	"github.com/derpl-del/gopro1/envcode/wrcode"
)

// FntReq for result.html
type combined struct {
	ListArticles []structcode.Article
	RsData
}

//EnvData for data
type EnvData struct {
}

//hpstruct
type hpstruct struct {
	ListArticles []structcode.Article
}

//TypeStruct a
type TypeStruct struct {
	ListType []structcode.TypePokemon
}

// RsData for result.html
type RsData struct {
	Name         string
	FrontDefault string
	BackDefault  string
	FrontShiny   string
	BackShiny    string
	DataBefore   int
	DataAfter    int
	Type1        string
	Type2        string
	structcode.Stat
}

//HomePage page
func HomePage(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "index.html")
	data := structcode.Articles
	hprs := hpstruct{data}
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, hprs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//HomeResult page
func HomeResult(w http.ResponseWriter, r *http.Request) {
	var validation1 bool
	var input1 = r.FormValue("pokemon")
	validation1 = dbcode.ValidationData(input1)
	dateNew1 := structcode.Articles
	fmt.Printf("The result validation is: %v\n", validation1)
	dateNew2 := ReturnData(validation1, input1)

	combinedNew := combined{dateNew1, dateNew2}
	b, _ := json.Marshal(dateNew2)
	wrcode.LoggingWrite(string(b))
	var filepath = path.Join("views", "result.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, combinedNew)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//fmt.Fprint(data)

}

//ReturnData for validation data
func ReturnData(validation bool, input1 string) RsData {
	if validation == true {
		data := structcode.GetPokeData(input1)
		structcode.GetType(data)
		dateStat := structcode.GetStat(data)
		typedata := structcode.ListType
		hprs := TypeStruct{typedata}
		//fmt.Println(hprs)
		intData, _ := strconv.Atoi(input1)
		prevIn := intData - 1
		nextIn := intData + 1
		values := InsDataEnv(hprs, intData, dateStat)
		dbcode.InsData(data.Name, data.Sprites.FrontDefault, data.Sprites.BackDefault, data.Sprites.FrontShiny, data.Sprites.BackShiny, prevIn, nextIn)
		dbcode.InsEnvTable("POKEMON_ENV", values, "")
		if len(hprs.ListType) >= 2 {
			dateNew := RsData{data.Name, data.Sprites.FrontDefault, data.Sprites.BackDefault, data.Sprites.FrontShiny, data.Sprites.BackShiny, prevIn, nextIn, hprs.ListType[0].Type, hprs.ListType[1].Type, dateStat}
			return dateNew
		}
		dateNew := RsData{data.Name, data.Sprites.FrontDefault, data.Sprites.BackDefault, data.Sprites.FrontShiny, data.Sprites.BackShiny, prevIn, nextIn, hprs.ListType[0].Type, "", dateStat}
		return dateNew
	}
	_, result2, result3, result4, result5, result6, result7, result8, result9, result10, result11, result12, result13, result14, result15, result16 := dbcode.GetData(input1)
	//fmt.Println(result2, result3, result4, result5, result6, result7, result8, result9, result10, result11, result12, result13, result14, result15, result16)
	stat := structcode.Stat{HP: result11, ATK: result12, DEF: result13, SPATK: result14, SPDEF: result15, SPD: result16}
	dateNew := RsData{result2, result3, result4, result5, result6, result7, result8, result9, result10, stat}
	return dateNew
}

//GetData page
func GetData(w http.ResponseWriter, r *http.Request) {
	data := structcode.GetValue()
	fmt.Println(data.Pokemon)
	fmt.Fprintf(w, "Hi")
	for i := 0; i < len(data.Pokemon); i++ {
		fmt.Println(data.Pokemon[i].EntryNo)
		input1 := data.Pokemon[i].EntryNo
		data2 := structcode.GetPokeData(strconv.Itoa(input1))
		structcode.GetType(data2)
		dateStat := structcode.GetStat(data2)
		typedata := structcode.ListType
		hprs := TypeStruct{typedata}
		intData := input1
		prevIn := intData - 1
		nextIn := intData + 1
		values := InsDataEnv(hprs, intData, dateStat)
		dbcode.InsEnvTable("POKEMON_ENV", values, "")
		dbcode.InsData(data2.Name, data2.Sprites.FrontDefault, data2.Sprites.BackDefault, data2.Sprites.FrontShiny, data2.Sprites.BackShiny, prevIn, nextIn)
		if len(hprs.ListType) >= 2 {
			dateNew := RsData{data2.Name, data2.Sprites.FrontDefault, data2.Sprites.BackDefault, data2.Sprites.FrontShiny, data2.Sprites.BackShiny, prevIn, nextIn, hprs.ListType[0].Type, hprs.ListType[1].Type, dateStat}
			b, _ := json.Marshal(dateNew)
			wrcode.LoggingWrite(string(b))
		}
		dateNew := RsData{data2.Name, data2.Sprites.FrontDefault, data2.Sprites.BackDefault, data2.Sprites.FrontShiny, data2.Sprites.BackShiny, prevIn, nextIn, hprs.ListType[0].Type, "", dateStat}
		b, _ := json.Marshal(dateNew)
		wrcode.LoggingWrite(string(b))
	}
	fmt.Fprintf(w, " ---- Success")
}

//ReturnAllArticles A
func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(structcode.Articles)
}

//InsDataEnv a
func InsDataEnv(input TypeStruct, num int, stat structcode.Stat) string {
	if len(input.ListType) >= 2 {
		valuestring := fmt.Sprintf("'%v','%v','%v','%v','%v','%v','%v','%v','%v'", num, input.ListType[0].Type, input.ListType[1].Type, stat.HP, stat.ATK, stat.DEF, stat.SPATK, stat.SPDEF, stat.SPDEF)
		return valuestring
	}
	valuestring := fmt.Sprintf("'%v','%v','%v','%v','%v','%v','%v','%v','%v'", num, input.ListType[0].Type, "", stat.HP, stat.ATK, stat.DEF, stat.SPATK, stat.SPDEF, stat.SPDEF)
	return valuestring
}
