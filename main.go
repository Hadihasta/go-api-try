package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"gorm.io/gorm"
)

// javascript by default dapat mengerti format json namun tidak semua bahasa dapat mengerti json
// maka dari itu golang harus di beri tahu format ini tidak seperti javascript

type Book struct{
	Author string `json:"author"`
	Title string	`json:"title"`
	Publisher string	`json:"publisher"`
}

// declare data type sendiri bernama repository
type Repository struct{
	DB *gorm.DB
}


func(r *Repository) SetupRoutes(app *fiber.App){
	api := app.Group("/api")
	api.Post("/create_books",  r.CreateBook)
	api.Delete("delete_book/:id", r.DeleteBook)
	api.Get("/get_books/:id", r.GetBookById)
	api.Get("/books",r.GetBooks)
}

func main(){
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal.(err)
	}

	db, err := storage.NewConnetion(config)
	if err != nil {
		log.Fatal("could not load the database")
	}

	r := Repository{
		DB: db,
	}

// same like express in Js but faster
	app := fiber.New()
	// setup routes is struct method seperti function method di javascript
	r.SetupRoutes(app)
	app.Listen(":8080")
}