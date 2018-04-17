package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"text/template"

	_ "github.com/lib/pq"
)

type sqlInsert struct {
	DBName string
	Attrs  []string
	Vals   []string
}

func main() {
	db := ConnectDB()
	ID, status := CreateAccount(db, 87654321, 20.00)
	fmt.Print(ID)
	fmt.Print(status)

}

func stringsAddComma(stringList []string) {
	for i := 0; i < len(stringList)-1; i++ {
		stringList[i] += ","
	}
}

func stringsAddQuote(stringList []string) {
	for i := 0; i < len(stringList); i++ {
		stringList[i] = "'" + stringList[i] + "'"
	}
}

func stringsAddQuoteAndComma(stringList []string) {
	stringsAddQuote(stringList)
	stringsAddComma(stringList)
}

func sqlInsertPrepare(sql sqlInsert) {
	stringsAddComma(sql.Attrs)
	stringsAddQuoteAndComma(sql.Vals)
}

func sqlInsertToString(sql sqlInsert) (sqlStr string) {
	const tmpl = `INSERT INTO {{.DBName}}({{range .Attrs}}{{.}}{{end}}) VALUES({{range .Vals}}{{.}}{{end}})`

	t := template.New("sqlInsert template")

	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, sql)
	if err != nil {
		log.Fatal("Execute: ", err)
		return
	}

	sqlStr = tpl.String()

	return
}

func sqlInsertGenerate(sql sqlInsert) (sqlString string) {
	sqlInsertPrepare(sql)
	return sqlInsertToString(sql)
}

func ConnectDB() (db *sql.DB) {
	//connect, need a password for authen, set it in psql
	connStr := "user=postgres dbname=stockuserandorder password=tryit sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	checkErr(err)
	return
}

func ExecSql(db *sql.DB, sql string) (success bool) {
	stmt, err := db.Prepare(sql)
	checkErr(err)
	_, err = stmt.Exec()
	checkErr(err)
	success = setSuccess(err)
	if err != nil {
		success = false
	} else {
		success = true
	}
	return
}

func setSuccess(err error) bool {
	var success bool
	if err != nil {
		success = false
	} else {
		success = true
	}
	return success
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func CreateAccount(db *sql.DB, ID int, balance float64) (retID int, success bool) {
	sql := sqlInsert{
		DBName: "Account",
		Attrs:  []string{"account_ID", "balance"},
		Vals:   []string{strconv.Itoa(ID), strconv.FormatFloat(balance, 'f', 2, 64)},
	}
	fmt.Println(sql.Vals)
	sqlStr := sqlInsertGenerate(sql)
	success = ExecSql(db, sqlStr)
	return ID, success
}

func CreateSymbolToAccount(ID int, symbolName string, amount float64) (retID int, retSymbolName string, success bool) {
	retID = 12345
	retSymbolName = `SPY`
	success = true
	return
}

func CreateBuyOrderToAccount(ID int, symbolName string, amount float64, limit float64) (retID int, retSymbolName string, retAmount float64, retLimit float64, success bool) {
	retID = 12345
	retSymbolName = `SPY`
	retAmount = 20.0
	retLimit = 100.0
	success = true
	return
}

func CreateSellOrderToAccount(ID int, symbolName string, amount float64, limit float64) (retID int, retSymbolName string, retAmount float64, retLimit float64, success bool) {
	retID = 12345
	retSymbolName = `SPY`
	retAmount = 20.0
	retLimit = 100.0
	success = true
	return
}
