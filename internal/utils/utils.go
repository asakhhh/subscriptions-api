package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var ErrInvalidDate = fmt.Errorf("invalid date")

func StrToDate(date string) (time.Time, error) {
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

func DateToStr(date *time.Time) string {
	if date == nil {
		return ""
	}
	return fmt.Sprintf("%02d-%d", int(date.Month()), date.Year())
}

func OverlapMonths(l1, r1, l2, r2 time.Time) int32 {
	date1 := l1
	if !l2.IsZero() && l1.Compare(l2) < 0 {
		date1 = l2
	}
	if r1.IsZero() && r2.IsZero() {
		now := time.Now()
		r1 = now
		r2 = now
	} else if r1.IsZero() {
		r1 = r2
	} else if r2.IsZero() {
		r2 = r1
	}
	date2 := r1
	if r1.Compare(r2) > 0 {
		date2 = r2
	}
	m1 := int32(date1.Month())
	y1 := int32(date1.Year())
	m2 := int32(date2.Month())
	y2 := int32(date2.Year())

	if y2 < y1 || (y2 == y1 && m2 < m1) {
		return 0
	}
	return (y2-y1)*12 + (m2 - m1 + 1)
}
