package utils

import (
	"album/backend/internal/domain"
	"fmt"
	"os"

	"github.com/rwcarlsen/goexif/exif"
)

func ExtractMeta(path string) (*domain.Photo, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	x, err := exif.Decode(f)
	if err != nil {
		return nil, err
	}

	// Récupérer les valeurs safely (some tags may be missing)
	var (
		dateStr        string
		camMakeStr     string
		camModelStr    string
		altitudeStr    string
		orientationStr string
		lat, long      float64
	)

	if d, err := x.DateTime(); err == nil {
		dateStr = d.String()
	}

	if t, err := x.Get(exif.Make); err == nil && t != nil {
		camMakeStr = t.String()
		fmt.Println(camMakeStr)
	}
	if t, err := x.Get(exif.Model); err == nil && t != nil {
		camModelStr = t.String()
		fmt.Println(camModelStr)
	}
	if t, err := x.Get(exif.GPSAltitude); err == nil && t != nil {
		altitudeStr = t.String()
		fmt.Println(altitudeStr)
	}
	if t, err := x.Get(exif.Orientation); err == nil && t != nil {
		orientationStr = t.String()
		fmt.Println(orientationStr)
	}

	if la, lo, err := x.LatLong(); err == nil {
		lat = la
		long = lo
		fmt.Println(lat, long)
	}

	photoMeta := &domain.Photo{
		DateTaken:   dateStr,
		CameraMake:  camMakeStr,
		CameraModel: camModelStr,
		Latitude:    lat,
		Longitude:   long,
		Altitude:    altitudeStr,
		Orientation: orientationStr,
	}

	return photoMeta, nil
}
