package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Temisaputra/warOnk/core/dto"
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
	mockProductRepo.On("GetAllProduct", mock.Anything, &dto.Pagination{
		Page:      1,
		PageSize:  10,
		Keyword:   "",
		OrderBy:   "",
		OrderType: "",
	}).Return(mockResponse, mockMeta, nil)

	// 5. Setup router dan request
	r := mux.NewRouter()
	r.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		resp, _, err := mockProductRepo.GetAllProduct(r.Context(), &dto.Pagination{
			Page:      1,
			PageSize:  10,
			Keyword:   "",
			OrderBy:   "",
			OrderType: "",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(resp)
	}).Methods("GET")

	req, _ := http.NewRequest("GET", "/products", nil)
	rr := httptest.NewRecorder()

	// 6. Jalankan handler-nya
	r.ServeHTTP(rr, req)

	// 7. Assert status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// 8. Decode response body
	var got []*dto.ProductResponse // ← ini array, karena response-nya slice
	err := json.Unmarshal(rr.Body.Bytes(), &got)
	assert.NoError(t, err)

	// 9. Assert isi response
	assert.Equal(t, mockResponse[0].ProductsId, got[0].ProductsId)
	assert.Equal(t, mockResponse[0].ProductName, got[0].ProductName)
	assert.Equal(t, mockResponse[0].SellingPrice, got[0].SellingPrice)
	assert.Equal(t, mockResponse[0].PurchasePrice, got[0].PurchasePrice)
	assert.Equal(t, mockResponse[0].ProductStock, got[0].ProductStock)

	mockProductRepo.AssertExpectations(t)
}
