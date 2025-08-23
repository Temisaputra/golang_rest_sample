// @title       My HTTP API
// @version     1.0
// @description Ini dokumentasi Swagger untuk aplikasi HTTP
// @host localhost:8085
// @BasePath /api/war-onk

package main

import (
	"fmt"
	"os"

	"github.com/Temisaputra/warOnk/cmd"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	// Jalankan Cobra root command
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
