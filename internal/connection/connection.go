package connection

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Connect ...
func Connect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "database.sqlite3")
	if err != nil {
		log.Panicln("failed to connect database", err)
	}

	return db
}
