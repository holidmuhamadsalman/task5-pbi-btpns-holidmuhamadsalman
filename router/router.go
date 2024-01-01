package router

import (
	"task5-pbi-btpns-holidmuhamadsalman/controllers"
	"task5-pbi-btpns-holidmuhamadsalman/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/api/v1/users/register", controllers.Register)
	router.POST("/api/v1/users/login", controllers.Login)

	router.Use(middlewares.Auth)

	users := router.Group("/api/v1/users")
	{
		users.GET("/", controllers.GetUsers)
		users.GET("/:id", middlewares.AuthUser, controllers.GetUserById)
		users.PUT("/:id", middlewares.AuthUser, controllers.UpdateUser)
		users.DELETE("/:id", middlewares.AuthUser, controllers.DeleteUser)
	}

	photos := router.Group("/api/v1/photos")
	{
		photos.POST("/",middlewares.AuthUser, controllers.CreatePhoto)
		photos.GET("/", controllers.GetPhotos)
		photos.GET("/:id", middlewares.AuthPhoto, controllers.GetPhotoById)
		photos.PUT("/:id", middlewares.AuthUser, controllers.UpdatePhoto)
		photos.DELETE("/:id", middlewares.AuthUser, controllers.DeletePhoto)
	}

	return router
}