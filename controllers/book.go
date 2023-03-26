package controllers

import (
	"net/http"
	"net/url"
	"quiz3-rizqyep/entities"
	"quiz3-rizqyep/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController interface {
	GetAll(c *gin.Context)
	Insert(c *gin.Context)
	GetById(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type bookController struct {
	service services.BookService
}

func NewBookController(service services.BookService) BookController {
	return &bookController{
		service,
	}
}

func validateBookRequest(book entities.Book, c *gin.Context) (bool, map[string]string) {
	errors := make(map[string]string)
	if book.ReleaseYear < 1980 || book.ReleaseYear > 2021 {
		errors["release_year"] = "Book release year should be in 1980 - 2021 range"
	}

	_, err := url.ParseRequestURI(book.ImageUrl)
	if err != nil {
		errors["image_url"] = "Image URL should be a valid url!"
	}

	if len(errors) != 0 {

		return false, errors
	}
	return true, errors
}

func (controller *bookController) GetAll(c *gin.Context) {
	results, err := controller.service.GetAll()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"data":    nil,
			"message": "Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   nil,
		"data":    results,
		"message": "Successfully fetch books!",
	})
}

func (controller *bookController) Insert(c *gin.Context) {

	var book entities.Book

	c.ShouldBindJSON(&book)

	validated, errors := validateBookRequest(book, c)
	if !validated {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"validation_errors": errors,
		})
		return
	}
	err := controller.service.Insert(book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"data":    nil,
			"message": "Error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"error":   nil,
		"message": "Successfully created a new book!",
	})

}

func (controller *bookController) GetById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	result, err := controller.service.GetById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"data":    nil,
			"message": "Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   nil,
		"message": "Successfully fetch a book!",
		"data":    result,
	})

}

func (controller *bookController) Update(c *gin.Context) {
	var book entities.Book
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.ShouldBindJSON(&book)
	validated, errors := validateBookRequest(book, c)
	if !validated {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"validation_errors": errors,
		})
		return
	}
	err = controller.service.Update(id, book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"data":    nil,
			"message": "Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   nil,
		"message": "Successfully updated a book!",
	})

}

func (controller *bookController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	err = controller.service.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"data":    nil,
			"message": "Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   nil,
		"message": "Successfully deleted a book!",
	})

}
