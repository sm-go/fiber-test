package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/smith-golang/fiber-test/book"
	"github.com/smith-golang/fiber-test/database"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("fail to connect to database")
	}
	fmt.Println("successfully connected to database")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func setupBookRoute(app *fiber.App) {
	v1 := app.Group("api/v1")
	v1.Get("/books", book.GetBooks)
	v1.Get("/book/:id", book.GetBook)
	v1.Post("/book", book.NewBook)
	v1.Put("/book/:id", book.EditBook)
	v1.Delete("/book/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()
	setupBookRoute(app)
	app.Listen(3000)
}
