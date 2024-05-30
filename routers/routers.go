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

	router.GET("/feeds", controllers.GetAllFeed)
	authGroup.POST("/feeds", controllers.InsertFeed)
	authGroup.PUT("/feeds/:id", controllers.UpdateFeed)
	authGroup.DELETE("/feeds/:id", controllers.DeleteFeed)
	router.GET("/feeds/:id", controllers.GetDetailFeed)

	authGroup.POST("/comments", controllers.InsertComment)
	authGroup.PUT("/comments/:id", controllers.UpdateComment)
	authGroup.DELETE("/comments/:id", controllers.DeleteComment)

	authGroup.POST("/likes", controllers.InsertLike)

	return router
}
