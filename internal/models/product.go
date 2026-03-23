package models

import "time"

type Product struct {
	ID        int       `json:"id"`
	BrandID   int       `json:"brand_id"`
	ShapeID   int       `json:"shape_id"`
	ModelName string    `json:"model_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Brand         *Brand         `json:"brand,omitempty"`
	BodyShape     *BodyShape     `json:"body_shape,omitempty"`
	Specification *Specification `json:"specification,omitempty"`
}
