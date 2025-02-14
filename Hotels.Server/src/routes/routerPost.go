package routes

import (
	"github.com/Hotels/src/controllers"
	"github.com/Hotels/src/models"
	"github.com/gofiber/fiber/v2"
)

func RouterPost(app fiber.Router) {
	app.Post("/users", func(c *fiber.Ctx) error {
		return controllers.CreateUser(c, "Users")
	})
	app.Post("/hotels", func(c *fiber.Ctx) error {
		return controllers.PostOne[models.Hotel](c, "Hotels")
	})
	app.Post("/habitaciones", func(c *fiber.Ctx) error {
		return controllers.PostOne[models.Habitacion](c, "Habitaciones")
	})
	app.Post("/registros", func(c *fiber.Ctx) error {
		return controllers.PostOne[models.Registro](c, "Registros")
	})
	app.Post("/UploadImage", func(c *fiber.Ctx) error {
		return controllers.PostImagen(c)
	})
}
