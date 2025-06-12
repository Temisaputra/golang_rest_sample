// @title       My HTTP API
// @version     1.0
// @description Ini dokumentasi Swagger untuk aplikasi HTTP
// @host localhost:8085
// @BasePath /api/war-onk

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
	"github.com/Temisaputra/warOnk/repository/products_repository"
	product_repository "github.com/Temisaputra/warOnk/repository/products_repository"
	"github.com/joho/godotenv"

	_ "github.com/Temisaputra/warOnk/docs" // Untuk swagger docs
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	cfg := config.Get()
	db := conn.InitPostgreSQL(cfg)
	db = db.Debug()

	// Jalankan migration
	if err := products_repository.Migrate(db); err != nil {
		log.Fatal("migration failed:", err)
	}

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
