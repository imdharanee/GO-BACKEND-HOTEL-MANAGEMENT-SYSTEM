package controllers

import (
	"HOTEL-MANAGEMENT/database"
	"HOTEL-MANAGEMENT/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var invoiceDB = database.GetInvoiceDB()

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {

		var invoices []models.Invoice

		err := invoiceDB.Find(&invoices)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, invoices)

	}
}
func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var myinvoice models.Invoice
		invoiceid := c.Param("invoice_id")
		if err := invoiceDB.Find("invoice_id=?", invoiceid).First(&myinvoice); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return

		}
		c.JSON(http.StatusOK, myinvoice)
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var existinginvoice models.Invoice
		var myinvoice models.Invoice

		if err := c.BindJSON(&myinvoice); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		if err := invoiceDB.Where("invoice_id=?", myinvoice.Invoice_id).First(&existinginvoice); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		existinginvoice.Updated_at = time.Now()

		if myinvoice.Due != nil {
			*existinginvoice.Due = *myinvoice.Due
		}
		if myinvoice.Payment_method != nil {
			*existinginvoice.Payment_method = *myinvoice.Payment_method
		}
		if myinvoice.Payment_status != nil {
			*existinginvoice.Payment_status = *myinvoice.Payment_status
		}

	}
}
