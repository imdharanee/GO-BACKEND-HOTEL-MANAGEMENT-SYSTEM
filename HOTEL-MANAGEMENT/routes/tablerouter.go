package routes

import (
	controllers "HOTEL-MANAGEMENT/controllers"

	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingroutes *gin.Engine) {
	incomingroutes.GET("/tables", controllers.GetTables())

	incomingroutes.GET("/tables/:table_id", controllers.GetTable())
	incomingroutes.POST("/tables", controllers.CreateTable())
	incomingroutes.PATCH("/tables/:table_id", controllers.UpdateTable())
}
