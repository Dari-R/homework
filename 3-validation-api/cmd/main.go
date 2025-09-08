package main

import (
	configs "3-validation-api/config"
	"3-validation-api/internal/email"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	email.NewEmailHandler(router, email.EmailHandlerDeps{Config: conf})
	server := http.Server{
		Addr:    ":8082",
		Handler: router,
	}
	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
