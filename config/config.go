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
		"DB_Username": "",
		"DB_Password": "",
		"DB_Port":     "",
		"DB_Host":     "",
		"DB_Name":     "",
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
