package email

import (
	"fmt"
	"net/http"
	"net/smtp"

	configs "3-validation-api/config"

	"github.com/jordan-wright/email"
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
		if req.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		e := email.NewEmail()
		e.From = "Jordan Wright <test@gmail.com>"
		e.To = []string{"test@example.com"}
		e.Bcc = []string{"test_bcc@example.com"}
		e.Cc = []string{"test_cc@example.com"}
		e.Subject = "Awesome Subject"
		e.Text = []byte("Text Body is, of course, supported!")
		e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")

		// Отправка письма через SMTP с проверкой ошибки
		err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "test@gmail.com", "password123", "smtp.gmail.com"))
		if err != nil {
			http.Error(w, "Failed to send email: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Email sent successfully"))
	}
}

func (handler EmailHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		// Получаем hash из URL: /verify/abc123
		hash := req.URL.Path[len("/verify/"):]
		if hash == "" {
			http.Error(w, "Hash not provided", http.StatusBadRequest)
			return
		}

		// В реальном проекте здесь проверка hash в базе или кеше
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Email verified with hash: %s", hash)))
	}
}
