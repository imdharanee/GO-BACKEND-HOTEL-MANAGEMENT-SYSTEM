package models

import (
	"time"
)

type Invoice struct {
	ID             uint    `gorm:"primaryKey" json:"_id"`
	Invoice_id     string  `json:"invoice_id"`
	Order_id       string  `json:"order_id"`
	Due            *int    `json:"due"`
	Payment_method *string `json:"payment_method" validate:"eq=CARD|eq=CASH|eq="`

	Payment_status *string   `json:"payment_status" validate:"required,eq=PENDING|eq=PAID"`
	Created_at     time.Time `json:"created_at"`
	Updated_at     time.Time `json:"updated_at"`
}
