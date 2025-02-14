package src

import (
	"fmt"

	"github.com/Hotels/src/middleware"
	"github.com/Hotels/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func AppFiber(port string) {

	app := fiber.New()

	store := session.New()

	routes.RouterMain(app)

	routes.RouterAuth(app, store)

	api := app.Group("/api", middleware.AuthMiddleware(store))

	routes.RouterGet(api)
	routes.RouterPost(api)
	routes.RouterPut(api)
	routes.RouterDelete(api)

	fmt.Print(app.Listen(":" + port))

}
