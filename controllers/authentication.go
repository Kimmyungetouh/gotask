package controllers

import (
	"TaskManager/helpers"
	"TaskManager/models"
	"TaskManager/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signin(context *gin.Context) {
	var input schemas.LoginInput
	db, dbError := helpers.GetDb()
	fmt.Println(input)
	helpers.HandleError(context, "Error when connecting to the database !", dbError)

	bindingError := context.ShouldBindJSON(&input)

	helpers.HandleError(context, "Input binding failed", bindingError)

	user := models.User{Username: input.Username, Password: input.Password}
	userGot, userCheckingError := models.CheckUserPass(user.Username, user.Password, db)
	if userCheckingError != nil {
		context.JSON(http.StatusNotFound, schemas.LoginResponse{Content: "Credentials not found", Error: userCheckingError.Error()})
		return
	}

	token, tokenCreationError := helpers.CreateToken(userGot.ID)
	if tokenCreationError != nil {
		context.JSON(http.StatusExpectationFailed, gin.H{
			"detail": "Error when creating token",
			"err":    tokenCreationError.Error(),
		})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{"token": token})
		return
	}

}

func Logout(context *gin.Context) {

}

func Signup(context *gin.Context) {
	var input schemas.LoginInput
	db, _err := helpers.GetDb()

	if _err != nil {
		panic(_err.Error())
	}

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"detail": "Something is wrong with your inputs",
			"err":    err.Error(),
		})
		return
	} else {
		user := models.User{Username: input.Username, Password: input.Password}
		if user.Validate("create") {
			userSaved, err := user.Save(db)
			helpers.HandleError(context, "Registration Failed", err)
			token, tokenErr := helpers.CreateToken(userSaved.ID)
			helpers.HandleError(context, "Token creation failed !", tokenErr)
			context.JSON(http.StatusOK, gin.H{
				"token": token,
			})
			return
		}
	}
}

func AllUsers(context *gin.Context) {
	db, _err := helpers.GetDb()
	var users []models.User

	helpers.HandleSimpleError(_err)

	/*_users :=*/
	db.Find(&users)

	context.JSON(http.StatusOK, gin.H{
		"data": users,
	})

	return
}
