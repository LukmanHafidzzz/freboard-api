package services

import (
	"database/sql"
	"guitar-api/internal/models"
)

type ProductService struct {
	DB *sql.DB
}

func (service *ProductService) GetAllProducts() ([]models.Product, error) {
	rows, err := service.DB.Query(`SELECT p.id, p.model_name, b.id, b.brand_name, bs.id, bs.shape_name FROM products p JOIN brands b ON p.brand_id = b.id JOIN body_shapes bs ON p.shape_id = bs.id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []models.Product
	for rows.Next() {
		var product models.Product
		product.Brand = &models.Brand{}
		product.BodyShape = &models.BodyShape{}
		rows.Scan(&product.ID, &product.ModelName, &product.Brand.ID, &product.Brand.BrandName, &product.BodyShape.ID, &product.BodyShape.ShapeName)
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}
