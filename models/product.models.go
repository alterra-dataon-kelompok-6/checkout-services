package models

import "time"

type Product struct {
	Id          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	CategoryId  uint      `json:"category_id"`
	Name        string    `json:"name"`
	Stock       uint      `json:"stock"`
	Price       uint      `json:"price"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	Category    Category  `json:"category"`
}
