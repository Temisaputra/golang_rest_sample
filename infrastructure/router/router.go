package router

import (
	"net/http"

	"github.com/Temisaputra/warOnk/delivery/handler"
	_ "github.com/Temisaputra/warOnk/docs" // wajib untuk register doc
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handlers struct {
	ProductHandler *handler.ProductHandler
}

// NewRouter bikin router dan register semua endpoint
func NewRouter(handlers *Handlers) http.Handler {
	router := mux.NewRouter()

	api := router.PathPrefix("/api/war-onk").Subrouter()

	// Product endpoints
	api.HandleFunc("/products", handlers.ProductHandler.GetAllProduct).Methods("GET")
	api.HandleFunc("/product/{id}", handlers.ProductHandler.GetProductByID).Methods("GET")
	api.HandleFunc("/product-create", handlers.ProductHandler.CreateProduct).Methods("POST")
	api.HandleFunc("/product-update/{id}", handlers.ProductHandler.UpdateProduct).Methods("PUT")
	api.HandleFunc("/product-delete/{id}", handlers.ProductHandler.DeleteProduct).Methods("DELETE")

	// Swagger endpoint
	router.PathPrefix("/api/war-onk/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/api/war-onk/swagger/doc.json"),
	))

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
