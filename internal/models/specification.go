package models

import "time"

type Specification struct {
	ID                int       `json:"id"`
	ProductID         int       `json:"product_id,omitempty"`
	BodyMaterial      string    `json:"body_material,omitempty"`
	NeckMaterial      string    `json:"neck_material,omitempty"`
	FretboardMaterial string    `json:"fretboard_material,omitempty"`
	NumberOfFrets     int       `json:"number_of_frets,omitempty"`
	ScaleLength       float64   `json:"scale_length,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`

	Product *Product `json:"product,omitempty"`
}
