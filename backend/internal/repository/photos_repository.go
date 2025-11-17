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

func (r *photosRepository) GetAllPhotos(id string) ([]*domain.Photo, error) {
	row, err := r.db.Query("SELECT id, filename, filepath, date_taken, camera_make, camera_model, latitude, longitude FROM photos WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var photos []*domain.Photo
	for row.Next() {
		var photo domain.Photo
		if err := row.Scan(&photo.ID, &photo.Filename, &photo.Path, &photo.DateTaken, &photo.CameraMake, &photo.CameraModel, &photo.Latitude, &photo.Longitude); err != nil {
			return nil, err
		}
		photos = append(photos, &photo)
	}
	return photos, nil
}
