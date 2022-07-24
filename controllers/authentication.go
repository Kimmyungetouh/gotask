package controllers

import (
	"TaskManager/globals"
	"TaskManager/helpers"
	"TaskManager/schemas"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func Login(context *gin.Context) {
	session := sessions.Default(context)
	user := session.Get(globals.UserKey)
	if user != nil {
		context.JSON(http.StatusBadRequest, schemas.LoginResponse{
			Content: "You are already logged in",
			User:    user,
		})

		return
	}
	switch context.Request.Method {
	case "GET":
		context.JSON(http.StatusOK, schemas.LoginResponse{Content: "Can login", User: user})
		break
	case "POST":
		db, _ := gorm.Open(sqlite.Open("task_manager.db"), &gorm.Config{})
		username := context.PostForm("username")
		password := context.PostForm("password")

		if helpers.UserPassFilled(username, password) && helpers.CheckUserPass(username, password, db) {
			context.JSON(http.StatusOK, schemas.LoginResponse{
				Content: "Login successful",
				User:    user,
			})

			context.JSON(http.StatusNotFound, schemas.LoginResponse{
				Content: "Credentials not found",
				User:    user,
			})

		}
		return

	}

}

func Logout(context *gin.Context) {

}

func Signup(context *gin.Context) {

}
