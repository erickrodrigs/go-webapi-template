package repository

import (
	"github.com/erickrodrigs/go-webapi-template/src/domain/model"
	"gorm.io/gorm"
)

// BookRepositoryDb ...
type BookRepositoryDb struct {
	Database *gorm.DB
}

// Register ...
func (repo BookRepositoryDb) Register(book *model.Book) error {
	err := repo.Database.Create(book).Error

	if err != nil {
		return err
	}

	return nil
}

// All ...
func (repo BookRepositoryDb) All() []*model.Book {
	var books []*model.Book
	repo.Database.Find(&books)
	return books
}
