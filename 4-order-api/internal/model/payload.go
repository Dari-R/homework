package model

import "github.com/lib/pq"

type ProductCreateRequest struct {
	Name        string         `json:"name" validate:"required"`
	Age         int            `json:"age" validate:"required"`
	Description string         `json:"description"`
	Images      pq.StringArray `json:"image"`
}

type ProductUpdateRequest struct {
	Name        string         `json:"name" validate:"required,name"`
	Age         int            `json:"age" validate:"required,age"`
	Description string         `json:"description"`
	Images      pq.StringArray `json:"image"`
	Hash        string         `json:"hash"`
}
