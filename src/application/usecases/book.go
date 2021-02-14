package usecases

import "github.com/erickrodrigs/go-webapi-template/src/domain/model"

// BookUseCase ...
type BookUseCase struct {
	BookRepository model.BookRepositoryInterface
}

// CreateNewBook ...
func (usecase *BookUseCase) CreateNewBook(title string, author string, description string, year int) (*model.Book, error) {
	book := model.NewBook(title, author, description, year)
	err := usecase.BookRepository.Register(book)

	if err != nil {
		return nil, err
	}

	return book, err
}

// GetAllBooks ...
func (usecase *BookUseCase) GetAllBooks() []*model.Book {
	return usecase.BookRepository.All()
}
