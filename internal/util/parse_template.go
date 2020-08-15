package util

import (
	"bytes"
	"html/template"
	"log"
)

// ParseTemplate parse variables in to given file.
func ParseTemplate(filename string, data interface{}) []byte {
	t, err := template.ParseFiles(filename)
	if err != nil {
		log.Fatalln("cannot read template file:", err.Error())
	}

	buffer := &bytes.Buffer{}

	t.Execute(buffer, data)

	return buffer.Bytes()
}
