package model

type BookRequest struct {
	Author []AuthorTable `json:"authors"`
	BookDetails
}

type BookDetails struct {
	Title           string  `json:"title"`
	ISBN            string  `json:"isbn"`
	ISBN10          string  `json:"isbn10,omitempty"`
	ListPrice       float64 `json:"listPrice"`
	PublicationYear int     `json:"publicationYear"`
	ImageURL        string  `json:"imageURL,omitempty"`
	Edition         string  `json:"edition,omitempty"`
	PublisherName   string  `json:"publisherName"`
}

type RequestModelAuthorBook struct {
	FirstName  []string `form:"first_name"`
	LastName   []string `form:"last_name"`
	MiddleName []string `form:"middle_name"`
	BookDetails
}
