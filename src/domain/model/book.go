package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// BookRepositoryInterface ...
type BookRepositoryInterface interface {
	Register(book *Book) error
	All() []*Book
}

// Book ...
type Book struct {
	ID          string    `json:"id"`
	Title       string    `json:"title" gorm:"column:title;type:varchar(255)"`
	Author      string    `json:"author" gorm:"column:author;type:varchar(255)"`
	Description string    `json:"description" gorm:"column:description;type:varchar(255)"`
	Year        int       `json:"year" gorm:"column:year;type:integer"`
	CreatedAt   time.Time `json:"created_at"`
}

// NewBook ...
func NewBook(title string, author string, description string, year int) *Book {
	book := Book{
		Title:       title,
		Author:      author,
		Description: description,
		Year:        year,
	}

	book.ID = uuid.NewV4().String()
	book.CreatedAt = time.Now()

	return &book
}
