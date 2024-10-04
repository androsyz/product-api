package controllers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/androsyahreza/product-api/src/models"
	"github.com/gin-gonic/gin"
)

func FilterProductsByCategory(c *gin.Context) {
	var products []models.Product
	category := c.Query("category")

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		pageSize = 5
	}

	dbQuery := models.DB.Limit(pageSize).Offset((page - 1) * pageSize)
	if category != "" {
		dbQuery = dbQuery.Where("category = ?", category)
	}
	if err := dbQuery.Find(&products).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products", "message": err.Error()})
		return
	}

	var totalProducts int64
	dbCountQuery := models.DB.Model(&models.Product{})
	if category != "" {
		dbCountQuery = dbCountQuery.Where("category = ?", category)
	}
	if err := dbCountQuery.Count(&totalProducts).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to count products", "message": err.Error()})
		return
	}

	totalPages := int(math.Ceil(float64(totalProducts) / float64(pageSize)))

	c.JSON(http.StatusOK, gin.H{
		"products":       products,
		"page":           page,
		"page_size":      pageSize,
		"total_pages":    totalPages,
		"total_products": totalProducts,
	})
}

func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := models.DB.Create(&product).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Product not found", "message": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "message": err.Error()})
		return
	}

	if err := models.DB.Save(&product).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product successfully updated", "product": product})
}

func Delete(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Product not found", "message": err.Error()})
		return
	}

	if err := models.DB.Delete(&product).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product successfully deleted"})
}
