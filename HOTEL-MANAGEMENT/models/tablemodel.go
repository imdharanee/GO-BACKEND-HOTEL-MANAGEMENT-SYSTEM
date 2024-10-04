package models

import (
	"time"
)

type Table struct {
	ID           uint      `gorm:"primaryKey" json:"_id"`
	Tot_guests   *int      `json:"tot_guests" validate:"required"`
	Table_member *int      `json:"table_member" validate:"required"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
	Table_id     string    `josn:"table_id"`
}
