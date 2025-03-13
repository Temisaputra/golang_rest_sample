package main

import (
	"context"
	"log"
	"sync"

	"github.com/Temisaputra/warOnk/cmd/api/rest"
	"github.com/Temisaputra/warOnk/cmd/api/rest/handler"
	"github.com/Temisaputra/warOnk/config"
	"github.com/Temisaputra/warOnk/core/module"
	conn "github.com/Temisaputra/warOnk/pkg/conn/postgre"
	category_repository "github.com/Temisaputra/warOnk/repository/category_repository"
	product_repository "github.com/Temisaputra/warOnk/repository/products_repository"
	sales_repository "github.com/Temisaputra/warOnk/repository/sales_repository"
	unit_repository "github.com/Temisaputra/warOnk/repository/unit_repository"
)

func main() {
	cfg := config.Get()
	db := conn.InitPostgreSQL(cfg)
	db = db.Debug()

	var (
		productRepo    = product_repository.New(cfg, db)
		categoryRepo   = category_repository.New(cfg, db)
		unitRepo       = unit_repository.New(cfg, db)
		productUsecase = module.NewProductUsecase(productRepo, categoryRepo, unitRepo)
		salesRepo      = sales_repository.New(cfg, db)
		salesUsecase   = module.NewSalesUsecase(salesRepo)
	)

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		ctx := context.Background()

		defer conn.DbClose(db)

		productHandler := handler.NewProductHandler(productUsecase)
		salesHandler := handler.NewSalesHandler(salesUsecase)

		handlers := &rest.Handlers{
			ProductHandler: productHandler,
			SalesHandler:   salesHandler,
		}

		err := rest.Run(ctx, *cfg, handlers)
		if err != nil {
			log.Fatal(err)
		}

		wg.Done()
	}()

	wg.Wait()
}
