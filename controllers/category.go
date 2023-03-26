package controllers

import (
	"net/http"
	"quiz3-rizqyep/entities"
	"quiz3-rizqyep/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	GetAll(c *gin.Context)
	Insert(c *gin.Context)
	GetById(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type categoryController struct {
	service services.CategoryService
}

func NewCategoryController(service services.CategoryService) CategoryController {
	return &categoryController{
		service,
	}
}

func (controller *categoryController) GetAll(c *gin.Context) {
	err, results := controller.service.GetAll()

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
		"message": "Successfully fetch categories!",
	})
}

func (controller *categoryController) Insert(c *gin.Context) {
	var category entities.Category

	c.ShouldBindJSON(&category)
	err := controller.service.Insert(category)
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
		"message": "Successfully created a new category!",
	})

}

func (controller *categoryController) GetById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	err, result := controller.service.GetById(id)
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
		"message": "Successfully fetch a category!",
		"data":    result,
	})

}

func (controller *categoryController) Update(c *gin.Context) {
	var category entities.Category
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.ShouldBindJSON(&category)
	err = controller.service.Update(id, category)
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
		"message": "Successfully updated a category!",
	})

}

func (controller *categoryController) Delete(c *gin.Context) {
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
		"message": "Successfully deleted a category!",
	})

}
