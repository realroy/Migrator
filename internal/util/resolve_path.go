package util

import (
	"log"
	"os"
	p "path"
)

// ResolvePath ...
func ResolvePath(path string) string {
	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatalln("cannot get working directory:", err)
	}

	return p.Join(workingDirectory, path)
}
