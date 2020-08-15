package migrate

var defaultTemplate string = `
package main

import (
	"log"
	"github.com/realroy/migrator/internal/connection"
	"github.com/realroy/migrator/migration"
		
	"gopkg.in/gormigrate.v1"
)
		
func main() {
	db := connection.Connect()
		
	defer db.Close()
		
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{{ range .MigrationFilenames }}
      {{.}},
    {{ end }}
	})
		
	if err := m.Migrate(); err != nil {
		log.Fatalln("Could not migrate:", err)
	}
		
	log.Printf("Migration did run successfully")
}	
`