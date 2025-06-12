package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Temisaputra/warOnk/core/dto"
	"github.com/Temisaputra/warOnk/pkg/helper"
	"github.com/gorilla/mux"
)

type productUsecase interface {
	GetAllProduct(ctx context.Context, pagination *dto.Pagination) (produtcs []*dto.ProductResponse, meta dto.Meta, err error)
	GetProductByID(ctx context.Context, id int) (product *dto.ProductResponse, err error)
	CreateProduct(ctx context.Context, params *dto.ProductRequest) error
	UpdateProduct(ctx context.Context, params *dto.ProductRequest, id int) error
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
// @Tags hello v1
// @Summary Get All Product
// @Description Get All Product
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param keyword query string false "Keyword for search"
// @Param order_by query string false "Order by field"
// @Param order_type query string false "Order type (asc/desc)"
// @Success 200 {object} helper.Response{data=[]dto.ProductResponse,meta=dto.Meta}
// @Failure 400 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /products [get]
func (h *ProductHandler) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))

	params := &dto.Pagination{
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

	response.Code = http.StatusOK
	response.Message = "success"
	response.Meta = &meta
	response.Data = data

	helper.WriteResponse(w, nil, &response)
}

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

	response.Code = http.StatusOK
	response.Message = "success"
	response.Data = data

	helper.WriteResponse(w, nil, &response)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var params dto.ProductRequest
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

	response.Code = http.StatusOK
	response.Message = "success"

	helper.WriteResponse(w, nil, &response)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)

	if idInt == 0 {
		helper.WriteResponse(w, helper.NewErrBadRequest("id is required"), nil)
		return
	}

	var params dto.ProductRequest
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

	response.Code = http.StatusOK
	response.Message = "success"

	helper.WriteResponse(w, nil, &response)
}

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

	response.Code = http.StatusOK
	response.Message = "success"

	helper.WriteResponse(w, nil, &response)
}
