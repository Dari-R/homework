package verify

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
	configs "3-validation-api/config"
)

type VerifyHandler struct {
	*configs.Config
	hashes map[string]bool
}

// Конструктор VerifyHandler
func NewVerifyHandler(router *http.ServeMux, cfg *configs.Config) {
	handler := &VerifyHandler{
		Config: cfg,
		hashes: make(map[string]bool),
	}

	// Регистрация маршрутов
	router.HandleFunc("/send", handler.Send())
	router.HandleFunc("/verify/", handler.Verify()) // для hash
}

// POST /send
func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		// Для демонстрации email можно захардкодить или брать из body запроса
		to := "user@example.com"

		// Генерация hash
		hash := "abc123"
		handler.hashes[hash] = false

		verifyLink := fmt.Sprintf("http://localhost:8080/verify/%s", hash)

		e := email.NewEmail()
		e.From = fmt.Sprintf("App <%s>", handler.Email)
		e.To = []string{to}
		e.Subject = "Подтверждение Email"
		e.Text = []byte("Нажмите для подтверждения: " + verifyLink)
		e.HTML = []byte(fmt.Sprintf("<h1>Подтвердите email: <a href='%s'>%s</a></h1>", verifyLink, verifyLink))

		auth := smtp.PlainAuth("", handler.Email, handler.Password, strings.Split(handler.Address, ":")[0])

		if err := e.Send(handler.Address, auth); err != nil {
			log.Printf("Ошибка при отправке email: %v", err)
			http.Error(w, "Не удалось отправить письмо", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Письмо с подтверждением отправлено"))
	}
}

// GET /verify/{hash}
func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		hash := strings.TrimPrefix(req.URL.Path, "/verify/")
		if hash == "" {
			http.Error(w, "Hash not provided", http.StatusBadRequest)
			return
		}

		_, ok := handler.hashes[hash]
		if !ok {
			http.Error(w, "Invalid verification link", http.StatusBadRequest)
			return
		}

		handler.hashes[hash] = true
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Email подтверждён с hash: %s", hash)))
	}
}
