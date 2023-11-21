package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kyon1313/books/database"
	"github.com/kyon1313/books/model"
)

func BodyParser(c *fiber.Ctx, in interface{}) error {
	err := c.BodyParser(in)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}
	return err
}

// Check if an ISBN-10 is valid
func isValidISBN10(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	isbn = strings.ReplaceAll(isbn, " ", "")

	if len(isbn) != 10 {
		return false
	}

	checksum := 0
	for i, char := range isbn {
		if i == 9 && char == 'X' {
			checksum += 10
		} else {
			digit, err := strconv.Atoi(string(char))
			if err != nil {
				return false
			}
			checksum += digit * (10 - i)
		}
	}

	return checksum%11 == 0
}

// Check if an ISBN-13 is valid
func isValidISBN13(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	isbn = strings.ReplaceAll(isbn, " ", "")

	if len(isbn) != 13 {
		return false
	}

	checksum := 0
	for i, char := range isbn {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}
		if i%2 == 0 {
			checksum += digit
		} else {
			checksum += digit * 3
		}
	}

	return checksum%10 == 0
}

func UnmarshalFormRequest(form *model.RequestModelAuthorBook) ([]byte, error) {
	var authors []model.AuthorTable

	for i := range form.FirstName {
		author := model.AuthorTable{
			FirstName:  form.FirstName[i],
			LastName:   form.LastName[i],
			MiddleName: form.MiddleName[i],
		}
		authors = append(authors, author)
	}

	bookRequest := model.BookRequest{
		Author: authors,
		BookDetails: model.BookDetails{
			Title:           form.Title,
			ISBN:            form.ISBN,
			ISBN10:          form.ISBN10,
			ListPrice:       form.ListPrice,
			PublicationYear: form.PublicationYear,
			ImageURL:        form.ImageURL,
			Edition:         form.Edition,
			PublisherName:   form.PublisherName,
		},
	}

	addBookRequest, err := json.Marshal(bookRequest)
	if err != nil {
		return []byte{}, errors.New("marshall failed")
	}
	return addBookRequest, nil
}

func SelectBookId(isbn string) (int, error) {
	var bookId int
	err := database.DB.Raw("select book_id from book_tables where  isbn13=?", isbn).Scan(&bookId).Error
	if err != nil {
		return 0, err
	}
	return bookId, nil
}

func DeleteBookAuthor(book_id int) error {
	var bookAuthors model.AuthorBook
	err := database.DB.Where("book_id=?", book_id).Delete(&bookAuthors).Error
	if err != nil {
		return err
	}
	var book model.BookTable

	err = database.DB.Where("book_id=?", book_id).Delete(&book).Error
	if err != nil {
		return err
	}
	return nil
}

func SelectAuthorBookByBookId(book_id int) {
	var authorBook []model.AuthorTable
	database.DB.Find(&authorBook)

}

func AddBooks(book model.BookDetails) (*model.BookTable, error) {
	if len(book.ISBN) != 10 && len(book.ISBN) != 13 {
		return &model.BookTable{}, errors.New("invalid ISBN")
	}

	if len(book.ISBN) == 10 {
		if !isValidISBN10(book.ISBN) {
			fmt.Println("not valid isbn10")
			return &model.BookTable{}, errors.New("invalid ISBN-10")
		}
	}
	if len(book.ISBN) == 13 {
		if !isValidISBN13(book.ISBN) {
			fmt.Println("not valid isbn13")
			return &model.BookTable{}, errors.New("invalid ISBN-13")

		}
	}

	publisherID := GetPublisherID(book.PublisherName)
	constructedBook := bookConstructor(book, publisherID)

	return constructedBook, nil
}

func GetPublisherID(pubName string) uint {
	publisher := &model.Publisher{}
	query := database.DB.Debug().Where("publisher_name=?", pubName).Find(publisher)
	if query.RowsAffected > 0 {
		database.DB.Debug().Save(publisher)
		return publisher.PublisherID

	} else {
		publisher.PublisherName = pubName
		database.DB.Debug().Create(publisher)

	}
	return publisher.PublisherID
}

func bookConstructor(bookParam model.BookDetails, pubID uint) *model.BookTable {
	book := &model.BookTable{
		Title:           bookParam.Title,
		ListPrice:       bookParam.ListPrice,
		PublicationYear: bookParam.PublicationYear,
		PublisherID:     pubID,
		ImageURL:        bookParam.ImageURL,
		Edition:         bookParam.Edition,
	}

	if len(bookParam.ISBN) == 13 {
		book.ISBN13 = bookParam.ISBN
	}
	if len(bookParam.ISBN) == 10 {
		book.ISBN10 = bookParam.ISBN
	}
	database.DB.Debug().Create(book)

	return book
}

func UpdatebookConstructor(bookParam model.BookDetails, isbnPAram string, pubID uint) (*model.BookTable, error) {
	var books model.BookTable

	result := database.DB.Debug().Where("isbn10=? or isbn13=?", isbnPAram, isbnPAram).Find(&books).RowsAffected
	if result == 0 {
		return &model.BookTable{}, errors.New("no data fetched")
	}

	book := &model.BookTable{
		Title:           bookParam.Title,
		ListPrice:       bookParam.ListPrice,
		ISBN13:          bookParam.ISBN,
		ISBN10:          bookParam.ISBN10,
		PublicationYear: bookParam.PublicationYear,
		PublisherID:     pubID,
		ImageURL:        bookParam.ImageURL,
		Edition:         bookParam.Edition,
	}

	database.DB.Debug().Where("book_id", books.BookID).Updates(&book).Find(&book)
	return book, nil
}

func AddAuthors(authors []model.AuthorTable) []uint {
	authorsId := []uint{}
	for _, author := range authors {
		authorNameToLowerCase(&author)
		query := database.DB.Debug().Where("first_name=? and last_name=?", author.FirstName, author.LastName).Find(&author)
		if query.RowsAffected > 0 {
			authorsId = append(authorsId, author.AuthorID)
		} else {
			database.DB.Debug().Save(&author)
			authorsId = append(authorsId, author.AuthorID)
		}
	}
	return authorsId

}

func AddAuthorBook(authorsId []uint, bookId uint) {
	for _, v := range authorsId {
		var authorBook model.AuthorBook
		database.DB.Raw("insert into author_books (auth_id,book_id) values(?,?)", v, bookId).Create(&authorBook)
	}
}

func authorNameToLowerCase(author *model.AuthorTable) *model.AuthorTable {
	return &model.AuthorTable{
		AuthorID:   author.AuthorID,
		FirstName:  strings.ToLower(author.FirstName),
		LastName:   strings.ToLower(author.LastName),
		MiddleName: strings.ToLower(author.MiddleName),
	}
}

func UpdateAuthorNames(authors []model.AuthorTable) []int {
	var authID []int

	for _, v := range authors {
		author := authorNameToLowerCase(&v)
		row := database.DB.Debug().Where("first_name=? and last_name=?", author.FirstName, author.LastName).Find(&author).RowsAffected

		if row == 0 {
			database.DB.Debug().Create(&author)
			authID = append(authID, int(author.AuthorID))

		} else {
			authID = append(authID, int(author.AuthorID))
		}
	}
	return authID
}
