package controllers

import (
	"HOTEL-MANAGEMENT/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var menus []models.Menu

		if err := menudb.Find(&menus); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, menus)

	}
}
func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		menuid := c.Param("menu_id")
		var mymenu models.Menu
		if err := menudb.Where("menu_id=?", menuid).First(&mymenu); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error while reading the menu"})
			return
		}
		c.JSON(http.StatusOK, mymenu)

	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newmenu models.Menu

		if err := c.BindJSON(&newmenu); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		newmenu.Created_at = time.Now()
		newmenu.Updated_at = time.Now()
		newmenu.Menu_id = uuid.New().String()

		if err := menudb.Create(&newmenu); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, newmenu)

	}
}
func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newmenu models.Menu
		var existingmenu models.Menu

		if err := c.BindJSON(&newmenu); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		menuid := c.Param("menu_id")

		if err := menudb.Where("menu_id=?", menuid).First(&existingmenu); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		if newmenu.Category != " " {
			existingmenu.Category = newmenu.Category
		}
		existingmenu.Created_at = time.Now()
		existingmenu.Updated_at = time.Now()

		if err := menudb.Save(&existingmenu); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cant save the db"})
			return
		}
		c.JSON(http.StatusOK, existingmenu)

	}

}
