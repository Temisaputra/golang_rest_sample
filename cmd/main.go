package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/Temisaputra/warOnk/cmd/rest"
	"github.com/Temisaputra/warOnk/cmd/rest/handler"
	"github.com/Temisaputra/warOnk/config"
	"github.com/Temisaputra/warOnk/core/module"
	conn "github.com/Temisaputra/warOnk/pkg/conn/postgre"
	product_repository "github.com/Temisaputra/warOnk/repository/products_repository"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	cfg := config.Get()
	db := conn.InitPostgreSQL(cfg)
	db = db.Debug()

	var (
		productRepo    = product_repository.New(cfg, db)
		productUsecase = module.NewProductUsecase(productRepo)
	)

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		ctx := context.Background()

		defer conn.DbClose(db)

		productHandler := handler.NewProductHandler(productUsecase)

		handlers := &rest.Handlers{
			ProductHandler: productHandler,
		}

		err := rest.Run(ctx, *cfg, handlers)
		if err != nil {
			log.Fatal(err)
		}

		wg.Done()
	}()

	wg.Wait()
}
