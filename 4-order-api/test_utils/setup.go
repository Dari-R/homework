// test_utils/setup.go
package test_utils

import (
	config "4-order-api/config"
	"4-order-api/internal/model"
	"4-order-api/internal/user"
	"4-order-api/pkg/res/db"
	"os"
)

type TestSuite struct {
	Config            *config.Config
	Database          *db.Db
	UserRepository    *user.UserRepository
	ProductRepository *model.ProductRepository
	TestUserPhone     string
	TestProductIDs    []uint
}

func SetupTestSuite() *TestSuite {
	// Загружаем тестовую конфигурацию
	testConfig := &config.Config{
		Db: config.DbConfig{
			Dsn: getTestDSN(),
		},
		Auth: config.AuthConfig{
			Secret: "test_secret_key",
		},
	}

	// Подключаемся к тестовой базе
	database := db.NewDb(testConfig)

	// Создаем репозитории
	userRepo := user.NewUserRepository(database)
	productRepo := model.NewProductRepository(database)

	return &TestSuite{
		Config:            testConfig,
		Database:          database,
		UserRepository:    userRepo,
		ProductRepository: productRepo,
		TestUserPhone:     "79990001122",
	}
}

func (suite *TestSuite) Cleanup() {
	// Очищаем базу данных после тестов
	suite.Database.DB.Exec("DELETE FROM order_products")
	suite.Database.DB.Exec("DELETE FROM orders")
	suite.Database.DB.Exec("DELETE FROM products")
	suite.Database.DB.Exec("DELETE FROM users")
	suite.Database.DB.Exec("DELETE FROM sessions")
}

func (suite *TestSuite) CreateTestUser() (*user.User, error) {
	testUser := &user.User{
		Phone: suite.TestUserPhone,
	}
	return suite.UserRepository.Create(testUser)
}

func (suite *TestSuite) CreateTestProducts() ([]uint, error) {
	products := []model.Product{
		{
			Name:        "Test Product 1",
			Age:         0,
			Description: "Test Description 1",
			Hash:        "testhash1",
		},
		{
			Name:        "Test Product 2",
			Age:         0,
			Description: "Test Description 2",
			Hash:        "testhash2",
		},
	}

	var productIDs []uint
	for _, product := range products {
		createdProduct, err := suite.ProductRepository.Create(&product)
		if err != nil {
			return nil, err
		}
		productIDs = append(productIDs, createdProduct.ID)
	}

	suite.TestProductIDs = productIDs
	return productIDs, nil
}

func getTestDSN() string {
	if dsn := os.Getenv("TEST_DSN"); dsn != "" {
		return dsn
	}
	return "host=localhost user=postgres password=postgres dbname=order_api_test port=5432 sslmode=disable"
}
