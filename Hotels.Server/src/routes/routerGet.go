package routes

import (
	"github.com/Hotels/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouterGet(app fiber.Router) {
	app.Get("/users", func(c *fiber.Ctx) error {
		return controllers.GetAll(c, "Users")
	})
	app.Get("/roles", func(c *fiber.Ctx) error {
		return controllers.GetRols(c)
	})
	app.Get("/hotels", func(c *fiber.Ctx) error {
		return controllers.GetAll(c, "Hotels")
	})
	app.Get("/habitaciones", func(c *fiber.Ctx) error {
		return controllers.GetAll(c, "Habitaciones")
	})
	app.Get("/registros", func(c *fiber.Ctx) error {
		return controllers.GetAll(c, "Registros")
	})
	app.Get("/users/:id", func(c *fiber.Ctx) error {
		return controllers.GetOne(c, "Users")
	})
	app.Get("/hotels/:id", func(c *fiber.Ctx) error {
		return controllers.GetOne(c, "Hotels")
	})
	app.Get("/habitaciones/:id", func(c *fiber.Ctx) error {
		return controllers.GetOne(c, "Habitaciones")
	})
	app.Get("/registros/:id", func(c *fiber.Ctx) error {
		return controllers.GetOne(c, "Registros")
	})
}
