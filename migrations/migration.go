package main

import (
	"github.com/rifkyrizkita/book_management/database"
	"github.com/rifkyrizkita/book_management/models"
)

func init() {
	database.InitDB()
}

func main() {
	database.DB.AutoMigrate(&models.User{}, models.Book{}, models.Borrowing{}, models.Category{})
}
