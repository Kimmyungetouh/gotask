package helpers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"html"
	"net/http"
	"strings"
)

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func UserPassFilled(username, password string) bool {
	return strings.Trim(username, "") != "" && strings.Trim(password, " ") != ""
}

func Prepare(stringToPrepare string) string {
	return html.EscapeString(strings.TrimSpace(stringToPrepare))
}

func HandleError(context *gin.Context, detail string, err error) {
	if err != nil {
		context.JSON(http.StatusExpectationFailed, gin.H{
			"detail": detail,
			"error":  err.Error(),
		})
	}
	return
}

func HandleSimpleError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func CurrentUser(context *gin.Context) string {
	userID := ExtractToken(context)
	return userID
}
