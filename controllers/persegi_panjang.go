package controllers

import (
	"net/http"
	"quiz3-rizqyep/entities"
	"quiz3-rizqyep/services"

	"github.com/gin-gonic/gin"
)

type PersegiPanjangController interface {
	Hitung(c *gin.Context)
}

type persegiPanjangController struct {
	service services.PersegiPanjangService
}

func NewPersegiPanjangController(service services.PersegiPanjangService) PersegiPanjangController {
	return &persegiPanjangController{
		service: service,
	}
}

func (controller *persegiPanjangController) Hitung(c *gin.Context) {
	var request entities.PersegiPanjangQuery

	c.ShouldBindQuery(&request)
	var result int
	if request.Hitung == "luas" {
		result = controller.service.HitungLuas(request)
	} else {
		result = controller.service.HitungKeliling(request)
	}

	c.JSON(http.StatusOK, gin.H{
		"perhitungan": request.Hitung,
		"result":      result,
	})
}
