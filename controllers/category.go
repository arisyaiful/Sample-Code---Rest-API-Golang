package controllers

import (
	"net/http"
	"sample-api/models"

	"github.com/gin-gonic/gin"
)

func FindCategory(c *gin.Context) {
	var category []models.Category
	models.DB.Find(&category)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

func CreateCategory(c *gin.Context) {
	var input models.CreateCategory

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := models.Category{Name: input.Name}

	models.DB.Create(&category)

	c.JSON(http.StatusCreated, gin.H{"data": category})
}

func FindCategoryById(c *gin.Context) {
	var category models.Category

	if err := models.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

func UpdateCategory(c *gin.Context) {
	var category models.Category
	var input models.UpdateCategory

	if err := models.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&category).Updates(models.Category{Name: input.Name})

	c.JSON(http.StatusOK, gin.H{"data": category})
}

func DeleteCategory(c *gin.Context) {
	var category models.Category

	if err := models.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
		return
	}

	models.DB.Delete(&category)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
