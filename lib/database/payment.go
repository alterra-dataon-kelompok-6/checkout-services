package database

import (
	"checkout-services/config"
	"checkout-services/models"
	"strconv"
	"time"
)

func CreatePaymentId() string {
	UniqueID := models.Uuid{}
	config.DB.Raw("SELECT uuid() as uuid").Scan(&UniqueID)

	return UniqueID.Uuid
}

func GetPaymentMethod() (interface{}, error) {
	var paymentMethod []models.Payment_method

	if err := config.DB.Find(&paymentMethod).Error; err != nil {
		return nil, err
	}

	return paymentMethod, nil
}

func GetPaymentHistory(userId int) (interface{}, error) {
	var payment []models.Payment_history
	query := config.DB.Raw("SELECT payment_id, payment_method_name, amount, payed_at FROM payments LEFT JOIN payment_methods ON payments.payment_method = payment_methods.payment_method_id WHERE user_id = ? AND payment_status = 1", userId).Scan(&payment)

	if err := query.Error; err != nil {
		return nil, err
	}

	if query.RowsAffected == 0 {
		return false, nil
	}

	return payment, nil
}

func GetPaymentDetails(paymentId string) (interface{}, error) {
	var paymentDetails []models.Payment_item
	query := config.DB.Where("payment_id = ?", paymentId).Find(&paymentDetails)

	if err := query.Error; err != nil {
		return nil, err
	}

	if query.RowsAffected == 0 {
		return false, nil
	}

	return paymentDetails, nil
}

func GetPendingPayment(userId int) (interface{}, error) {
	var payment []models.Payment
	query := config.DB.Where("user_id = ? AND payment_status = 0", userId).Find(&payment)

	if err := query.Error; err != nil {
		return nil, err
	}

	if query.RowsAffected == 0 {
		return false, nil
	}

	return payment, nil
}

func PostPayment(userId int, payloadData map[string]string) (interface{}, error) {
	payment := models.Payment{}
	payment_method, _ := strconv.Atoi(payloadData["payment_method"])
	amount, _ := strconv.Atoi(payloadData["amount"])
	check := config.DB.Where("payment_id = ? AND user_id = ? AND amount = ? AND payment_status = 0", payloadData["payment_id"], userId, amount).Find(&payment)
	isValid := CheckPaymentMethod(payment_method)

	if !isValid {
		// jika payment method tidak valid
		return "Payment method invalid", nil
	}

	if check.Error != nil {
		return nil, check.Error

	} else if check.RowsAffected == 0 {
		checkPayStatus := config.DB.Where("payment_id = ? AND user_id = ? AND amount = ? AND payment_status = 1", payloadData["payment_id"], userId, amount).Find(&payment)

		if checkPayStatus.RowsAffected > 0 {
			return "This bill has been paid", nil
		}

		return "The bill you want to pay was not found", nil
	}

	paymentUpdate := models.Payment{
		Payment_method: payment_method,
		Payed_at:       time.Now(),
		Payment_status: 1,
	}

	config.DB.Where("payment_id = ?", payloadData["payment_id"]).Updates(&paymentUpdate)

	return paymentUpdate, nil
}

func CheckPaymentMethod(id int) bool {
	paymentMethod := models.Payment_method{}
	query := config.DB.Where("payment_method_id = ? ", id).Find(&paymentMethod)

	// if query.RowsAffected == 0 {
	// 	return false
	// }
	// return true

	return query.RowsAffected != 0
}
