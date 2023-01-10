package product

import (
	"api_rest/internal/domain"
	errorpkg "api_rest/pkg/error"
	"encoding/json"
	"os"
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
	lastID := pr.products.Products[len(pr.products.Products)-1].Id
	product.Id = lastID + 1

	pr.products.Products = append(pr.products.Products, product)

	return product, nil
}

func (pr *ProductRepository) UpdateProduct(id int, product domain.Product) (updatedProduct domain.Product, err error) {
	updated := false

	for index := range pr.products.Products {
		if pr.products.Products[index].Id == id {
			updatedProduct = product
			updatedProduct.Id = id
			pr.products.Products[index] = updatedProduct
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, ErrNotFound
	}

	return updatedProduct, nil
}

func (pr *ProductRepository) UpdatePatchProduct(id int, productPatch domain.ProductPatch) (updatedProduct domain.Product, err error) {
	updated := false

	for index := range pr.products.Products {
		if pr.products.Products[index].Id == id {
			if productPatch.Name != "" {
				pr.products.Products[index].Name = productPatch.Name
			} else if productPatch.CodeValue != "" {
				pr.products.Products[index].CodeValue = productPatch.CodeValue
			} else if productPatch.Expiration != "" {
				pr.products.Products[index].Expiration = productPatch.Expiration
			} else if productPatch.Price != 0 {
				pr.products.Products[index].Price = productPatch.Price
			} else if productPatch.Quantity != 0 {
				pr.products.Products[index].Quantity = productPatch.Quantity
			} else if productPatch.IsPublished != false {
				pr.products.Products[index].IsPublished = productPatch.IsPublished
			}
			updatedProduct = pr.products.Products[index]
			updatedProduct.Id = id
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, ErrNotFound
	}

	return updatedProduct, nil

}

func (pr *ProductRepository) DeleteProduct(id int) (err error) {
	deleted := false
	var index int

	for i := range pr.products.Products {
		if pr.products.Products[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return ErrNotFound
	}

	pr.products.Products = append(pr.products.Products[:index], pr.products.Products[index+1:]...)

	return nil
}

func (pr *ProductRepository) ExistProduct(pName string) bool {
	for _, p := range pr.products.Products {
		if p.Name == pName {
			return true
		}
	}

	return false
}

func (pr *ProductRepository) UniqueCodeValue(pCodeValue string) bool {
	for _, p := range pr.products.Products {
		if p.CodeValue == pCodeValue {
			return true
		}
	}

	return false
}
