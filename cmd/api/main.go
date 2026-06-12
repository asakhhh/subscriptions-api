package main

import (
	"fmt"
	"log"
	"subs-app/internal/config"
	"subs-app/internal/database"
)

func main() {
	cfg := config.Load()

	db, err := database.New(cfg)
	if err != nil {
		log.Fatalf("database init error: %v", err)
	}
	fmt.Println(db)
}
