package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/derpl-del/gopro1/envcode/chcode"
	"github.com/derpl-del/gopro1/envcode/pagecode"
	"github.com/derpl-del/gopro1/envcode/structcode"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
)

func main() {
	chcode.DeleteFile()
	c := cron.New()
	c.AddFunc("5 * * * * *", chcode.DeleteFile)
	c.Start()
	Funchandler()
}

//Funchandler func
func Funchandler() {
	fmt.Println("morning")
	//pagecode.Mux()
	structcode.SearchFunc()
	task()
	r := mux.NewRouter()
	r.HandleFunc("/", pagecode.HomePage)
	r.HandleFunc("/result", pagecode.HomeResult)
	r.HandleFunc("/return", pagecode.ReturnAllArticles)
	r.HandleFunc("/getdata", pagecode.GetData)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("hdmonochrome"))))
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", r)
}

func task() {
	currentTime := time.Now()
	logtime := currentTime.Format("2006-01-02 15:04:05.000000")
	fmt.Println(logtime)
}
