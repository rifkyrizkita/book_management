package middlewares

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rifkyrizkita/book_management/helpers"
	"github.com/rifkyrizkita/book_management/web/requests"
)

var v = validator.New()

func ValidatorLogin(c *fiber.Ctx) error {
	var body requests.LoginBody
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	v.RegisterValidation("password", helpers.PasswordValidator)
	if err := v.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Next()
}
func ValidatorRegister(c *fiber.Ctx) error {
	var body requests.RegisterBody
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	v.RegisterValidation("password", helpers.PasswordValidator)
	if err := v.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Next()
}

func ValidatorUpdateProfile(c *fiber.Ctx) error {
	var body requests.UpdateProfileBody
	if err := v.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Next()
}

func ValidatorUpdatePassword(c *fiber.Ctx) error {
	var body requests.UpdatePasswordBody
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	v.RegisterValidation("password", helpers.PasswordValidator)
	if err := v.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Next()
}
func ValidatorForgetPassword(c *fiber.Ctx) error {
	var body requests.ForgetPasswordBody
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if err := v.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Next()
}
func ValidatorResetPassword(c *fiber.Ctx) error {
	var body requests.ResetPasswordBody
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	v.RegisterValidation("password", helpers.PasswordValidator)
	if err := v.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Next()
}

func ValidatorAddNewBook(c *fiber.Ctx) error {
	var body requests.AddNewBookBody
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if err := v.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Next()
}

func ValidatorAddCategory(c *fiber.Ctx) error {
	var body requests.AddCategoryRequest
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if err := v.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Next()
}
