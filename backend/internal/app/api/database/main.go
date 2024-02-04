package database

import (
	"backend/internal/app/api/pkg/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d TimeZone=%s sslmode=disable",
		config.DatabaseHost,
		config.DatabaseUser,
		config.DatabasePassword,
		config.DatabaseName,
		config.DatabasePort,
		config.DatabaseTimeZone,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); if err != nil {
		panic(err.Error())
	}
}

func Get() *gorm.DB {
	return db
}
