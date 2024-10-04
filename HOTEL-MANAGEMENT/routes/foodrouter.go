package routes

import (
	controllers "HOTEL-MANAGEMENT/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingroutes *gin.Engine) {
	incomingroutes.GET("/foods", controllers.GetFoods())

	incomingroutes.GET("/foods/:food_id", controllers.GetFood())
	incomingroutes.POST("/foods", controllers.CreateFood())
	incomingroutes.PATCH("/foods/:food_id", controllers.UpdateFood())

}
