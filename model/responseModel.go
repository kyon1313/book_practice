package model

type Book struct {
	Title           string  `json:"title"`
	Authors         string  `json:"authors"`
	ISBN13          string  `json:"isbn13"`
	ISBN10          string  `json:"isbn10"`
	PublicationYear int     `json:"publicationYear"`
	PublisherName   string  `json:"publisherName"`
	Edition         string  `json:"edition"`
	Price           float64 `json:"price"`
	ImageURL        string  `json:"imageUrl,omitempty"`
}

type AuthorsBookResponse struct {
	Title           string        `json:"title"`
	Authors         []AuthorTable `json:"authors"`
	ISBN13          string        `json:"isbn13,omitempty"`
	ISBN10          string        `json:"isbn10,omitempty"`
	PublicationYear int           `json:"publication_year"`
	PublisherName   string        `json:"publisher_name"`
	Edition         string        `json:"edition,omitempty"`
	Price           float64       `json:"price"`
	ImageURL        string        `json:"imageUrl,omitempty"`
}

func ConstructAuthorBooksResponse(book Book, author []AuthorTable) *AuthorsBookResponse {
	return &AuthorsBookResponse{
		Title:           book.Title,
		Authors:         author,
		ISBN13:          book.ISBN13,
		ISBN10:          book.ISBN10,
		PublicationYear: book.PublicationYear,
		PublisherName:   book.PublisherName,
		Edition:         book.Edition,
		Price:           book.Price,
		ImageURL:        book.ImageURL,
	}
}
