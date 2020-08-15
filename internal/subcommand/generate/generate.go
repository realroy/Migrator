package generate

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/realroy/migrator/internal/util"
)

// Execute Generate migration
func Execute(a *Args) {
	timestamp := time.Now().Unix()
	parseTimestamp := strconv.Itoa(int(timestamp))
	migrationName := a.Name + parseTimestamp

	filename := fmt.Sprintf("./migration/%s.go", migrationName)

	_, err := os.Stat("./migration")
	if os.IsNotExist(err) {
		os.Mkdir("./migration", 0777)
	}

	log.Printf("generating migration %s ...\n", filename)

	template := defaultTemplate
	content := util.ParseTemplate(template, data{
		VariableName: migrationName,
		Timestamp:    parseTimestamp,
	})

	err = ioutil.WriteFile(filename, content, 0777)
	if err != nil {
		log.Fatalln("generate migration failed:", err)
	}

}
