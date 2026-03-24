package services

import (
	"database/sql"
	"guitar-api/internal/models"
)

type BrandService struct {
	DB *sql.DB
}

func (service *BrandService) GetAllBrands() ([]models.Brand, error) {
	rows, err := service.DB.Query("SELECT id, brand_name FROM brands")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var brands []models.Brand
	for rows.Next() {
		var brand models.Brand
		rows.Scan(&brand.ID, &brand.BrandName)
		brands = append(brands, brand)
	}
	return brands, nil
}