package main

import (
	"HOTEL-MANAGEMENT/middleware"
	"HOTEL-MANAGEMENT/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()

	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)

	routes.MenuRoutes(router)

	routes.TableRoutes(router)

	routes.OrderRoutes(router)

	routes.InvoiceRoutes(router)

	router.Run(":" + port)

}
