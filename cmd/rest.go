package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Temisaputra/warOnk/delivery/handler"
	repository "github.com/Temisaputra/warOnk/internal/infrastructure/db"
	"github.com/Temisaputra/warOnk/internal/infrastructure/router"
	usecase "github.com/Temisaputra/warOnk/internal/usecase"
	"github.com/Temisaputra/warOnk/pkg/auth"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "Run REST API server",
	Run: func(cmd *cobra.Command, args []string) {
		deps := InitDependencies() // 🔑 ambil dari bootstrap.go
		defer func() {
			sqlDB, _ := deps.DB.DB()
			sqlDB.Close()

			// flush zap buffer
			deps.Logger.Sync()
		}()

		// Inject dependency

		// Product
		productRepo := repository.NewProductRepo(deps.DB)
		transactionRepo := repository.NewTransactionRepo(deps.DB)
		productUC := usecase.NewProductUsecase(productRepo, transactionRepo)
		productHandler := handler.NewProductHandler(productUC)

		// User
		userRepo := repository.NewUserRepo(deps.DB)

		// JWT Service
		jwtService := auth.NewJwtService(*deps.Cfg, *deps.Logger, userRepo)

		// Auth
		authRepo := repository.NewAuthRepo(deps.DB)
		authUC := usecase.NewAuthUsecase(authRepo, userRepo, transactionRepo, jwtService)
		authHandler := handler.NewAuthHandler(authUC)

		handlers := &router.Handlers{
			ProductHandler: productHandler,
			AuthHandler:    authHandler,
			Logger:         deps.Logger,
			JwtService:     jwtService,
		}

		r := router.NewRouter(handlers)

		server := &http.Server{
			Addr:    fmt.Sprintf(":%s", deps.Cfg.HTTPPort),
			Handler: r,
		}

		// Run server
		go func() {
			deps.Logger.Info("Starting REST server", zap.String("port", deps.Cfg.HTTPPort))
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				deps.Logger.Fatal("Failed to start server", zap.Error(err))
			}
		}()

		// Graceful shutdown
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		deps.Logger.Info("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			deps.Logger.Fatal("Server forced to shutdown", zap.Error(err))
		}
		deps.Logger.Info("Server exited properly")
	},
}

func init() {
	rootCmd.AddCommand(restCmd)
}
