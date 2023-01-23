package product

import (
	"api_rest/internal/domain"
	errorpkg "api_rest/pkg/error"
	"api_rest/pkg/store"
	"log"
)

type ProductRepository struct {
	store    store.Store
	products []domain.Product
}

var (
	errOpen            = &errorpkg.CustomError{Msg: "Error: no se puede abrir el archivo"}
	ErrParser          = &errorpkg.CustomError{Msg: "Error: no se puede parsear el objeto"}
	ErrItemExist       = &errorpkg.CustomError{Msg: "El objeto a registrar ya existe"}
	ErrCodeValueRepeat = &errorpkg.CustomError{Msg: "El codigo del producto a agregar debe ser unico"}
	ErrDateExp         = &errorpkg.CustomError{Msg: "La fecha de expiracion del producto no es valida"}
	ErrNotFound        = &errorpkg.CustomError{Msg: "Objeto no encontrado"}
)

func NewProductRepository(store store.Store) *ProductRepository {
	pr := &ProductRepository{
		store: store,
	}

	products, err := store.Get()

	if err != nil {
		log.Fatal("No se ha podido establecer conexion con la BD")

	}
	pr.products = products
	// pr.ReadFile()
	return pr
}

// func (pr *ProductRepository) ReadFile() (err error) {
// 	data, err := os.ReadFile("../products.json")
// 	if err != nil {
// 		return errOpen
// 	}

// 	err = json.Unmarshal(data, &pr.products)
// 	if err != nil {
// 		return ErrParser
// 	}

// 	return err
// }

func (pr *ProductRepository) GetProducts() (products []domain.Product, err error) {
	return pr.products, nil
}

func (pr *ProductRepository) GetProduct(id int) (product domain.Product, err error) {
	var foundProduct domain.Product
	for _, product := range pr.products {
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
	for _, product := range pr.products {
		if query != 0 && product.Price > query {
			filteredProduct = append(filteredProduct, product)
		}
	}
	return filteredProduct, nil
}

func (pr *ProductRepository) AddProduct(product domain.Product) (newProduct domain.Product, err error) {
	lastID := pr.products[len(pr.products)-1].Id
	product.Id = lastID + 1

	pr.products = append(pr.products, product)
	pr.store.Set(pr.products)

	return product, nil
}

func (pr *ProductRepository) UpdateProduct(id int, product domain.Product) (updatedProduct domain.Product, err error) {
	updated := false

	for index := range pr.products {
		if pr.products[index].Id == id {
			updatedProduct = product
			updatedProduct.Id = id
			pr.products[index] = updatedProduct
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, ErrNotFound
	}
	pr.store.Set(pr.products)

	return updatedProduct, nil
}

func (pr *ProductRepository) UpdatePatchProduct(id int, productPatch domain.ProductPatch) (updatedProduct domain.Product, err error) {
	updated := false

	for index := range pr.products {
		if pr.products[index].Id == id {
			if productPatch.Name != "" {
				pr.products[index].Name = productPatch.Name
			} else if productPatch.CodeValue != "" {
				pr.products[index].CodeValue = productPatch.CodeValue
			} else if productPatch.Expiration != "" {
				pr.products[index].Expiration = productPatch.Expiration
			} else if productPatch.Price != 0 {
				pr.products[index].Price = productPatch.Price
			} else if productPatch.Quantity != 0 {
				pr.products[index].Quantity = productPatch.Quantity
			} else if productPatch.IsPublished != false {
				pr.products[index].IsPublished = productPatch.IsPublished
			}
			updatedProduct = pr.products[index]
			updatedProduct.Id = id
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, ErrNotFound
	}
	pr.store.Set(pr.products)

	return updatedProduct, nil

}

func (pr *ProductRepository) DeleteProduct(id int) (err error) {
	deleted := false
	var index int

	for i := range pr.products {
		if pr.products[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return ErrNotFound
	}

	pr.products = append(pr.products[:index], pr.products[index+1:]...)

	pr.store.Set(pr.products)

	return nil
}

func (pr *ProductRepository) ExistProduct(pName string) bool {
	for _, p := range pr.products {
		if p.Name == pName {
			return true
		}
	}

	return false
}

func (pr *ProductRepository) UniqueCodeValue(pCodeValue string) bool {
	for _, p := range pr.products {
		if p.CodeValue == pCodeValue {
			return true
		}
	}

	return false
}
