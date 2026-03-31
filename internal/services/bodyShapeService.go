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
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return bodyShapes, nil
}

func (service *BodyShapeService) GetBodyShapeById(id int) (*models.BodyShape, error) {
	var bodyShape models.BodyShape
	err := service.DB.QueryRow("SELECT id, shape_name FROM body_shapes WHERE id = ?", id).Scan(&bodyShape.ID, &bodyShape.ShapeName)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &bodyShape, nil
}