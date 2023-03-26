package controllers

import (
	"quiz3-rizqyep/services"
	"sync"
)

type controllersPool struct {
	SegitigaSamaSisiController
	PersegiController
	PersegiPanjangController
	LingkaranController
	CategoryController
}

var serviceInstance = services.InitServiceInstance()
var controllerInstance *controllersPool
var once sync.Once

func InitControllerInstance() *controllersPool {
	once.Do(func() {
		controllerInstance = NewControllerInstance()
	})
	return controllerInstance
}

func NewControllerInstance() *controllersPool {
	return &controllersPool{
		SegitigaSamaSisiController: NewSegitigaSamaSisiController(serviceInstance.SegitigaSamaSisiService),
		PersegiController:          NewPersegiController(serviceInstance.PersegiService),
		PersegiPanjangController:   NewPersegiPanjangController(serviceInstance.PersegiPanjangService),
		LingkaranController:        NewLingkaranController(serviceInstance.LingkaranService),
		CategoryController:         NewCategoryController(serviceInstance.CategoryService),
	}
}
