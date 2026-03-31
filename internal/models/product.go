package models

import "time"

type Product struct {
	ID        int       `json:"id"`
	BrandID   int       `json:"brand_id,omitempty"`
	ShapeID   int       `json:"shape_id,omitempty"`
	ModelName string    `json:"model_name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	Brand         *Brand         `json:"brand,omitempty"`
	BodyShape     *BodyShape     `json:"body_shape,omitempty"`
	Specification *Specification `json:"specification,omitempty"`
}
