package migrations

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rifkyrizkita/book_management/database"
	"github.com/rifkyrizkita/book_management/models"
)

func Migration(c *fiber.Ctx) error {

	if err := database.DB.AutoMigrate(&models.User{}, &models.Book{}, &models.Borrowing{}, &models.Category{}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Migration failed",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Migration successful",
	})
}
