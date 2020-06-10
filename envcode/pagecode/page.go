package pagecode

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strconv"

	"github.com/derpl-del/gopro1/envcode/chcode"
	"github.com/derpl-del/gopro1/envcode/dbcode"
	"github.com/derpl-del/gopro1/envcode/structcode"
	"github.com/derpl-del/gopro1/envcode/structdata"
	"github.com/derpl-del/gopro1/envcode/wrcode"
)

//HomePage page
func HomePage(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "index.html")
	data := structcode.Articles
	hprs := structdata.HpStruct{ListArticles: data}
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
	chcode.DeleteSche()
	validationcache := chcode.FileExist(input1)
	//fmt.Println(validationcache)
	if validationcache == false {
		validation1 = dbcode.ValidationData(input1)
		dateNew1 := structcode.Articles
		//fmt.Printf("The result validation is: %v\n", validation1)
		dateNew2 := ReturnData(validation1, input1)
		combinedNew := structdata.Combined{RsData: dateNew2, ListArticles: dateNew1}
		b, _ := json.Marshal(dateNew2)
		wrcode.LoggingWrite(string(b))
		b2, _ := json.Marshal(combinedNew)
		wrcode.WriteFileExcel(combinedNew)
		//fmt.Println(string(b2))
		chcode.MakeCache(string(b2), input1)
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
		chcode.DeleteSche()
	} else {
		struct1 := chcode.ReadFile(input1)
		wrcode.LoggingWrite("Cache")
		wrcode.WriteFileExcel(struct1)
		b, _ := json.Marshal(struct1.RsData)
		wrcode.LoggingWrite(string(b))
		var filepath = path.Join("views", "result.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, struct1)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		//fmt.Fprint(data)
		chcode.DeleteSche()
	}

}

//ReturnData for validation data
func ReturnData(validation bool, input1 string) structdata.RsData {
	if validation == true {
		data := structcode.GetPokeData(input1)
		structcode.GetType(data)
		dateStat := structcode.GetStat(data)
		typedata := structcode.ListType
		hprs := structdata.TypeStruct{ListType: typedata}
		//fmt.Println(hprs)
		wrcode.LoggingWrite("API")
		intData, _ := strconv.Atoi(input1)
		prevIn := intData - 1
		nextIn := intData + 1
		values := InsDataEnv(hprs, intData, dateStat)
		dbcode.InsData(data.Name, data.Sprites.FrontDefault, data.Sprites.BackDefault, data.Sprites.FrontShiny, data.Sprites.BackShiny, prevIn, nextIn)
		dbcode.InsEnvTable("POKEMON_ENV", values, "")
		if len(hprs.ListType) >= 2 {
			dateNew := structdata.RsData{Id: input1, Name: data.Name, FrontDefault: data.Sprites.FrontDefault, BackDefault: data.Sprites.BackDefault, FrontShiny: data.Sprites.FrontShiny, BackShiny: data.Sprites.BackShiny, DataBefore: prevIn, DataAfter: nextIn, Type1: hprs.ListType[0].Type, Type2: hprs.ListType[1].Type, Stat: dateStat}
			return dateNew
		}
		dateNew := structdata.RsData{Id: input1, Name: data.Name, FrontDefault: data.Sprites.FrontDefault, BackDefault: data.Sprites.BackDefault, FrontShiny: data.Sprites.FrontShiny, BackShiny: data.Sprites.BackShiny, DataBefore: prevIn, DataAfter: nextIn, Type1: hprs.ListType[0].Type, Type2: "", Stat: dateStat}
		return dateNew
	}
	_, result2, result3, result4, result5, result6, result7, result8, result9, result10, result11, result12, result13, result14, result15, result16 := dbcode.GetData(input1)
	//fmt.Println(result2, result3, result4, result5, result6, result7, result8, result9, result10, result11, result12, result13, result14, result15, result16)
	stat := structcode.Stat{HP: result11, ATK: result12, DEF: result13, SPATK: result14, SPDEF: result15, SPD: result16}
	dateNew := structdata.RsData{Id: input1, Name: result2, FrontDefault: result3, BackDefault: result4, FrontShiny: result5, BackShiny: result6, DataBefore: result7, DataAfter: result8, Type1: result9, Type2: result10, Stat: stat}
	wrcode.LoggingWrite("DataBase")
	return dateNew
}

//GetData page
func GetData(w http.ResponseWriter, r *http.Request) {
	data2 := structcode.GetValue()
	//fmt.Println(data.Pokemon)
	fmt.Fprintf(w, "Hi")
	for i := 0; i < len(data2.Pokemon); i++ {
		//fmt.Println(data.Pokemon[i].EntryNo)
		input1 := data2.Pokemon[i].EntryNo
		data := structcode.GetPokeData(strconv.Itoa(input1))
		structcode.GetType(data)
		dateStat := structcode.GetStat(data)
		typedata := structcode.ListType
		hprs := structdata.TypeStruct{ListType: typedata}
		intData := input1
		prevIn := intData - 1
		nextIn := intData + 1
		values := InsDataEnv(hprs, intData, dateStat)
		dbcode.InsEnvTable("POKEMON_ENV", values, "")
		dbcode.InsData(data.Name, data.Sprites.FrontDefault, data.Sprites.BackDefault, data.Sprites.FrontShiny, data.Sprites.BackShiny, prevIn, nextIn)
		if len(hprs.ListType) >= 2 {
			dateNew := structdata.RsData{Id: strconv.Itoa(input1), Name: data.Name, FrontDefault: data.Sprites.FrontDefault, BackDefault: data.Sprites.BackDefault, FrontShiny: data.Sprites.FrontShiny, BackShiny: data.Sprites.BackShiny, DataBefore: prevIn, DataAfter: nextIn, Type1: hprs.ListType[0].Type, Type2: hprs.ListType[1].Type, Stat: dateStat}
			b, _ := json.Marshal(dateNew)
			wrcode.LoggingWrite(string(b))
		}
		dateNew := structdata.RsData{Id: strconv.Itoa(input1), Name: data.Name, FrontDefault: data.Sprites.FrontDefault, BackDefault: data.Sprites.BackDefault, FrontShiny: data.Sprites.FrontShiny, BackShiny: data.Sprites.BackShiny, DataBefore: prevIn, DataAfter: nextIn, Type1: hprs.ListType[0].Type, Type2: "", Stat: dateStat}
		b, _ := json.Marshal(dateNew)
		wrcode.LoggingWrite(string(b))
	}
	fmt.Fprintf(w, " ---- Success")
}

//ReturnAllArticles A
func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(structcode.Articles)
}

//InsDataEnv a
func InsDataEnv(input structdata.TypeStruct, num int, stat structcode.Stat) string {
	if len(input.ListType) >= 2 {
		valuestring := fmt.Sprintf("'%v','%v','%v','%v','%v','%v','%v','%v','%v'", num, input.ListType[0].Type, input.ListType[1].Type, stat.HP, stat.ATK, stat.DEF, stat.SPATK, stat.SPDEF, stat.SPDEF)
		return valuestring
	}
	valuestring := fmt.Sprintf("'%v','%v','%v','%v','%v','%v','%v','%v','%v'", num, input.ListType[0].Type, "", stat.HP, stat.ATK, stat.DEF, stat.SPATK, stat.SPDEF, stat.SPDEF)
	return valuestring
}
