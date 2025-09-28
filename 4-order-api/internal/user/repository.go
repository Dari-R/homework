package user

import (
	"4-order-api/pkg/res/db"
)

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	res := repo.Database.DB.Create(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (repo *UserRepository) FindByPhone(phone string) (*User, error) {
	var user User
	res := repo.Database.DB.First(&user, "phone = ?", phone)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
