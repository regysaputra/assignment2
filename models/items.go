package models

import "time"

type Items struct {
	ItemID      uint      `gorm:"primary_key" json:"item_id"`
	ItemCode    string    `json:"item_code"`
	Description string    `json:"description"`
	Quantity    uint      `json:"quantity"`
	OrderID     uint      `json:"order_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Orders      Orders    `json:"-"`
}
