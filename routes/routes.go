package routes

import (
	"producer/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/api/v1/policy/insert", controller.InsertPolicy)
	app.Get("/api/v1/policy/find/:id", controller.GetPolicy)
	app.Put("/api/v1/policy/update/:id", controller.UpdatePolicy)
	app.Delete("/api/v1/policy/delete/:id", controller.DeletePolicy)

	app.Post("/api/v1/claim/publish", controller.SendPublishClaim)
}
