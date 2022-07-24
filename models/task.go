package models

import (
	"TaskManager/helpers"
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID          uint      `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `gorm:"not null;unique"`
	Done        bool      `gorm:"default:false"`
	CheckListID uint      `gorm:"not null"`
	CreatedAt   time.Time `gorm:"default: CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default: CURRENT_TIMESTAMP" json:"updated_at"`
}

func (task *Task) PrepareForSave() {
	task.Title = helpers.Prepare(task.Title)
}

func (task *Task) CreateTask(checkListID uint, db *gorm.DB) (*Task, error) {
	var err error
	err = db.Debug().Create(task).Error
	if err == nil {
		return task, err
	}

	return &Task{}, err
}

func (task *Task) UpdateTask(_task Task, db *gorm.DB) (*Task, error) {
	var err error
	err = db.Debug().Model(task).Updates(_task).Error
	return task, err
}

func (task *Task) DeleteTask(db *gorm.DB) (*Task, error) {
	var err error
	err = db.Debug().Delete(task).Error

	return task, err
}
