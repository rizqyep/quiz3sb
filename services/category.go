package services

import (
	"quiz3-rizqyep/entities"
	"quiz3-rizqyep/repository"
)

type CategoryService interface {
	GetAll() (error, []entities.Category)
	Insert(category entities.Category) error
	GetById(id int) (error, entities.Category)
	Update(id int, category entities.Category) error
	Delete(id int) error
}

type categoryService struct {
	repository repository.CategoryRepository
}

func NewCategoryService(repository repository.CategoryRepository) CategoryService {
	return &categoryService{
		repository,
	}
}

func (service *categoryService) GetAll() (err error, results []entities.Category) {
	return service.repository.GetAll()
}

func (service *categoryService) GetById(id int) (err error, result entities.Category) {
	return service.repository.GetById(id)
}

func (service *categoryService) Insert(category entities.Category) (err error) {
	return service.repository.Insert(category)
}

func (service *categoryService) Update(id int, category entities.Category) (err error) {
	return service.repository.Update(id, category)
}

func (service *categoryService) Delete(id int) (err error) {
	return service.repository.Delete(id)
}
