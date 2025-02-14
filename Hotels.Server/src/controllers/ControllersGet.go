package controllers

import (
	"github.com/Hotels/src/db"
	"github.com/Hotels/src/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAll(c *fiber.Ctx, collection string) error {
	c.Accepts("application/json")

	result := db.FindAll(collection)
	if result == nil {
		return c.Status(404).JSON(Message{
			Msg: "No records found",
		})
	}
	return c.Status(200).JSON(result)
}

func GetOne(c *fiber.Ctx, collection string) error {
	c.Accepts("application/json")

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(405).JSON(err)
	}

	result := db.FindById(collection, objID)
	if result == nil {
		return c.Status(404).JSON(Message{
			Msg: "No records found",
		})
	}
	return c.Status(200).JSON(result)
}

func GetRols(c *fiber.Ctx) error {
	c.Accepts("application/json")
	return c.Status(200).JSON(models.Roles)

}
