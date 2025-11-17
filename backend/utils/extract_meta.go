// package utils

// import (
// 	"album/backend/internal/domain"
// 	"os"

// 	"github.com/rwcarlsen/goexif/exif"
// 	"github.com/rwcarlsen/goexif/tiff"
// )

// func safeTagString(t *tiff.Tag) string {
// 	if t == nil {
// 		return ""
// 	}

// 	// StringVal() est SAFE — ne panique jamais
// 	v, err := t.StringVal()
// 	if err == nil {
// 		return v
// 	}

// 	return ""
// }

// func ExtractMeta(path string) (*domain.Photo, error) {

// 	f, err := os.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer f.Close()

// 	x, err := exif.Decode(f)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var (
// 		dateStr        string
// 		camMakeStr     string
// 		camModelStr    string
// 		altitudeStr    string
// 		orientationStr string
// 		lat, long      float64
// 	)

// 	if d, err := x.DateTime(); err == nil {
// 		dateStr = d.String()
// 	}

// 	if t, err := x.Get(exif.Make); err == nil {
// 		camMakeStr = safeTagString(t)
// 	}
// 	if t, err := x.Get(exif.Model); err == nil {
// 		camModelStr = safeTagString(t)
// 	}
// 	if t, err := x.Get(exif.GPSAltitude); err == nil {
// 		altitudeStr = safeTagString(t)
// 	}
// 	if t, err := x.Get(exif.Orientation); err == nil {
// 		orientationStr = safeTagString(t)
// 	}

// 	if la, lo, err := x.LatLong(); err == nil {
// 		lat = la
// 		long = lo
// 	}

// 	return &domain.Photo{
// 		DateTaken:   dateStr,
// 		CameraMake:  camMakeStr,
// 		CameraModel: camModelStr,
// 		Latitude:    lat,
// 		Longitude:   long,
// 		Altitude:    altitudeStr,
// 		Orientation: orientationStr,
// 	}, nil
// }

package utils

import (
	"album/backend/internal/domain"
	"os"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

// safeTagString retourne la valeur d'un tag EXIF en toute sécurité
func safeTagString(t *tiff.Tag) string {
	if t == nil {
		return ""
	}
	val, err := t.StringVal() // NE PANIC JAMAIS
	if err != nil {
		return ""
	}
	return val
}

// ExtractMeta extrait les métadonnées EXIF d'une photo et les retourne sous forme de Photo
func ExtractMeta(path string) (*domain.Photo, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	x, err := exif.Decode(f)
	if err != nil {
		// Pas d'EXIF ? On retourne une photo vide mais pas de panic
		return &domain.Photo{
			Filename: path,
			Path:     path,
		}, nil
	}

	dateStr := ""
	if d, err := x.DateTime(); err == nil {
		dateStr = d.String()
	}

	camMakeStr := safeTagString(getTag(x, exif.Make))
	camModelStr := safeTagString(getTag(x, exif.Model))
	altitudeStr := safeTagString(getTag(x, exif.GPSAltitude))
	orientationStr := safeTagString(getTag(x, exif.Orientation))

	lat, long := 0.0, 0.0
	if la, lo, err := x.LatLong(); err == nil {
		lat, long = la, lo
	}

	return &domain.Photo{
		Filename:    path,
		Path:        path,
		DateTaken:   dateStr,
		CameraMake:  camMakeStr,
		CameraModel: camModelStr,
		Altitude:    altitudeStr,
		Orientation: orientationStr,
		Latitude:    lat,
		Longitude:   long,
	}, nil
}

// getTag récupère un tag EXIF ou nil si absent
func getTag(x *exif.Exif, field exif.FieldName) *tiff.Tag {
	t, err := x.Get(field)
	if err != nil || t == nil {
		return nil
	}
	return t
}
