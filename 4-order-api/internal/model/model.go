package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Age         int
	Description string
	Images      pq.StringArray `gorm:"type:text[]"`
}

func NewLink(name string, age int, description string, image pq.StringArray) *Product {
	return &Product{
		Name:        name,
		Age:         age,
		Description: description,
		Images:       pq.StringArray(image),
	}
}

// var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// func StringRunes(n int) string {
// 	b := make([]rune, n)
// 	for i := range b {
// 		b[i] = letterRunes[rand.Intn((len(letterRunes)))]
// 	}
// 	return string(b)
// }
