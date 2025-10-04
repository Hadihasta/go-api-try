package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	// ambil dari folder models & storage ini kalau tidak import otomatis bisa di tambahkan manual
	// run go mod tidy 
	"github.com/go-api-try/models"
	"github.com/go-api-try/storage"

	"gorm.io/gorm"
)

// javascript by default dapat mengerti format json namun tidak semua bahasa dapat mengerti json
// maka dari itu golang harus di beri tahu format ini tidak seperti javascript

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

// declare data type sendiri bernama repository
type Repository struct {
	DB *gorm.DB
}

// declare struct method sendiri sebagai function for createbook  dengan r sebagai repository
func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}

	// convert json yang di terima dengan fitur si fiber.ctx
	err := context.BodyParser{&book}

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	// r ada acces ke db sesuai yang di declare di bawah
	// buat data buku ke database dan jika tidak kirim message error
	err := r.DB.Create(&book).Error
	if err != nil{
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message":"cannot create book"}
		)
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"book has been created"
	})
	// return nya kenapa nil karena di atas kita declare expect error
	return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error{
	// ini [] data type slice
	bookModels := &[]models.Books{}

	err := r.DB.Find(bookModels).Error

	if err != nil{
		// jika error return errornya
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message":"could not get books"}
		)
		return err
	}
// jika tidak ada eror maka berhasil
context.Status(http.StatusOK).JSON(&fiber.Map{
	"message":"books fetched succesfully",
	"data": bookModels,
})
// balikin nil nya karna kalau tidak sama dengan nill masuk di atas
return nil

}

// fiber framework sangat membantu untuk mendapatkan HTTP status tapi jika ingin menjadi golang developer
// yang baik maka ada baiknya untuk mengetahui setiap bentuk dasar tanpa menggunakan framework apapun
// dan membuatnya dengan apa yang golang sediakan
func (r *Repository) DeleteBook(context *fiber.Ctx) error{
	bookModel := models.Books{}
	id := context,Params("id")
	if id == ""{
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message":"id cannot be empty"
		})
		return nil
	}

	err := r.DB.Delete(bookModel, id)

	if err != nil{
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message":"could not delete books",
		})
		return err.Error
	}

	context.Status(http.StatusOK).JSON(&fiber.Map({
		"message": "books delete succesfully"
	}))
	// nil maksudnya adalah error
	return nil
}
// *GREAT NOTE: setiap membuat program jangan mencoba untuk berasumsi, cukup buka dokumentasi
// dan ambil semua yang kamu butuhkan


func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("delete_book/:id", r.DeleteBook)
	api.Get("/get_books/:id", r.GetBookById)
	api.Get("/books", r.GetBooks)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal.(err)
	}



	config :=  &storage.Config{
	 Host: os.Getenv("DB_HOST"),
	 Port: os.Getenv("DB_PORT"),
	 Password: os.Getenv("DB_PASS"),
	 User: os.Getenv("DB_USER"),
	 DBName: os.Getenv("DB_SSLMODE"),
	 SSLMode: os.Getenv("DB_NAME"),
	}
	// dari package storage do postgres.go
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
