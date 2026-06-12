package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Subscription struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	Name string `gorm:"size:255;not null;index" json:"name"`
	Price decimal.Decimal `gorm:"type:decimal(7, 2);not null" json:"price"`
	StartDate time.Time `gorm:"type:date;not null" json:"start_date"`
	EndDate time.Time `gorm:"type:date" json:"end_date"`
}
