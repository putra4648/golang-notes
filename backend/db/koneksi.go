package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DB() (*gorm.DB, error) {
	return gorm.Open((postgres.New(postgres.Config{
		DSN:                  "host=localhost user=user password=user123 dbname=notes_db port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	})), &gorm.Config{})
}
