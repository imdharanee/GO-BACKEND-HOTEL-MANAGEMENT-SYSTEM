package routes

import (
	"github.com/gin-gonic/gin"

	controllers "HOTEL-MANAGEMENT/controllers"
)

func UserRoutes(incomingroutes *gin.Engine) {
	incomingroutes.GET("/users", controllers.GetUsers())
	incomingroutes.GET("/users/:user_id", controllers.GetUser())
	incomingroutes.POST("/users/signup", controllers.Signup())
	incomingroutes.POST("/users/login", controllers.Login())

}
