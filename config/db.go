package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, error) {
	dsn := "host=localhost user=scholak password=postgres dbname=go_crud_1 port=5432 sslmode=disable TimeZone=Europe/Istanbul"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}