package dtos

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID        uuid.UUID `json:"id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	UserID    uuid.UUID `json:"user_id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	Name      string    `json:"service_name" example:"Yandex Plus"`
	Price     int32     `json:"price" example:"400"`
	StartDate string    `json:"start_date" example:"07-2025"`
	EndDate   string    `json:"end_date,omitempty" example:"12-2025"`
}

type CreateResponse struct {
	ID uuid.UUID `json:"id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
}

type UpdateRequest struct {
	Name      string `json:"service_name" example:"Yandex Plus"`
	Price     int32  `json:"price" example:"450"`
	StartDate string `json:"start_date" example:"07-2025"`
	EndDate   string `json:"end_date" example:"12-2025"`

	UpdateName      bool `json:"update_service_name" example:"false"`
	UpdatePrice     bool `json:"update_price" example:"true"`
	UpdateStartDate bool `json:"update_start_date" example:"false"`
	UpdateEndDate   bool `json:"update_end_date" example:"false"`
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
	Total int32           `json:"total" example:"2400"`
}
