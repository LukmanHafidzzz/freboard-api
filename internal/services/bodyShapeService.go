package services

import (
	"database/sql"
	"guitar-api/internal/models"
	"log"
	"time"
)

type BodyShapeService struct {
	DB *sql.DB
}

func (service *BodyShapeService) GetAllBodyShape() ([]models.BodyShape, error) {
	start := time.Now()
	rows, err := service.DB.Query("SELECT id, shape_name FROM body_shapes")
	log.Printf("DB query took: %v", time.Since(start))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bodyShapes []models.BodyShape
	for rows.Next() {
		var bodyShape models.BodyShape
		rows.Scan(&bodyShape.ID, &bodyShape.ShapeName)
		bodyShapes = append(bodyShapes, bodyShape)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return bodyShapes, nil
}

func (service *BodyShapeService) GetBodyShapeById(id int) (*models.BodyShape, error) {
	var bodyShape models.BodyShape
	start := time.Now()
	err := service.DB.QueryRow("SELECT id, shape_name FROM body_shapes WHERE id = ?", id).Scan(&bodyShape.ID, &bodyShape.ShapeName)
	log.Printf("DB query took: %v", time.Since(start))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &bodyShape, nil
}

func (service *BodyShapeService) GetAllProductsByBodyShapeId(id int) ([]models.Product, error) {
	start := time.Now()
	rows, err := service.DB.Query(`SELECT bs.id, bs.shape_name, p.id, p.model_name, b.id, b.brand_name FROM products p JOIN body_shapes bs ON p.shape_id = bs.id JOIN brands b ON p.brand_id = b.id WHERE bs.id = ?`, id)
	log.Printf("DB query took: %v", time.Since(start))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []models.Product
	for rows.Next() {
		var product models.Product
		product.Brand = &models.Brand{}
		product.BodyShape = &models.BodyShape{}
		rows.Scan(&product.BodyShape.ID, &product.BodyShape.ShapeName, &product.ID, &product.ModelName, &product.Brand.ID, &product.Brand.BrandName)
		products = append(products, product)
	}
	return products, nil
}