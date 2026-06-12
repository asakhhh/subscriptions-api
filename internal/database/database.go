package database

import (
	"fmt"
	"subs-app/internal/config"
	"subs-app/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(cfg config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DatabaseConnStr()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, fmt.Errorf("open gorm connection: %w", err)
	}
	if err = db.AutoMigrate(&models.Subscription{}); err != nil {
		return nil, fmt.Errorf("auto migrate models: %w", err)
	}
	return db, nil
}
