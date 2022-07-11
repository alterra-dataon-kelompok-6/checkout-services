package models

type Checkout struct {
	Checkout_id   int        `json:"checout_id"`
	Product       string     `json:"product"`
	Item_total    int        `json:"item_total"`
	Amount        int        `json:"amount"`
	Checkout_item []CartItem `json:"items" gorm:"foreignKey:Cart_id;references:Checkout_id"`
}
