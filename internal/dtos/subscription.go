package dtos

import (
	"github.com/google/uuid"
)

type Subscription struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Name      string    `json:"service_name"`
	Price     int32     `json:"price"`
	StartDate string    `json:"start_date"`
	EndDate   string    `json:"end_date"`
}

type CreateResponse struct {
	ID uuid.UUID `json:"id"`
}

type UpdateRequest struct {
	Name      string    `json:"service_name"`
	Price     int32     `json:"price"`
	StartDate string    `json:"start_date"`
	EndDate   string    `json:"end_date"`

	UpdateName      bool `json:"update_service_name"`
	UpdatePrice     bool `json:"update_price"`
	UpdateStartDate bool `json:"update_start_date"`
	UpdateEndDate   bool `json:"update_end_date"`
}
