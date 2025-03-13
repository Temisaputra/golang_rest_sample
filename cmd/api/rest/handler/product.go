package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Temisaputra/warOnk/core/entity"
	"github.com/Temisaputra/warOnk/pkg/helper"
	"github.com/gorilla/mux"
)

type productUsecase interface {
	GetAllProduct(ctx context.Context, pagination *entity.Pagination) (produtcs []*entity.ProductResponse, meta entity.Meta, err error)
	GetProductByID(ctx context.Context, id int) (product *entity.ProductResponse, err error)
	CreateProduct(ctx context.Context, params *entity.ProductRequest) error
	UpdateProduct(ctx context.Context, params *entity.ProductRequest, id int) error
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

func (h *ProductHandler) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))

	params := &entity.Pagination{
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
	var params entity.ProductRequest
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

	var params entity.ProductRequest
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
