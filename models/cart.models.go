package models

import "time"

type Cart struct {
	Id         uint       `json:"id"`
	CustomerId uint       `json:"customer_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  time.Time  `json:"deleted_at"`
	CartItems  []CartItem `json:"cart_items"`
}

type CartItem struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	CartId    uint      `json:"cart_id"`
	ProductId uint      `json:"product_id"`
	Qty       uint      `json:"qty"`
}
