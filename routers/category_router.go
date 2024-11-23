package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rifkyrizkita/book_management/controllers"
	"github.com/rifkyrizkita/book_management/middlewares"
)

func CategoryRouters(category fiber.Router) {
	// post routers
	category.Post("/", middlewares.ValidatorAddCategory, controllers.AddCategory)
	// delete routers
	category.Delete("/:id", controllers.DeleteCategory)
	// get routers
	category.Get("/", controllers.GetAllCategories)
	
}
