package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kyon1313/books/handlers"
	"github.com/kyon1313/books/model"
	"github.com/kyon1313/books/templates"
)

func Routes(app *fiber.App) {
	endpoints := app.Group("/endpoint/v1")
	endpoints.Post("/addBook", handlers.HandleAddAuthorBook)
	endpoints.Get("/getBooks", handlers.HandleGetAuthorBooks)
	endpoints.Get("/getBookByIbsn/:isbn", handlers.HandleGetAuthorBookByIsbn)
	endpoints.Get("/deleteBookByIbsn/:isbn", handlers.HandleDeleteAuthorBook) //deletiona
	endpoints.Get("/deleteAuthorOfBook/:author_id/:book_id", handlers.DeleteAuthorOfBook)
	endpoints.Post("/updateAuthorBook/:isbn", handlers.HandleUpdatAauthorBook)

	//update book

	app.Post("/test", func(c *fiber.Ctx) error {
		form := model.RequestModelAuthorBook{}
		c.BodyParser(&form)
		fmt.Println("authors:", form.FirstName)
		return c.JSON(form)

	})

	template := app.Group("/template")
	template.Post("/submit", templates.TemplateAddAuthorBook)
	template.Get("/deleteBook/:isbn", templates.TemplateDeleteAuthorBook)
	template.Post("/updateBook/:isbn", templates.TemplateUpdateAuthorBook)

	renderPages := app.Group("/render")
	renderPages.Get("/", templates.RenderIndexTablePage)
	renderPages.Get("/flipBook", templates.RenderFlipBookPage)

}
