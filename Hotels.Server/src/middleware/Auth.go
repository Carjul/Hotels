package middleware

import (
	"github.com/Hotels/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Could not get session",
				"error":   err.Error(),
			})
		}

		tokenString := sess.Get("token")
		if tokenString == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
				"error":   "No token in session",
			})
		}

		token, err := jwt.Parse(tokenString.(string), func(token *jwt.Token) (interface{}, error) {
			return []byte(utils.API_SECRET), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
				"error":   err.Error(),
			})
		}

		return c.Next()
	}
}
