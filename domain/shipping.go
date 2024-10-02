package domain

import (
	"time"
)

type Address struct {
	ID           int64  `json:"id"`
	AddressName  string `json:"address_name"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zip          string `json:"zip"`
	Country      string `json:"country"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
	CustomerID   int64  `json:"customer_id"`
}

type Shipping struct {
	ID            int64          `json:"id"`
	CustomerID    int64          `json:"customer_id"`
	Status        string         `json:"status"`
	ShippingItems []ShippingItem `json:"shipping_items"`
	CreatedAt     int64          `json:"created_at"`
	UpdatedAt     int64          `json:"updated_at"`
	AddressID     int64          `json:"address_id"`
	OrderID       int64          `json:"order_id"`
}

type ShippingItem struct {
	ProductCode string  `json:"product_code"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    int32   `json:"quantity"`
	CreatedAt   int64   `json:"created_at"`
	UpdatedAt   int64   `json:"updated_at"`
	ShippingID  int64   `json:"shipping_id"`
}

func NewShipping(customerId int64, shippingItems []ShippingItem) Shipping {
	return Shipping{
		CreatedAt:     time.Now().Unix(),
		Status:        "Pending",
		CustomerID:    customerId,
		ShippingItems: shippingItems,
	}
}

func (o *Shipping) TotalPrice() float32 {
	var totalPrice float32
	for _, shipItem := range o.ShippingItems {
		totalPrice += shipItem.UnitPrice * float32(shipItem.Quantity)
	}
	return totalPrice
}

func NewAddress(customerId int64, addressName string, addressLine1 string, addressLine2 string, city string, state string, zip string, country string) Address {
	return Address{
		CreatedAt:    time.Now().Unix(),
		CustomerID:   customerId,
		AddressName:  addressName,
		AddressLine1: addressLine1,
		AddressLine2: addressLine2,
		City:         city,
		State:        state,
		Zip:          zip,
		Country:      country,
	}
}
