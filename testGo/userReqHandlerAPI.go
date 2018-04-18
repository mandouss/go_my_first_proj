package main

import (
	"database/sql"
	"strconv"
	"regexp"
)



func CreateAccountHandler(db *sql.DB, ID_str string, balance_str string) (xmlRes string) {
	if(isUnsignedInt(ID_str)==false || isNonNegativeFloat(balance_str)==false) {
		return generateErrorXmlNode("ID or balance is invalid",[]string{},[]string{})
	}
	ID64, _ := strconv.ParseInt(ID_str,10,64)
	ID := int(ID64)
	balance, _ := strconv.ParseFloat(balance_str, 64)
	success := CreateAccount(db,ID,balance)
	if success == false {
		return generateErrorXmlNode("Account already exists",[]string{"id"},[]string{ID_str})
	}
	return generateXmlNode("created","",[]string{"id"},[]string{ID_str})
}

func CreateSymbolHandler(db *sql.DB, sym_name string, ID_str string, amount_str string) (xmlRes string) {
	//check
	//* is symbol name valid --> return xml::error
	//* is account_ID valid --> return xml::error
	//* is amount valid --> return xml::error
	if isAlphaNumeric(sym_name) == false || isUnsignedInt(ID_str) == false || isNonNegativeFloat(amount_str) == false {
		return generateErrorXmlNode("failed to add symbol",[]string{},[]string{})
	}
	//in DB, check
	//* account_ID exists --> return xml::error
	if isAccountIDExists(db, ID_str) == false {
		return generateErrorXmlNode("Account ID does not exist",[]string{},[]string{})
	}
	//* symbol_exists or not --<
	//if exists, add extra amount to this symbol in this account
	if isSymbolExists(db,ID_str,sym_name) == true {
		sym_id := getSymIDFromSymName(db,sym_name)
		UpdateToAccountShare(db,[]string{"share=share"+"+"+amount_str},[]string{"account_id="+wrapWithSingleQuote(ID_str), "symbol_id="+wrapWithSingleQuote(IntToString(sym_id))})
		return generateXmlNode("created","Update symbol successfully",[]string{"symbol_name"},[]string{sym_name})
	}
	//if not exist, add new symbol with amount to this account
	//return xml::success
	CreateSymbolToAccount(db,stringToInt(ID_str),sym_name,stringToFloat(amount_str))
	return generateXmlNode("created","Create symbol successfully",[]string{"symbol_name"},[]string{sym_name})


}

func isUnsignedInt(input string)(bool) {
	_, err := strconv.ParseUint(input, 10, 64)
	if err!=nil {
		return false
	}
	return true
}

func isFloat(input string)(bool) {
	_, err := strconv.ParseFloat(input, 64)
	if err!=nil {
		return false
	}
	return true
}

func isNonNegativeFloat(input string)(bool) {
	number, err := strconv.ParseFloat(input, 64)
	if err!=nil {
		return false
	} else if number < 0 {
		return false
	}
	return true
}

func isAlphaNumeric(input string)(bool) {
	match, _ := regexp.MatchString("^[A-Za-z0-9]+$", input)
	return match
}

func stringToInt(input string)(res int) {
	res,_ = strconv.Atoi(input)
	return
}

func stringToFloat(input string)(res float64) {
	res,_ = strconv.ParseFloat(input, 64)
	return
}