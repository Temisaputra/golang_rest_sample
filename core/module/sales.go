package module

import (
	"context"
	"fmt"

	"github.com/Temisaputra/warOnk/core/entity"
	"github.com/Temisaputra/warOnk/core/repository"
)

type salesUsecase struct {
	salesRepo repository.SalesRepository
}

func NewSalesUsecase(salesRepository repository.SalesRepository) *salesUsecase {
	return &salesUsecase{
		salesRepo: salesRepository,
	}
}

func (u *salesUsecase) CreateSales(ctx context.Context, sales *entity.SalesRequest) (err error) {
	tx, err := u.salesRepo.BeginTx(ctx, nil)
	if err != nil {
		errMsg := fmt.Errorf("[SalesUsecase.CreateSales] error begin tx: %w", err)
		return errMsg
	}

	u.salesRepo.UseTx(tx)

	headerID, err := u.salesRepo.CreateSalesHeader(ctx, sales.SalesHeader)
	if err != nil {
		tx.Rollback()
		errMsg := fmt.Errorf("[SalesUsecase.CreateSales] error create sales header: %w", err)
		return errMsg
	}

	if len(*sales.SalesDetail) == 0 {
		tx.Rollback()
		errMsg := fmt.Errorf("[SalesUsecase.CreateSales] sales detail is empty")
		return errMsg
	}

	for _, detail := range *sales.SalesDetail {
		detail.IdSalesHeader = headerID
		err := u.salesRepo.CreateSalesDetail(ctx, &detail)
		if err != nil {
			tx.Rollback()
			errMsg := fmt.Errorf("[SalesUsecase.CreateSales] error create sales detail: %w", err)
			return errMsg
		}
	}

	err = tx.Commit()
	if err != nil {
		errMsg := fmt.Errorf("[SalesUsecase.CreateSales] error commit tx: %w", err)
		return errMsg
	}

	return
}
