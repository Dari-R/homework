package main

import (
	configs "3-validation-api/config"
	verify "3-validation-api/internal/email"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()

	// Подключаем verify-модуль
	verify.NewVerifyHandler(router, conf)

	server := http.Server{
		Addr:    ":8082",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8082")
	server.ListenAndServe()
}
