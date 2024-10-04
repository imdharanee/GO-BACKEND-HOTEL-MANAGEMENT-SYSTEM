package models

import (
	"time"
)

type Note struct {
	ID   uint   `gorm:"primaryKey" json:"_id"`
	Text string `json:"text"`

	Title      string    `json:"title"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Note_it    string    `json:"note_id"`
}
