// internal/order/order.go
package order

import (
	"4-order-api/internal/model"
	"4-order-api/internal/user"
	"fmt"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID   uint            `gorm:"not null"`
	User     user.User       `gorm:"foreignKey:UserID"`
	Products []model.Product `gorm:"many2many:order_products;"`
	Status   string          `gorm:"default:'pending'"`
	Total    float64         `gorm:"type:decimal(10,2)"`
}

type OrderProduct struct {
	OrderID   uint    `gorm:"primaryKey"`
	ProductID uint    `gorm:"primaryKey"`
	Quantity  int     `gorm:"not null;default:1"`
	Price     float64 `gorm:"type:decimal(10,2)"`
}

func NewOrder(userID uint, products []model.Product) *Order {
	// Расчет общей суммы
	var total float64
	for _, product := range products {
		fmt.Println(product)
		total += 100.0 
	}

	return &Order{
		UserID:   userID,
		Products: products,
		Status:   "pending",
		Total:    total,
	}
}
