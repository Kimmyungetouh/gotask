package controllers

import (
	"TaskManager/helpers"
	"TaskManager/models"
	"TaskManager/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Login(context *gin.Context) {
	var input schemas.LoginInput
	db, _err := helpers.GetDb(os.Getenv("DEFAULT_DATABASE"))

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, schemas.LoginResponse{
			Content: "Bad request",
			User:    nil,
		})
		return
	}

	if _err != nil {
		context.JSON(http.StatusExpectationFailed, schemas.LoginResponse{
			Content: "Error when connecting to the database !",
			User:    nil,
		})
	}

	switch context.Request.Method {
	case "GET":
		user := models.User{Username: input.Username, Password: input.Password}
		userGot, err := helpers.CheckUserPass(user.Username, user.Password, db)
		if err != nil {
			context.JSON(http.StatusNotFound, schemas.LoginResponse{Content: "Credentials not found", User: user})
		}
		token, _err := helpers.CreateToken(userGot.ID)
		if _err != nil {
			context.JSON(http.StatusExpectationFailed, gin.H{
				"detail": "Error when creating token",
				"err":    _err.Error(),
			}
			return
		}
		context.JSON(http.StatusOK, gin.H{"token": token})
		break
	case "POST":
		if err := context.ShouldBindJSON(&input); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"detail": "Something is wrong with your inputs",
				"err":    err.Error(),
			})
			return
		}
		user := models.User{Username: input.Username, Password: input.Password}
		if user.Validate("create") {
			var err error
			_, err = user.Save(db)
			if err != nil {
				context.JSON(http.StatusExpectationFailed, gin.H{
					"detail": "Registration Failed",
					"err":    err.Error(),
				})
				return
			}
		}

		return

	}
	return

}

func Logout(context *gin.Context) {

}

func Signup(context *gin.Context) {

}
