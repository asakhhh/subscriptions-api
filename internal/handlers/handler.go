package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"subs-app/internal/dtos"
	"subs-app/internal/models"
	"subs-app/internal/repositories"
	"subs-app/internal/services"
	"subs-app/internal/utils"
	"time"

	"github.com/google/uuid"
)

type Handler struct {
	service services.Service
}

func NewHandler(service services.Service) *Handler {
	return &Handler{service: service}
}

// CreateSub creates a new subscription.
//
// @Summary      Create subscription
// @Description  Create a subscription record. Dates use MM-YYYY format. End date is optional.
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        body  body      dtos.Subscription  true  "Subscription payload"
// @Success      201   {object}  dtos.CreateResponse
// @Failure      400   {object}  ErrorResponse
// @Failure      409   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /create_subscription [post]
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
	startDate, err := utils.StrToDate(body.StartDate)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	var endDate *time.Time = nil
	if body.EndDate != "" {
		endDate = &time.Time{}
		*endDate, err = utils.StrToDate(body.EndDate)
		if err != nil || endDate.Before(startDate) {
			respondError(w, http.StatusBadRequest, "invalid date")
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

// GetSub returns a subscription by ID.
//
// @Summary      Get subscription
// @Description  Get a single subscription by UUID.
// @Tags         subscriptions
// @Produce      json
// @Param        id    query     string  true  "Subscription UUID"
// @Success      200   {object}  dtos.Subscription
// @Failure      400   {object}  ErrorResponse
// @Failure      404   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /subscriptions [get]
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
		StartDate: utils.DateToStr(&response.StartDate),
		EndDate:   utils.DateToStr(response.EndDate),
	})
}

// UpdateSub partially updates a subscription.
//
// @Summary      Update subscription
// @Description  Update selected fields. Set update_* flags to true for fields you want to change.
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        id    query     string             true  "Subscription UUID"
// @Param        body  body      dtos.UpdateRequest true  "Update payload"
// @Success      200
// @Failure      400   {object}  ErrorResponse
// @Failure      404   {object}  ErrorResponse
// @Failure      409   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /subscriptions [put]
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
	startDate := time.Time{}
	if body.StartDate != "" {
		startDate, err = utils.StrToDate(body.StartDate)
		if err != nil {
			respondError(w, http.StatusBadRequest, "invalid date")
			return
		}
	} else if body.UpdateStartDate {
		respondError(w, http.StatusBadRequest, "invalid date")
		return
	}
	var endDate *time.Time = nil
	if body.EndDate != "" {
		endDate = &time.Time{}
		*endDate, err = utils.StrToDate(body.EndDate)
		if err != nil || endDate.Before(startDate) {
			respondError(w, http.StatusBadRequest, "invalid date")
			return
		}
	}
	if body.UpdatePrice && body.Price < 0 {
		respondError(w, http.StatusBadRequest, "invalid price")
		return
	}
	err = h.service.UpdateSub(id, &dtos.UpdateObject{
		Name:            body.Name,
		Price:           body.Price,
		StartDate:       startDate,
		EndDate:         endDate,
		UpdateName:      body.UpdateName,
		UpdatePrice:     body.UpdatePrice,
		UpdateStartDate: body.UpdateStartDate,
		UpdateEndDate:   body.UpdateEndDate,
	})
	if err != nil {
		h.handleError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteSub removes a subscription.
//
// @Summary      Delete subscription
// @Description  Soft-delete a subscription by UUID.
// @Tags         subscriptions
// @Param        id    query  string  true  "Subscription UUID"
// @Success      204
// @Failure      400   {object}  ErrorResponse
// @Failure      404   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /subscriptions [delete]
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

// AggregateSubs calculates total subscription cost for a period.
//
// @Summary      Aggregate subscriptions
// @Description  Sum monthly prices for subscriptions overlapping the selected period. Use list_subs=true to also return matching records.
// @Tags         subscriptions
// @Produce      json
// @Param        user_id       query  string  false  "Filter by user UUID"
// @Param        service_name  query  string  false  "Filter by service name"
// @Param        min_date      query  string  false  "Period start (MM-YYYY)"
// @Param        max_date      query  string  false  "Period end (MM-YYYY)"
// @Param        list_subs     query  string  false  "Include subscriptions in response (true/false/1/0)"
// @Success      200   {object}  dtos.AggregateResult
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /subscriptions/aggregate [get]
func (h *Handler) AggregateSubs(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	userIDStr := params.Get("user_id")
	userID := uuid.UUID{}
	var err error = nil
	if userIDStr != "" {
		userID, err = uuid.Parse(userIDStr)
		if err != nil {
			respondError(w, http.StatusBadRequest, "user_id doesn't follow UUID format")
			return
		}
	}

	name := params.Get("service_name")

	minDateStr := params.Get("min_date")
	minDate := time.Time{}
	if minDateStr != "" {
		minDate, err = utils.StrToDate(minDateStr)
		if err != nil {
			respondError(w, http.StatusBadRequest, "invalid date")
			return
		}
	}

	maxDateStr := params.Get("max_date")
	maxDate := time.Time{}
	if maxDateStr != "" {
		maxDate, err = utils.StrToDate(maxDateStr)
		if err != nil || maxDate.Before(minDate) {
			respondError(w, http.StatusBadRequest, "invalid date")
			return
		}
	}

	listSubsStr := strings.ToLower(params.Get("list_subs"))
	listSubs := false
	switch listSubsStr {
	case "true", "1":
		listSubs = true
	case "false", "0", "":
		listSubs = false
	default:
		respondError(w, http.StatusBadRequest, "invalid list_subs value")
		return
	}

	response, err := h.service.AggregateSubs(&dtos.FilterObject{
		UserID:   userID,
		Name:     name,
		MinDate:  minDate,
		MaxDate:  maxDate,
		ListSubs: listSubs,
	})
	if err != nil {
		h.handleError(w, err)
		return
	}
	respondJSON(w, http.StatusOK, response)
}

func (h *Handler) handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, utils.ErrInvalidDate):
		respondError(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, repositories.ErrTimeRangeOverlap):
		respondError(w, http.StatusConflict, err.Error())
	case errors.Is(err, repositories.ErrSubNotFound):
		respondError(w, http.StatusNotFound, err.Error())
	default:
		respondError(w, http.StatusInternalServerError, "internal server error")
	}
}
