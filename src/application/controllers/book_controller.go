package controllers

import (
	"net/http"

	"github.com/erickrodrigs/go-webapi-template/src/application/usecases"
	"github.com/erickrodrigs/go-webapi-template/src/infrastructure/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateBookParams ...
type CreateBookParams struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Year        int    `json:"year"`
}

// BookController ...
type BookController struct {
	Database    *gorm.DB
	BookUseCase *usecases.BookUseCase
}

// Index ...
func (controller *BookController) Index(c *gin.Context) {
	books := controller.BookUseCase.GetAllBooks()
	c.JSON(200, gin.H{
		"books": books,
	})
}

// Create ...
func (controller *BookController) Create(c *gin.Context) {
	var json CreateBookParams
	err := c.ShouldBindJSON(&json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	title := json.Title
	author := json.Author
	description := json.Description
	year := json.Year

	book, err := controller.BookUseCase.CreateNewBook(title, author, description, year)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"book": book,
	})
}

// NewBookController ...
func NewBookController(db *gorm.DB) *BookController {
	bookController := BookController{}
	bookRepository := repository.BookRepositoryDb{
		Database: db,
	}

	bookController.BookUseCase = &usecases.BookUseCase{
		BookRepository: bookRepository,
	}

	return &bookController
}
