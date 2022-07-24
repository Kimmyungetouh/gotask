package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
}

type CheckList struct {
	gorm.Model
	Name    string `gorm:"not null"`
	OwnerID uint   `gorm:"not null"`
}

type Task struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Done        bool   `gorm:"default:false"`
	CheckListID uint   `gorm:"not null"`
}
