package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/kyon1313/books/database"
	"github.com/kyon1313/books/helper"
	"github.com/kyon1313/books/model"
	queries "github.com/kyon1313/books/rawQueries.go"
)

func HandleAddAuthorBook(c *fiber.Ctx) error {
	bookRequest := &model.BookRequest{}
	helper.BodyParser(c, bookRequest)

	book, err := helper.AddBooks(bookRequest.BookDetails)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": err.Error(),
			"data":    nil,
		})
	}
	authorsID := helper.AddAuthors(bookRequest.Author)

	helper.AddAuthorBook(authorsID, book.BookID)
	return c.Status(200).JSON(&fiber.Map{
		"message": "data inserted",
		"data":    bookRequest,
	})
}

func HandleDeleteAuthorBook(c *fiber.Ctx) error {
	isbnParam := c.Params("isbn")

	book_id, err := helper.SelectBookId(isbnParam)
	if err != nil {
		return c.Status(204).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	err = helper.DeleteBookAuthor(book_id)
	if err != nil {
		return c.Status(204).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(&fiber.Map{
		"message": "authors book deleted",
	})
}

func HandleGetAuthorBooks(c *fiber.Ctx) error {
	var book []model.Book
	database.DB.Raw(queries.GetAllAuthorBooksQuery).Scan(&book)
	var bookResponse []model.AuthorsBookResponse

	for _, v := range book {
		var authors []model.AuthorTable

		err := json.Unmarshal([]byte(v.Authors), &authors)
		if err != nil {
			return c.JSON(&fiber.Map{
				"message": err.Error(),
			})
		}

		bookResponse = append(bookResponse, *model.ConstructAuthorBooksResponse(v, authors))
	}

	return c.JSON(bookResponse)

}

func HandleGetAuthorBookByIsbn(c *fiber.Ctx) error {
	isbnParam := c.Params("isbn")

	book, err := queries.QueryAuthorBookByIsbn(isbnParam)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": err,
			"data":    nil,
		})
	}

	var authors []model.AuthorTable
	err = json.Unmarshal([]byte(book.Authors), &authors)
	if err != nil {
		return c.JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	constructBook := model.ConstructAuthorBooksResponse(book, authors)
	return c.JSON(constructBook)
}

func DeleteAuthorOfBook(c *fiber.Ctx) error {
	book_id, _ := c.ParamsInt("book_id")
	author_id, _ := c.ParamsInt("author_id")
	var authorBook model.AuthorBook

	row := database.DB.Debug().Where("auth_id=? and book_id=?", author_id, book_id).Delete(&authorBook).RowsAffected
	if row == 0 {
		return c.JSON(&fiber.Map{
			"message": "no author deleted",
		})
	}
	return c.JSON(&fiber.Map{
		"message": "author of book deleted",
	})
}

func HandleUpdatAauthorBook(c *fiber.Ctx) error {
	isbnParam := c.Params("isbn")
	bookRequest := &model.BookRequest{}

	helper.BodyParser(c, bookRequest)

	authorsIds := helper.UpdateAuthorNames(bookRequest.Author)

	publisherId := helper.GetPublisherID(bookRequest.BookDetails.PublisherName)
	bookConstruct, err := helper.UpdatebookConstructor(bookRequest.BookDetails, isbnParam, publisherId)
	if err != nil {
		return c.JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	for _, v := range authorsIds {
		authorBook := &model.AuthorBook{}
		row := database.DB.Debug().Where("auth_id=? and book_id=?", v, bookConstruct.BookID).Find(&authorBook).RowsAffected
		if row == 0 {
			database.DB.Debug().Raw("insert into author_books (auth_id,book_id) values(?,?)", v, bookConstruct.BookID).Create(&authorBook)
		}
	}

	return c.JSON(&fiber.Map{
		"message": "data updated",
	})
}
