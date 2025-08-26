package router

import (
	"net/http"

	"github.com/Temisaputra/warOnk/delivery/handler"
	"github.com/Temisaputra/warOnk/delivery/middleware"
	_ "github.com/Temisaputra/warOnk/docs" // wajib untuk register doc
	"github.com/Temisaputra/warOnk/pkg/auth"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

type Handlers struct {
	ProductHandler *handler.ProductHandler
	UserHandler    *handler.UserHandler
	Logger         *zap.Logger
	JwtService     auth.JwtService
}

// NewRouter bikin router dan register semua endpoint
func NewRouter(handlers *Handlers) http.Handler {
	router := mux.NewRouter()
	authMW := middleware.NewAuthMiddleware(handlers.JwtService)

	router.Use(middleware.LoggingMiddleware(handlers.Logger)) // <- inject logger

	// Swagger endpoint
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	api := router.PathPrefix("/api/war-onk").Subrouter()

	// ---------------- Protected ----------------
	protected := api.PathPrefix("").Subrouter()
	protected.Use(authMW.Authorization)
	// Product endpoints
	protected.HandleFunc("/products", handlers.ProductHandler.GetAllProduct).Methods("GET")
	protected.HandleFunc("/product/{id}", handlers.ProductHandler.GetProductByID).Methods("GET")
	protected.HandleFunc("/product-create", handlers.ProductHandler.CreateProduct).Methods("POST")
	protected.HandleFunc("/product-update/{id}", handlers.ProductHandler.UpdateProduct).Methods("PUT")
	protected.HandleFunc("/product-delete/{id}", handlers.ProductHandler.DeleteProduct).Methods("DELETE")

	// User endpoints
	protected.HandleFunc("/users", handlers.UserHandler.GetAllUsers).Methods("GET")
	protected.HandleFunc("/user/{id}", handlers.UserHandler.GetUserByID).Methods("GET")
	protected.HandleFunc("/user-create", handlers.UserHandler.CreateUser).Methods("POST")
	protected.HandleFunc("/user-update/{id}", handlers.UserHandler.UpdateUser).Methods("PUT")
	protected.HandleFunc("/user-delete/{id}", handlers.UserHandler.DeleteUser).Methods("DELETE")

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
