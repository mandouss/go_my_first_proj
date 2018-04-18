package main

import (
	"log"
	"bytes"
	"text/template"
)

type xmlNode struct {
	Header string
	Keys  []string
	Vals  []string
	Content string
	KVPairs []string
}

func generateErrorXmlNode(content string, keys []string, vals []string)(xmlStr string) {
	xml := xmlNode{
		Header:"error",
		Keys:keys,
		Vals:vals,
		Content:content,
	}
	return xmlNodeToString(xml)
}

func generateXmlNode(header string, content string, keys []string, vals []string)(xmlStr string) {
	xml := xmlNode{
		Header:header,
		Keys:keys,
		Vals:vals,
		Content:content,
	}
	return xmlNodeToString(xml)
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
}

func xmlNodeToString(xml xmlNode) (xmlStr string) {
	const tmpl = `<{{.Header}}{{range .KVPairs}} {{.}}{{end}}>{{.Content}}</{{.Header}}>`

	fillXmlNodeWithKVPairs(&xml)

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


