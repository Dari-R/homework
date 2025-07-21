package api

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/google/uuid"
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
	cfg *config.Config
}

func NewApi(cfg *config.Config) ApiClient {
	return &apiImpl{cfg: cfg}
}

func (a *apiImpl) PrintKey() {
	fmt.Println("API Key:", a.cfg.Key)
}

func (a *apiImpl) Create(filePath, name string) {
	s := storage.NewStorage("main.json")

	list, err := s.Load()
	if err != nil {
		color.Red("Ошибка при загрузке bins: %v", err)
	}
	id := uuid.New().String()
	bin := bins.CreateBin(id, name, false)

	list.Bins = append(list.Bins, bin)

	err = s.Save(list)
	if err != nil {
		color.Red("Ошибка при сохранении bins: %v", err)
	}

	color.Green("Bin создан: ID = %s", id)
}

func (a *apiImpl) Update(id, newName string) {
	s := storage.NewStorage("main.json")
	list, err := s.Load()
	if err != nil {
		color.Red("Ошибка при загрузке bins: %v", err)
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
		color.Red("Bin с ID = %s не найден", id)
		return
	}

	err = s.Save(list)
	if err != nil {
		color.Red("Ошибка при сохранении bins: %v", err)
		return
	}

	color.Green("Bin обновлён: ID = %s, новое имя = %s", id, newName)
}

func (a *apiImpl) Delete(id string) {
	s := storage.NewStorage("main.json")
	list, err := s.Load()
	if err != nil {
		color.Red("Ошибка при загрузке bins: %v", err)
		return
	}

	newBins := make([]bins.Bin, 0)
	found := false

	for _, bin := range list.Bins {
		if bin.Id == id {
			found = true
			continue
		}
		newBins = append(newBins, bin)
	}

	if !found {
		color.Red("Bin с ID = %s не найден", id)
		return
	}

	list.Bins = newBins
	err = s.Save(list)
	if err != nil {
		color.Red("Ошибка при сохранении bins: %v", err)
		return
	}

	color.Green("Bin удалён: ID = %s", id)
}

func (a *apiImpl) Get(id string) {
	s := storage.NewStorage("main.json")
	list, err := s.Load()
	if err != nil {
		color.Red("Ошибка при загрузке bins: %v", err)
		return
	}

	for _, bin := range list.Bins {
		if bin.Id == id {
			color.Green("Bin найден:")
			color.Cyan("ID: %s", bin.Id)
			color.Cyan("Name: %s", bin.Name)
			color.Cyan("Private: %v", bin.Private)
			color.Cyan("CreatedAt: %s", bin.CreatedAt.Format(time.RFC3339))
			return
		}
	}

	color.Red("Bin с ID = %s не найден", id)
}


func (a *apiImpl) List() {
	s := storage.NewStorage("main.json")
	list, err := s.Load()
	if err != nil {
		color.Red("Ошибка при загрузке bins: %v", err)
		return
	}

	if len(list.Bins) == 0 {
		color.Yellow("Список bin'ов пуст")
		return
	}

	color.Cyan("Список всех bin'ов:")
	for _, bin := range list.Bins {
		color.Green("ID: %s | Name: %s | Private: %v | CreatedAt: %s",
			bin.Id, bin.Name, bin.Private, bin.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}
