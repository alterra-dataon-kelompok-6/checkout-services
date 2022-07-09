package database

import (
	"checkout-services/config"
	"checkout-services/models"
)

func CheckCart(userId int) int {
	cart := models.Cart{
		User_id: userId,
	}

	query := config.DB.Where("user_id = ?", userId).Find(&cart)

	if query.RowsAffected == 0 {
		config.DB.Create(&cart)
		CheckCart(userId)
	}

	return int(cart.Cart_id)
}
