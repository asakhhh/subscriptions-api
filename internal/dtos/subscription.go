package dtos

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	UserID    uuid.UUID `json:"user_id"`
	Name      string    `json:"name"`
	Price     int32     `json:"price"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
