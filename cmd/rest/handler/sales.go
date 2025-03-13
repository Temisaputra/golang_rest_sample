package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Temisaputra/warOnk/core/entity"
	"github.com/Temisaputra/warOnk/pkg/helper"
)

type salesUsecase interface {
	CreateSales(ctx context.Context, params *entity.SalesRequest) error
}

type SalesHandler struct {
	salesUsecase salesUsecase
}

func NewSalesHandler(salesUsecase salesUsecase) *SalesHandler {
	return &SalesHandler{
		salesUsecase: salesUsecase,
	}
}

func (s *SalesHandler) CreateSales(w http.ResponseWriter, r *http.Request) {
	var sales entity.SalesRequest
	var err error
	err = json.NewDecoder(r.Body).Decode(&sales)
	if err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	err = s.salesUsecase.CreateSales(r.Context(), &sales)
	if err != nil {
		helper.WriteResponse(w, err, nil)
		return
	}

	var response helper.Response

	response.Code = http.StatusOK
	response.Message = "success"

	helper.WriteResponse(w, nil, &response)
}
