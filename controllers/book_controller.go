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
	// Ambil ID pengguna dari token JWT
	// id, ok := c.Locals("user").(jwt.MapClaims)["sub"].(float64)
	// if !ok {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token not valid"})
	// }

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
	// Ambil ID pengguna dari token JWT
	// id, ok := c.Locals("user").(jwt.MapClaims)["sub"].(float64)
	// if !ok {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token not valid"})
	// }

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
	bookId := c.Params("id")
	var book models.Book
	err := database.DB.Where("id = ?", bookId).Take(&book).Error
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

	// Cek apakah buku ada
	var book models.Book
	if err := database.DB.First(&book, bookId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	// Cek stok buku
	if book.Stock == 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Book out of stock"})
	}

	// Buat data peminjaman
	borrowing := models.Borrowing{
		UserID:     uint(id),
		BookID:     uint(bookId),
		BorrowDate: string(time.Now().Format("2006-01-02")),
		Status:     "borrowed",
	}

	// Simpan data peminjaman ke database
	if err := database.DB.Create(&borrowing).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Kurangi stok buku
	book.Stock--
	if err := database.DB.Save(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Respon sukses
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Book borrowed successfully",
	})
}

func ReturnBook(c *fiber.Ctx) error {
	// Mendapatkan ID pengguna dari token
	id, ok := c.Locals("user").(jwt.MapClaims)["sub"].(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token not valid"})
	}

	// Mendapatkan ID peminjaman dari parameter
	borrowingId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid borrowing ID"})
	}

	// Cari data peminjaman berdasarkan ID
	var borrowing models.Borrowing
	if err := database.DB.First(&borrowing, borrowingId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Borrowing record not found"})
	}

	// Pastikan peminjaman milik user yang sedang login
	if borrowing.UserID != uint(id) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You are not authorized to return this book"})
	}

	// Cek apakah buku sudah dikembalikan
	if borrowing.Status == "returned" {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Book already returned"})
	}

	// Update status peminjaman dan tanggal pengembalian
	now := string(time.Now().Format("2006-01-02"))
	borrowing.Status = "returned"
	borrowing.ReturnDate = &now

	if err := database.DB.Save(&borrowing).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Tambahkan kembali stok buku
	var book models.Book
	if err := database.DB.First(&book, borrowing.BookID).Error; err == nil {
		book.Stock++
		database.DB.Save(&book)
	}

	// Respon sukses
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book returned successfully",
	})
}
