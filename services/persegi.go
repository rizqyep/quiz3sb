package services

import (
	"quiz3-rizqyep/entities"
)

type PersegiService interface {
	HitungKeliling(request entities.PersegiQuery) int
	HitungLuas(request entities.PersegiQuery) int
}

type persegiService struct {
}

func NewPersegiService() PersegiService {
	return &persegiService{}
}

func (service *persegiService) HitungKeliling(request entities.PersegiQuery) int {
	return 4 * request.Sisi
}

func (service *persegiService) HitungLuas(request entities.PersegiQuery) int {
	return request.Sisi * request.Sisi
}
