package wrcode

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

//LoggingWrite for write
func LoggingWrite(input string) {
	currentTime := time.Now()
	logtime := currentTime.Format("2006-01-02 15:04:05.000000")
	logdata := "##########LOGNEW##########\n"
	mydata := []byte(logdata)
	logTittle := "log/" + currentTime.Format("20060102") + "access_log.log"
	if FileExist(logTittle) {

	} else {
		err := ioutil.WriteFile(logTittle, mydata, 0777)
		// handle this error
		if err != nil {
			// print it out
			fmt.Println(err)
		}
	}
	f, err := os.OpenFile(logTittle, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	logtext := logtime + " : " + input + "\n" + "###############\n"
	if _, err = f.WriteString(logtext); err != nil {
		panic(err)
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
