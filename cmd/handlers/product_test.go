package handlers_test

import (
	"api_rest/cmd/routes"
	"api_rest/internal/domain"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func creatServerForProductsHandler() *gin.Engine {
	router := gin.Default()
	routerInit := routes.NewRouter(router)
	routerInit.SetProduct()

	return router
}

func Test_GetProducts(t *testing.T) {
	//Arrange
	server := creatServerForProductsHandler()
	request := httptest.NewRequest(http.MethodGet, "/products", nil)
	response := httptest.NewRecorder()

	//Act
	server.ServeHTTP(response, request)

	//Response Body
	body, err := io.ReadAll(response.Body)

	if err != nil {
		t.Fatal(err)
	}

	//Assert
	assert.Equal(t, http.StatusOK, response.Code)
	assert.True(t, len(body) > 0)
}

func Test_AddProduct(t *testing.T) {
	//Arrange
	server := creatServerForProductsHandler()
	product := domain.Product{
		Name:       "a",
		Quantity:   5,
		CodeValue:  "AA",
		Expiration: "11/05/2024",
		Price:      50.0,
	}

	productMarshal, _ := json.Marshal(product)

	log.Println(string(productMarshal))

	request := httptest.NewRequest(http.MethodPost, "/products/", bytes.NewBuffer(productMarshal))
	response := httptest.NewRecorder()
	// log.Println(response)

	//Act
	server.ServeHTTP(response, request)

	//Assert
	assert.Equal(t, http.StatusCreated, response.Code)
}
