package controllers

import (
	"net/http"
	"quiz3-rizqyep/entities"
	"quiz3-rizqyep/services"

	"github.com/gin-gonic/gin"
)

type SegitigaSamaSisiController interface {
	Hitung(c *gin.Context)
}

type segitigaSamaSisiController struct {
	service services.SegitigaSamaSisiService
}

func NewSegitigaSamaSisiController(service services.SegitigaSamaSisiService) SegitigaSamaSisiController {
	return &segitigaSamaSisiController{
		service: service,
	}
}

func (controller *segitigaSamaSisiController) Hitung(c *gin.Context) {
	var request entities.SegitigaQuery

	c.ShouldBindQuery(&request)
	var result float64
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
