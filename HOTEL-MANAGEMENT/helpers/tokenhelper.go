package helpers

import (
	"HOTEL-MANAGEMENT/database"
	"HOTEL-MANAGEMENT/models"
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Email string

	Name string

	Uid string

	jwt.StandardClaims
}

var userDB = database.GetUserDB()

var secretkey = "helloooo"

func GenerateTokens(email string, uid string) (string, string, error) {
	claims := &SignedDetails{

		Email: email,
		Uid:   uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshclaims := &SignedDetails{

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretkey))
	if err != nil {
		return "", "", err
	}

	refreshtoken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims).SignedString([]byte(secretkey))
	if err != nil {
		return "", "", err
	}

	return token, refreshtoken, err
}

func UpdateAllTokens(signedtoken *string, signedrefreshtoken *string, userid string) error {

	var user models.User

	err := userDB.Where("user_id=?", userid).First(&user)

	if err != nil {
		newuser := models.User{
			User_id:       userid,
			Token:         signedtoken,
			Refresh_token: signedrefreshtoken,
			Updated_at:    time.Now(),
		}
		if err := userDB.Create(&newuser); err != nil {
			log.Panic("Error while creating new user")
			return err.Error
		}

	} else {

		user.Token = signedtoken
		user.Refresh_token = signedrefreshtoken
		user.Updated_at = time.Now()

		if err := userDB.Save(&user); err != nil {

			log.Panic("Error while updating the user:", err)
			return err.Error
		}
	}
	return nil

}

func ValidateTokens(signedtoken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(

		signedtoken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretkey), nil
		},
	)

	claims, ok := token.Claims.(*SignedDetails)

	if !ok {

		msg = fmt.Sprintf(err.Error(), "Token is Invalid")

		return nil, msg
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {

		msg = fmt.Sprintf(err.Error(), "Token is Expired")

		return nil, msg
	}
	return claims, msg

}
