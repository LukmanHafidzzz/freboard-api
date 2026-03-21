package models

import "time"

type Brand struct {
	ID        int       `json:"id"`
	BrandName string    `json:"brand_name"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Products []Product `json:"products,omitempty"`
}
