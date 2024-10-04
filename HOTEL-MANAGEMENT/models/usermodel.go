package models

import (
	"time"
)

type User struct {
	ID   uint    `gorm:"primaryKey" json:"_id"`
	Name *string `json:"name" validate:"required,min=2,max=100"`

	Email         *string   `json:"email" validate:"required"`
	Password      *string   `json:"password" validated:"required min=6"`
	Token         *string   `json:"token"`
	Refresh_token *string   `json:"refresh_token"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
	User_id       string    `json:"user_id"`
}
