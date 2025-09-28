// internal/order/payload.go
package order

type OrderResponse struct {
	ID        uint                  `json:"id"`
	Status    string                `json:"status"`
	Total     float64               `json:"total"`
	Products  []ProductOrderInfo    `json:"products"`
	CreatedAt string                `json:"created_at"`
}

type ProductOrderInfo struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Price       float64 `json:"price"`
}