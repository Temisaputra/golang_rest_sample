package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Temisaputra/warOnk/delivery/presenter"
	"github.com/Temisaputra/warOnk/delivery/presenter/request"
	"github.com/Temisaputra/warOnk/delivery/presenter/response"
	"github.com/Temisaputra/warOnk/pkg/helper"
	"github.com/gorilla/mux"
)

type productUsecase interface {
	GetAllProduct(ctx context.Context, pagination *request.Pagination) (produtcs []*presenter.ProductResponse, meta response.Meta, err error)
	GetProductByID(ctx context.Context, id int) (product *presenter.ProductResponse, err error)
	CreateProduct(ctx context.Context, params *presenter.ProductRequest) error
	UpdateProduct(ctx context.Context, params *presenter.ProductRequest, id int) error
	DeleteProduct(ctx context.Context, id int) error
}

type ProductHandler struct {
	productUsecase productUsecase
}

func NewProductHandler(productUsecase productUsecase) *ProductHandler {
	return &ProductHandler{
		productUsecase: productUsecase,
	}
}

// GetAllProduct godoc
// @Tags Product
// @Summary Get All Product
// @Description Get All Product
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param keyword query string false "Keyword for search"
// @Param order_by query string false "Order by field"
// @Param order_type query string false "Order type (asc/desc)"
// @Success 200 {object} helper.Response{data=[]presenter.ProductResponse,meta=response.Meta}
// @Failure 400 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /products [get]
func (h *ProductHandler) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))

	params := &request.Pagination{
		Keyword:   r.URL.Query().Get("keyword"),
		OrderBy:   r.URL.Query().Get("order_by"),
		OrderType: r.URL.Query().Get("order_type"),
		Page:      page,
		PageSize:  pageSize,
	}

	data, meta, err := h.productUsecase.GetAllProduct(r.Context(), params)
	if err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	var response helper.Response

	response.StatusCode = http.StatusOK
	response.Message = "success"
	response.Meta = &meta
	response.Data = data

	helper.WriteResponse(w, nil, &response)
}

// GetProductByID godoc
// @Tags Product
// @Summary Get Product by ID
// @Description Get Product by ID
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} helper.Response{data=presenter.ProductResponse}
// @Failure 400 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /product/{id} [get]
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)
	if idInt == 0 {
		helper.WriteResponse(w, helper.NewErrBadRequest("id is required"), nil)
		return
	}
	data, err := h.productUsecase.GetProductByID(r.Context(), idInt)
	if err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	var response helper.Response

	response.StatusCode = http.StatusOK
	response.Message = "success"
	response.Data = data

	helper.WriteResponse(w, nil, &response)
}

// CreateProduct godoc
// @Tags Product
// @Summary Create a new product
// @Description Create a new product
// @Accept json
// @Produce json
// @Param request body presenter.ProductRequest true "Product data"
// @Success 201 {object} helper.Response{data=presenter.ProductResponse}
// @Failure 400 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Security BearerAuth
// @Router /product-create [post]
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var params presenter.ProductRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.productUsecase.CreateProduct(r.Context(), &params)
	if err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	var response helper.Response

	response.StatusCode = http.StatusCreated
	response.Message = "success"

	helper.WriteResponse(w, nil, &response)
}

// UpdateProduct godoc
// @Tags Product
// @Summary Update a product
// @Description Update a product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param request body presenter.ProductRequest true "Product data"
// @Success 200 {object} helper.Response{data=presenter.ProductResponse}
// @Failure 400 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)

	if idInt == 0 {
		helper.WriteResponse(w, helper.NewErrBadRequest("id is required"), nil)
		return
	}

	var params presenter.ProductRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.productUsecase.UpdateProduct(r.Context(), &params, idInt)
	if err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	var response helper.Response

	response.StatusCode = http.StatusOK
	response.Message = "success"

	helper.WriteResponse(w, nil, &response)
}

// DeleteProduct godoc
// @Tags Product
// @Summary Delete a product
// @Description Delete a product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} helper.Response{data=presenter.ProductResponse}
// @Failure 400 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)

	if idInt == 0 {
		helper.WriteResponse(w, helper.NewErrBadRequest("id is required"), nil)
		return
	}

	err := h.productUsecase.DeleteProduct(r.Context(), idInt)
	if err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	var response helper.Response

	response.StatusCode = http.StatusOK
	response.Message = "success"

	helper.WriteResponse(w, nil, &response)
}
