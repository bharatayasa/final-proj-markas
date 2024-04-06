package router

import (
	"github.com/bharatayasa/final-project/controller"
	"github.com/gofiber/fiber/v2"
)

func RouterDatabaseBackup(app *fiber.App) {
	app.Get("/", controller.GetDatabaseLatest)
}