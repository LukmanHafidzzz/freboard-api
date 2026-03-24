package models

import "time"

type Brand struct {
	ID        int       `json:"id"`
	BrandName string    `json:"brand_name,omitempty"`
	Country   string    `json:"country,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	Products []Product `json:"products,omitempty"`
}
