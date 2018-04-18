package main

import (
	"database/sql"
	"fmt"
)



func InsertToAccount(db *sql.DB, account_ID int, balance float64) (success bool) {
	return InsertToTable(db, `account`, []string{`account_ID`, `balance`}, []string{IntToString(account_ID), Float64ToString(balance)})
}

func InsertToSymbol(db *sql.DB, symbol_name string) (success bool) {
	return InsertToTable(db,`symbol`,[]string{`symbol_name`},[]string{symbol_name})
}

func InsertToAccountShare(db *sql.DB, account_ID int, symbol_ID int, share float64) (success bool) {
	return InsertToTable(db,`accountShare`,[]string{`account_ID`, `symbol_ID`, `share`},[]string{IntToString(account_ID), IntToString(symbol_ID), Float64ToString(share)})
}

func InsertToContractIDToAccountID(db *sql.DB, account_ID int) (success bool) {
	return InsertToTable(db, `contractIDToAccountID`, []string{`account_ID`}, []string{IntToString(account_ID)})
}

func InsertToContract(db *sql.DB, contract_ID int, account_ID int, symbol_ID int, price float64, amount float64, contract_type string) (success bool) {
	return InsertToTable(db, `contract`, []string{`contract_ID`, `account_ID`, `symbol_ID`, `price`, `amount`, `contract_type`}, []string{IntToString(contract_ID), IntToString(account_ID), IntToString(symbol_ID), Float64ToString(price), Float64ToString(amount), contract_type})
}

func InsertToExecutedContract(db *sql.DB, contract_ID int, account_ID int, symbol_ID int, price float64, amount float64) (success bool) {
	return InsertToTable(db, `executedContract`, []string{`contract_ID`, `account_ID`, `symbol_ID`, `price`, `amount`}, []string{IntToString(contract_ID), IntToString(account_ID), IntToString(symbol_ID), Float64ToString(price), Float64ToString(amount)})
}

func InsertToCancelledContract(db *sql.DB, contract_ID int, account_ID int, symbol_ID int, price float64, amount float64) (success bool) {
	return InsertToTable(db, `cancelledContract`, []string{`contract_ID`, `account_ID`, `symbol_ID`, `price`, `amount`}, []string{IntToString(contract_ID), IntToString(account_ID), IntToString(symbol_ID), Float64ToString(price), Float64ToString(amount)})
}





func InsertToTable(db *sql.DB, tableName string, attrs []string, vals []string) (success bool) {
	sql := sqlInsert{
		TableName: tableName,
		Attrs:  attrs,
		Vals:   vals,
	}
	fmt.Println(sql.Attrs,sql.Vals)
	success = ExecSqlInsert(db, sql)
	return
}



