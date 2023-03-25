package services

import (
	"quiz3-rizqyep/entities"
)

type PersegiPanjangService interface {
	HitungKeliling(request entities.PersegiPanjangQuery) int
	HitungLuas(request entities.PersegiPanjangQuery) int
}

type persegiPanjangService struct {
}

func NewPersegiPanjangService() PersegiPanjangService {
	return &persegiPanjangService{}
}

func (service *persegiPanjangService) HitungKeliling(request entities.PersegiPanjangQuery) int {
	return (2 * request.Panjang) + (2 * request.Lebar)
}

func (service *persegiPanjangService) HitungLuas(request entities.PersegiPanjangQuery) int {
	return request.Panjang * request.Lebar
}
