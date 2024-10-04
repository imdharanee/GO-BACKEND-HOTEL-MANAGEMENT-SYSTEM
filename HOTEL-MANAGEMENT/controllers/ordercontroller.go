package controllers

import (
	"HOTEL-MANAGEMENT/database"
	"HOTEL-MANAGEMENT/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var orderdb = database.GetOrderDB()

var tabledb = database.GetTableDB()

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {

		var orders models.Order

		if err := orderdb.Find(&orders); err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, orders)

	}
}
func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

		order_id := c.Param("order_id")

		var myorder models.Order

		if err := orderdb.Find("order_id=?", order_id).Find(&myorder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cant find the order with tht given order id"})
			return
		}
		c.JSON(http.StatusOK, myorder)
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var existingorder models.Order

		var tab models.Table

		if err := c.BindJSON(&existingorder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		if existingorder.Table_id != nil {
			if err := tabledb.Where("table_id=?", existingorder.Table_id).Find(&tab); err != nil {
				c.JSON(http.StatusInternalServerError, err)
				return
			}
		}
		existingorder.Created_at = time.Now()
		existingorder.Updated_at = time.Now()

		if err := orderdb.Save(&existingorder); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, existingorder)

	}
}
func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

		var existingorder models.Order

		var myorder models.Order

		orderid := c.Param("order_id")

		if err := c.BindJSON(&myorder); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		if err := orderdb.Where("order_id=?", orderid).First(&existingorder); err != nil {

			c.JSON(http.StatusInternalServerError, err)
			return
		}

		if myorder.Table_id != nil {

			var tab models.Table
			if err := tabledb.Where("table_id=?", existingorder.Table_id).First(&tab); err != nil {
				c.JSON(http.StatusInternalServerError, err)
				return
			}
			existingorder.Table_id = myorder.Table_id

		}
		existingorder.Updated_at = time.Now()

		if err := orderdb.Save(&existingorder); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, existingorder)

	}
}
