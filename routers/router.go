package routers

import (
	"quiz3-rizqyep/controllers"
	"quiz3-rizqyep/middleware"

	"github.com/gin-gonic/gin"
)

func RouteHandlers(r *gin.Engine) {
	controllerInstance := controllers.InitControllerInstance()

	bangunDatar := r.Group("/bangun-datar")
	bangunDatar.Use(middleware.HitungMiddleware())
	{
		bangunDatar.GET("/segitiga-samasisi", controllerInstance.SegitigaSamaSisiController.Hitung)
		bangunDatar.GET("/persegi", controllerInstance.PersegiController.Hitung)
		bangunDatar.GET("/persegi-panjang", controllerInstance.PersegiPanjangController.Hitung)
		bangunDatar.GET("/lingkaran", controllerInstance.LingkaranController.Hitung)
	}

	categories := r.Group("/categories")
	categories.GET("/", controllerInstance.CategoryController.GetAll)
	categories.POST("/", controllerInstance.CategoryController.Insert)
	categories.PUT("/:id", controllerInstance.CategoryController.Update)
	categories.GET("/:id", controllerInstance.GetById)
	categories.DELETE("/:id", controllerInstance.Delete)

}
