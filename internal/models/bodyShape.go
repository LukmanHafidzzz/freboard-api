package models

import "time"

type BodyShape struct {
	ID        int       `json:"id"`
	ShapeName string    `json:"shape_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Products []Product `json:"products,omitempty"`
}
