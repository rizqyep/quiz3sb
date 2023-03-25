package controllers

import (
	"net/http"
	"quiz3-rizqyep/entities"
	"quiz3-rizqyep/services"

	"github.com/gin-gonic/gin"
)

type LingkaranController interface {
	Hitung(c *gin.Context)
}

type lingkaranController struct {
	service services.LingkaranService
}

func NewLingkaranController(service services.LingkaranService) LingkaranController {
	return &lingkaranController{
		service: service,
	}
}

func (controller *lingkaranController) Hitung(c *gin.Context) {
	var request entities.LingkaranQuery

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
