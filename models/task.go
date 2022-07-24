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

/*func (task *Task) GetTasksByCheckList(checkListID uint, db *gorm.DB) (*[]Task, error) {
	var err error
	var tasks *[]Task
	err = db.Debug().Where(&Task{CheckListID: checkListID}).Find(tasks).Error
	return tasks, err
}*/

func (task *Task) UpdateTask(taskID uint, db *gorm.DB) (*Task, error) {
	var err error
	err = db.Debug().Where(&Task{ID: taskID}).UpdateColumns(map[string]interface{}{
		"title":      task.Title,
		"updated_at": time.Now(),
	}).Error
	return task, err
}

func (task *Task) DeleteTask(taskID uint, db *gorm.DB) (*Task, error) {
	var err error
	err = db.Debug().Where(&Task{ID: taskID}).Delete(task).Error

	return task, err
}
