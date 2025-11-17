package handlers

import (
	"album/backend/internal/domain"
	"net/http"
)

func Router(photoService domain.PhotosService) http.Handler {
	h := NewPhotosHandler(photoService)

	mux := http.NewServeMux()
	mux.HandleFunc("/photos", h.GetAllPhotos)

	return mux
}
