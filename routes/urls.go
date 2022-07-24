package routes

import (
	"TaskManager/controllers"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(g *gin.RouterGroup) {
	g.GET("/", controllers.Home)
	g.GET("/login", controllers.Login)
	g.POST("/login", controllers.Login)
	g.GET("/signup", controllers.Signup)

}

func PrivateRoutes(g *gin.RouterGroup) {
	g.GET("/logout", controllers.Logout)
	g.GET("/checklists", controllers.Checklists)
	g.GET("/checklist", controllers.ChecklistDetail)
	g.POST("/checklist", controllers.ChecklistCreate)
	g.PUT("/checklist", controllers.ChecklistUpdate)
	g.DELETE("/checklist", controllers.ChecklistDelete)
	g.GET("tasks", controllers.Tasks)
	g.GET("/task", controllers.TaskDetail)
	g.POST("/task", controllers.TaskCreate)
	g.PUT("/task", controllers.TaskUpdate)
	g.DELETE("/task", controllers.TaskDelete)
}
