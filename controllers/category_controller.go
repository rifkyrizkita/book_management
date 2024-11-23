package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rifkyrizkita/book_management/database"
	"github.com/rifkyrizkita/book_management/models"
	"github.com/rifkyrizkita/book_management/web/requests"
)

func AddCategory(c *fiber.Ctx) error {
	var body requests.AddCategoryRequest

	// Parse request body
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var existingCategory models.Category
	if err := database.DB.Where("category = ?", body.Category).First(&existingCategory).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Category with this category already exists",
		})
	}

	// Buat kategori baru
	category := models.Category{
		Category: body.Category,
	}

	if err := database.DB.Create(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create category",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "Category created successfully",
	})
}

func GetAllCategories(c *fiber.Ctx) error {
    var categories []models.Category

    if err := database.DB.Find(&categories).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to fetch categories",
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "categories": categories,
    })
}

func DeleteCategory(c *fiber.Ctx) error {
    categoryId, err := c.ParamsInt("id")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid category ID",
        })
    }

    // Cek apakah kategori ada di database
    var category models.Category
    if err := database.DB.First(&category, categoryId).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Category not found",
        })
    }

    // Hapus kategori
    if err := database.DB.Delete(&category).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete category",
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Category deleted successfully",
    })
}

