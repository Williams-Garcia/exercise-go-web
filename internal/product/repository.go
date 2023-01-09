package product

import (
	"api_rest/internal/domain"
	errorpkg "api_rest/pkg/error"
	"encoding/json"
	"log"
	"os"
	"time"
)

type ProductRepository struct {
	products domain.ListProducts
}

var (
	errOpen            = &errorpkg.CustomError{Msg: "Error: no se puede abrir el archivo"}
	ErrParser          = &errorpkg.CustomError{Msg: "Error: no se puede parsear el objeto"}
	ErrItemExist       = &errorpkg.CustomError{Msg: "El objeto a registrar ya existe"}
	ErrCodeValueRepeat = &errorpkg.CustomError{Msg: "El codigo del producto a agregar debe ser unico"}
	ErrDateExp         = &errorpkg.CustomError{Msg: "La fecha de expiracion del producto no es valida"}
	ErrNotFound        = &errorpkg.CustomError{Msg: "Objeto no encontrado"}
)

func NewProductRepository() *ProductRepository {
	pr := &ProductRepository{}
	pr.ReadFile()
	log.Println("SI PASA EL READFILE")
	return pr
}

func (pr *ProductRepository) ReadFile() (err error) {
	data, err := os.ReadFile("../products.json")
	if err != nil {
		return errOpen
	}

	err = json.Unmarshal(data, &pr.products.Products)
	if err != nil {
		return ErrParser
	}

	return err
}

func (pr *ProductRepository) GetProducts() (products []domain.Product, err error) {
	log.Println("SI LO USA")
	return pr.products.Products, nil
}

func (pr *ProductRepository) GetProduct(id int) (product domain.Product, err error) {
	var foundProduct domain.Product
	for _, product := range pr.products.Products {
		if product.Id == id {
			foundProduct = product
			break
		}
	}

	if foundProduct.Id != id {
		return domain.Product{}, ErrNotFound
	}

	return foundProduct, nil
}

func (pr *ProductRepository) SearchProduct(query float64) (products []domain.Product, err error) {
	var filteredProduct []domain.Product
	for _, product := range pr.products.Products {
		if query != 0 && product.Price > query {
			filteredProduct = append(filteredProduct, product)
		}
	}
	return filteredProduct, nil
}

func (pr *ProductRepository) AddProduct(product domain.Product) (newProduct domain.Product, err error) {
	if existProduct(product.Name, pr.products.Products) {
		return domain.Product{}, ErrItemExist
	}

	if uniqueCodeValue(product.CodeValue, pr.products.Products) {
		return domain.Product{}, ErrCodeValueRepeat
	}

	expDate, err := parseDate(product.Expiration)
	if err != nil {
		return domain.Product{}, err
	}

	if validDate(expDate) {
		return domain.Product{}, ErrDateExp
	}

	lastID := pr.products.Products[len(pr.products.Products)-1].Id
	product.Id = lastID + 1

	pr.products.Products = append(pr.products.Products, product)

	return product, nil
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
