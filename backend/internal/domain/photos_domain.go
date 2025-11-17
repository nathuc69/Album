package domain

type Photo struct {
	ID          string  `json:"id"`
	Filename    string  `json:"filename"`
	Path        string  `json:"path"`
	DateTaken   string  `json:"datetaken"`
	CameraMake  string  `json:"cameramake"`
	CameraModel string  `json:"cameramodel"`
	Orientation string  `json:"orientation"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Altitude    string  `json:"altitude"`
}

type PhotosRepository interface {
	AddPhotos(filepath string, photo *Photo) error
	GetAllPhotos(id string) ([]*Photo, error)
}

type PhotosService interface {
	AddPhotos(filepath string, photo *Photo) error
	GetAllPhotos(id string) ([]*Photo, error)
}
