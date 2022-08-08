package models

import (
	"TaskManager/helpers"
	"gorm.io/gorm"
	"strings"
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"not null;unique" json:"username"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt time.Time `gorm:"default: CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default: CURRENT_TIMESTAMP" json:"updated_at"`
}

func (user *User) BeforeSave(db *gorm.DB) error {
	user.PrepareToSave()
	hashedPassword, err := helpers.Hash(user.Password)
	if err == nil {
		user.Password = string(hashedPassword)
	}

	return err
}

func (user *User) PrepareToSave() {
	user.Username = helpers.Prepare(user.Username)
	user.Password = helpers.Prepare(user.Password)
}

func (user *User) Validate(action string) bool {
	switch strings.ToLower(action) {
	case "update":
		if user.Username == "" && user.Password == "" {
			return false
		}
		break
	case "create":
		if user.Username == "" || user.Password == "" {
			return false
		}
		break
	case "login":
		if user.Username == "" || user.Password == "" {
			return false
		}
	default:
		return false
	}
	return true
}

func (user *User) Save(db *gorm.DB) (*User, error) {
	err := user.BeforeSave(db)
	if err != nil {
		return nil, err
	}
	err = db.Create(user).Error
	return user, err
}

func (user *User) GetCheckLists(db *gorm.DB) (*[]CheckList, error) {
	var checkLists *[]CheckList
	err := db.Debug().Model(CheckList{UserID: user.ID}).Find(checkLists).Error
	return checkLists, err
}

func (user *User) GetCompletedCheckLists(db *gorm.DB) (*[]CheckList, error) {
	var checkLists *[]CheckList
	err := db.Debug().Model(CheckList{UserID: user.ID, Complete: true}).Find(checkLists).Error
	return checkLists, err
}

func (user *User) GetUncompletedChecklists(db *gorm.DB) (*[]CheckList, error) {
	var checkLists *[]CheckList
	err := db.Debug().Model(CheckList{UserID: user.ID, Complete: false}).Find(checkLists).Error
	return checkLists, err
}

func (user *User) UpdateUser(_user User, db *gorm.DB) (*User, error) {
	err := user.BeforeSave(db)
	if err != nil {
		return nil, err
	}
	err = db.Debug().Model(user).Updates(_user).Error
	return user, err
}

func (user *User) DeleteUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Delete(user).Error
	return nil, err
}

func CheckUserPass(username string, password string, db *gorm.DB) (User, error) {
	hashPassword, _ := helpers.Hash(password)
	var user User
	err := db.Debug().Model(User{Username: username, Password: string(hashPassword)}).Find(&user).Error

	return user, err
}

func LoginCheck(username, password string, db *gorm.DB) (string, error) {

	var token string

	user, err := CheckUserPass(username, password, db)
	if err != nil {
		token = ""
	}
	token, err = helpers.CreateToken(user.ID)

	return token, err
}
