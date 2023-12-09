package controllers

import (
	"api-geo-location/dtos"
	"api-geo-location/services"
	"strconv"

	"github.com/kataras/iris/v12"
)

type PointController struct {
	pointService *services.PointService
}

func NewPointController() *PointController {
	return &PointController{
		pointService: services.NewPointService(),
	}
}

func (pointController *PointController) Create(ctx iris.Context) {
	pointCreateDto, _ := ctx.Values().Get("model").(*dtos.CreatePointDto)

	data, err := pointController.pointService.Create(pointCreateDto)

	if err != nil {
		ctx.StopWithProblem(err.Status, iris.NewProblem().Title(err.Description))
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(data)

}

func (pointController *PointController) GetAll(ctx iris.Context) {
	var points, err = pointController.pointService.GetAll()

	if err != nil {
		ctx.StopWithProblem(err.Status, iris.NewProblem().Title(err.Description))
	}

	ctx.JSON(points)
}

func (pointController *PointController) GetOne(ctx iris.Context) {

	id, _ := strconv.ParseUint(ctx.Params().Get("id"), 10, 32)

	var points, err = pointController.pointService.GetOne(uint(id))

	if err != nil {
		ctx.StopWithProblem(err.Status, iris.NewProblem().Title(err.Description).Status(err.Status))
	}

	ctx.JSON(points)
}

func (pointController *PointController) Update(ctx iris.Context) {
	id, _ := strconv.ParseUint(ctx.Params().Get("id"), 10, 32)

	pointUpdated, _ := ctx.Values().Get("model").(*dtos.EditPointDto)

	data, err := pointController.pointService.Update(uint(id), pointUpdated)

	if err != nil {
		ctx.StopWithProblem(err.Status, iris.NewProblem().Title(err.Description))
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(data)

}

func (pointController *PointController) Delete(ctx iris.Context) {

	id, _ := strconv.ParseUint(ctx.Params().Get("id"), 10, 32)

	_, err := pointController.pointService.Delete(uint(id))

	if err != nil {
		ctx.StopWithProblem(err.Status, iris.NewProblem().Title(err.Description).Status(err.Status))
	}

	ctx.StatusCode(iris.StatusNoContent)
}
