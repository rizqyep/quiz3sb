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
	categories.Use(gin.BasicAuth(gin.Accounts{
		"admin":  "password",
		"editor": "secret",
	}))
	categories.GET("/", controllerInstance.CategoryController.GetAll)
	categories.POST("/", controllerInstance.CategoryController.Insert)
	categories.PUT("/:id", controllerInstance.CategoryController.Update)
	categories.GET("/:id", controllerInstance.CategoryController.GetById)
	categories.DELETE("/:id", controllerInstance.CategoryController.Delete)

	books := r.Group("/books")
	books.Use(gin.BasicAuth(gin.Accounts{
		"admin":  "password",
		"editor": "secret",
	}))
	books.GET("/", controllerInstance.BookController.GetAll)
	books.POST("/", controllerInstance.BookController.Insert)
	books.PUT("/:id", controllerInstance.BookController.Update)
	books.GET("/:id", controllerInstance.BookController.GetById)
	books.DELETE("/:id", controllerInstance.BookController.Delete)

}
