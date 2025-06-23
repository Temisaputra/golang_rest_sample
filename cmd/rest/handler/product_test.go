package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Temisaputra/warOnk/core/dto"
	"github.com/Temisaputra/warOnk/core/module"
	_ "github.com/Temisaputra/warOnk/core/module"
	products_repository_mock "github.com/Temisaputra/warOnk/repository/product_repository_mock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetProduct_Success(t *testing.T) {
	// 1. Buat mock
	mockProductRepo := new(products_repository_mock.MockProductRepository)
	// 2. Dummy response
	mockResponse := []*dto.ProductResponse{
		{
			ProductsId:    1,
			ProductName:   "Test Product",
			SellingPrice:  100.0,
			PurchasePrice: 80.0,
			ProductStock:  50,
		},
	}

	mockMeta := dto.Meta{
		Page:      1,
		PageSize:  10,
		TotalData: 1,
		TotalPage: 1,
	}

	// 3. Set expectation
	mockProductRepo.On("GetAllProduct", mock.Anything, mock.AnythingOfType("*dto.Pagination")).Return(mockResponse, mockMeta, nil)

	productUC := module.NewProductUsecase(mockProductRepo)
	handler := NewProductHandler(productUC)

	// Setup HTTP test
	req := httptest.NewRequest("GET", "/products?page=1&page_size=10", nil)
	rec := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/products", handler.GetAllProduct).Methods("GET")
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var body map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &body)
	assert.NoError(t, err)

	// Optional: assert content
	data := body["data"].([]interface{})[0].(map[string]interface{})
	assert.Equal(t, "Test Product", data["product_name"])

	mockProductRepo.AssertExpectations(t)
}
