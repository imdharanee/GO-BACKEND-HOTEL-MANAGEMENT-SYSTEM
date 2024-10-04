package routes

import (
	controllers "HOTEL-MANAGEMENT/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingroutes *gin.Engine) {
	incomingroutes.GET("/menus", controllers.GetMenus())

	incomingroutes.GET("/menus/:menu_id", controllers.GetMenu())
	incomingroutes.POST("/menus", controllers.CreateMenu())
	incomingroutes.PATCH("/menus/:menu_id", controllers.UpdateMenu())
}
