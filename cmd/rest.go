package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Temisaputra/warOnk/delivery/handler"
	"github.com/Temisaputra/warOnk/infrastructure/config"
	repository "github.com/Temisaputra/warOnk/infrastructure/db"
	"github.com/Temisaputra/warOnk/infrastructure/router"
	usecase "github.com/Temisaputra/warOnk/internal/usecase"
	"github.com/spf13/cobra"
)

var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "Run REST API server",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Get()

		// Inisialisasi database
		db := InitPostgreSQL(cfg)
		defer DbClose(db)

		// Dependency injection
		productRepo := repository.NewProductRepo(db)
		transactionRepo := repository.NewTransactionRepo(db)
		productUC := usecase.NewProductUsecase(productRepo, transactionRepo)
		productHandler := handler.NewProductHandler(productUC)

		// Compose handlers
		handlers := &router.Handlers{
			ProductHandler: productHandler,
		}

		// Buat router
		router := router.NewRouter(handlers)

		// Buat HTTP server
		server := &http.Server{
			Addr:    fmt.Sprintf(":%s", cfg.HTTPPort),
			Handler: router,
		}

		// Jalankan server di goroutine supaya bisa graceful shutdown
		go func() {
			log.Printf("Server is starting on port %s", cfg.HTTPPort)
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Printf("Failed to start server: %v", err)
			}
		}()

		// Tangkap interrupt signal untuk shutdown
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Println("Shutting down server...")

		// Buat context dengan timeout untuk shutdown
		ctxShutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctxShutdown); err != nil {
			log.Printf("Server forced to shutdown: %v", err)
		}

		log.Println("Server exited properly")
	},
}

func init() {
	rootCmd.AddCommand(restCmd)
}
