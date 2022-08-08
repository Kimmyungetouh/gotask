package models

import (
	"TaskManager/helpers"
	"gorm.io/gorm"
	"time"
)

type CheckList struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	UserID    uint      `gorm:"not null" json:"owner_id"`
	Complete  bool      `gorm:"default: false" json:"complete"`
	CreatedAt time.Time `gorm:"default: CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default: CURRENT_TIMESTAMP" json:"updated_at"`
}

func (checkList *CheckList) SaveCheckList(db *gorm.DB) error {
	err := db.Debug().Create(checkList).Error
	return err
}

func (checkList *CheckList) UpdateCheckList(_checkList CheckList, db *gorm.DB) (*CheckList, error) {
	err := db.Debug().Model(checkList).Updates(_checkList).Error

	return checkList, err
}

func (checkList *CheckList) GetTasks(db *gorm.DB) (*[]Task, error) {
	var tasks *[]Task
	err := db.Debug().Model(Task{CheckListID: checkList.ID}).Find(tasks).Error
	return tasks, err
}

func (checkList *CheckList) Delete(db *gorm.DB) error {
	err := db.Debug().Delete(checkList).Error
	return err
}

func (checkList *CheckList) DeleteCascade(db *gorm.DB) error {
	var tasks *[]Task
	err := db.Model(Task{CheckListID: checkList.ID}).Find(tasks).Error
	helpers.HandleSimpleError(err)
	for task := range *tasks {
		err = db.Delete(&task).Error
		helpers.HandleSimpleError(err)
	}

	err = db.Debug().Delete(checkList).Error

	return nil
}
