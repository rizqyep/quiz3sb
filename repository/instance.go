package repository

import "sync"

type repositoryPool struct {
	CategoryRepository
	BookRepository
}

var repositoryInstance *repositoryPool
var once *sync.Once

func NewRepository() *repositoryPool {
	return &repositoryPool{
		CategoryRepository: NewCategoryRepository(),
		BookRepository:     NewBookRepository(),
	}
}

func InitRepository() *repositoryPool {
	once.Do(func() {
		repositoryInstance = NewRepository()
	})
	return repositoryInstance
}
