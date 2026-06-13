package services

import "subs-app/internal/repositories"

type Service interface {
}

type service struct {
	repo repositories.Repository
}

func NewService(repo repositories.Repository) Service {
	return &service{repo: repo}
}
