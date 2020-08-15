package generate

var defaultTemplate string = `package migration

import (
	"gopkg.in/gormigrate.v1"
	"github.com/jinzhu/gorm"
)

var {{.VariableName}} = &gormigrate.Migration {
	ID: "{{.Timestamp}}",
	Migrate: func(tx *gorm.DB) error {
		// it's a good pratice to copy the struct inside the function,
		// so side effects are prevented if the original struct changes during the time
		type Model struct {
			gorm.Model
		}

		return tx.AutoMigrate(&Model{}).Error
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.DropTable("model").Error
	},
}
`
