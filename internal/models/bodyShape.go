package models

import "time"

type BodyShape struct {
	ID        int       `json:"id"`
	ShapeName string    `json:"shape_name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	Products []Product `json:"products,omitempty"`
}
