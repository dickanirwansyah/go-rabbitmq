package routes

import (
	"producer/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/api/v1/policy/insert", controller.InsertPolicy)

}
