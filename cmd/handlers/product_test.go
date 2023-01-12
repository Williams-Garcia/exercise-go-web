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

func Test_GetProduct(t *testing.T) {
	//Arrange
	server := creatServerForProductsHandler()
	request := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	response := httptest.NewRecorder()

	//Act
	server.ServeHTTP(response, request)

	//Response Body
	body, err := io.ReadAll(response.Body)

	if err != nil {
		t.Fatal(err)
	}

	//Assert
	assert.Equal(t, http.StatusFound, response.Code)
	assert.True(t, len(body) > 0)
}

func Test_DeleteProduct(t *testing.T) {
	//Arrange
	server := creatServerForProductsHandler()
	request := httptest.NewRequest(http.MethodDelete, "/products/10", nil)
	response := httptest.NewRecorder()

	//Act
	server.ServeHTTP(response, request)

	//Response Body
	_, err := io.ReadAll(response.Body)

	if err != nil {
		t.Fatal(err)
	}

	//Assert
	assert.Equal(t, http.StatusNoContent, response.Code)
}

func Test_AddProduct(t *testing.T) {
	//Arrange
	server := creatServerForProductsHandler()
	product := domain.Product{
		Id:         500,
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
	assert.Equal(t, response.Header().Get("Content-Type"), "application/json; charset=utf-8")
	// assert.Equal(t, http.StatusCreated, response.Code)
}

func Test_GetProductNotFound(t *testing.T) {
	//Arrange
	server := creatServerForProductsHandler()
	request := httptest.NewRequest(http.MethodGet, "/products/1000", nil)
	response := httptest.NewRecorder()

	//Act
	server.ServeHTTP(response, request)

	//Response Body
	body, err := io.ReadAll(response.Body)

	if err != nil {
		t.Fatal(err)
	}

	//Assert
	assert.Equal(t, http.StatusNotFound, response.Code)
	assert.True(t, len(body) > 0)
}

func Test_GetProductFailedId(t *testing.T) {
	//Arrange
	server := creatServerForProductsHandler()
	request := httptest.NewRequest(http.MethodGet, "/products/AWQSD", nil)
	response := httptest.NewRecorder()

	//Act
	server.ServeHTTP(response, request)

	//Response Body
	body, err := io.ReadAll(response.Body)

	if err != nil {
		t.Fatal(err)
	}

	//Assert
	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.True(t, len(body) > 0)
}

func Test_GetProductFailedToken(t *testing.T) {
	//Arrange
	server := creatServerForProductsHandler()
	request := httptest.NewRequest(http.MethodGet, "/products/12", nil)
	// request.Header.Clone().Set("token", "1232142")
	response := httptest.NewRecorder()

	//Act
	server.ServeHTTP(response, request)

	//Response Body
	body, err := io.ReadAll(response.Body)

	if err != nil {
		t.Fatal(err)
	}

	//Assert
	assert.Equal(t, http.StatusUnauthorized, response.Code)
	assert.True(t, len(body) > 0)
}
