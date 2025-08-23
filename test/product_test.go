package module

import (
	"context"
	"errors"
	"testing"

	"github.com/Temisaputra/warOnk/delivery/presenter"
	"github.com/Temisaputra/warOnk/delivery/presenter/request"
	"github.com/Temisaputra/warOnk/delivery/presenter/response"
	productUsecase "github.com/Temisaputra/warOnk/internal/usecase"
	repositoryMock "github.com/Temisaputra/warOnk/shared/mock/repository"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/mock/gomock"
)

func TestGetAllProduct(t *testing.T) {
	Convey("Test get list product", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoMock := repositoryMock.NewMockProductRepository(ctrl)
		transactionMock := repositoryMock.NewMockTransactionRepository(ctrl)

		uc := productUsecase.NewProductUsecase(repoMock, transactionMock)

		var (
			ctx     = context.Background()
			errResp = errors.New("error")

			products = []*presenter.ProductResponse{
				{
					ProductsId:    1,
					ProductName:   "Test Product",
					SellingPrice:  100,
					PurchasePrice: 80,
					ProductStock:  10,
				},
				{
					ProductsId:    2,
					ProductName:   "Another Product",
					SellingPrice:  200,
					PurchasePrice: 150,
					ProductStock:  5,
				},
			}

			pagination = &request.Pagination{
				Page:      1,
				PageSize:  10,
				Keyword:   "",
				OrderBy:   "",
				OrderType: "",
			}
			meta = response.Meta{
				Page:      1,
				PageSize:  10,
				TotalData: 2,
				TotalPage: 1,
			}
		)

		Convey("resp err when get all product", func() {
			repoMock.EXPECT().GetAllProduct(ctx, pagination).Return(nil, response.Meta{}, errResp).Times(1)

			resp, metaResp, err := uc.GetAllProduct(ctx, pagination)

			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
			So(metaResp, ShouldResemble, response.Meta{})
		})

		Convey("resp err when product not found", func() {
			repoMock.EXPECT().GetAllProduct(ctx, pagination).Return(nil, response.Meta{}, nil).Times(1)

			resp, metaResp, err := uc.GetAllProduct(ctx, pagination)

			So(err, ShouldBeNil)
			So(resp, ShouldBeNil)
			So(metaResp, ShouldResemble, response.Meta{})
		})

		Convey("resp success when get all product", func() {
			repoMock.EXPECT().GetAllProduct(ctx, pagination).Return(products, meta, nil).Times(1)

			resp, metaResp, err := uc.GetAllProduct(ctx, pagination)

			So(err, ShouldBeNil)
			So(resp, ShouldHaveLength, 2)
			So(metaResp.TotalData, ShouldEqual, 2)
			So(metaResp.PageSize, ShouldEqual, 10)
			So(metaResp.Page, ShouldEqual, 1)
		})

		Convey("resp err when get all product with pagination", func() {
			pagination.Page = 0 // Invalid page number
			repoMock.EXPECT().GetAllProduct(ctx, pagination).Return(nil, response.Meta{}, errResp).Times(1)

			resp, metaResp, err := uc.GetAllProduct(ctx, pagination)

			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
			So(metaResp, ShouldResemble, response.Meta{})
		})

		Convey("resp success when get all product with pagination", func() {
			pagination.Page = 1 // Valid page number
			repoMock.EXPECT().GetAllProduct(ctx, pagination).Return(products, meta, nil).Times(1)

			resp, metaResp, err := uc.GetAllProduct(ctx, pagination)

			So(err, ShouldBeNil)
			So(resp, ShouldHaveLength, 2)
			So(metaResp.TotalData, ShouldEqual, 2)
			So(metaResp.PageSize, ShouldEqual, 10)
			So(metaResp.Page, ShouldEqual, 1)
		})
	})

}

func TestCreateProduct(t *testing.T) {
	Convey("Test CreateProduct usecase", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		productRepoMock := repositoryMock.NewMockProductRepository(ctrl)
		trxRepoMock := repositoryMock.NewMockTransactionRepository(ctrl)
		uc := productUsecase.NewProductUsecase(productRepoMock, trxRepoMock)

		ctx := context.Background()
		params := &presenter.ProductRequest{
			ProductName:   "New Product",
			SellingPrice:  100,
			PurchasePrice: 80,
			ProductStock:  10,
		}

		// ===========================================
		Convey("Success create product", func() {
			// Mock WithTransaction
			trxRepoMock.EXPECT().WithTransaction(ctx, gomock.Any()).DoAndReturn(
				func(ctx context.Context, fn func(ctx context.Context) error) error {
					return fn(ctx) // langsung eksekusi callback
				},
			).Times(1)

			productRepoMock.EXPECT().CreateProduct(ctx, params).Return(nil).Times(1)

			err := uc.CreateProduct(ctx, params)
			So(err, ShouldBeNil)
		})

		// ===========================================
		Convey("Error in CreateProduct inside transaction", func() {
			trxRepoMock.EXPECT().WithTransaction(ctx, gomock.Any()).DoAndReturn(
				func(ctx context.Context, fn func(ctx context.Context) error) error {
					return fn(ctx)
				},
			).Times(1)

			productRepoMock.EXPECT().CreateProduct(ctx, params).Return(errors.New("db error")).Times(1)

			err := uc.CreateProduct(ctx, params)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "db error")
		})

		// ===========================================
		Convey("Error in WithTransaction itself", func() {
			trxRepoMock.EXPECT().WithTransaction(ctx, gomock.Any()).Return(errors.New("transaction error")).Times(1)

			err := uc.CreateProduct(ctx, params)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "transaction error")
		})

	})
}
