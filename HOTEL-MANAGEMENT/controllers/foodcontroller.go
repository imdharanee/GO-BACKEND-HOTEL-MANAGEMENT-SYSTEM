package controllers

import (
	"HOTEL-MANAGEMENT/database"
	"HOTEL-MANAGEMENT/models"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	fooddb = database.GetFoodDB()
	menudb = database.GetMenuDB()
)

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

		var foods []models.Food

		if err := fooddb.Find(&foods); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, foods)

	}
}
func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		obj := c.Query("food_id")

		if obj != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "food not found"})
			return
		}
		c.JSON(http.StatusOK, obj)
	}
}
func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

		var updfood models.Food

		foodid := c.Param("food_id")

		if err := c.BindJSON(&updfood); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var myfood models.Food

		if err := fooddb.Where("food_id=?", foodid).First(&myfood); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "food not found "})
			return
		}

		if myfood.Name != nil {
			myfood.Name = updfood.Name
		}
		if myfood.Price != nil {
			myfood.Price = updfood.Price
		}
		if myfood.Menu_id != nil {
			var updmenu models.Menu

			if err := menudb.Where("menu_id=?", myfood.Menu_id).First(&updmenu); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "food is not found in the menu"})
				return
			}
			myfood.Menu_id = updfood.Menu_id
		}
		myfood.Updated_at = time.Now()

		if err := fooddb.Save(&myfood); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cant update your food"})
		}

		c.JSON(http.StatusOK, updfood)

	}
}
func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var food models.Food

		var menu models.Menu

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		menuid := food.Menu_id

		if err := menudb.Where("menu_id=?", menuid).First(&menu); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "menu not found"})
			return
		}
		food.Created_at = time.Now()
		food.Updated_at = time.Now()
		if err := fooddb.Create(&food).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Food item was not created"})
			return
		}
		c.JSON(http.StatusOK, food)

	}

}
