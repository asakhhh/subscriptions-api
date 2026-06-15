package dtos

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Name      string    `json:"service_name"`
	Price     int32     `json:"price"`
	StartDate string    `json:"start_date"`
	EndDate   string    `json:"end_date,omitempty"`
}

type CreateResponse struct {
	ID uuid.UUID `json:"id"`
}

type UpdateRequest struct {
	Name      string `json:"service_name"`
	Price     int32  `json:"price"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`

	UpdateName      bool `json:"update_service_name"`
	UpdatePrice     bool `json:"update_price"`
	UpdateStartDate bool `json:"update_start_date"`
	UpdateEndDate   bool `json:"update_end_date"`
}

type UpdateObject struct {
	Name      string
	Price     int32
	StartDate time.Time
	EndDate   *time.Time

	UpdateName      bool
	UpdatePrice     bool
	UpdateStartDate bool
	UpdateEndDate   bool
}

type FilterObject struct {
	UserID   uuid.UUID
	Name     string
	MinDate  time.Time
	MaxDate  time.Time
	ListSubs bool
}

type AggregateResult struct {
	Subs  []*Subscription `json:"subscriptions,omitempty"`
	Total int32           `json:"total"`
}
