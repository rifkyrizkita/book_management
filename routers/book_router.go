package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rifkyrizkita/book_management/controllers"
	"github.com/rifkyrizkita/book_management/middlewares"
)

func BookRouters(book fiber.Router) {
	// post routers
	book.Post("/", middlewares.VerifyToken, middlewares.AdminAuth, middlewares.ValidatorAddNewBook, middlewares.UploadFile("BIMG", ""), controllers.AddNewBook)
	book.Post("/borrow/:id", middlewares.VerifyToken, controllers.BorrowBook)
	// patch routers
	book.Patch("/return/:id", middlewares.VerifyToken, controllers.ReturnBook)
	// delete routers
	book.Delete("/:id", middlewares.VerifyToken, middlewares.AdminAuth, controllers.DeleteBook)
	// get routers
	book.Get("/", controllers.GetAllBooks)
	book.Get("/:id", controllers.GetBookById)
}