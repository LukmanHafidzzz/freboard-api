package models

import "time"

type Specification struct {
	ID                int       `json:"id"`
	ProductID         int       `json:"product_id"`
	BodyMaterial      string    `json:"body_material"`
	NeckMaterial      string    `json:"neck_material"`
	FretboardMaterial string    `json:"fretboard_material"`
	NumberOfFrets     int       `json:"number_of_frets"`
	ScaleLength       float64   `json:"scale_length"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`

	Product *Product `json:"product,omitempty"`
}
