package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rifkyrizkita/book_management/migrations"
)

func MigrationRouters(migration fiber.Router) {
	// post routers
	migration.Post("/", migrations.Migration)
}
