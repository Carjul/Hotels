package routes

import (
	"github.com/Hotels/src/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func RouterAuth(app *fiber.App, store *session.Store) {

	app.Post("/api/login", func(c *fiber.Ctx) error {
		return controllers.Login(c, store)
	})
	app.Post("/api/logout", func(c *fiber.Ctx) error {
		return controllers.Logout(c, store)
	})
	app.Post("/api/reset-password", func(c *fiber.Ctx) error {
		return controllers.SendMailRecuperacion(c)
	})
	app.Get("/api/reset-password", func(c *fiber.Ctx) error {
		return controllers.GetPassword(c)
	})
	app.Put("/api/reset-password/:id", func(c *fiber.Ctx) error {
		return controllers.PutPassword(c, "Users")
	})
}
