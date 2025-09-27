package main

import (
	"4-order-api/config"
	"4-order-api/pkg/res/db"
	"4-order-api/internal/model"
	"fmt"
	"net/http"
)

func main() {
	conf := config.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	productRepository := model.NewProductRepository(db)
	model.NewProductHandler(router, model.ProductHandlerDeps{ProductRepository: productRepository})
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()

}
