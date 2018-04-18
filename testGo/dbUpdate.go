package main

import (
	"database/sql"
	"fmt"
)

func UpdateToAccountShare(db *sql.DB, changes []string, conditions []string) (success bool) {
	return UpdateToTable(db, "accountShare", changes, conditions)
}

func UpdateToTable(db *sql.DB, tableName string, changes []string, conditions []string) (success bool) {
	sql := sqlUpdate{
		TableName: tableName,
		Changes:  changes,
		Conditions:   conditions,
	}
	fmt.Println(sql.Changes,sql.Conditions)
	success = ExecSqlUpdate(db, sql)
	return
}