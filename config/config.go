package config

import (
	"checkout-services/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "admin123",
		"DB_Port":     "3306",
		"DB_Host":     "db-dataon-echo.cwx9bmizehkf.ap-southeast-1.rds.amazonaws.com",
		"DB_Name":     "kelompok_6_checkout",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func InitMigrate() {
	DB.AutoMigrate(&models.Cart_item{})
	DB.AutoMigrate(&models.Cart{})
	DB.AutoMigrate(&models.Checkout{})
	DB.AutoMigrate(&models.Payment{})
	DB.AutoMigrate(&models.Payment_history{})
	DB.AutoMigrate(&models.Payment_item{})
	DB.AutoMigrate(&models.Payment_method{})
	DB.AutoMigrate(&models.Response{})
	DB.AutoMigrate(&models.Uuid{})
}
