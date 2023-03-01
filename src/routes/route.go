package routes

import (
	"github.com/androsyahreza/product-api/src/controllers"
	"github.com/gin-gonic/gin"
)

func ServerRoutes() {
	router := gin.Default()
	products := router.Group("/api/products")
	{
		products.GET("/api/products", controllers.Index)
		products.GET("/api/product/:id", controllers.Show)
		products.POST("/api/product", controllers.Create)
		products.PUT("/api/product/:id", controllers.Update)
		products.DELETE("/api/product", controllers.Delete)
	}
	_ = router.Run()
}
