package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/rifkyrizkita/book_management/database"
	"github.com/rifkyrizkita/book_management/routers"
)

func init() {
	godotenv.Load()
	database.InitDB()
}

func main() {
	app := fiber.New()

	// middlewares
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	
	api := app.Group("/api")

	user := api.Group("/user")
	routers.UserRouters(user)
	book := api.Group("/book")
	routers.BookRouters(book)
	category := api.Group("/category")
	routers.CategoryRouters(category)
	
	log.Fatal(app.Listen(os.Getenv("PORT")))
}