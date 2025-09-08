package email

import (
	"fmt"
	"net/http"

	configs "3-validation-api/config"
)

type EmailHandlerDeps struct {
	*configs.Config
}

type EmailHandler struct {
	*configs.Config
}

func NewEmailHandler(router *http.ServeMux, deps EmailHandlerDeps) {
	handler := &EmailHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler EmailHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Send")
	}
}

func (handler EmailHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Verify")
	}
}
