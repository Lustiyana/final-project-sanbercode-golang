package routers

import (
	"github.com/gin-gonic/gin"
	"simple-social-media/controllers"

	"simple-social-media/middleware"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	authGroup := router.Group("/")
	authGroup.Use(middleware.AuthMiddleware())

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	authGroup.GET("/feeds", controllers.GetAllFeed)
	router.POST("/feeds", controllers.InsertFeed)
	router.PUT("/feeds/:id", controllers.UpdateFeed)
	router.DELETE("/feeds/:id", controllers.DeleteFeed)
	router.GET("/feeds/:id", controllers.GetDetailFeed)

	return router
}
