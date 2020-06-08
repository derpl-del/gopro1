package chcode

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
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
func FileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
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
