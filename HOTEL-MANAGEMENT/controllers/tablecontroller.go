package controllers

import (
	"HOTEL-MANAGEMENT/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tables []models.Table

		if err := tabledb.Find(&tables); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, tables)
	}
}
func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {

		var mytable models.Table

		tableid := c.Param("table_id")

		if err := tabledb.Where("table_id=?", tableid).First(&mytable); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, mytable)
	}
}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var existingtable models.Table

		var mytable models.Table

		if err := c.BindJSON(mytable); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		if err := tabledb.Create(&mytable); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		existingtable.Created_at = time.Now()
		existingtable.Updated_at = time.Now()
		existingtable.Table_id = uuid.New().String()
		c.JSON(http.StatusOK, existingtable)
	}
}
func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var existingtable models.Table

		var incoming models.Table

		if err := c.BindJSON(&incoming); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		tableid := c.Param("table_id")
		if err := tabledb.Where("table_id=?", tableid).First(&existingtable); err != nil {

			c.JSON(http.StatusInternalServerError, err)
			return
		}

		existingtable.Updated_at = time.Now()

		if incoming.Table_id != "" {
			existingtable.Table_id = incoming.Table_id

		}
		if existingtable.Table_member != nil {
			existingtable.Table_member = incoming.Table_member

		}
		c.JSON(http.StatusOK, existingtable)

	}
}
