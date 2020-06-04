package main

import (
	"fmt"
	"net/http"

	"github.com/derpl-del/gopro1/envcode/pagecode"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("morning")
	//pagecode.Mux()
	r := mux.NewRouter()
	r.HandleFunc("/", pagecode.HomePage)
	r.HandleFunc("/result", pagecode.HomeResult)
	r.HandleFunc("/getdata", pagecode.GetData)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("hdmonochrome"))))
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", r)
}
