package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	userapp "GO-Crud/internal/modules/user/application"
	userinfra "GO-Crud/internal/modules/user/infrastructure"
	userhttp "GO-Crud/internal/modules/user/transport/http"
	"GO-Crud/internal/platform/db/postgres"
)

func main() {
	db, err := postgres.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepository := userinfra.NewUserRepository(db)

	userService := userapp.NewUserService(userRepository)

	userHandler := userhttp.NewHandler(userService)

	router := mux.NewRouter()

	userhttp.RegisterRoutes(router, userHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Server started on :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
