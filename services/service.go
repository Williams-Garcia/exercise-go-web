package services

import (
	"api_rest/models"
)

func GetProductService(id int, products []models.Product) models.Product {
	var foundProduct models.Product
	for _, product := range products {
		if product.Id == id {
			foundProduct = product
			break
		}
	}

	return foundProduct
}

func SearchProductService(query float64, products []models.Product) []models.Product {

	var filteredProduct []models.Product
	for _, product := range products {
		if query != 0 && product.Price > query {
			filteredProduct = append(filteredProduct, product)
		}
	}
	return filteredProduct
}
