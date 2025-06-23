package module

import (
	"context"
	"errors"
	"testing"

	"github.com/Temisaputra/warOnk/core/dto"
	repositoryMock "github.com/Temisaputra/warOnk/shared/mock/repository"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/mock/gomock"
)

func TestGetAllProduct(t *testing.T) {
	Convey("Test get list product", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoMock := repositoryMock.NewMockProductRepository(ctrl)

		uc := NewProductUsecase(repoMock)

		var (
			ctx     = context.Background()
			errResp = errors.New("error")

			products = []*dto.ProductResponse{
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

			pagination = &dto.Pagination{
				Page:      1,
				PageSize:  10,
				Keyword:   "",
				OrderBy:   "",
				OrderType: "",
			}
			meta = dto.Meta{
				Page:      1,
				PageSize:  10,
				TotalData: 2,
				TotalPage: 1,
			}
		)

		Convey("resp err when get all product", func() {
			repoMock.EXPECT().GetAllProduct(ctx, pagination).Return(nil, dto.Meta{}, errResp).Times(1)

			resp, metaResp, err := uc.GetAllProduct(ctx, pagination)

			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
			So(metaResp, ShouldResemble, dto.Meta{})
		})

		Convey("resp err when product not found", func() {
			repoMock.EXPECT().GetAllProduct(ctx, pagination).Return(nil, dto.Meta{}, nil).Times(1)

			resp, metaResp, err := uc.GetAllProduct(ctx, pagination)

			So(err, ShouldBeNil)
			So(resp, ShouldBeNil)
			So(metaResp, ShouldResemble, dto.Meta{})
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
			repoMock.EXPECT().GetAllProduct(ctx, pagination).Return(nil, dto.Meta{}, errResp).Times(1)

			resp, metaResp, err := uc.GetAllProduct(ctx, pagination)

			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
			So(metaResp, ShouldResemble, dto.Meta{})
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
