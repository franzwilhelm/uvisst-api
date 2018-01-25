package models

import (
	"time"

	"github.com/franzwilhelm/uvisst-api/db"
)

// Note ...
type Note struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Data      string    `json:"data"`
}

// GetAllNotes ...
func GetAllNotes() (notes []Note, err error) {
	return notes, db.Default.Order("id DESC").Find(&notes).Error
}

// Create ...
func (n *Note) Create() error {
	return db.Default.Create(&n).Error
}
