package dbcode

import (
	"database/sql"
	"fmt"
	"strconv"

	//for framework
	_ "github.com/godror/godror"
	//for framework
	"github.com/derpl-del/gopro1/envcode/wrcode"
)

//GetSysDate a
func GetSysDate() {

	db, err := sql.Open("godror", "testing/welcome1@xe")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select MODEL_TYPE,time_stamp from XXVALS where rownum = '1'")

	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var thedate string
	var themodel string
	for rows.Next() {

		rows.Scan(&themodel, &thedate)
	}
	fmt.Printf("The date is: %s\n", thedate)
}

//InsData insert data
func InsData(input1 string, input2 string, input3 string, input4 string, input5 string, input6 int, input7 int) {
	var id int
	db, err := sql.Open("godror", "testing/welcome1@xe")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	id = input6 + 1
	values1 := ValidationData(strconv.Itoa(id))
	if values1 == true {
		statementSQL := fmt.Sprintf("INSERT INTO POKEMON_NEW VALUES ('%v','%s', '%s', '%s', '%s', '%s', '%d', '%d')", id, input1, input2, input3, input4, input5, input6, input7)
		//statementSQL := "INSERT INTO POKEMON VALUES (" + input1 + "," + input2 + "," + input3 + "," + input4 + "," + input5 + "," + string(input6) + "," + string(input7) + ")"
		//fmt.Printf("The query is: %s\n", statementSQL)
		rows, err := db.Query(statementSQL)

		if err != nil {
			fmt.Println("Error running query")
			fmt.Println(err)
			return
		}
		defer rows.Close()
	}
}

//ValidationData validation data
func ValidationData(input1 string) bool {
	db, err := sql.Open("godror", "testing/welcome1@xe")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer db.Close()
	statementSQL := fmt.Sprintf("select id from POKEMON_NEW WHERE id = '%s' and rownum = '1'", input1)
	//statementSQL := "INSERT INTO POKEMON VALUES (" + input1 + "," + input2 + "," + input3 + "," + input4 + "," + input5 + "," + string(input6) + "," + string(input7) + ")"
	//fmt.Printf("The query is: %s\n", statementSQL)
	rows, err := db.Query(statementSQL)

	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return false
	}
	defer rows.Close()
	var thename string
	for rows.Next() {
		rows.Scan(&thename)
	}
	if string(thename) == input1 {
		return false
	}
	return true
}

//GetData a
func GetData(input1 string) (string, string, string, string, string, string, int, int, string, string, int, int, int, int, int, int) {
	//fmt.Println("GetData")
	db, err := sql.Open("godror", "testing/welcome1@xe")
	if err != nil {
		fmt.Println(err)
		return "null", "null", "null", "null", "null", "null", 0, 0, "null", "null", 0, 0, 0, 0, 0, 0
	}
	defer db.Close()

	statementSQL := fmt.Sprintf("select X.ID,X.NAME,X.FRONTDEFAULT,X.BACKDEFAULT,X.FRONTSHINY,X.BACKSHINY,X.DATABEFORE,X.DATAAFTER,Z.TYPE1,Z.TYPE2,Z.HP,Z.ATK,Z.DEF,Z.SPATK,Z.SPDEF,Z.SPD from POKEMON_NEW X join POKEMON_ENV Z on X.ID = Z.ID where X.id = '%v' and rownum = '1'", input1)
	//statementSQL := "INSERT INTO POKEMON VALUES (" + input1 + "," + input2 + "," + input3 + "," + input4 + "," + input5 + "," + string(input6) + "," + string(input7) + ")"
	//fmt.Printf("The query is: %s\n", statementSQL)
	rows, err := db.Query(statementSQL)

	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return "null", "null", "null", "null", "null", "null", 0, 0, "null", "null", 0, 0, 0, 0, 0, 0
	}
	defer rows.Close()

	var theRs1 string
	var theRs2 string
	var theRs3 string
	var theRs4 string
	var theRs5 string
	var theRs6 string
	var theRs7 int
	var theRs8 int
	var theRs9 string
	var theRs10 string
	var theRs11 int
	var theRs12 int
	var theRs13 int
	var theRs14 int
	var theRs15 int
	var theRs16 int
	for rows.Next() {

		rows.Scan(&theRs1, &theRs2, &theRs3, &theRs4, &theRs5, &theRs6, &theRs7, &theRs8, &theRs9, &theRs10, &theRs11, &theRs12, &theRs13, &theRs14, &theRs15, &theRs16)
	}
	//fmt.Printf("The data is: %s,%s,%s,%s,%s,%s,%v,%v,\n", theRs1, theRs2, theRs3, theRs4, theRs5, theRs6, theRs7, theRs8)
	return theRs1, theRs2, theRs3, theRs4, theRs5, theRs6, theRs7, theRs8, theRs9, theRs10, theRs11, theRs12, theRs13, theRs14, theRs15, theRs16
}

//InsEnvTable Insert
func InsEnvTable(table string, value string, condition string) {
	db, err := sql.Open("godror", "testing/welcome1@xe")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	statementSQL := fmt.Sprintf("INSERT INTO %v VALUES (%v) %v", table, value, condition)
	wrcode.LoggingWrite(statementSQL)
	rows, err := db.Query(statementSQL)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()
}
