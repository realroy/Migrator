package util

import (
	"bytes"
	"log"
	"strings"
	"text/template"
)

// ParseTemplate parse variables in to given file.
func ParseTemplate(text string, data interface{}) []byte {
	lines := strings.Split(text, "\n")

	firstLineIsEmpty := len(strings.Trim(lines[0], "")) == 0
	if firstLineIsEmpty {
		text = strings.Join(lines[1:], "\n")
	}

	t, err := template.New("temp").Parse(text)
	if err != nil {
		log.Fatalln("cannot parse template", err.Error())
	}

	buffer := &bytes.Buffer{}

	t.Execute(buffer, data)

	return buffer.Bytes()
}
