package handlers

import (
	"net/http"
	"subs-app/internal/services"
)

type Handler struct {
	service services.Service
}

func NewHandler(service services.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateSub(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) GetSub(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) UpdateSub(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) DeleteSub(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) AggregateSubs(w http.ResponseWriter, r *http.Request) {}
