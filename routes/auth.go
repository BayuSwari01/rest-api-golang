package routes

import (
	"rest-api-golang/controllers"
	"rest-api-golang/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine)  {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/test", controllers.Test)
	r.GET("/testAuth", middleware.JWTAuthMiddleware(), controllers.TestAuth)
}