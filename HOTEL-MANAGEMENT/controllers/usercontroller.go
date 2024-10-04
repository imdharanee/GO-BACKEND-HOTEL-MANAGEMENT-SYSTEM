package controllers

import (
	"HOTEL-MANAGEMENT/database"
	"HOTEL-MANAGEMENT/helpers"
	"HOTEL-MANAGEMENT/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var userDB = database.GetUserDB()

func GetUsers() gin.HandlerFunc {

	return func(c *gin.Context) {
		var users []models.User

		if err := c.BindJSON(&users); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userid := c.Param("user_id")

		var user models.User

		if err := userDB.Where("user_id=?", userid).First(&user); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return

		}
		c.JSON(http.StatusOK, user)

	}
}
func Signup() gin.HandlerFunc {

	return func(c *gin.Context) {
		var myuser models.User
		if err := c.BindJSON(&myuser); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		email := myuser.Email
		uid := myuser.User_id

		var existing models.User
		if err := userDB.Where("email=?", email).First(&existing); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "mail already exists"})
			return
		}
		token, refreshtoken, _ := helpers.GenerateTokens(*email, uid)

		existing.Token = &token
		existing.Refresh_token = &refreshtoken
		existing.Created_at = time.Now()
		existing.Updated_at = time.Now()

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if err := c.BindJSON(&user); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var count int64
		if err := userDB.Where("email=?", user.Email).Count(&count); err != nil {

			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error while checking the email"})
			return
		}
		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email already exists"})
			return
		}
		password := HashPass(*user.Password)

		user.Password = &password

		user.Created_at = time.Now()
		user.Updated_at = time.Now()
		user.User_id = uuid.New().String()

		token, refreshtoken, _ := helpers.GenerateTokens(*user.Email, user.User_id)

		user.Token = &token

		user.Refresh_token = &refreshtoken

		if err := userDB.Create(&user); err != nil {
			msg := "User was not created"
			c.JSON(http.StatusInternalServerError, msg)
			return
		}
		c.JSON(http.StatusOK, user)

	}
}

func HashPass(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}
func VerifyPass(userpass string, providepass string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providepass), []byte(userpass))
	check := true
	msg := ""

	if err != nil {
		msg = "login or password is incorrect"
		check = false
	}
	return check, msg
}
