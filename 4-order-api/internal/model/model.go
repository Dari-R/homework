package model

import (
	"math/rand"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Age         int
	Description string
	Images      pq.StringArray `gorm:"type:text[]"`
	Hash        string         `json:"hash" gorm:"uniqueIndex"`
}

func NewProduct(name string, age int, description string, image pq.StringArray) *Product {
	product:= &Product{
		Name:        name,
		Age:         age,
		Description: description,
		Images:      pq.StringArray(image),
		Hash: StringRunes(6),
	}
	product.GenerateHash()
	return product
}
func (product *Product) GenerateHash() {
	product.Hash = StringRunes(6)
}


var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func StringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn((len(letterRunes)))]
	}
	return string(b)
}
