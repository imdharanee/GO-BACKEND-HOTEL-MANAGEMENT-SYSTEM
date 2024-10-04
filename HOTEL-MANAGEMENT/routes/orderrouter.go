package routes

import (
	"github.com/gin-gonic/gin"

	controllers "HOTEL-MANAGEMENT/controllers"
)

func OrderRoutes(incomingroutes *gin.Engine) {
	incomingroutes.GET("/orders", controllers.GetOrders())

	incomingroutes.GET("/orders/:order_id", controllers.GetOrder())
	incomingroutes.POST("/orders", controllers.CreateOrder())
	incomingroutes.PATCH("/menus/:menu_id", controllers.UpdateOrder())
}
