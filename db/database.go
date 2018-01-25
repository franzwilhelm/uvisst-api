package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Default is the connection reference
var Default *gorm.DB

// Connect creates the connection reference
func Connect(dialect string, conn string) (err error) {
	if Default != nil {
		return nil
	}
	Default, err = gorm.Open(dialect, conn)
	Default.LogMode(true)
	return
}
