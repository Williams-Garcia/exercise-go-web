package product

import (
	"api_rest/internal/domain"
	errorpkg "api_rest/pkg/error"

	"time"
)

var (
	ErrItemExist       = &errorpkg.CustomError{Msg: "El objeto a registrar ya existe"}
	ErrCodeValueRepeat = &errorpkg.CustomError{Msg: "El codigo del producto a agregar debe ser unico"}
	ErrDateExp         = &errorpkg.CustomError{Msg: "La fecha de expiracion del producto no es valida"}
)

func GetProductService(id int, products []domain.Product) domain.Product {
	var foundProduct domain.Product
	for _, product := range products {
		if product.Id == id {
			foundProduct = product
			break
		}
	}

	return foundProduct
}

func SearchProductService(query float64, products []domain.Product) []domain.Product {

	var filteredProduct []domain.Product
	for _, product := range products {
		if query != 0 && product.Price > query {
			filteredProduct = append(filteredProduct, product)
		}
	}
	return filteredProduct
}

func existProduct(pName string, products []domain.Product) bool {
	for _, p := range products {
		if p.Name == pName {
			return true
		}
	}

	return false
}

func uniqueCodeValue(pCodeValue string, products []domain.Product) bool {
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

func AddProduct(p domain.Product, products []domain.Product) (newProduct domain.Product, err error) {
	if existProduct(p.Name, products) {
		return domain.Product{}, ErrItemExist
	}

	if uniqueCodeValue(p.CodeValue, products) {
		return domain.Product{}, ErrCodeValueRepeat
	}

	expDate, err := parseDate(p.Expiration)
	if err != nil {
		return domain.Product{}, err
	}

	if validDate(expDate) {
		return domain.Product{}, ErrDateExp
	}

	lastID := products[len(products)-1].Id

	p.Id = lastID + 1

	return p, nil
}
