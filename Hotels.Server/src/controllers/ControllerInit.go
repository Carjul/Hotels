package controllers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Msg string `json:"message"`
}

func Init(c *fiber.Ctx) error {
	c.Accepts("application/json")
	return c.JSON(Message{
		Msg: "Hola, api hotels!",
	})
}

func RedirectToClient(c *fiber.Ctx) error {
	if strings.HasPrefix(c.Path(), "/api") {
		return c.Next()
	}
	// For all other routes, serve index.html
	return c.SendFile("./public/dist/index.html")
}
