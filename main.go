package main

import (
	"TaskManager/helpers"
	"TaskManager/middleware"
	"TaskManager/models"
	"TaskManager/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db := helpers.ConnectDatabase()
	db = models.RunMigrations(db)
	// Loading env variables

	//envError := godotenv.Load(".env")
	//helpers.HandleSimpleError(envError)

	router := gin.Default()

	prefix := router.Group("/api/")
	public := prefix
	routes.PublicRoutes(public)

	private := prefix
	private.Use(middleware.JWTAuthMiddleware())
	routes.PrivateRoutes(private)

	router.Run(":8000")
}
