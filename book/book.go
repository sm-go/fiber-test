package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/smith-golang/fiber-test/database"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
	gorm.Model
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	err := c.JSON(books)
	if err != nil {
		panic("error return json")
	}
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	book := new(Book)
	db.First(&book, id)
	if book.ID < 1 {
		c.JSON("Record not found with " + id)
	}
	c.JSON(book)
}

func NewBook(c *fiber.Ctx) {
	db := database.DBConn
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
	}
	db.Create(&book)
	c.JSON(book)
}

func EditBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	req := new(Book)
	if err := c.BodyParser(req); err != nil {
		c.Status(503).Send(err)
	}
	book := new(Book)
	db.First(&book, id)
	book.Title = req.Title
	book.Author = req.Author
	book.Rating = req.Rating
	db.Save(&book)
	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No record found!")
		return
	}
	db.Delete(&book)
	c.Send("Book deleted!")
}
