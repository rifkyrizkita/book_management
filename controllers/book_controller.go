package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rifkyrizkita/book_management/database"
	"github.com/rifkyrizkita/book_management/models"
	"github.com/rifkyrizkita/book_management/web/requests"
)

func AddNewBook(c *fiber.Ctx) error {

	filename := c.Locals("filename").(string)

	var body requests.AddNewBookBody
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	book := models.Book{
		Title:       body.Title,
		Author:      body.Author,
		Synopsis:    body.Synopsis,
		ISBN:        body.ISBN,
		PublishedAt: body.PublishedAt,
		Publisher:   body.Publisher,
		Stock:       uint(body.Stock),
		Image:       &filename,
		CategoryID:  body.CategoryID,
	}

	err = database.DB.Create(&book).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response := fiber.Map{
		"message": "Book added successfully!",
		"result":  book,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func DeleteBook(c *fiber.Ctx) error {

	bookId := c.Params("id")

	if err := database.DB.Where("id = ?", bookId).Delete(&models.Book{}).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Book deleted successfully!"})
}

func GetAllBooks(c *fiber.Ctx) error {
	title := c.Query("title")
	author := c.Query("author")
	categoryID := c.Query("category_id")
	var books []models.Book

	query := database.DB

	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	if author != "" {
		query = query.Where("author LIKE ?", "%"+author+"%")
	}
	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	err := query.Joins("Category").Find(&books).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"result": books})
}

func GetBookById(c *fiber.Ctx) error {
	bookId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
	}

	var book models.Book

	err = database.DB.Preload("Category").Where("id = ?", bookId).Take(&book).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"result": book})
}

func BorrowBook(c *fiber.Ctx) error {
	id, ok := c.Locals("user").(jwt.MapClaims)["sub"].(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token not valid"})
	}

	bookId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
	}

	var book models.Book
	if err := database.DB.First(&book, bookId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	if book.Stock == 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Book out of stock"})
	}

	borrowing := models.Borrowing{
		UserID:     uint(id),
		BookID:     uint(bookId),
		BorrowDate: string(time.Now().Format("2006-01-02")),
		Status:     "borrowed",
	}

	if err := database.DB.Create(&borrowing).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	book.Stock--
	if err := database.DB.Save(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Book borrowed successfully",
	})
}

func ReturnBook(c *fiber.Ctx) error {

	

	borrowingId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid borrowing ID"})
	}

	var borrowing models.Borrowing
	if err := database.DB.First(&borrowing, borrowingId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Borrowing record not found"})
	}

	
	if borrowing.Status == "returned" {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Book already returned"})
	}

	now := string(time.Now().Format("2006-01-02"))
	borrowing.Status = "returned"
	borrowing.ReturnDate = &now

	if err := database.DB.Save(&borrowing).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var book models.Book
	if err := database.DB.First(&book, borrowing.BookID).Error; err == nil {
		book.Stock++
		database.DB.Save(&book)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book returned successfully",
	})
}

func GetBorrowedBooksByUserId(c *fiber.Ctx) error {

	id, ok := c.Locals("user").(jwt.MapClaims)["sub"].(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token not valid"})
	}

	var borrowings []models.Borrowing

	err := database.DB.Preload("Book").Where("user_id = ? AND status = ?", uint(id), "borrowed").Find(&borrowings).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching borrowed books"})
	}

	if len(borrowings) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "No books currently borrowed"})
	}

	var borrowedBooks []fiber.Map
	for _, borrowing := range borrowings {
		borrowedBooks = append(borrowedBooks, fiber.Map{
			"id":          borrowing.Book.ID,
			"title":       borrowing.Book.Title,
			"author":      borrowing.Book.Author,
			"borrow_date": borrowing.BorrowDate,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Books currently borrowed",
		"books":   borrowedBooks,
	})
}

func GetAllBorrowedBooks(c *fiber.Ctx) error {
	var borrowings []models.Borrowing

	err := database.DB.Preload("Book").Preload("User").Where("status = ?", "borrowed").Find(&borrowings).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching borrowed books"})
	}

	if len(borrowings) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "No books currently borrowed"})
	}

	var borrowedBooks []fiber.Map
	for _, borrowing := range borrowings {
		borrowedBooks = append(borrowedBooks, fiber.Map{
			"borrow_id":     borrowing.ID,
			"user_id":       borrowing.UserID,
			"user_name":     borrowing.User.Username,
			"user_email":    borrowing.User.Email,
			"book_id":       borrowing.Book.ID,
			"book_title":    borrowing.Book.Title,
			"book_author":   borrowing.Book.Author,
			"borrow_date":   borrowing.BorrowDate,
			"return_date":   borrowing.ReturnDate,
			"borrow_status": borrowing.Status,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "All borrowed books",
		"borrowings": borrowedBooks,
	})
}
