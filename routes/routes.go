package routes

import (
	"net/http"

	"github.com/emmanuelYohore/api-rest-crud-go/controllers"
	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "home",
		})
	})

	router.GET("/api/products", controllers.GetProducts)
	router.GET("/api/product/:id", controllers.GetProductByID)
	router.POST("/api/product", controllers.CreateProduct)
	router.PUT("/api/product/:id", controllers.UpdateProduct)
	router.DELETE("/api/product/:id", controllers.DeleteProduct)

	router.Run(":8080")
}
