// tests/order_e2e_test.go
package tests

import (
	"4-order-api/internal/model"
	"4-order-api/internal/order"
	"4-order-api/internal/user"
	"4-order-api/pkg/jwt"
	"4-order-api/test_utils"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type OrderE2ETestSuite struct {
	suite.Suite
	suite *test_utils.TestSuite
}

func TestOrderE2ESuite(t *testing.T) {
	suite.Run(t, new(OrderE2ETestSuite))
}

func (s *OrderE2ETestSuite) SetupSuite() {
	s.suite = test_utils.SetupTestSuite()

	// Мигрируем тестовую базу
	s.migrateTestDatabase()
}

func (s *OrderE2ETestSuite) TearDownSuite() {
	if s.suite != nil {
		s.suite.Cleanup()
	}
}

func (s *OrderE2ETestSuite) SetupTest() {
	// Очищаем данные перед каждым тестом
	s.suite.Cleanup()
}

func (s *OrderE2ETestSuite) migrateTestDatabase() {
	// Автомиграция для тестов
	s.suite.Database.DB.AutoMigrate(
		&user.User{},
		&model.Product{},
		&order.Order{},
		&order.OrderProduct{},
	)
}

func (s *OrderE2ETestSuite) TestCreateOrder_E2E() {
	// Подготовка данных
	user, err := s.suite.CreateTestUser()
	assert.NoError(s.T(), err)

	productIDs, err := s.suite.CreateTestProducts()
	assert.NoError(s.T(), err)

	// Создаем тестовый сервер
	router := s.createTestRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	// Получаем JWT токен
	token := s.getTestToken(user.Phone)

	// Создаем заказ
	orderRequest := map[string]interface{}{
		"product_ids": productIDs,
	}

	requestBody, _ := json.Marshal(orderRequest)
	req, err := http.NewRequest("POST", server.URL+"/order", bytes.NewBuffer(requestBody))
	assert.NoError(s.T(), err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(s.T(), err)
	defer resp.Body.Close()

	// Проверяем ответ
	assert.Equal(s.T(), http.StatusCreated, resp.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(s.T(), err)

	assert.Contains(s.T(), response, "order_id")
	assert.Equal(s.T(), "pending", response["status"])
	assert.Equal(s.T(), "Order created successfully", response["message"])
}

func (s *OrderE2ETestSuite) TestGetMyOrders_E2E() {
	// Подготовка данных
	user, err := s.suite.CreateTestUser()
	assert.NoError(s.T(), err)

	productIDs, err := s.suite.CreateTestProducts()
	assert.NoError(s.T(), err)

	// Создаем тестовый сервер
	router := s.createTestRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	token := s.getTestToken(user.Phone)

	// Сначала создаем заказ
	orderRequest := map[string]interface{}{
		"product_ids": productIDs,
	}

	requestBody, _ := json.Marshal(orderRequest)
	req, err := http.NewRequest("POST", server.URL+"/order", bytes.NewBuffer(requestBody))
	assert.NoError(s.T(), err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(s.T(), err)
	resp.Body.Close()

	// Теперь получаем список заказов
	req, err = http.NewRequest("GET", server.URL+"/my-orders", nil)
	assert.NoError(s.T(), err)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err = client.Do(req)
	assert.NoError(s.T(), err)
	defer resp.Body.Close()

	assert.Equal(s.T(), http.StatusOK, resp.StatusCode)

	var orders []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&orders)
	assert.NoError(s.T(), err)

	assert.Greater(s.T(), len(orders), 0)
	assert.Equal(s.T(), "pending", orders[0]["status"])
}

func (s *OrderE2ETestSuite) createTestRouter() *http.ServeMux {
	router := http.NewServeMux()

	// Инициализируем хендлеры с тестовыми зависимостями
	orderRepository := order.NewOrderRepository(s.suite.Database)

	orderHandlerDeps := order.OrderHandlerDeps{
		OrderRepository:   orderRepository,
		UserRepository:    s.suite.UserRepository,
		ProductRepository: s.suite.ProductRepository,
	}

	order.NewOrderHandler(router, orderHandlerDeps)

	return router
}

func (s *OrderE2ETestSuite) getTestToken(phone string) string {
	// Создаем JWT токен для тестов
	jwtService := jwt.NewJWT(s.suite.Config.Auth.Secret)
	token, err := jwtService.Create(jwt.JWTData{Phone: phone})
	if err != nil {
		panic("Failed to create test token: " + err.Error())
	}
	return token
}
