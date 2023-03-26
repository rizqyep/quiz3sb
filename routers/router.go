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
	categories.GET("/:id", controllerInstance.CategoryController.GetById)

	categories.Use(gin.BasicAuth(gin.Accounts{
		"admin":  "password",
		"editor": "secret",
	}))

	categories.POST("/", controllerInstance.CategoryController.Insert)
	categories.PUT("/:id", controllerInstance.CategoryController.Update)
	categories.DELETE("/:id", controllerInstance.CategoryController.Delete)

	books := r.Group("/books")
	books.GET("/", controllerInstance.BookController.GetAll)
	books.GET("/:id", controllerInstance.BookController.GetById)
	r.Use(gin.BasicAuth(gin.Accounts{
		"admin":  "password",
		"editor": "secret",
	}))

	books.POST("/", controllerInstance.BookController.Insert)
	books.PUT("/:id", controllerInstance.BookController.Update)

	books.DELETE("/:id", controllerInstance.BookController.Delete)

}
