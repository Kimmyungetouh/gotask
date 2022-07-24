package main

import "gorm.io/gorm"
import "gorm.io/driver/sqlite"
import "TaskManager/models"

func main() {
	db, err := gorm.Open(sqlite.Open("task_manager.db"), &gorm.Config{})
	if err != nil {
		panic(any("Error when connecting to the database !"))
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.CheckList{})
	db.AutoMigrate(&models.Task{})
}
