package handlers

import (
	"album/backend/internal/domain"
	"encoding/json"
	"net/http"
)

type PhotosHandler struct {
	Service domain.PhotosService
}

func NewPhotosHandler(service domain.PhotosService) *PhotosHandler {
	return &PhotosHandler{Service: service}
}

func (h *PhotosHandler) GetAllPhotos(w http.ResponseWriter, r *http.Request) {
	photos, err := h.Service.GetAllPhotos("1")
	if err != nil {
		http.Error(w, "Error retrieving photos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(photos)
}
