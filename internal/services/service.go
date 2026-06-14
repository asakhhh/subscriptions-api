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
	UpdateSub(id uuid.UUID, sub *dtos.UpdateRequest) error
	DeleteSub(id uuid.UUID) error
}

type service struct {
	repo repositories.Repository
}

func NewService(repo repositories.Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateSub(sub *models.Subscription) (*dtos.CreateResponse, error) {
	return nil, nil
}

func (s *service) GetSub(id uuid.UUID) (*models.Subscription, error) {
	return nil, nil
}

func (s *service) UpdateSub(id uuid.UUID, sub *dtos.UpdateRequest) error {
	return nil
}

func (s *service) DeleteSub(id uuid.UUID) error {
	return nil
}
