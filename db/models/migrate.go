package models

import "github.com/franzwilhelm/uvisst-api/db"

// Migrate does migrations
func Migrate() error {
	return db.Default.AutoMigrate(&Note{}).Error
}
