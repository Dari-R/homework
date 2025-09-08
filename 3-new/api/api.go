package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"main.go/bins"
	"main.go/config"
	"main.go/storage"
)

type ApiClient interface {
	Create(filePath, name string)
	Update(id, filePath string)
	Delete(id string)
	Get(id string)
	List()
}

type apiImpl struct {
	cfg    *config.Config
	client *http.Client
}

func NewApi(cfg *config.Config) ApiClient {
	return &apiImpl{
		cfg: cfg,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (a *apiImpl) makeRequest(method, url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", a.cfg.Key)

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return a.client.Do(req)
}

func (a *apiImpl) Create(filePath, name string) {
	// Чтение файла
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		color.Red("Ошибка чтения файла: %v", err)
		return
	}

	// Валидация JSON
	var jsonData map[string]interface{}
	if err := json.Unmarshal(fileData, &jsonData); err != nil {
		color.Red("Ошибка валидации JSON: %v", err)
		return
	}

	// Создание запроса
	resp, err := a.makeRequest(
		"POST",
		"https://api.jsonbin.io/v3/b",
		bytes.NewBuffer(fileData),
		map[string]string{"X-Bin-Name": name},
	)
	if err != nil {
		color.Red("Ошибка создания bin: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(resp.Body)
		color.Red("Ошибка API (%d): %s", resp.StatusCode, string(bodyBytes))
		return
	}

	// Обработка ответа
	var result struct {
		Metadata struct {
			ID        string    `json:"id"`
			CreatedAt time.Time `json:"createdAt"`
		} `json:"metadata"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		color.Red("Ошибка декодирования ответа: %v", err)
		return
	}

	// Сохранение локально
	bin := bins.Bin{
		Id:        result.Metadata.ID,
		Name:      name,
		Private:   false,
		CreatedAt: result.Metadata.CreatedAt,
	}

	s := storage.NewStorage("main.json")
	list, _ := s.Load()
	list.Bins = append(list.Bins, bin)
	if err := s.Save(list); err != nil {
		color.Red("Ошибка сохранения: %v", err)
		return
	}

	color.Green("Bin успешно создан! ID: %s", result.Metadata.ID)
}

func (a *apiImpl) Update(id, newName string) {
	// Сначала получаем текущий bin
	resp, err := a.makeRequest(
		"GET",
		fmt.Sprintf("https://api.jsonbin.io/v3/b/%s", id),
		nil,
		nil,
	)
	if err != nil {
		color.Red("Ошибка получения bin: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		color.Red("Ошибка API (%d): %s", resp.StatusCode, string(bodyBytes))
		return
	}

	// Декодируем текущие данные
	var currentData struct {
		Record   map[string]interface{} `json:"record"`
		Metadata struct {
			ID string `json:"id"`
		} `json:"metadata"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&currentData); err != nil {
		color.Red("Ошибка декодирования: %v", err)
		return
	}

	// Обновляем только имя
	currentData.Record["name"] = newName

	// Подготавливаем данные для обновления
	body, err := json.Marshal(currentData.Record)
	if err != nil {
		color.Red("Ошибка сериализации: %v", err)
		return
	}

	// Отправляем запрос на обновление
	updateResp, err := a.makeRequest(
		"PUT",
		fmt.Sprintf("https://api.jsonbin.io/v3/b/%s", id),
		bytes.NewBuffer(body),
		map[string]string{"X-Bin-Name": newName},
	)
	if err != nil {
		color.Red("Ошибка обновления: %v", err)
		return
	}
	defer updateResp.Body.Close()

	if updateResp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(updateResp.Body)
		color.Red("Ошибка API (%d): %s", updateResp.StatusCode, string(bodyBytes))
		return
	}

	// Обновляем локальное хранилище
	s := storage.NewStorage("main.json")
	list, err := s.Load()
	if err != nil {
		color.Red("Ошибка загрузки хранилища: %v", err)
		return
	}

	updated := false
	for i, bin := range list.Bins {
		if bin.Id == id {
			list.Bins[i].Name = newName
			updated = true
			break
		}
	}

	if !updated {
		color.Yellow("Bin не найден в локальном хранилище, добавляем...")
		list.Bins = append(list.Bins, bins.Bin{
			Id:   id,
			Name: newName,
		})
	}

	if err := s.Save(list); err != nil {
		color.Red("Ошибка сохранения: %v", err)
		return
	}

	color.Green("Bin успешно обновлен! ID: %s, Новое имя: %s", id, newName)
}

func (a *apiImpl) Delete(id string) {
	// Удаление из API
	resp, err := a.makeRequest(
		"DELETE",
		fmt.Sprintf("https://api.jsonbin.io/v3/b/%s", id),
		nil,
		nil,
	)
	if err != nil {
		color.Red("Ошибка удаления bin: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		color.Red("Ошибка API (%d): %s", resp.StatusCode, string(bodyBytes))
		return
	}

	// Удаление из локального хранилища
	s := storage.NewStorage("main.json")
	list, err := s.Load()
	if err != nil {
		color.Red("Ошибка загрузки хранилища: %v", err)
		return
	}

	newBins := make([]bins.Bin, 0, len(list.Bins)-1)
	for _, bin := range list.Bins {
		if bin.Id != id {
			newBins = append(newBins, bin)
		}
	}

	list.Bins = newBins
	if err := s.Save(list); err != nil {
		color.Red("Ошибка сохранения: %v", err)
		return
	}

	color.Green("Bin успешно удален! ID: %s", id)
}

func (a *apiImpl) Get(id string) {
	resp, err := a.makeRequest(
		"GET",
		fmt.Sprintf("https://api.jsonbin.io/v3/b/%s", id),
		nil,
		nil,
	)
	if err != nil {
		color.Red("Ошибка получения bin: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		color.Red("Ошибка API (%d): %s", resp.StatusCode, string(bodyBytes))
		return
	}

	var result struct {
		Record   map[string]interface{} `json:"record"`
		Metadata struct {
			ID        string    `json:"id"`
			CreatedAt time.Time `json:"createdAt"`
		} `json:"metadata"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		color.Red("Ошибка декодирования ответа: %v", err)
		return
	}

	color.Green("\nИнформация о Bin:")
	color.Cyan("ID: %s", result.Metadata.ID)
	color.Cyan("Создан: %s", result.Metadata.CreatedAt.Format("2006-01-02 15:04:05"))
	color.Cyan("Содержимое:")
	prettyJSON, _ := json.MarshalIndent(result.Record, "", "  ")
	fmt.Println(string(prettyJSON))
}

func (a *apiImpl) List() {
	s := storage.NewStorage("main.json")
	list, err := s.Load()
	if err != nil {
		color.Red("Ошибка загрузки хранилища: %v", err)
		return
	}

	if len(list.Bins) == 0 {
		color.Yellow("Нет сохраненных bins")
		return
	}

	color.Green("\nСписок всех Bins:")
	for _, bin := range list.Bins {
		color.Cyan("ID: %s | Имя: %s | Создан: %s",
			bin.Id,
			bin.Name,
			bin.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}
