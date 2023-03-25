package controllers

import (
	"net/http"
	"quiz3-rizqyep/entities"
	"quiz3-rizqyep/services"

	"github.com/gin-gonic/gin"
)

type PersegiController interface {
	Hitung(c *gin.Context)
}

type persegiController struct {
	service services.PersegiService
}

func NewPersegiController(service services.PersegiService) PersegiController {
	return &persegiController{
		service: service,
	}
}

func (controller *persegiController) Hitung(c *gin.Context) {
	var request entities.PersegiQuery

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
