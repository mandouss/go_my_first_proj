package main

import (
	_ "github.com/lib/pq";
	"fmt"
)



func main() {
	db := ConnectDB()

	runFlag := []bool{true,false}

	if runFlag[0] {
		success1 := CreateOrderToAccount(db, 12345, "SPY", 20.00, -100.00)
		success2 := CreateOrderToAccount(db, 12345, "PEK", 20.00, -100.00)
		fmt.Println(success1, success2)
	}
	if runFlag[1] {
		sql := sqlQuery{
			Attrs: []string{"Symbol.SYMBOL_NAME", "Contract.AMOUNT"},
			TableNames: []string{"Contract", "Symbol"},
			Conditions: []string{"Contract.SYMBOL_ID=Symbol.SYMBOL_ID", "CONTRACT_ID='1'"},
		}
		rows, success := QueryFromTable(db,sql.Attrs,sql.TableNames,sql.Conditions)
		if success {
			for rows.Next() {
				var symbol_name string
				var amount float64
				err := rows.Scan(&symbol_name, &amount)
				checkErr(err)
				fmt.Println(symbol_name, amount)
			}
		}
	}


}




