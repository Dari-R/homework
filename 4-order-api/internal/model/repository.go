package model

import (
	"4-order-api/pkg/res/db"

	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	Database *db.Db
}

func NewProductRepository(database *db.Db) *ProductRepository {
	return &ProductRepository{
		Database: database,
	}
}

func (repo *ProductRepository) GetById(id uint) (*Product, error) {
	var product Product
	res := repo.Database.DB.First(&product, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &product, nil
}

func (repo *ProductRepository) Create(product *Product) (*Product, error) {
	res := repo.Database.DB.Create(product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (repo *ProductRepository) GetByHash(hash string) (*Product, error) {
	var product Product
	res := repo.Database.DB.First(&product, "hash = ?", hash)
	if res.Error != nil {
		return nil, res.Error
	}
	return &product, nil
}

func (repo *ProductRepository) Update(product *Product) (*Product, error) {
	res := repo.Database.DB.Clauses(clause.Returning{}).Updates(product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (repo *ProductRepository) Delete(id uint) error {
	res := repo.Database.DB.Delete(&Product{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
