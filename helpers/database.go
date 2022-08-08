package helpers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDb() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=taskmanager port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func ConnectDatabase() *gorm.DB {
	db, err := GetDb()
	if err != nil {
		panic(any("Error when connecting to the database !"))
	}

	return db
}
