package wrcode

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/derpl-del/gopro1/envcode/structdata"
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

//WriteFileExcel Func
func WriteFileExcel(input structdata.Combined) {
	f := excelize.NewFile()
	// Set value of a cell.
	f.SetCellValue("Sheet1", "A1", "Name")
	f.SetCellValue("Sheet1", "A2", input.Name)
	f.SetCellValue("Sheet1", "B1", "Type1")
	f.SetCellValue("Sheet1", "B2", input.RsData.Type1)
	f.SetCellValue("Sheet1", "C1", "Type2")
	f.SetCellValue("Sheet1", "C2", input.RsData.Type2)
	f.SetCellValue("Sheet1", "D1", "HP")
	f.SetCellValue("Sheet1", "D2", input.RsData.HP)
	f.SetCellValue("Sheet1", "E1", "ATK")
	f.SetCellValue("Sheet1", "E2", input.RsData.ATK)
	f.SetCellValue("Sheet1", "F1", "DEF")
	f.SetCellValue("Sheet1", "F2", input.RsData.DEF)
	f.SetCellValue("Sheet1", "G1", "SPATK")
	f.SetCellValue("Sheet1", "G2", input.RsData.SPATK)
	f.SetCellValue("Sheet1", "H1", "SPDEF")
	f.SetCellValue("Sheet1", "H2", input.RsData.SPDEF)
	f.SetCellValue("Sheet1", "I1", "SPD")
	f.SetCellValue("Sheet1", "I2", input.RsData.SPD)
	dataid := input.DataBefore + 1
	// Save xlsx file by the given path.
	TittleFile := "FileDownload/pokemon" + strconv.Itoa(dataid) + "_file.xlsx"
	if err := f.SaveAs(TittleFile); err != nil {
		fmt.Println(err)
	}
}
