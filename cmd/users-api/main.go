package main

import (
	"os"
	"fmt"
	"log"
	"net/http"

	userHandler "github.com/GCrel/go-microservices-docker/internal/users/adapters/http"
	"github.com/GCrel/go-microservices-docker/internal/users/adapters/repository"
	"github.com/GCrel/go-microservices-docker/internal/users/core/domain"
	"github.com/GCrel/go-microservices-docker/internal/users/core/service"
	"github.com/GCrel/go-microservices-docker/pkg/database"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	dbHost := os.Getenv("USERS_DB_HOST")
    dbPort := os.Getenv("USERS_DB_PORT")
    dbUser := os.Getenv("USERS_DB_USER")
    dbPass := os.Getenv("USERS_DB_PASSWORD")
    dbName := os.Getenv("USERS_DB_NAME")
    apiPort := os.Getenv("API_PORT")
    if apiPort == "" {
    	apiPort = ":8080"
    }

	db := database.InitDB(dbUser, dbPass, dbName, dbHost, dbPort)
	log.Println("Ejecutando migraciones de la base de datos...")
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Fatalf("No se pudieron ejecutar las migraciones: %v", err)
	}
	log.Println("Migraciones completadas.")

	userRepo := postgres.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	handler := userHandler.NewUserHandler(userService)

	router := mux.NewRouter()
	router.HandleFunc("/users", handler.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", handler.GetUser).Methods(http.MethodGet)

	fmt.Printf("Servidor de usuarios escuchando en http://localhost%s\n", apiPort)
	if err := http.ListenAndServe(apiPort, router); err != nil {
		log.Fatalf("El servidor falló al iniciar: %v", err)
	}
}