package routes

import (
	"github.com/Hotels/src/controllers"
	"github.com/Hotels/src/models"
	"github.com/gofiber/fiber/v2"
)

func RouterPut(app fiber.Router) {
	app.Put("/users/:id", func(c *fiber.Ctx) error {
		return controllers.PutOne[models.Users](c, "Users")
	})
	app.Put("/hotels/:id", func(c *fiber.Ctx) error {
		return controllers.PutOne[models.Hotel](c, "Hotels")
	})
	app.Put("/habitaciones/:id", func(c *fiber.Ctx) error {
		return controllers.PutOne[models.Habitacion](c, "Habitaciones")
	})
	app.Put("/registros/:id", func(c *fiber.Ctx) error {
		return controllers.PutOne[models.Registro](c, "Registros")
	})
}
