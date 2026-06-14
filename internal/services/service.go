package services

import (
	"subs-app/internal/dtos"
	"subs-app/internal/models"
	"subs-app/internal/repositories"

	"github.com/google/uuid"
)

type Service interface {
	CreateSub(sub *models.Subscription) (*dtos.CreateResponse, error)
	GetSub(id uuid.UUID) (*models.Subscription, error)
	UpdateSub(id uuid.UUID, sub *dtos.UpdateObject) error
	DeleteSub(id uuid.UUID) error
}

type service struct {
	repo repositories.Repository
}

func NewService(repo repositories.Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateSub(sub *models.Subscription) (*dtos.CreateResponse, error) {
	err := s.repo.CheckUserServiceDateRange(sub.UserID, sub.Name, sub.StartDate, sub.EndDate)
	if err != nil {
		return nil, err
	}
	id, err := s.repo.CreateSub(sub)
	if err != nil {
		return nil, err
	}
	return &dtos.CreateResponse{ID: id}, nil
}

func (s *service) GetSub(id uuid.UUID) (*models.Subscription, error) {
	return s.repo.GetSubById(id)
}

func (s *service) UpdateSub(id uuid.UUID, upd *dtos.UpdateObject) error {
	if upd.UpdateStartDate || upd.UpdateEndDate {
		sub, err := s.repo.GetSubById(id)
		if err != nil {
			return err
		}
		if upd.UpdateName {
			sub.Name = upd.Name
		}
		if upd.UpdateStartDate {
			sub.StartDate = upd.StartDate
		}
		if upd.UpdateEndDate {
			sub.EndDate = upd.EndDate
		}
		err = s.repo.CheckUserServiceDateRange(sub.UserID, sub.Name, sub.StartDate, sub.EndDate)
		if err != nil {
			return err
		}
	}
	return s.repo.UpdateSubById(id, upd)
}

func (s *service) DeleteSub(id uuid.UUID) error {
	return s.repo.DeleteSubById(id)
}
