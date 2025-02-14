package controllers

import (
	"github.com/Hotels/src/db"
	"github.com/Hotels/src/models"
	"github.com/Hotels/src/utils"
	"github.com/gofiber/fiber/v2"
)

func PostOne[T any](c *fiber.Ctx, collection string) error {

	data := new(T)

	if err := c.BodyParser(data); err != nil {
		return c.Status(406).JSON(err)
	}

	if result := db.Create(collection, data); result == "Error al crear el documento" {
		return c.Status(404).JSON(Message{
			Msg: result,
		})
	} else {
		return c.Status(200).JSON(Message{
			Msg: result,
		})
	}

}

func CreateUser(c *fiber.Ctx, collection string) error {
	c.Accepts("application/json")

	data := new(models.Users)

	if err := c.BodyParser(data); err != nil {
		return c.Status(406).JSON(err)
	}

	hash, er := utils.HashPassword(data.Contraseña)
	if er != nil {
		return c.Status(401).JSON(Message{
			Msg: "Error al crear el documento por contraseña",
		})

	}
	data.Contraseña = hash
	if result := db.Create(collection, data); result == "Error al crear el documento" {
		return c.Status(404).JSON(Message{
			Msg: result,
		})
	} else {
		return c.Status(200).JSON(Message{
			Msg: result,
		})
	}
}
