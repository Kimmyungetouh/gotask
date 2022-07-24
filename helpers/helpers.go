package helpers

import (
	"TaskManager/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func UserPassFilled(username, password string) bool {
	return strings.Trim(username, "") != "" && strings.Trim(password, " ") != ""
}

func CheckUserPass(username string, password string, db *gorm.DB) bool {
	hashPassword, _ := Hash(password)
	result := db.Find(models.User{Username: username, Password: string(hashPassword)})
	if result.Error == nil {
		return true
	}
	return false
}

func Prepare(stringToPrepare string) string {
	return html.EscapeString(strings.TrimSpace(stringToPrepare))
}

func CreateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	return "", err
}