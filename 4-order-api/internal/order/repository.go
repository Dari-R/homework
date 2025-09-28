// internal/order/repository.go
package order

import (
	"4-order-api/pkg/res/db"

	"gorm.io/gorm"
)

type OrderRepository struct {
	Database *db.Db
}

func NewOrderRepository(database *db.Db) *OrderRepository {
	return &OrderRepository{
		Database: database,
	}
}

func (repo *OrderRepository) Create(order *Order) (*Order, error) {
	err := repo.Database.DB.Transaction(func(tx *gorm.DB) error {
		// Создаем заказ
		if err := tx.Create(order).Error; err != nil {
			return err
		}

		// Связываем продукты с заказом
		for _, product := range order.Products {
			orderProduct := OrderProduct{
				OrderID:   order.ID,
				ProductID: product.ID,
				Quantity:  1,
				Price:     100.0, // Примерная цена
			}
			if err := tx.Create(&orderProduct).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return order, nil
}

func (repo *OrderRepository) GetByID(id uint) (*Order, error) {
	var order Order
	result := repo.Database.DB.
		Preload("User").
		Preload("Products").
		First(&order, id)
	
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

func (repo *OrderRepository) GetByUserID(userID uint) ([]Order, error) {
	var orders []Order
	result := repo.Database.DB.
		Preload("Products").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&orders)
	
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

func (repo *OrderRepository) GetByUserPhone(phone string) ([]Order, error) {
	var orders []Order
	result := repo.Database.DB.
		Preload("Products").
		Joins("JOIN users ON users.id = orders.user_id").
		Where("users.phone = ?", phone).
		Order("orders.created_at DESC").
		Find(&orders)
	
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}