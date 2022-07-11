package controllers

import (
	"bytes"
	"checkout-services/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Body struct {
	Stock uint `json:"stock"`
}

type ResponseCart struct {
	Data   Cart `json:"data"`
	Status bool `json:"status"`
}

type ResponseProduct struct {
	Data   Product `json:"data"`
	Status bool    `json:"status"`
}

type Cart struct {
	Id         uint       `json:"id"`
	CustomerId uint       `json:"customer_id"`
	CreatedAt  string     `json:"created_at"`
	UpdatedAt  string     `json:"updated_at"`
	DeletedAt  string     `json:"deleted_at"`
	CartItems  []CartItem `json:"cart_items"`
}

type CartItem struct {
	Id        uint   `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	CartId    uint   `json:"cart_id"`
	ProductId uint   `json:"product_id"`
	Qty       uint   `json:"qty"`
}

type Product struct {
	Id          uint     `json:"id"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	DeletedAt   string   `json:"deleted_at"`
	CategoryId  uint     `json:"category_id"`
	Name        string   `json:"name"`
	Stock       uint     `json:"stock"`
	Price       uint     `json:"price"`
	Image       string   `json:"image"`
	Description string   `json:"description"`
	Category    Category `json:"category"`
}

type Category struct {
	CategoryId uint   `json:"category_id"`
	Category   string `json:"category"`
}

func getQty(token string) ([]uint, []uint) {
	var data ResponseCart
	var productId, qty []uint

	url := "http://54.179.213.175:8089/carts"

	// create a Bearer string by appending string access token
	// var bearer = "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTc1NDkyMjUsInVzZXJJZCI6MX0.btcMQNtoqVcvpM_fBh5SPh4mBwJ85K50kKlsv7bKIs4"
	var bearer = "Bearer " + token

	// create a new request using http
	req, _ := http.NewRequest("GET", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &data)

	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	for index := range data.Data.CartItems {
		productId = append(productId, data.Data.CartItems[index].ProductId)
		qty = append(qty, data.Data.CartItems[index].Qty)
	}

	return productId, qty
}

func getCurrentStock(productId string) uint {
	var data ResponseProduct

	url := "http://54.179.213.175:8088/products/" + productId

	// create a new request using http
	req, _ := http.NewRequest("GET", url, nil)

	// send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &data)

	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	stock := data.Data.Stock

	return stock
}

func getPrice(productId string) uint {
	var data ResponseProduct

	url := "http://54.179.213.175:8088/products/" + productId

	// create a new request using http
	req, _ := http.NewRequest("GET", url, nil)

	// send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &data)

	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	price := data.Data.Price

	return price
}

func updateStock(productId string, update uint) {
	body := Body{
		Stock: update,
	}

	// initialize http client
	client := &http.Client{}

	// marshal User to json
	json, err := json.Marshal(body)

	if err != nil {
		log.Println("Error in marshal json.\n[ERROR] -", err)
	}

	// set the HTTP method, url, request body
	req, err := http.NewRequest(http.MethodPut, "http://54.179.213.175:8088/products/"+productId, bytes.NewBuffer(json))

	if err != nil {
		log.Println("Error in request.\n[ERROR] -", err)
	}

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error get response.\n[ERROR] -", err)
	}

	defer resp.Body.Close()
}

func deleteCart(token string) {
	url := "http://54.179.213.175:8089/carts"

	// create a Bearer string by appending string access token
	// var bearer = "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTc1NDkyMjUsInVzZXJJZCI6MX0.btcMQNtoqVcvpM_fBh5SPh4mBwJ85K50kKlsv7bKIs4"
	var bearer = "Bearer " + token

	// create a new request using http
	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		log.Println("Error in request.\n[ERROR] -", err)
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// send req using http client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}

	defer resp.Body.Close()
}

// func GetCheckoutTotalController(c echo.Context) error {
// 	userId := middlewares.ExtractTokenUserId(c)
// 	checkouts, err := database.GetCheckoutTotal(userId)

// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}

// 	if checkouts == false {
// 		return c.JSON(http.StatusOK, models.Response{
// 			Status:  "success",
// 			Message: "your shopping cart is empty",
// 		})
// 	}

// 	return c.JSON(http.StatusOK, models.Response{
// 		Status:  "success",
// 		Message: "success get checkout data from your shopping cart",
// 		Data:    checkouts,
// 	})
// }

// func GetCheckoutByIdController(c echo.Context) error {
// 	userId := middlewares.ExtractTokenUserId(c)
// 	cartItemId, _ := strconv.Atoi(c.Param("id"))
// 	checkout, err := database.GetCheckoutTotalById(cartItemId, userId)

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
// 		Message: "success get checkout data from your shopping cart",
// 		Data:    checkout,
// 	})
// }

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
	authToken := c.Request().Header.Get("Authorization")

	// url := "http://54.179.213.175:8089/carts"

	// Create a Bearer string by appending string access token
	// var bearer = "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTc1NDkyMjUsInVzZXJJZCI6MX0.btcMQNtoqVcvpM_fBh5SPh4mBwJ85K50kKlsv7bKIs4"
	// var bearer = "Bearer " + authToken

	// Create a new request using http
	// req, _ := http.NewRequest("DELETE", url, nil)

	// add authorization header to the req
	// req.Header.Add("Authorization", bearer)

	// Send req using http Client
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// log.Println("Error on response.\n[ERROR] -", err)
	// }
	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// log.Println("Error while reading the response bytes:", err)
	// }
	// log.Println(string([]byte(body)))

	// if err != nil {
	// return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// }

	// if req == nil {
	// 	return c.JSON(http.StatusBadRequest, models.Response{
	// 		Status:  "fail",
	// 		Message: "your requested data was not found",
	// 	})
	// }

	var (
		stock, price []uint
		total        uint
	)

	productId, qty := getQty(authToken)

	for i := 0; i < len(qty); i++ {
		stock = append(stock, getCurrentStock(strconv.FormatUint(uint64(productId[i]), 10)))
		price = append(price, getPrice(strconv.FormatUint(uint64(productId[i]), 10)))
	}

	for i := 0; i < len(qty); i++ {
		updateStock(strconv.FormatUint(uint64(productId[i]), 10), stock[i]-qty[i])
	}

	for i := 0; i < len(qty); i++ {
		total += price[i] * qty[i]
	}

	deleteCart(authToken)

	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "success checkout product from your shopping cart",
		Total:   total,
	})
}
