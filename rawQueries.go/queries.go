package queries

import (
	"errors"
	"fmt"

	"github.com/kyon1313/books/database"
	"github.com/kyon1313/books/model"
)

const (
	GetAllAuthorBooksQuery = `SELECT b.title,
		JSONB_AGG(
			JSONB_BUILD_OBJECT(
				'firstName', INITCAP(a.first_name),
				'middleName', INITCAP(a.middle_name),
				'lastName', INITCAP(a.last_name)
			)
		) AS authors,
		b.isbn13,
		b.isbn10,
		b.publication_year,
		p.publisher_name,
		b.edition,
		b.image_url,
		b.list_price AS price
	FROM book_tables b
	INNER JOIN
	author_books ab ON b.book_id = ab.book_id
	INNER JOIN
	author_tables a ON a.author_id = ab.auth_id
	INNER JOIN
	publishers p ON p.publisher_id = b.publisher_id
	WHERE
	 b.book_id = ab.book_id
	GROUP BY
	 b.title,
	 b.isbn13,
	 b.isbn10,
	 b.publication_year,
	 p.publisher_name,
	 b.edition,
	 b.image_url,
	 b.list_price`
)

func QueryAuthorBookByIsbn(isbnParam string) (model.Book, error) {
	book := model.Book{}
	query := fmt.Sprintf(`SELECT b.title,
	JSONB_AGG(
		JSONB_BUILD_OBJECT(
			'firstName', INITCAP(a.first_name),
			'middleName', INITCAP(a.middle_name),
			'lastName', INITCAP(a.last_name)
		)
	) AS authors,
	b.isbn13,
	b.isbn10,
	b.publication_year,
	p.publisher_name,
	b.edition,
	b.image_url,
	b.list_price AS price
FROM book_tables b
INNER JOIN
author_books ab ON b.book_id = ab.book_id
INNER JOIN
author_tables a ON a.author_id = ab.auth_id
INNER JOIN
publishers p ON p.publisher_id = b.publisher_id
WHERE
b.book_id = ab.book_id and (b.isbn10='%s' or b.isbn13='%s')
GROUP BY
 b.title,
 b.isbn13,
 b.isbn10,
 b.publication_year,
 p.publisher_name,
 b.edition,
 b.image_url,
 b.list_price`, isbnParam, isbnParam)

	row := database.DB.Raw(query).Find(&book).RowsAffected
	if row == 0 {
		return model.Book{}, errors.New("no data fetched")
	}

	return book, nil

}
