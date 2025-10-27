package controllers

import (
	"net/http"
	"strconv"

	"github.com/emmanuelYohore/api-rest-crud-go/models"
	"github.com/emmanuelYohore/api-rest-crud-go/services"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products, err := services.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erreur récupération des products",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"produits": products,
	})
}

func GetProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := services.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "produits pas trouvé",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})

}

func CreateProduct(c *gin.Context) {
	var newProduct models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := services.CreateProduct(&newProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erreur création de produit",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": newProduct,
	})
}

func UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := services.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "produits pas trouvé",
		})
		return

	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	services.UpdateProduct(product)
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := services.DeleteProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "produits pas trouvé",
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"message": "produits supprimé",
	})
}
