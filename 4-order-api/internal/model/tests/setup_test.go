// tests/setup_test.go
package tests

import (
	"4-order-api/internal/model"
	"4-order-api/internal/user"
	"4-order-api/test_utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataPreparation(t *testing.T) {
	suite := test_utils.SetupTestSuite()
	defer suite.Cleanup()

	// Тест создания пользователя
	user, err := suite.CreateTestUser()
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, suite.TestUserPhone, user.Phone)

	// Тест создания продуктов
	productIDs, err := suite.CreateTestProducts()
	assert.NoError(t, err)
	assert.Len(t, productIDs, 2)

	// Проверяем что продукты создались
	product1, err := suite.ProductRepository.GetById(productIDs[0])
	assert.NoError(t, err)
	assert.Equal(t, "Test Product 1", product1.Name)

	product2, err := suite.ProductRepository.GetById(productIDs[1])
	assert.NoError(t, err)
	assert.Equal(t, "Test Product 2", product2.Name)
}

func TestDatabaseCleanup(t *testing.T) {
	suite := test_utils.SetupTestSuite()

	// Создаем тестовые данные
	_, err := suite.CreateTestUser()
	assert.NoError(t, err)

	_, err = suite.CreateTestProducts()
	assert.NoError(t, err)

	// Проверяем что данные создались
	var userCount int64
	suite.Database.DB.Model(&user.User{}).Count(&userCount)
	assert.Equal(t, int64(1), userCount)

	var productCount int64
	suite.Database.DB.Model(&model.Product{}).Count(&productCount)
	assert.Equal(t, int64(2), productCount)

	// Очищаем
	suite.Cleanup()

	// Проверяем что данные удалились
	suite.Database.DB.Model(&user.User{}).Count(&userCount)
	assert.Equal(t, int64(0), userCount)

	suite.Database.DB.Model(&model.Product{}).Count(&productCount)
	assert.Equal(t, int64(0), productCount)
}
