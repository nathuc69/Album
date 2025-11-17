package utils

import (
	"album/backend/internal/domain"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

// safeTagString retourne la valeur d'un tag EXIF sans panique
func safeTagString(t *tiff.Tag) string {
	if t == nil {
		return ""
	}
	val, err := t.StringVal()
	if err != nil {
		return ""
	}
	return val
}

// getTag récupère un tag EXIF ou nil si absent
func getTag(x *exif.Exif, field exif.FieldName) *tiff.Tag {
	t, err := x.Get(field)
	if err != nil || t == nil {
		return nil
	}
	return t
}

// ExtractMeta récupère les métadonnées EXIF pour un fichier donné
func ExtractMeta(path string) (*domain.Photo, error) {
	ext := strings.ToLower(filepath.Ext(path))
	photo := &domain.Photo{
		Filename: filepath.Base(path),
		Path:     path,
	}

	switch ext {
	case ".jpg", ".jpeg", ".tiff", ".JPG", ".JPEG", ".TIFF":
		f, err := os.Open(path)
		if err != nil {
			return photo, err
		}
		defer f.Close()

		x, err := exif.Decode(f)
		if err != nil {
			// Pas d'EXIF : retourne quand même photo minimale
			return photo, nil
		}

		if d, err := x.DateTime(); err == nil {
			photo.DateTaken = d.String()
		}
		photo.CameraMake = safeTagString(getTag(x, exif.Make))
		photo.CameraModel = safeTagString(getTag(x, exif.Model))
		photo.Altitude = safeTagString(getTag(x, exif.GPSAltitude))
		photo.Orientation = safeTagString(getTag(x, exif.Orientation))
		if la, lo, err := x.LatLong(); err == nil {
			photo.Latitude = la
			photo.Longitude = lo
		}

	case ".heic":
		// TODO: Ajouter une lib go-heif si tu veux récupérer metadata HEIC
		// Exemple : https://pkg.go.dev/github.com/pixiv/go-heif
	case ".png":
		// Les PNG n'ont quasiment jamais d'EXIF, on peut ignorer
	default:
		// Fichiers non-supportés
		return nil, nil
	}

	return photo, nil
}

// WalkFolder parcourt un dossier et extrait toutes les métadonnées
func WalkFolder(folderPath string) ([]*domain.Photo, error) {
	var photos []*domain.Photo

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Ne pas arrêter le Walk si un fichier est inaccessible
			fmt.Println("Warning:", err)
			return nil
		}
		if info.IsDir() {
			return nil
		}

		photo, err := ExtractMeta(path)
		if err != nil {
			fmt.Println("Erreur lors de l'extraction :", path, err)
			return nil
		}
		if photo != nil {
			photos = append(photos, photo)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return photos, nil
}
