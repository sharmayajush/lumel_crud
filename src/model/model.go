package model

import (
	"time"
)

type Customer struct {
	ID      string `gorm:"primaryKey;" json:"id"`
	Name    string `gorm:"type:varchar(100);not null" json:"name"`
	Email   string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Address string `gorm:"type:text" json:"address"`
}

// Product model
type Product struct {
	ID       string  `gorm:"primaryKey;" json:"id"`
	Name     string  `gorm:"type:varchar(100);not null" json:"name"`
	Category string  `gorm:"type:varchar(100);not null" json:"category"`
	Price    float64 `gorm:"not null" json:"price"`
}

// Order model
type Order struct {
	ID            uint64    `gorm:"primaryKey;" json:"id"`
	CustomerID    string    `gorm:"not null" json:"customer_id"`
	ProductID     string    `gorm:"not null" json:"product_id"`
	Quantity      int       `gorm:"not null" json:"quantity"`
	Discount      float64   `gorm:"default:0" json:"discount"`
	ShippingCost  float64   `gorm:"not null" json:"shilpping_cost"`
	PaymentMethod string    `gorm:"type:varchar(50);not null" json:"payment_method"`
	DateOfSale    time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"date_of_sale"`
	Region        string    `gorm:"type:varchar(50);not null" json:"region"`
}
