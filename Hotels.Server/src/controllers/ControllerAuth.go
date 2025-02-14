package controllers

import (
	"fmt"
	"os"
	"time"

	"github.com/Hotels/src/db"
	"github.com/Hotels/src/models"
	"github.com/Hotels/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Login(c *fiber.Ctx, store *session.Store) error {

	var userBody models.Users
	var userDB models.Users

	resultUser := db.InstanceDB.Database.Collection("Users")

	if err := c.BodyParser(&userBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not parse JSON User",
		})
	}

	if err := resultUser.FindOne(c.Context(), bson.M{"correo": userBody.Correo}).Decode(&userDB); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Aquí deberías validar el usuario contra una base de datos
	if utils.CheckPassword(userBody.Contraseña, userDB.Contraseña) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	// Crear el token JWT
	token, err := utils.GenerateJWT(userDB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not generate token",
		})
	}

	// Guardar el token en la sesión
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get session",
		})
	}
	sess.Set("token", token)
	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not save session",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

func Logout(c *fiber.Ctx, store *session.Store) error {

	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get session",
		})
	}

	// Eliminar la sesión
	if err := sess.Destroy(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not destroy session",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Logged out successfully",
	})

}

func SendMailRecuperacion(c *fiber.Ctx) error {
	var userBody models.Users
	var userDB models.Users

	var dominio = os.Getenv("DOMINIO")

	if dominio == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "No se ha configurado el dominio",
		})

	}

	resultUser := db.InstanceDB.Database.Collection("Users")

	if err := c.BodyParser(&userBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not parse JSON User",
		})

	}
	if err := resultUser.FindOne(c.Context(), bson.M{"correo": userBody.Correo}).Decode(&userDB); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "User not found",
		})

	}
	var resetLink string = dominio + "/api/reset-password?id=" + userDB.ID.Hex()

	html := fmt.Sprintf(`
	<div style="font-family: Arial, sans-serif; background-color: #f4f4f4; padding: 20px;">
	  <div style="max-width: 600px; margin: 0 auto; background-color: #ffffff; border-radius: 8px; overflow: hidden; box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);">
		  <div style="background-color: #4CAF50; padding: 20px; text-align: center;">
			  <h1 style="color: #ffffff; font-size: 24px;">Recuperación de Contraseña</h1>
		  </div>
		  <div style="padding: 20px; text-align: center;">
			  <h2 style="color: #333333; font-size: 20px;">Hola, %s</h2>
			  <p style="color: #555555; font-size: 16px;">
				  Parece que solicitaste un cambio de contraseña. Haz clic en el siguiente botón para cambiar tu contraseña:
			  </p>
			  <a href="%s" style="display: inline-block; background-color: #4CAF50; color: #ffffff; text-decoration: none; padding: 12px 24px; border-radius: 4px; font-size: 16px; margin-top: 20px;">
				  Cambiar contraseña
			  </a>
			  <p style="color: #999999; font-size: 14px; margin-top: 30px;">
				  Si no solicitaste este cambio, por favor ignora este correo.
			  </p>
		  </div>
		  <div style="background-color: #f4f4f4; padding: 10px; text-align: center; color: #888888; font-size: 12px;">
			  
		  </div>
	  </div>
  </div>
  `, userDB.Nombre, resetLink)
	result := utils.SendMail(userDB.Correo, "Recuperación de contraseña", html)

	return c.JSON(fiber.Map{
		"message": result,
	})

}

func GetPassword(c *fiber.Ctx) error {
	return c.SendFile("./public/reset-password.html")
}

func PutPassword(c *fiber.Ctx, collection string) error {
	c.Accepts("application/json")

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(405).JSON(err)
	}

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

	result := db.InstanceDB.Database.Collection("Users")

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"contraseña": data.Contraseña, "updated_at": time.Now()}}

	_, er = result.UpdateOne(c.Context(), filter, update)
	if er != nil {
		return c.Status(404).JSON(Message{
			Msg: "Error al actualizar el documento",
		})
	}

	return c.Status(200).JSON(Message{
		Msg: "Contraseña actualizada exitosamente",
	})
}
