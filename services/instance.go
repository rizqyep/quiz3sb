package services

import (
	"quiz3-rizqyep/repository"
	"sync"
)

type servicesPool struct {
	SegitigaSamaSisiService
	PersegiPanjangService
	PersegiService
	LingkaranService
	CategoryService
	BookService
}

var repositoryInstance = repository.NewRepository()
var serviceInstance *servicesPool
var once sync.Once

func InitServiceInstance() *servicesPool {
	once.Do(func() {
		serviceInstance = NewServiceInstance()
	})
	return serviceInstance
}

func NewServiceInstance() *servicesPool {
	return &servicesPool{
		SegitigaSamaSisiService: NewSegitigaSamaSisiService(),
		PersegiPanjangService:   NewPersegiPanjangService(),
		PersegiService:          NewPersegiService(),
		LingkaranService:        NewLingkaranService(),
		CategoryService:         NewCategoryService(repositoryInstance.CategoryRepository),
		BookService:             NewBookService(repositoryInstance.BookRepository),
	}
}
