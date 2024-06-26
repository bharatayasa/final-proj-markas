package router

import (
	"github.com/bharatayasa/final-project/controller"
	"github.com/bharatayasa/final-project/middleware"
	"github.com/gofiber/fiber/v2"
)

func RouterDatabaseBackup(app *fiber.App) {
	app.Use(middleware.AuthMiddleware)

	app.Get("/", controller.GetDatabaseLatest)
	app.Get("/:database_name", controller.GetDatabaseByDbName)
	app.Post("/:database_name", controller.InsertNewData)
	app.Get("/:id/download", controller.DownloadFile)
}
