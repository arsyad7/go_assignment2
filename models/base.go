package models

import "time"

type BaseOrder struct {
	ID        uint       `gorm:"primary_key" json:"order_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

type BaseItem struct {
	ID        uint       `gorm:"primary_key" json:"lineItemId"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}
