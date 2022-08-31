package entity

import "time"

// Product struct
type Product struct {
	ID          int64      `json:"id"`
	ProductType int        `json:"product_type,omitempty"`
	Name        string     `json:"name"`
	Description string     `json:"description,omitempty"`
	SKU         string     `json:"sku,omitempty"`
	Price       float64    `json:"price,omitempty"`
	DestroyedAt *time.Time `json:"destroyed_at,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
