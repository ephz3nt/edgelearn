package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"go-edge/restapi-fiber/database"
)

type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

func NewBook(c *fiber.Ctx) {
	db := database.DBConn
	var book = new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
	}
	if (book.Title == "") || (book.Author == "") {
		c.Status(500).Send("name/author不能为空")
		return
	}
	db.Create(&book)
	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	db := database.DBConn
	id := c.Params("id")
	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No Book Found with ID")
		return
	}
	db.Delete(&book)
	c.Send("Book Successfully deleted")
}
