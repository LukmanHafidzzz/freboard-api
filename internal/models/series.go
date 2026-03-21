package models

import "time"

type Series struct {
	ID        int       `json:"id"`
	Name      string    `json:"series_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Products []Product `json:"products,omitempty"`
}
