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

func PostCheckoutController(c echo.Context) error {
	authToken := c.Request().Header.Get("Authorization")

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

	if total == 0 {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "bad request",
			Message: "cart is empty",
			Total:   total,
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "success checkout product from your shopping cart",
		Total:   total,
	})
}

func getQty(token string) ([]uint, []uint) {
	var data models.ResponseCart
	var productId, qty []uint

	url := "http://54.179.213.175:8089/carts"

	// create a Bearer string by appending string access token
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
	var data models.ResponseProduct

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
	var data models.ResponseProduct

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
	body := models.Body{
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

func deleteCart(token string) error {
	url := "http://54.179.213.175:8089/carts"

	// create a Bearer string by appending string access token
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

	return err
}
