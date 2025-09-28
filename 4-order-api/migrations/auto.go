package main

import (
	"4-order-api/internal/model"
	"4-order-api/internal/session"
	"4-order-api/internal/user"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())

	db.AutoMigrate(&model.Product{}, &user.User{}, &session.Session{})
}
