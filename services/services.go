package services

import (
	"github.com/emmanuelYohore/api-rest-crud-go/database"
	"github.com/emmanuelYohore/api-rest-crud-go/models"
)

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	result := database.DB.Find(&products)
	return products, result.Error
}
func GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	result := database.DB.First(&product, id)
	return &product, result.Error
}

func CreateProduct(product *models.Product) error {
	return database.DB.Create(product).Error
}

func UpdateProduct(product *models.Product) error {
	return database.DB.Save(product).Error
}

func DeleteProduct(id uint) error {
	return database.DB.Delete(&models.Product{}, id).Error
}


