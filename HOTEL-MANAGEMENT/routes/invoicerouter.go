package routes

import (
	"github.com/gin-gonic/gin"

	controllers "HOTEL-MANAGEMENT/controllers"
)

func InvoiceRoutes(incomingroutes *gin.Engine) {
	incomingroutes.GET("/invoices", controllers.GetInvoices())

	incomingroutes.GET("/invoices/:invoice_id", controllers.GetInvoice())
	incomingroutes.POST("/invoices", controllers.CreateInvoice())
	incomingroutes.PATCH("/invoices/:invoices_id", controllers.UpdateInvoice())

}
