package migrate

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/realroy/migrator/internal/util"
)

// Execute database migration from migration files
func Execute(a *Args) {
	files, err := ioutil.ReadDir(defaultMigrationDir)
	if err != nil {
		log.Fatalln("cannot read dir", err)
	}

	filenames := extractFilenames(files)

	content := util.ParseTemplate(defaultTemplateFilename, data{
		MigrationFilenames: filenames,
	})

	err = ioutil.WriteFile(defaultFileName, content, 0777)
	if err != nil {
		log.Fatalln("cannot write file", err)
	}

	err = executeMigratorFile()
	if err != nil {
		log.Fatalln("Migration failed", err)
	}

	log.Println("Migration successfully")

	err = os.Remove(defaultFileName)
	if err != nil {
		log.Fatalln("cannot remove migrator file", err)
	}
}

func extractFilenames(files []os.FileInfo) []string {
	filenames := make([]string, 0)

	for _, file := range files {
		name := strings.ReplaceAll(file.Name(), ".go", "")
		variableName := fmt.Sprintf("migration.%s", name)

		filenames = append(filenames, variableName)
	}

	return filenames
}

func executeMigratorFile() error {
	cmd := exec.Command("go", "run", defaultFileName)

	_, err := cmd.Output()

	return err
}
