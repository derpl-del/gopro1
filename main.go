package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"text/template"

	"github.com/derpl-del/gopro1/envcode/structcode"
	"github.com/derpl-del/gopro1/envcode/wrcode"
	"github.com/gorilla/mux"
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

func main() {
	fmt.Println("morning")
	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/result", homeResult)
	//r.HandleFunc("/getdata", getData)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("hdmonochrome"))))
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", r)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "index.html")
	//data := function.getData()
	data := structcode.GetValue()
	fmt.Println(data)
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

func homeResult(w http.ResponseWriter, r *http.Request) {
	var validation1 bool
	var input1 = r.FormValue("pokemon")
	//validation1 = L3.ValidationData(input1)
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
	//if validation == true {
	data := structcode.GetPokeData(input1)
	intData, _ := strconv.Atoi(input1)
	prevIn := intData - 1
	nextIn := intData + 1
	dateNew := RsData{data.Name, data.Sprites.FrontDefault, data.Sprites.BackDefault, data.Sprites.FrontShiny, data.Sprites.BackShiny, prevIn, nextIn}
	//L3.InsData(data.Name, data.Sprites.FrontDefault, data.Sprites.BackDefault, data.Sprites.FrontShiny, data.Sprites.BackShiny, prevIn, nextIn)
	return dateNew
	//}
	//_, result2, result3, result4, result5, result6, result7, result8 := L3.GetData(input1)
	//dateNew := RsData{result2, result3, result4, result5, result6, result7, result8}
	//return dateNew
}
