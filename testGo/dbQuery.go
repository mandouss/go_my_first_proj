package main

import (
	"database/sql"
	"fmt"
)
func getLastContractIDbyAccountID(db *sql.DB, accountID int)(contractID int) {
	rows, success := QueryFromTable(db, []string{"contract_id"}, []string{"ContractIDToAccountID"}, []string{"account_ID="+wrapWithSingleQuote(IntToString(accountID))})
	if success {
		for rows.Next() {
			err := rows.Scan(&contractID)
			checkErr(err)
		}
	}
	rows.Close()
	return
}

func getSymIDFromSymName(db *sql.DB, symbolName string)(symbol_id int){
	rows, success := QueryFromTable(db, []string{"symbol_id"}, []string{"symbol"}, []string{"symbol_name="+wrapWithSingleQuote(symbolName)})
	if success {
		for rows.Next() {
			err := rows.Scan(&symbol_id)
			checkErr(err)
		}
	}
	rows.Close()
	return
}

func QueryFromTable(db *sql.DB,  attrs []string, tableNames []string, conditions []string) (rows *sql.Rows, success bool) {
	sql := sqlQuery{
		Attrs:  attrs,
		TableNames: tableNames,
		Conditions: conditions,
	}
	fmt.Println(sql.Attrs,sql.TableNames,sql.Conditions)

	rows, success = DoQuery(db, sql)

	return
}
