package routes

import (
	"github.com/androsyz/product-api/src/controllers"
	"github.com/gin-gonic/gin"
)

func ServerRoutes() {
	router := gin.Default()
	products := router.Group("/api/products")
	{
		products.GET("/", controllers.FilterProductsByCategory)
		products.GET("/:id", controllers.Show)
		products.POST("/", controllers.Create)
		products.PUT("/:id", controllers.Update)
		products.DELETE("/:id", controllers.Delete)
	}
	_ = router.Run()
}
