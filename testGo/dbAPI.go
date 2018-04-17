package main

import (
	"database/sql"
	"strconv"
)

func CreateAccount(db *sql.DB, ID int, balance float64) (success bool) {
	success = InsertToAccount(db, ID, balance)
	return
}

func CreateSymbolToAccount(db *sql.DB, ID int, symbolName string, amount float64) (success bool) {
	//query sym_id by sym_name
	symbol_id := getSymIDFromSymName(db, symbolName)
	//if not found:s_id==0
	if symbol_id == 0 {
		//	insert symbol
		InsertToSymbol(db, symbolName)
		// query sym_id again
		symbol_id = getSymIDFromSymName(db, symbolName)
	}
	//insert accountShare
	success = InsertToAccountShare(db,12345,symbol_id,10.00)

	return success
}

func CreateOrderToAccount(db *sql.DB, accountID int, symbolName string, amount float64, limit float64) (success bool) {

	//insert ID to 'contractIDtoAccountID'
	InsertToContractIDToAccountID(db, accountID)
	//get the last contractID by accountID
	contractID := getLastContractIDbyAccountID(db, accountID)
	//get symbolID by sym_name from 'symbol'
	symbolID := getSymIDFromSymName(db, symbolName)
	if symbolID == 0 {
		//if symbol_name not exist : insert sym_name
		InsertToSymbol(db,symbolName)
		//	get symbolID by sym_name from 'symbol'
		symbolID = getSymIDFromSymName(db, symbolName)
	}

	//create contract
	success = InsertToContract(db,contractID,accountID,symbolID,limit,amount,`open`)
	return
}


func IntToString(input int) string {
	return strconv.Itoa(input)
}

func Float64ToString(input float64) string {
	return strconv.FormatFloat(input, 'f', 2, 64)
}