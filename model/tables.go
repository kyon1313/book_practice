package model

type AuthorTable struct {
	AuthorID   uint   `json:"-" gorm:"primaryKey"`
	FirstName  string `json:"firstName" gorm:"not null"`
	LastName   string `json:"lastName" gorm:"not null"`
	MiddleName string `json:"middleName"`
}

type BookTable struct {
	BookID          uint    `json:"-" gorm:"primaryKey;index"`
	Title           string  `json:"title"`
	ISBN13          string  `json:"isbn13"`
	ISBN10          string  `json:"isbn10"`
	ListPrice       float64 `json:"listPrice"`
	PublicationYear int     `json:"publicationYear"`
	ImageURL        string  `json:"imageURL,omitempty"`
	Edition         string  `json:"edition,omitempty"`
	PublisherID     uint    `json:"publisherId" gorm:"index"`
}
type Publisher struct {
	PublisherID   uint   `json:"publisherId" gorm:"primaryKey"`
	PublisherName string `json:"publisherName"`
}

type AuthorBook struct {
	AuthorBookId uint        `json:"-" gorm:"primaryKey;index"`
	AuthId       uint        `json:"authId" gorm:"index"`
	Author       AuthorTable `gorm:"foreignKey:AuthId"`
	BookId       uint        `json:"bookId" gorm:"index"`
	Book         BookTable   `gorm:"foreignKey:BookId"`
}
