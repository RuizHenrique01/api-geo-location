package dtos

type CreatePointDto struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude" validate:"required,latitude,number"`
	Longitude   float64 `json:"longitude" validate:"required,longitude,number"`
}

type EditPointDto struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude" validate:"required,latitude,number"`
	Longitude   float64 `json:"longitude" validate:"required,longitude,number"`
}
