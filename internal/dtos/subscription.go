package dtos

import (
	"github.com/google/uuid"
)

type Subscription struct {
	UserID    uuid.UUID `json:"user_id"`
	Name      string    `json:"name"`
	Price     int32     `json:"price"`
	StartDate string    `json:"start_date"`
	EndDate   string    `json:"end_date"`
}
