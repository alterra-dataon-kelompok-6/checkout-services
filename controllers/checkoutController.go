package controllers

import (
	"checkout-services/lib/database"
	"checkout-services/middlewares"
	"checkout-services/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCheckoutTotalController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	checkouts, err := database.GetCheckoutTotal(userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if checkouts == false {
		return c.JSON(http.StatusOK, models.Response{
			Status:  "success",
			Message: "your shopping cart is empty",
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "success get checkout data from your shopping cart",
		Data:    checkouts,
	})
}

func GetCheckoutByIdController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	cartItemId, _ := strconv.Atoi(c.Param("id"))
	checkout, err := database.GetCheckoutTotalById(cartItemId, userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if checkout == false {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: "your requested data was not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "success get checkout data from your shopping cart",
		Data:    checkout,
	})
}

// func PostCheckoutController(c echo.Context) error {
// 	url := "http://54.179.213.175:8089/carts"

// 	// Create a Bearer string by appending string access token
// 	var bearer = "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTc1NDkyMjUsInVzZXJJZCI6MX0.btcMQNtoqVcvpM_fBh5SPh4mBwJ85K50kKlsv7bKIs4"

// 	// Create a new request using http
// 	req, err := http.NewRequest("GET", url, nil)

// 	// add authorization header to the req
// 	req.Header.Add("Authorization", bearer)

// 	// Send req using http Client
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Println("Error on response.\n[ERROR] -", err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Println("Error while reading the response bytes:", err)
// 	}
// 	log.Println(string([]byte(body)))

// 	userId := middlewares.ExtractTokenUserId(c)
// 	// userId :=
// 	cartItemId := c.FormValue("cart_id")

// 	// voucherCode := c.FormValue("voucher_code")

// 	// validasi jika menggunakan kode voucher
// 	// if voucherCode != "" {
// 	// 	validateVoucher, errString := database.ValidateUserVoucher(userId, voucherCode)
// 	// 	if !validateVoucher {
// 	// 		return c.JSON(http.StatusBadRequest, models.Response{
// 	// 			Status:  "fail",
// 	// 			Message: errString,
// 	// 		})
// 	// 	}
// 	// }

// 	if cartItemId == "all" || cartItemId == "" {
// 		cartItemId = "0"
// 	}

// 	cartItemIdInt, _ := strconv.Atoi(cartItemId)
// 	checkout, err := database.CheckoutItem(cartItemIdInt, userId)

// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}

// 	if checkout == false {
// 		return c.JSON(http.StatusBadRequest, models.Response{
// 			Status:  "fail",
// 			Message: "your requested data was not found",
// 		})
// 	}

// 	return c.JSON(http.StatusOK, models.Response{
// 		Status:  "success",
// 		Message: "success checkout product from your shopping cart",
// 		Data:    checkout,
// 	})
// }

func PostCheckoutController(c echo.Context) error {
	authToken := c.Request().Header.Get("Token")

	url := "http://54.179.213.175:8089/carts"

	// Create a Bearer string by appending string access token
	// var bearer = "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTc1NDkyMjUsInVzZXJJZCI6MX0.btcMQNtoqVcvpM_fBh5SPh4mBwJ85K50kKlsv7bKIs4"
	var bearer = "Bearer " + authToken

	// Create a new request using http
	req, err := http.NewRequest("DELETE", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println(string([]byte(body)))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if req == nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: "your requested data was not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "success checkout product from your shopping cart",
		Data:    body,
	})
}
