package models

import (
	"time"
)

type Order struct {
	BaseOrder
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item
}
