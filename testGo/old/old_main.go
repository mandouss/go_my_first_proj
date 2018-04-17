package main
import (
	"fmt"
	"database/sql"
	"strings"

	_ "github.com/lib/pq"
	"strconv"
)

func main() {


	//connect, need a password for authen, set it in psql
	connStr := "user=postgres dbname=stockuserandorder password=tryit sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	checkErr(err) 

	//插入数据, use := here, to declare stmt
    // stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
    // checkErr(err)

    // stmt.Exec("astaxie", "研发部门", "2012-12-09")

    //use = here, because stmt and err have been declared
	stmt, err := db.Prepare("INSERT INTO Account (ACCOUNT_ID, BALANCE) VALUES($1,$2) RETURNING ACCOUNT_ID")
    checkErr(err)

    stmt.Exec("111111", "2200.00")

	rows, err := db.Query("SELECT * FROM Account")
	for rows.Next() {
        var account_id string
        var balance float64
        err = rows.Scan(&account_id, &balance)
        checkErr(err)
        fmt.Println(account_id)
        fmt.Println(balance)
    }

	// rows, err := db.Query("SELECT * FROM account");
	// fmt.Println(rows);

	db.Close()

	fmt.Println("end")
}

func CreateAccountXmlMaker(xmlHeader string, ID int, msg string) {
	return fmt.Sprint("<", xmlHeader, " id=\"", strconv.Itoa(int), "\">", errMsg, "<//", xmlHeader, ">" );
}


func CreateAccountXmlMakerForError(ID int, errMsg string)(xmlnode string) {
	xmlHeader = `error`;
	return CreateAccountXmlMaker(xmlHeader, ID, errMsg);
}

func CreateAccount ( ID int, balance float64)(xmlnode string){
	var xml_headers [2]string {
		`error`,
		`created`
	}
	var err_msgs [2]string = {
		`Account ID must be numberSequence`,
		`balance must be non-negative`,
		`Account already exists`
	}
	xmlnode = ``;
	if(ID < 0) {
		return CreateAccountXmlMakerForError(xml_headers[0], ID, err_msgs[0]);
	} else if (balance < 0) {
		return CreateAccountXmlMakerForError(xml_headers[0], ID, err_msgs[1]);
	}

	//connect, need a password for authen, set it in psql
	connStr := "user=postgres dbname=stockuserandorder password=tryit sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	checkErr(err)

	//插入数据, use := here, to declare stmt
	// stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
	// checkErr(err)

	// stmt.Exec("astaxie", "研发部门", "2012-12-09")

	//use = here, because stmt and err have been declared
	stmt, err := db.Prepare("INSERT INTO Account (ACCOUNT_ID, BALANCE) VALUES($1,$2) RETURNING ACCOUNT_ID")
	checkErr(err)

	stmt.Exec(strconv.Itoa(ID), strconv.FormatFloat(balance, 'f', 2, 64));
	if err != nil {
		return CreateAccountXmlMakerForError(xml_headers[0], ID, err_msgs[2]);
	}
	return CreateAccountXmlMakerForSuccess(xml_headers[1], ID)


}



func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}