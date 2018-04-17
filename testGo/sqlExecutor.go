package main

import (
	"fmt"
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() (db *sql.DB) {
	//connect, need a password for authen, set it in psql
	connStr := "user=postgres dbname=stockuserandorder password=tryit sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	checkErr(err)
	return
}

func ExecSql(db *sql.DB, sql sqlInsert) (success bool) {
	sqlStr := sqlInsertGenerate(sql)
	fmt.Println(sqlStr)
	stmt, err := db.Prepare(sqlStr)
	checkErr(err)
	_, err = stmt.Exec()
	checkErr(err)
	success = setSuccess(err)
	return
}

func DoQuery(db *sql.DB, sql sqlQuery) (rows *sql.Rows, success bool) {
	sqlStr := sqlQueryGenerate(sql)
	rows, err := db.Query(sqlStr)
	checkErr(err)
	success = setSuccess(err)
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
