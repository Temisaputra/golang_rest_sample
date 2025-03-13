package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Temisaputra/warOnk/cmd/rest/handler"
	"github.com/Temisaputra/warOnk/config"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Handlers struct {
	ProductHandler *handler.ProductHandler
	SalesHandler   *handler.SalesHandler
}

func Run(ctx context.Context, cfg config.Config, handlers *Handlers) error {
	router := mux.NewRouter()

	fmt.Printf("App Name: %s\n", cfg.AppName)

	// middleware := middleware.NewMiddleware(cfg)

	app := router.PathPrefix("/api/war-onk").Subrouter()
	// app.Use(middleware.Authorization)

	//Product
	app.HandleFunc("/products", handlers.ProductHandler.GetAllProduct).Methods("GET")
	app.HandleFunc("/product/{id}", handlers.ProductHandler.GetProductByID).Methods("GET")
	app.HandleFunc("/product-create", handlers.ProductHandler.CreateProduct).Methods("POST")
	app.HandleFunc("/product-update/{id}", handlers.ProductHandler.UpdateProduct).Methods("PUT")
	app.HandleFunc("/product-delete/{id}", handlers.ProductHandler.DeleteProduct).Methods("PUT")

	//Sales
	app.HandleFunc("/sales/create", handlers.SalesHandler.CreateSales).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"POST", "GET", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Mode"},
		MaxAge:             60, // 1 minutes
		AllowCredentials:   true,
		OptionsPassthrough: false,
		Debug:              false,
	})

	httpHandler := c.Handler(router)

	err := startServer(ctx, httpHandler, cfg)
	if err != nil {
		return err
	}

	return nil
}

func startServer(ctx context.Context, httpHandler http.Handler, cfg config.Config) error {
	errChan := make(chan error)

	go func() {
		errChan <- startHTTP(ctx, httpHandler, cfg)
	}()

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func startHTTP(ctx context.Context, httpHandler http.Handler, cfg config.Config) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.HTTPPort),
		Handler: httpHandler,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("failed to start server: ", err)
		}
	}()

	log.Printf("%s is starting at port: %s", cfg.AppName, cfg.HTTPPort)
	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	<-interruption

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("failed to shutdown: %s", err)
		return err
	}

	log.Println("server is shutting down")
	return nil
}
