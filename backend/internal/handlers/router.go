package handlers

import (
	"album/backend/internal/domain"
	"net/http"
)

func Router(photoService domain.PhotosService) http.Handler {
	h := NewPhotosHandler(photoService)

	mux := http.NewServeMux()
	mux.HandleFunc("/photos", h.GetAllPhotos)

	// Wrap with CORS middleware
	return corsMiddleware(mux)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
