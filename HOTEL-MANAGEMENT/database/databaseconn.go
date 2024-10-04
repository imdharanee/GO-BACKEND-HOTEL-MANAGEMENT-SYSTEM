package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var foodDB *gorm.DB

var menuDB *gorm.DB

var orderDB *gorm.DB

var tableDB *gorm.DB

var userDB *gorm.DB

var invoiceDB *gorm.DB

func foodDBinstance() {
	dsn := "yourusername:yourpassword@tcp(127.0.0.1:3306)/yourdbname?charset=utf8mb4&parseTime=True&loc=Local"
	dbfood, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	foodDB = dbfood

}
func menuDBinstance() {
	dsn := "yourusername:yourpassword@tcp(127.0.0.1:3306)/yourdbname?charset=utf8mb4&parseTime=True&loc=Local"
	dbmenu, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	menuDB = dbmenu

}
func orderDBinstance() {
	dsn := "yourusername:yourpassword@tcp(127.0.0.1:3306)/yourdbname?charset=utf8mb4&parseTime=True&loc=Local"
	dborder, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	orderDB = dborder

}
func tableDBinstance() {
	dsn := "yourusername:yourpassword@tcp(127.0.0.1:3306)/yourdbname?charset=utf8mb4&parseTime=True&loc=Local"
	dbtable, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	orderDB = dbtable

}
func userDBinstance() {
	dsn := "yourusername:yourpassword@tcp(127.0.0.1:3306)/yourdbname?charset=utf8mb4&parseTime=True&loc=Local"
	dbuser, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	userDB = dbuser

}
func invoiceDBinstance() {
	dsn := "yourusername:yourpassword@tcp(127.0.0.1:3306)/yourdbname?charset=utf8mb4&parseTime=True&loc=Local"
	dbuser, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	userDB = dbuser

}

func GetFoodDB() *gorm.DB {
	foodDBinstance()
	return foodDB
}
func GetMenuDB() *gorm.DB {
	menuDBinstance()
	return menuDB
}
func GetOrderDB() *gorm.DB {
	orderDBinstance()
	return orderDB
}
func GetTableDB() *gorm.DB {
	tableDBinstance()
	return tableDB
}
func GetUserDB() *gorm.DB {
	userDBinstance()
	return userDB
}

func GetInvoiceDB() *gorm.DB {
	invoiceDBinstance()
	return invoiceDB
}
