package controllers

import (
	"github.com/Hotels/src/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteOne(c *fiber.Ctx, collection string) error {
	c.Accepts("application/json")

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(405).JSON(err)
	}

	result := db.DeleteById(collection, objID)

	if result == "Error al eliminar el documento" {
		return c.Status(404).JSON(Message{
			Msg: result,
		})
	}
	return c.Status(200).JSON(Message{
		Msg: result,
	})
}
