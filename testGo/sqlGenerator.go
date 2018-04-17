package main

import (
	"log"
	"bytes"

	"text/template"
)

type sqlInsert struct {
	TableName string
	Attrs  []string
	Vals   []string
}

type sqlQuery struct {
	Attrs []string
	TableNames []string
	Conditions  []string
}

func sqlQueryGenerate(sql sqlQuery) (sqlString string) {
	sqlQueryPrepare(sql)
	return sqlQueryToString(sql)
}

func sqlQueryPrepare(sql sqlQuery) {
	stringsAddComma(sql.Attrs)
	stringsAddComma(sql.TableNames)
	stringsAddAnd(sql.Conditions)
}

func sqlQueryToString(sql sqlQuery) (sqlStr string) {
	const tmpl = `SELECT {{range .Attrs}}{{.}}{{end}} FROM {{range .TableNames}}{{.}}{{end}} WHERE {{range .Conditions}}{{.}}{{end}}`

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

func sqlInsertPrepare(sql sqlInsert) {
	stringsAddComma(sql.Attrs)
	stringsAddQuoteAndComma(sql.Vals)
}

func sqlInsertToString(sql sqlInsert) (sqlStr string) {
	const tmpl = `INSERT INTO {{.TableName}}({{range .Attrs}}{{.}}{{end}}) VALUES({{range .Vals}}{{.}}{{end}})`

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





func stringsAddQuoteAndComma(stringList []string) {
	stringsAddQuote(stringList)
	stringsAddComma(stringList)
}

func stringsAddAnd(stringList []string) {
	for i := 0; i < len(stringList)-1; i++ {
		stringList[i] += " AND "
	}
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

func wrapWithSingleQuote(str string) (string){
	str = "'" + str + "'"
	return str
}


