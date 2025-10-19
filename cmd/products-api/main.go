package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	productHandler "github.com/GCrel/go-microservices-docker/internal/products/adapters/http"
	postgres "github.com/GCrel/go-microservices-docker/internal/products/adapters/repository"
	"github.com/GCrel/go-microservices-docker/internal/products/core/domain"
	"github.com/GCrel/go-microservices-docker/internal/products/core/service"
	"github.com/GCrel/go-microservices-docker/pkg/database"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	dbHost := os.Getenv("PRODUCTS_DB_HOST")
	dbPort := os.Getenv("PRODUCTS_DB_PORT")
	dbUser := os.Getenv("PRODUCTS_DB_USER")
	dbPass := os.Getenv("PRODUCTS_DB_PASSWORD")
	dbName := os.Getenv("PRODUCTS_DB_NAME")
	apiPort := os.Getenv("API_PORT")
	if apiPort == "" {
		apiPort = ":8080"
	}
	usersApiURL := os.Getenv("USERS_API_URL")
	if usersApiURL == "" {
		usersApiURL = "http://localhost:8080"
	}

	db := database.InitDB(dbUser, dbPass, dbName, dbHost, dbPort)

	log.Println("Ejecutando migraciones de la base de datos para productos...")
	if err := db.AutoMigrate(&domain.Product{}); err != nil {
		log.Fatalf("No se pudieron ejecutar las migraciones: %v", err)
	}
	log.Println("Migraciones de productos completadas.")

	productRepo := postgres.NewProductRepository(db)
	productService := service.NewProductService(productRepo, usersApiURL)
	handler := productHandler.NewProductHandler(productService)

	router := mux.NewRouter()
	router.HandleFunc("/products", handler.CreateProduct).Methods(http.MethodPost)
	router.HandleFunc("/products/{id}", handler.GetProduct).Methods(http.MethodGet)

	fmt.Printf("Servidor de productos escuchando en http://localhost%s\n", apiPort)
	if err := http.ListenAndServe(apiPort, router); err != nil {
		log.Fatalf("El servidor de productos falló al iniciar: %v", err)
	}
}
