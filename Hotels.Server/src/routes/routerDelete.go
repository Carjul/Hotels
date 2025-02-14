package routes

import (
	"github.com/Hotels/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouterDelete(app fiber.Router) {
	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		return controllers.DeleteOne(c, "Users")
	})
	app.Delete("/hotels/:id", func(c *fiber.Ctx) error {
		return controllers.DeleteOne(c, "Hotels")
	})
	app.Delete("/habitaciones/:id", func(c *fiber.Ctx) error {
		return controllers.DeleteOne(c, "Habitaciones")
	})
	app.Delete("/registros/:id", func(c *fiber.Ctx) error {
		return controllers.DeleteOne(c, "Registros")
	})
}
