package main

import (
	"4-order-api/config"
	"4-order-api/internal/model"
	"4-order-api/middleware"
	"4-order-api/pkg/res/db"
	"fmt"
	"net/http"
)

func main() {
	conf := config.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	productRepository := model.NewProductRepository(db)
	model.NewProductHandler(router, model.ProductHandlerDeps{ProductRepository: productRepository})
	handlerWithMiddleware := middleware.LoggingMiddleware(router)
	server := http.Server{
		Addr:    ":8081",
		Handler: handlerWithMiddleware,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()

}
