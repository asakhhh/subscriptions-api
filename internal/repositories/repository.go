package repositories

import (
	"errors"
	"subs-app/internal/dtos"
	"subs-app/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrSubNotFound = errors.New("subscription not found")
var ErrTimeRangeOverlap = errors.New("subscription time range overlaps with other")

type Repository interface {
	CreateSub(sub *models.Subscription) (uuid.UUID, error)
	GetSubById(id uuid.UUID) (*models.Subscription, error)
	CheckUserServiceDateRange(userID uuid.UUID, name string, min time.Time, max time.Time) error
	UpdateSubById(id uuid.UUID, sub *dtos.UpdateObject) error
	DeleteSubById(id uuid.UUID) error
	// todo: FilterSubs()
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateSub(sub *models.Subscription) (uuid.UUID, error) {
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

func (r *repository) CheckUserServiceDateRange(userID uuid.UUID, name string, min time.Time, max time.Time) error {
	var sub models.Subscription
	query := r.db.Model(&models.Subscription{}).
		Where("user_id = ?", userID).
		Where("name = ?", name).
		Where("deleted_at IS NULL").
		Where("end_date IS NULL OR end_date >= ?", min.Format("2006-01-02"))
	if !max.IsZero() {
		query = query.Where("start_date <= ?", max.Format("2006-01-02"))
	}
	if err := query.First(&sub).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil // we need empty range
		}
		return err
	}
	return ErrTimeRangeOverlap
}

func (r *repository) UpdateSubById(id uuid.UUID, sub *dtos.UpdateObject) error {
	columns := []string{}
	if sub.UpdateName {
		columns = append(columns, "name")
	}
	if sub.UpdatePrice {
		columns = append(columns, "price")
	}
	if sub.UpdateStartDate {
		columns = append(columns, "start_date")
	}
	if sub.UpdateEndDate {
		columns = append(columns, "end_date")
	}
	result := r.db.Model(&models.Subscription{ID: id}).Select(columns).Updates(map[string]any{
		"name":       sub.Name,
		"price":      sub.Price,
		"start_date": sub.StartDate.Format("2006-01-02"),
		"end_date":   sub.EndDate.Format("2006-01-02"),
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrSubNotFound
	}
	return nil
}

func (r *repository) DeleteSubById(id uuid.UUID) error {
	result := r.db.Delete(&models.Subscription{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrSubNotFound
	}
	return nil
}
