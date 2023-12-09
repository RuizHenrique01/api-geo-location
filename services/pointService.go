package services

import (
	"api-geo-location/database"
	"api-geo-location/dtos"
	"api-geo-location/errors"
	"api-geo-location/models"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type PointService struct {
	db *gorm.DB
}

func NewPointService() *PointService {
	return &PointService{
		db: database.GetInstance(),
	}
}

func (pointService *PointService) Create(createPointDto *dtos.CreatePointDto) (*models.Point, *errors.ErrorAPI) {

	pointCreated := models.Point{
		Name:        createPointDto.Name,
		Description: createPointDto.Description,
		Latitude:    createPointDto.Latitude,
		Longitude:   createPointDto.Longitude,
	}

	result := pointService.db.Create(&pointCreated)

	if result.Error != nil {
		return nil, &errors.ErrorAPI{Description: "Falha ao criar novo ponto!", Status: iris.StatusBadRequest}
	}

	return &pointCreated, nil
}

func (pointService *PointService) GetAll() ([]models.Point, *errors.ErrorAPI) {
	var points []models.Point

	result := pointService.db.Find(&points)

	if result.Error != nil {
		return nil, &errors.ErrorAPI{Description: "Falha ao listar pontos!", Status: iris.StatusBadRequest}
	}

	return points, nil
}

func (pointService *PointService) GetOne(id uint) (*models.Point, *errors.ErrorAPI) {
	var point models.Point

	result := pointService.db.First(&point, id)

	if result.Error != nil {
		return nil, &errors.ErrorAPI{Description: "Ponto não encontrado!", Status: iris.StatusNotFound}
	}

	return &point, nil
}

func (pointService *PointService) Update(id uint, editedPoint *dtos.EditPointDto) (*models.Point, *errors.ErrorAPI) {
	var point *models.Point

	point, err := pointService.GetOne(id)

	if err != nil {
		return point, err
	}

	result := pointService.db.Model(&point).Updates(models.Point{
		Name:        editedPoint.Name,
		Description: editedPoint.Description,
		Latitude:    editedPoint.Latitude,
		Longitude:   editedPoint.Longitude,
	})

	if result.Error != nil {
		return nil, &errors.ErrorAPI{Description: "Falha ao atualizar novo ponto!", Status: iris.StatusBadRequest}
	}

	return point, nil
}

func (pointService *PointService) Delete(id uint) (*models.Point, *errors.ErrorAPI) {
	var point *models.Point

	point, err := pointService.GetOne(id)

	if err != nil {
		return point, err
	}

	result := pointService.db.Delete(&point, id)

	if result.Error != nil {
		return nil, &errors.ErrorAPI{Description: "Falha ao deletar ponto!", Status: iris.StatusBadRequest}
	}

	return point, nil
}
