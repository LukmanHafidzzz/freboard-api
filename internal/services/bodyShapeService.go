package services

import (
	"database/sql"
	"guitar-api/internal/models"
)

type BodyShapeService struct {
	DB *sql.DB
}

func (service *BodyShapeService) GetAllBodyShape() ([]models.BodyShape, error) {
	rows, err := service.DB.Query("SELECT id, shape_name FROM body_shapes")
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
	return bodyShapes, nil
}