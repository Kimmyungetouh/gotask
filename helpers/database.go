package helpers

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDb(_db string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(_db), &gorm.Config{})
	return db, err
}
