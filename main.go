package main

import (
	"fmt"
	"net/http"
	"path"
	"text/template"

	"github.com/derpl-del/gopro1/envcode/structcode"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("morning")
	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	//r.HandleFunc("/result", homeResult)
	//r.HandleFunc("/getdata", getData)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("hdmonochrome"))))
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", r)
}

func homePage(w http.ResponseWriter, r *http.Request) {
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
