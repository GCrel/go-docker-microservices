package domain

import "time"

type Product struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	SellerID    string    `json:"seller_id"`
	CreatedAt   time.Time `json:"created_at"`
}
