package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"subs-app/internal/dtos"
	"subs-app/internal/services"
	"time"

	"github.com/google/uuid"
)

type Handler struct {
	service services.Service
}

func NewHandler(service services.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateSub(w http.ResponseWriter, r *http.Request) {
	var body dtos.Subscription
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		respondError(w, http.StatusBadRequest, "invalid request body")
		return
	} else if body.UserID == uuid.Nil || body.Name == "" ||
		body.Price < 0 {
		respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	startDate, err := strToDate(body.StartDate)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	endDate := time.Time{}
	if body.EndDate != "" {
		endDate, err = strToDate(body.EndDate)
		if err != nil {
			respondError(w, http.StatusBadRequest, "invalid request body")
			return
		}
	}
	fmt.Println(startDate, endDate)

	// response, err := h.service.CreateSub()
}

func (h *Handler) GetSub(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		respondError(w, http.StatusBadRequest, "id not provided")
		return
	}
	// todo
}

func (h *Handler) UpdateSub(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		respondError(w, http.StatusBadRequest, "id not provided")
		return
	}
	// todo
}

func (h *Handler) DeleteSub(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		respondError(w, http.StatusBadRequest, "id not provided")
		return
	}
	// todo
}

func (h *Handler) AggregateSubs(w http.ResponseWriter, r *http.Request) {
	// todo
}
