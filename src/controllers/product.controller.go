package controllers

import (
	"log"
	"net/http"

	"github.com/androsyahreza/product-api/src/models"
	"github.com/gin-gonic/gin"
)

func FilterProductsByCategory(c *gin.Context) {
	var products []models.Product
	category := c.Query("category")

	if category != "" {
		err := models.DB.Where("category = ?", category).Find(&products).Error
		if err != nil {
			log.Fatal(err.Error())
		}
	} else {
		err := models.DB.Find(&products).Error
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
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
