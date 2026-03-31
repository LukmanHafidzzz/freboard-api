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
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return brands, nil
}

func (service *BrandService) GetBrandById(id int) (*models.Brand, error) {
	var brand models.Brand
	err := service.DB.QueryRow("SELECT id, brand_name FROM brands WHERE id = ?", id).Scan(&brand.ID, &brand.BrandName)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &brand, nil
}
