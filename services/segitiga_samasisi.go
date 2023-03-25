package services

import "quiz3-rizqyep/entities"

type SegitigaSamaSisiService interface {
	HitungKeliling(request entities.SegitigaQuery) float64
	HitungLuas(request entities.SegitigaQuery) float64
}

type segitigaSamaSisiService struct {
}

func NewSegitigaSamaSisiService() SegitigaSamaSisiService {
	return &segitigaSamaSisiService{}
}

func (service *segitigaSamaSisiService) HitungKeliling(request entities.SegitigaQuery) float64 {
	return request.Alas * 3
}

func (service *segitigaSamaSisiService) HitungLuas(request entities.SegitigaQuery) float64 {
	return request.Alas * request.Tinggi / 2
}
