package handlers

import (
	"fmt"
	"testing"
	"time"
)

func TestHelper(t *testing.T) {
	fmt.Println(strToDate("02-2025"))
	fmt.Println(strToDate("ab--c"))
	fmt.Println(strToDate("13-2025"))
	fmt.Println(dateToStr(time.Date(2025, 12, 2, 0, 0, 0, 0, time.UTC)))
	got, err := strToDate("-022025")
	if err != nil {
		t.Errorf("02-2025 gave error")
	}
	fmt.Println(got)
	t.FailNow()
}
