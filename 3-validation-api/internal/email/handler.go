package verify

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"

	configs "3-validation-api/config"
	"3-validation-api/pkg/res"

	"github.com/jordan-wright/email"
)

type VerifyHandler struct {
	*configs.Config
	storageFile string
}

type Storage struct {
	Hashes map[string]string `json:"hashes"`
}

// Загрузка данных из файла
func (handler *VerifyHandler) loadStorage() (*Storage, error) {
	data, err := os.ReadFile(handler.storageFile)
	if os.IsNotExist(err) {
		return &Storage{Hashes: make(map[string]string)}, nil
	}
	if err != nil {
		return nil, err
	}

	var storage Storage
	if err := json.Unmarshal(data, &storage); err != nil {
		return nil, err
	}
	if storage.Hashes == nil {
		storage.Hashes = make(map[string]string)
	}
	return &storage, nil
}

// Сохранение данных в файл
func (handler *VerifyHandler) saveStorage(storage *Storage) error {
	data, err := json.MarshalIndent(storage, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(handler.storageFile, data, 0644)
}

// Конструктор VerifyHandler
func NewVerifyHandler(router *http.ServeMux, cfg *configs.Config) {
	handler := &VerifyHandler{
		Config:      cfg,
		storageFile: "storage.json",
	}

	// Регистрация маршрутов
	router.HandleFunc("/send", handler.Send())
	router.HandleFunc("/verify/", handler.Verify()) // для hash
}

type Request struct {
	Email string `json:"email"`
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		var body Request
		if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		to := body.Email
		fmt.Println(to)

		if len(to) == 0 {
			http.Error(w, "Email is required", http.StatusBadRequest)
			return
		}

		defer req.Body.Close()

		hash := res.GenerateRandomHash(16)

		storage, err := handler.loadStorage()
		if err != nil {
			http.Error(w, "Storage error", http.StatusInternalServerError)
			return
		}

		storage.Hashes[hash] = to
		if err := handler.saveStorage(storage); err != nil {
			http.Error(w, "Save error", http.StatusInternalServerError)
			return
		}

		verifyLink := fmt.Sprintf("http://localhost:8082/verify/%s", hash)

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

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet{
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		hash := strings.TrimPrefix(req.URL.Path, "/verify/")
		if hash == "" {
			http.Error(w, "Hash not provided", http.StatusBadRequest)
			return
		}

        // Загружаем storage
        storage, err := handler.loadStorage()
        if err != nil {
            http.Error(w, "Storage error", http.StatusInternalServerError)
            return
        }

        email, exists := storage.Hashes[hash]
        verified := exists && email != ""
        if exists {
            delete(storage.Hashes, hash)
            handler.saveStorage(storage) 
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]bool{"verified": verified})
    }
}
