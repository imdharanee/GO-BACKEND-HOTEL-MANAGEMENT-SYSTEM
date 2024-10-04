package models

import (
	"time"
)

type Menu struct {
	ID         uint      `gorm:"primaryKey" json:"_id"`
	Name       string    `json:"name" validate:"required"`
	Category   string    `json:"category" validate:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Menu_id    string    `json:"menu_id"`
}
