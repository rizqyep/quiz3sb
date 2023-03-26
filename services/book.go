package services

import (
	"quiz3-rizqyep/entities"
	"quiz3-rizqyep/repository"
)

type BookService interface {
	GetAll() ([]entities.Book, error)
	Insert(book entities.Book) error
	GetById(id int) (entities.Book, error)
	Update(id int, book entities.Book) error
	Delete(id int) error
}

type bookService struct {
	repository repository.BookRepository
}

func NewBookService(repository repository.BookRepository) BookService {
	return &bookService{
		repository,
	}
}

func preProcessThickness(total_page int) string {
	if total_page > 200 {
		return "Tebal"
	} else if total_page > 100 {
		return "Sedang"
	} else {
		return "Tipis"
	}
}

func (service *bookService) GetAll() (result []entities.Book, err error) {
	return service.repository.GetAll()
}

func (service *bookService) GetById(id int) (result entities.Book, err error) {
	return service.repository.GetById(id)
}

func (service *bookService) Insert(book entities.Book) (err error) {
	book.Thickness = preProcessThickness(book.TotalPage)
	return service.repository.Insert(book)
}

func (service *bookService) Update(id int, book entities.Book) (err error) {
	book.Thickness = preProcessThickness(book.TotalPage)
	return service.repository.Update(id, book)
}

func (service *bookService) Delete(id int) (err error) {
	return service.repository.Delete(id)
}
