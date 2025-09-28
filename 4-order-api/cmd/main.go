// main.go
package main

import (
	"4-order-api/internal/auth"
	configs "4-order-api/config"
	"4-order-api/internal/model"
	"4-order-api/internal/session"
	"4-order-api/internal/user"
	"4-order-api/middleware"
	"4-order-api/pkg/res/db"
	"4-order-api/pkg/sms"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// Инициализация репозиториев
	userRepository := user.NewUserRepository(db)
	sessionRepository := session.NewSessionRepository(db)

	// Инициализация сервисов
	smsService := sms.NewSMSService()
	authService := auth.NewAuthService(userRepository, sessionRepository, smsService)

	// Регистрация хендлеров
	productRepository := model.NewProductRepository(db)
	model.NewProductHandler(router, model.ProductHandlerDeps{ProductRepository: productRepository})

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})

	// Middleware chain
	handlerWithLogging := middleware.LoggingMiddleware(router)
	handlerWithAuth := middleware.AuthMiddleware(handlerWithLogging, conf)

	server := http.Server{
		Addr:    ":8081",
		Handler: handlerWithAuth,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
