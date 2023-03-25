package services

import (
	"sync"
)

type servicesPool struct {
	SegitigaSamaSisiService
	PersegiPanjangService
	PersegiService
	LingkaranService
}

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
	}
}
