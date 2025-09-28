// internal/order/handler.go
package order

import (
	"4-order-api/internal/model"
	"4-order-api/internal/user"
	"4-order-api/pkg/req"
	"4-order-api/pkg/res"
	"net/http"
	"strconv"
)

type OrderHandlerDeps struct {
	OrderRepository   *OrderRepository
	UserRepository    *user.UserRepository
	ProductRepository *model.ProductRepository
}

type OrderHandler struct {
	OrderRepository   *OrderRepository
	UserRepository    *user.UserRepository
	ProductRepository *model.ProductRepository
}

func NewOrderHandler(router *http.ServeMux, deps OrderHandlerDeps) {
	handler := &OrderHandler{
		OrderRepository:   deps.OrderRepository,
		UserRepository:    deps.UserRepository,
		ProductRepository: deps.ProductRepository,
	}

	router.HandleFunc("POST /order", handler.CreateOrder())
	router.HandleFunc("GET /order/{id}", handler.GetOrderByID())
	router.HandleFunc("GET /my-orders", handler.GetMyOrders())
}

type CreateOrderRequest struct {
	ProductIDs []uint `json:"product_ids" validate:"required,min=1"`
}

type CreateOrderResponse struct {
	OrderID uint    `json:"order_id"`
	Status  string  `json:"status"`
	Total   float64 `json:"total"`
	Message string  `json:"message"`
}

func (handler *OrderHandler) CreateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userPhone := r.Context().Value("userPhone")
		if userPhone == nil {
			http.Error(w, "Authentication required", http.StatusUnauthorized)
			return
		}

		payload, err := req.HandleBody[CreateOrderRequest](&w, r)
		if err != nil {
			return
		}

		user, err := handler.UserRepository.FindByPhone(userPhone.(string))
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		var products []model.Product
		for _, productID := range payload.ProductIDs {
			product, err := handler.ProductRepository.GetById(productID)
			if err != nil {
				http.Error(w, "Product not found: "+strconv.FormatUint(uint64(productID), 10), http.StatusNotFound)
				return
			}
			products = append(products, *product)
		}

		// Создаем заказ
		order := NewOrder(user.ID, products)
		createdOrder, err := handler.OrderRepository.Create(order)
		if err != nil {
			http.Error(w, "Failed to create order: "+err.Error(), http.StatusInternalServerError)
			return
		}

		response := CreateOrderResponse{
			OrderID: createdOrder.ID,
			Status:  createdOrder.Status,
			Total:   createdOrder.Total,
			Message: "Order created successfully",
		}

		res.Json(w, response, http.StatusCreated)
	}
}

func (handler *OrderHandler) GetOrderByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Получаем телефон пользователя из контекста
		userPhone := r.Context().Value("userPhone")
		if userPhone == nil {
			http.Error(w, "Authentication required", http.StatusUnauthorized)
			return
		}

		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, "Invalid order ID", http.StatusBadRequest)
			return
		}

		// Находим заказ
		order, err := handler.OrderRepository.GetByID(uint(id))
		if err != nil {
			http.Error(w, "Order not found", http.StatusNotFound)
			return
		}

		// Проверяем, что заказ принадлежит пользователю
		user, err := handler.UserRepository.FindByPhone(userPhone.(string))
		if err != nil || user.ID != order.UserID {
			http.Error(w, "Access denied", http.StatusForbidden)
			return
		}

		res.Json(w, order, http.StatusOK)
	}
}

func (handler *OrderHandler) GetMyOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Получаем телефон пользователя из контекста
		userPhone := r.Context().Value("userPhone")
		if userPhone == nil {
			http.Error(w, "Authentication required", http.StatusUnauthorized)
			return
		}

		// Находим заказы пользователя
		orders, err := handler.OrderRepository.GetByUserPhone(userPhone.(string))
		if err != nil {
			http.Error(w, "Failed to get orders", http.StatusInternalServerError)
			return
		}

		res.Json(w, orders, http.StatusOK)
	}
}
