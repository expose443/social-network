package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	postgresData := ""
	db, err := gorm.Open(postgres.Open(postgresData))
	if err != nil {
		return nil, err
	}
	return db, nil
}
