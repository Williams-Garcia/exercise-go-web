package services

import (
	"api_rest/models"
	"time"
)

var (
	ErrItemExist       = &models.CustomError{Msg: "El objeto a registrar ya existe"}
	ErrCodeValueRepeat = &models.CustomError{Msg: "El codigo del producto a agregar debe ser unico"}
	ErrDateExp         = &models.CustomError{Msg: "La fecha de expiracion del producto no es valida"}
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

func existProduct(pName string, products []models.Product) bool {
	for _, p := range products {
		if p.Name == pName {
			return true
		}
	}

	return false
}

func uniqueCodeValue(pCodeValue string, products []models.Product) bool {
	for _, p := range products {
		if p.CodeValue == pCodeValue {
			return true
		}
	}

	return false
}

func validDate(date time.Time) bool {
	t := time.Now()

	return t.After(date)
}

func parseDate(date string) (time.Time, error) {
	parseDate, err := time.Parse("01/02/2006", date)
	if err != nil {
		return parseDate, err
	}
	return parseDate, nil
}

func AddProduct(p models.Product, products []models.Product) (newProduct models.Product, err error) {
	if existProduct(p.Name, products) {
		return models.Product{}, ErrItemExist
	}

	if uniqueCodeValue(p.CodeValue, products) {
		return models.Product{}, ErrCodeValueRepeat
	}

	expDate, err := parseDate(p.Expiration)
	if err != nil {
		return models.Product{}, err
	}

	if validDate(expDate) {
		return models.Product{}, ErrDateExp
	}

	lastID := products[len(products)-1].Id

	p.Id = lastID + 1

	return p, nil
}
