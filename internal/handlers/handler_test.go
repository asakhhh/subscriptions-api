package handlers

import (
	"encoding/json"
	"fmt"
	"strings"
	"subs-app/internal/dtos"
	"testing"
	"time"
)

func TestHelper(t *testing.T) {
	fmt.Println(strToDate("02-2025"))
	fmt.Println(strToDate("ab--c"))
	fmt.Println(strToDate("13-2025"))
	fmt.Println(dateToStr(time.Date(2025, 12, 2, 0, 0, 0, 0, time.UTC)))

	var body dtos.CreateResponse
	if err := json.NewDecoder(strings.NewReader(`
	{
		"id": "00000000-0000-0000-0000-000000000000"
	}
	`)).Decode(&body); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(body)
	}

	res, err := json.Marshal(struct{}{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(res))
	}

	t.FailNow()
}
