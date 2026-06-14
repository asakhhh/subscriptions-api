package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subscription struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	Name      string         `gorm:"size:255;not null;index" json:"service_name"`
	Price     int32          `gorm:"type:integer;not null" json:"price"`
	StartDate time.Time      `gorm:"type:date;not null" json:"start_date"`
	EndDate   time.Time      `gorm:"type:date" json:"end_date"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
