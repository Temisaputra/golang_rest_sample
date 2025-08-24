package router

import (
	"net/http"

	"github.com/Temisaputra/warOnk/delivery/handler"
	_ "github.com/Temisaputra/warOnk/docs" // wajib untuk register doc
	"github.com/Temisaputra/warOnk/infrastructure/middleware"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

type Handlers struct {
	ProductHandler *handler.ProductHandler
	Logger         *zap.Logger
}

// NewRouter bikin router dan register semua endpoint
func NewRouter(handlers *Handlers) http.Handler {
	router := mux.NewRouter()

	router.Use(middleware.LoggingMiddleware(handlers.Logger)) // <- inject logger

	// Swagger endpoint
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	api := router.PathPrefix("/api/war-onk").Subrouter()

	// Product endpoints
	api.HandleFunc("/products", handlers.ProductHandler.GetAllProduct).Methods("GET")
	api.HandleFunc("/product/{id}", handlers.ProductHandler.GetProductByID).Methods("GET")
	api.HandleFunc("/product-create", handlers.ProductHandler.CreateProduct).Methods("POST")
	api.HandleFunc("/product-update/{id}", handlers.ProductHandler.UpdateProduct).Methods("PUT")
	api.HandleFunc("/product-delete/{id}", handlers.ProductHandler.DeleteProduct).Methods("DELETE")

	// CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"POST", "GET", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Mode"},
		MaxAge:             60,
		AllowCredentials:   true,
		OptionsPassthrough: false,
	})

	return c.Handler(router)
}
