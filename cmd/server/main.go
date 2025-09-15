package main

import (
	"fmt"
	"net/http"
	"taskmanagerapi/internal/handlers"
)

func main() {
	// Create an Instance and add Keys to Map routes
	router := handlers.NewRouter()
	router.Handle("POST", "/users/signup", handlers.HandleSignup)
	router.Handle("POST", "/tasks/create", handlers.HandleCreateTask)

	// Server listening on port 3000
	fmt.Println("Server listening on port 3000...")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
