package helpers

import (
	"TaskManager/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"strings"
)

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func UserPassFilled(username, password string) bool {
	return strings.Trim(username, "") != "" && strings.Trim(password, " ") != ""
}

func CheckUserPass(username string, password string, db *gorm.DB) (models.User, error) {
	hashPassword, _ := Hash(password)
	var user models.User
	err := db.Debug().Model(models.User{Username: username, Password: string(hashPassword)}).Find(&user).Error

	return user, err
}

func Prepare(stringToPrepare string) string {
	return html.EscapeString(strings.TrimSpace(stringToPrepare))
}

func HandleError(err error) error {
	return err
}
