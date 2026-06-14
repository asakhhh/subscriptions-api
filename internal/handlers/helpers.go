package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var ErrInvalidDate = fmt.Errorf("invalid date")

type errorResponse struct {
	Message string `json:"message"`
}

func respondJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, errorResponse{Message: message})
}

func strToDate(date string) (time.Time, error) {
	arr := strings.Split(date, "-")
	if len(arr) != 2 {
		return time.Time{}, ErrInvalidDate
	}
	monthStr, yearStr := arr[0], arr[1]
	month, err := strconv.Atoi(monthStr)
	if err != nil || month < 1 || month > 12 {
		return time.Time{}, ErrInvalidDate
	}
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return time.Time{}, ErrInvalidDate
	}
	return time.Date(year, time.Month(month), 2, 0, 0, 0, 0, time.UTC), nil
}

func dateToStr(date time.Time) string {
	return fmt.Sprintf("%02d-%d", int(date.Month()), date.Year())
}
