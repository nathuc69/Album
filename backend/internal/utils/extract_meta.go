package utils

import (
	"album/backend/internal/domain"
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

	// Récupérer les valeurs
	date, _ := x.DateTime()
	camMake, _ := x.Get(exif.Make)
	camModel, _ := x.Get(exif.Model)
	// focal, _ := x.Get(exif.FocalLength)
	// iso, _ := x.Get(exif.ISOSpeedRatings)
	altitude, _ := x.Get(exif.GPSAltitude)
	orientation, _ := x.Get(exif.Orientation)

	lat, long, _ := x.LatLong()

	photoMeta := &domain.Photo{
		DateTaken:   date.String(),
		CameraMake:  camMake.String(),
		CameraModel: camModel.String(),
		Latitude:    lat,
		Longitude:   long,
		Altitude:    altitude.String(),
		Orientation: orientation.String(),
	}

	return photoMeta, nil
}
