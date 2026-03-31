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

func (service *ProductService) GetProductById(id int) (*models.Product, error) {
	var product models.Product
	product.Brand = &models.Brand{}
	product.BodyShape = &models.BodyShape{}
	product.Specification = &models.Specification{}
	err := service.DB.QueryRow(`SELECT p.id, p.model_name, b.id, b.brand_name, b.country, bs.id, bs.shape_name, s.body_material, s.neck_material, s.fretboard_material, s.number_of_frets, s.scale_length FROM products p JOIN brands b ON p.brand_id = b.id JOIN body_shapes bs ON p.shape_id = bs.id LEFT JOIN specifications s ON p.id = s.product_id WHERE p.id = ?`, id).Scan(&product.ID, &product.ModelName, &product.Brand.ID, &product.Brand.BrandName, &product.Brand.Country, &product.BodyShape.ID, &product.BodyShape.ShapeName, &product.Specification.BodyMaterial, &product.Specification.NeckMaterial, &product.Specification.FretboardMaterial, &product.Specification.NumberOfFrets, &product.Specification.ScaleLength)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &product, nil
}