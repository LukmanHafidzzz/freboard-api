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

func (service *BrandService) GetAllProductsByBrandId(id int) ([]models.Product, error) {
	rows, err := service.DB.Query(`SELECT b.id, b.brand_name, p.id, p.model_name, bs.id, bs.shape_name FROM products p JOIN brands b ON p.brand_id = b.id JOIN body_shapes bs ON p.shape_id = bs.id WHERE b.id = ?`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []models.Product
	for rows.Next() {
		var product models.Product
		product.Brand = &models.Brand{}
		product.BodyShape = &models.BodyShape{}
		rows.Scan(&product.Brand.ID, &product.Brand.BrandName, &product.ID, &product.ModelName, &product.BodyShape.ID, &product.BodyShape.ShapeName)
		products = append(products, product)
	}
	return products, nil
}
