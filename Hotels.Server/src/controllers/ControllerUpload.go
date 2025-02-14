package controllers

import (
	"sync"

	"github.com/Hotels/src/utils"
	"github.com/gofiber/fiber/v2"
)

func PostImagen(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	var UrlCloudinary string = ""

	files := form.File["image"]

	if len(files) > 0 {
		ImageFile := files[0]
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			UrlCloudinary = utils.UploadImage(ImageFile)
			wg.Done()
		}()
		wg.Wait()

	}

	if UrlCloudinary == "error al subir la imagen a cloudinary" || UrlCloudinary == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error al subir la imagen",
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(Message{Msg: UrlCloudinary})
}
