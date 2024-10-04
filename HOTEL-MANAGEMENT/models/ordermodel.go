package models

import (
	"time"
)

type Order struct {
	ID         uint      `gorm:"primaryKey" json:"_id"`
	Order_date time.Time `json:"order_date" validate:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"upated_at"`
	Order_id   string    `json:"order_id"`
	Table_id   *string   `json:"table_id"`
}
