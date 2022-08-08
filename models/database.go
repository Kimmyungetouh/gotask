package models

import "gorm.io/gorm"

func RunMigrations(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{}, &CheckList{}, &Task{})
	return db
}
