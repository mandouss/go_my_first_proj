package main

import (
	"log"
	"bytes"
	"text/template"
	"fmt"
)

type xmlNode struct {
	Header string
	Keys  []string
	Vals  []string
	Content string
	KVPairs []string
}

func wrapWithDoubleQuotes(str string) string {
	return "\"" + str + "\""
}

func fillXmlNodeWithKVPairs(xml *xmlNode) {

	for i := 0; i < len(xml.Keys); i++ {
		if(i < len(xml.Vals)) {
			KVpair := xml.Keys[i] + "=" + wrapWithDoubleQuotes(xml.Vals[i])
			xml.KVPairs = append(xml.KVPairs, KVpair)
		} else {
			KVpair := xml.Keys[i] + "=" + wrapWithDoubleQuotes("")
			xml.KVPairs = append(xml.KVPairs, KVpair)
		}
	}
	fmt.Println(xml.KVPairs)
}

func xmlNodeToString(xml xmlNode) (xmlStr string) {
	const tmpl = `<{{.Header}}{{range .KVPairs}} {{.}}{{end}}>{{.Content}}</{{.Header}}>`

	fillXmlNodeWithKVPairs(&xml)
	fmt.Println(xml.KVPairs)

	t := template.New("xmlNode template")

	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, xml)
	if err != nil {
		log.Fatal("Execute: ", err)
		return
	}

	xmlStr = tpl.String()

	return
}
