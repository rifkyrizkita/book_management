package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rifkyrizkita/book_management/database"
	"github.com/rifkyrizkita/book_management/models"
)

func AdminAuth(c *fiber.Ctx) error {
	id, ok := c.Locals("user").(jwt.MapClaims)["sub"].(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token not valid"})
	}
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	if !user.IsAdmin {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You are not authorized to access this resource"})
	}

	return c.Next()
}
