package repository

import (
	"album/backend/internal/domain"
	"database/sql"
)

type photosRepository struct {
	db *sql.DB
}

func NewPhotosRepository(db *sql.DB) domain.PhotosRepository {
	return &photosRepository{db: db}
}

func (r *photosRepository) AddPhotos(filepath string, photo *domain.Photo) error {
	_, err := r.db.Exec("INSERT INTO photos (filename, filepath, date_taken, camera_make, camera_model, latitude, longitude) VALUES(?,?,?,?,?,?,?)",
		filepath, filepath, photo.DateTaken, photo.CameraMake, photo.CameraModel, photo.Latitude, photo.Longitude)
	if err != nil {
		return err
	}
	return nil
}

func (r *photosRepository) GetAllPhotos() ([]*domain.Photo, error) {
	// Requête pour récupérer toutes les photos
	rows, err := r.db.Query(`
        SELECT 
            id, 
            filename, 
            filepath, 
            date_taken, 
            camera_make, 
            camera_model, 
            latitude, 
            longitude
        FROM photos
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photos []*domain.Photo
	for rows.Next() {
		var photo domain.Photo
		if err := rows.Scan(
			&photo.ID,
			&photo.Filename,
			&photo.Path,
			&photo.DateTaken,
			&photo.CameraMake,
			&photo.CameraModel,
			&photo.Latitude,
			&photo.Longitude,
		); err != nil {
			return nil, err
		}
		photos = append(photos, &photo)
	}

	// Vérifie si une erreur est survenue lors de l'itération
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return photos, nil
}
