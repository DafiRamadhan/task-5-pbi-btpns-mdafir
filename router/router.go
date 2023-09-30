package router

import (
    "github.com/gin-gonic/gin"
    "task-5-pbi-btpns-mdafir/controllers"
    "task-5-pbi-btpns-mdafir/middlewares"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.Use(middlewares.AuthMiddleware())

    userController := controllers.UserController{}
    photoController := controllers.PhotoController{}

    users := r.Group("/users")
    {
        users.POST("/register", userController.Register)
        users.POST("/login", userController.Login)
        users.PUT("/:userId", userController.Update)
        users.DELETE("/:userId", userController.Delete)
    }

    photos := r.Group("/photos")
    {
        photos.POST("/", photoController.Create)
        photos.GET("/", photoController.GetAll)
        photos.PUT("/:photoId", photoController.Update)
        photos.DELETE("/:photoId", photoController.Delete)
    }

    return r
}