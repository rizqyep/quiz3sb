package services

import (
	"math"
	"quiz3-rizqyep/entities"
)

type LingkaranService interface {
	HitungKeliling(request entities.LingkaranQuery) float64
	HitungLuas(request entities.LingkaranQuery) float64
}

type lingkaranService struct {
}

func NewLingkaranService() LingkaranService {
	return &lingkaranService{}
}

func (service *lingkaranService) HitungKeliling(request entities.LingkaranQuery) float64 {
	return 2 * math.Pi * request.JariJari
}

func (service *lingkaranService) HitungLuas(request entities.LingkaranQuery) float64 {
	return math.Pi * request.JariJari * request.JariJari
}
