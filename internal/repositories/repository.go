package repositories

import (
	"errors"
	"subs-app/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrSubNotFound = errors.New("subscription not found")

type Repository interface {
	Create(sub *models.Subscription) (uuid.UUID, error)
	GetSubById(id uuid.UUID) (*models.Subscription, error)
	UpdateSubById(sub *models.Subscription) error
	DeleteSubById(id uuid.UUID) error
	// todo: FilterSubs()
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(sub *models.Subscription) (uuid.UUID, error) {
	err := r.db.Create(sub).Error
	if err != nil {
		return uuid.Nil, err
	}
	return sub.ID, nil
}

func (r *repository) GetSubById(id uuid.UUID) (*models.Subscription, error) {
	var sub models.Subscription
	if err := r.db.First(&sub, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrSubNotFound
		}
		return nil, err
	}
	return &sub, nil
}

func (r *repository) UpdateSubById(sub *models.Subscription) error {
	// todo
	return nil
}

func (r *repository) DeleteSubById(id uuid.UUID) error {
	// todo
	return nil
}
