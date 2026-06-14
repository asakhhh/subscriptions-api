package handlers

import (
	"encoding/json"
	"net/http"
	"subs-app/internal/dtos"
	"subs-app/internal/models"
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
	response, err := h.service.CreateSub(&models.Subscription{
		ID:        body.ID,
		UserID:    body.UserID,
		Name:      body.Name,
		Price:     body.Price,
		StartDate: startDate,
		EndDate:   endDate,
	})
	if err != nil {
		h.handleError(w, err)
		return
	}
	respondJSON(w, http.StatusCreated, response)
}

func (h *Handler) GetSub(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if len(idStr) == 0 {
		respondError(w, http.StatusBadRequest, "id not provided")
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "id doesn't follow UUID format")
		return
	}
	response, err := h.service.GetSub(id)
	if err != nil {
		h.handleError(w, err)
		return
	}
	respondJSON(w, http.StatusOK, dtos.Subscription{
		ID:        response.ID,
		UserID:    response.UserID,
		Name:      response.Name,
		Price:     response.Price,
		StartDate: dateToStr(response.StartDate),
		EndDate:   dateToStr(response.EndDate),
	})
}

func (h *Handler) UpdateSub(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if len(idStr) == 0 {
		respondError(w, http.StatusBadRequest, "id not provided")
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "id doesn't follow UUID format")
		return
	}
	var body dtos.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	err = h.service.UpdateSub(id, &body)
	if err != nil {
		h.handleError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteSub(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if len(idStr) == 0 {
		respondError(w, http.StatusBadRequest, "id not provided")
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "id doesn't follow UUID format")
		return
	}
	err = h.service.DeleteSub(id)
	if err != nil {
		h.handleError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) AggregateSubs(w http.ResponseWriter, r *http.Request) {
	// todo
}

func (h *Handler) handleError(w http.ResponseWriter, err error) {
	// todo
}
