package config

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DbPostgresConn() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	pg, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	migrate, err := strconv.ParseBool(os.Getenv("RUN_MIGRATION"))
	if err != nil {
		migrate = false
	}

	if migrate {
		// if err := migration.PostgresMigration(pg); err != nil {
		// 	return nil, err
		// }
	}

	if os.Getenv("APP_MODE") == "development" {
		pg.Logger = logger.Default.LogMode(logger.Info)
	} else if os.Getenv("APP_MODE") == "production" {
		pg.Logger = logger.Default.LogMode(logger.Silent)
	}

	return pg, nil
}
