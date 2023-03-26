package repository

import "sync"

type repositoryPool struct {
	CategoryRepository
}

var repositoryInstance *repositoryPool
var once *sync.Once

func NewRepository() *repositoryPool {
	return &repositoryPool{
		CategoryRepository: NewCategoryRepository(),
	}
}

func InitRepository() *repositoryPool {
	once.Do(func() {
		repositoryInstance = NewRepository()
	})
	return repositoryInstance
}
