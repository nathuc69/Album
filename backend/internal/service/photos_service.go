package service

import (
	"album/backend/internal/domain"
)

type photosService struct {
	repo domain.PhotosRepository
}

func NewPhotosService(repo domain.PhotosRepository) domain.PhotosService {
	return &photosService{repo: repo}
}

func (s *photosService) AddPhotos(filepath string, photo *domain.Photo) error {
	return s.repo.AddPhotos(filepath, photo)
}

func (s *photosService) GetAllPhotos() ([]*domain.Photo, error) {
	return s.repo.GetAllPhotos()
}
