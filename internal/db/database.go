package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	postgresData := "host=localhost user=postgres password=secret dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(postgresData))
	if err != nil {
		return nil, err
	}
	return db, nil
}
