package util

import (
	"log"
	"os"
	p "path"
	"strings"
)

// ResolvePath ...
func ResolvePath(path string) string {
	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatalln("cannot get working directory:", err)
	}

	s := strings.Split(workingDirectory, "/")
	result := s[len(s)-1]
	result = p.Join(result, "./migration")

	return result
}
