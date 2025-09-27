package main

import (
	"4-order-api/config"
	"4-order-api/pkg/res/db"
	"fmt"
	"net/http"
)

func main() {
	conf := config.LoadConfig()
	_ = db.NewDb(conf)
	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()

}
