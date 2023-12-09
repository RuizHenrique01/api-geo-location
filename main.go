package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"

	"api-geo-location/database"
	"api-geo-location/models"
	"api-geo-location/routes"
)

func main() {
	app := iris.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Validator = validator.New()

	app.UseRouter(recover.New())

	db := database.GetInstance()

	db.AutoMigrate(&models.Point{})

	pointRoutes := routes.NewPointRoute()

	pointRoutes.Routes(app)

	app.Listen(":8000")
}
