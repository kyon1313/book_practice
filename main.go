package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/kyon1313/books/database"
	"github.com/kyon1313/books/endpoints"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.Migration()
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/ui", "./ui")

	app.Use(cors.New())

	endpoints.Routes(app)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}
