package chcode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/derpl-del/gopro1/envcode/structdata"
)

//MakeCache Func
func MakeCache(data string, PokeID string) {
	CacheTittle := "cachefile/pokemon" + PokeID + "_cache.json"
	mydata := []byte(data)
	err := ioutil.WriteFile(CacheTittle, mydata, 0777)
	if err != nil {
		// print it out
		fmt.Println(err)
	}
}

//FileExist validation
func FileExist(PokeID string) bool {
	CacheTittle := "cachefile/pokemon" + PokeID + "_cache.json"
	info, err := os.Stat(CacheTittle)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//DeleteSche validation
func DeleteSche() {
	DeleteFile()
	DeleteFileExcel()
}

//DeleteFile validation
func DeleteFile() {
	var cutoff = 4 * time.Minute
	fileInfo, err := ioutil.ReadDir("cachefile/")
	if err != nil {
		fmt.Println(err)
	}
	now := time.Now()
	for _, info := range fileInfo {
		if diff := now.Sub(info.ModTime()); diff > cutoff {
			namefile := "cachefile/" + info.Name()
			os.Remove(namefile)
		}
	}
}

//DeleteFileExcel validation
func DeleteFileExcel() {
	var cutoff = 4 * time.Minute
	fileInfo, err := ioutil.ReadDir("FileDownload/")
	if err != nil {
		fmt.Println(err)
	}
	now := time.Now()
	for _, info := range fileInfo {
		if diff := now.Sub(info.ModTime()); diff > cutoff {
			namefile := "FileDownload/" + info.Name()
			os.Remove(namefile)
		}
	}
}

//ReadFile func
func ReadFile(PokeID string) structdata.Combined {
	CacheTittle := "cachefile/pokemon" + PokeID + "_cache.json"
	jsonFile, err := os.Open(CacheTittle)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var struct1 structdata.Combined
	json.Unmarshal(byteValue, &struct1)
	return struct1
}
