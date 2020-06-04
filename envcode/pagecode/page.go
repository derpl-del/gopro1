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
	*structcode.Response
	RsData
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
}

//HomePage page
func HomePage(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "index.html")
	//data := function.getData()
	data := structcode.GetValue()
	//fmt.Println(data)
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//HomeResult page
func HomeResult(w http.ResponseWriter, r *http.Request) {
	var validation1 bool
	var input1 = r.FormValue("pokemon")
	validation1 = dbcode.ValidationData(input1)
	dateNew1 := structcode.GetValue()
	fmt.Printf("The result validation is: %v\n", validation1)
	dateNew2 := ReturnData(validation1, input1)
	combinedNew := combined{dateNew1, dateNew2}
	b, _ := json.Marshal(combinedNew)
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
		intData, _ := strconv.Atoi(input1)
		prevIn := intData - 1
		nextIn := intData + 1
		dateNew := RsData{data.Name, data.Sprites.FrontDefault, data.Sprites.BackDefault, data.Sprites.FrontShiny, data.Sprites.BackShiny, prevIn, nextIn}
		dbcode.InsData(data.Name, data.Sprites.FrontDefault, data.Sprites.BackDefault, data.Sprites.FrontShiny, data.Sprites.BackShiny, prevIn, nextIn)
		return dateNew
	}
	_, result2, result3, result4, result5, result6, result7, result8 := dbcode.GetData(input1)
	dateNew := RsData{result2, result3, result4, result5, result6, result7, result8}
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
		data := structcode.GetPokeData(strconv.Itoa(input1))
		intData := input1
		prevIn := intData - 1
		nextIn := intData + 1
		dateNew := RsData{data.Name, data.Sprites.FrontDefault, data.Sprites.BackDefault, data.Sprites.FrontShiny, data.Sprites.BackShiny, prevIn, nextIn}
		b, _ := json.Marshal(dateNew)
		wrcode.LoggingWrite(string(b))
		dbcode.InsData(data.Name, data.Sprites.FrontDefault, data.Sprites.BackDefault, data.Sprites.FrontShiny, data.Sprites.BackShiny, prevIn, nextIn)
	}
	fmt.Fprintf(w, " ---- Success")
}
