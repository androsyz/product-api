package controllers

import (
	"log"
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

	if category != "" {
		err := models.DB.Where("category = ?", category).Limit(pageSize).Offset((page - 1) * pageSize).Find(&products).Error
		if err != nil {
			log.Fatal(err.Error())
		}
	} else {
		err := models.DB.Find(&products).Limit(pageSize).Offset((page - 1) * pageSize).Find(&products).Error
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	var totalProducts int64
	if category != "" {
		err = models.DB.Model(&models.Product{}).Where("category = ?", category).Count(&totalProducts).Error
		if err != nil {
			log.Fatal(err.Error())
		}
	} else {
		err = models.DB.Find(&products).Count(&totalProducts).Limit(pageSize).Offset((page - 1) * pageSize).Find(&products).Error
		if err != nil {
			log.Fatal(err.Error())
		}
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
	} else {
		c.JSON(http.StatusOK, gin.H{"product": product})
	}
}

func Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.Model(&product).Where("id = ?", id).First(&product).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not found!"})
	} else {
		c.ShouldBindJSON(&product)
		models.DB.Save(&product)
		c.JSON(http.StatusOK, gin.H{"message": "Data successfully updated!"})
	}
}

func Delete(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.Model(&product).Where("id = ?", id).First(&product).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not found!"})
	} else {
		c.ShouldBindJSON(&product)
		models.DB.Delete(&product)
		c.JSON(http.StatusOK, gin.H{"message": "Data successfully deleted!"})
	}
}
