package params

import (
	"go_assignment2/models"
	"time"
)

type CreateOrder struct {
	OrderedAt    time.Time     `json:"orderedAt"`
	CustomerName string        `json:"customerName"`
	Items        []models.Item `json:"items"`
}
