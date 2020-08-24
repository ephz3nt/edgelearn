package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go-edge/restapi-fiber/book"
	"go-edge/restapi-fiber/database"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
	//return app
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Connection Opened to Database")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)

	app.Listen(3000)
	defer database.DBConn.Close()
}
