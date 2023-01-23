package routes

import (
	"api_rest/cmd/handlers"
	"api_rest/internal/product"
	"api_rest/internal/product/impl"
	"api_rest/pkg/store"

	"github.com/gin-gonic/gin"
)

type Router struct {
	en *gin.Engine
}

func NewRouter(en *gin.Engine) *Router {
	return &Router{en: en}
}

func (r *Router) SetRoutes() {
	r.SetPing()
	r.SetProduct()
}

// website
func (r *Router) SetProduct() {
	// instances
	db := *store.NewStore("../products.json")
	repo := product.NewProductRepository(db)
	service := *impl.NewProductService(repo)
	serviceImpl := handlers.NewProductHandler(service)

	routerProduct := r.en.Group("/products")
	// productHandler.ProductService = *impl.NewRepository()
	// productHandler.ProductService.ReadFile()
	routerProduct.GET("", serviceImpl.GetProducts())
	routerProduct.GET("/:id", serviceImpl.GetProduct())
	routerProduct.GET("/search", serviceImpl.SearchProduct())
	routerProduct.POST("/", serviceImpl.AddProduct())
	routerProduct.PUT("/:id", serviceImpl.UpdateProduct())
	routerProduct.PATCH("/:id", serviceImpl.UpdatePatchProduct())
	routerProduct.DELETE("/:id", serviceImpl.DeleteProduct())
}

func (r *Router) SetPing() {
	r.en.GET("/ping", handlers.Ping)
}
