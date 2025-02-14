package routes

import (
	"github.com/Hotels/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouterMain(app *fiber.App) {
	app.Use(controllers.RedirectToClient)
	app.Static("/", "./public/dist")
	app.Get("/api", controllers.Init)

}
