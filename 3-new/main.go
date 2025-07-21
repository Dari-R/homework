package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
	"main.go/api"
	"main.go/bins"
	"main.go/config"
	"main.go/file"
	"main.go/storage"
)

func main() {
	cfg := config.NewConfig()
	apiClient := api.NewApi(cfg)
	storage := storage.NewStorage()
	runApp(storage, apiClient)
	fmt.Println("Read bins")
	fileName := promtData("Enter the file name")
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		initial := bins.BinList{Bins: []bins.Bin{}}
		data, _ := json.Marshal(initial)
		file.WriteFile(data, fileName)

	}

	binList, err := bins.NewBinList(fileName)
	if err != nil {
		color.Red("Ошибка при загрузке списка: %v", err)
		return
	}

	fmt.Println("1 - Create bin\n2 - Show bins")
	var menu int
	fmt.Scan(&menu)

	switch menu {
	case 1:
		createBin(binList, fileName, storage)
	case 2:
		data, err := file.ReadFile(fileName)
		if err != nil {
			color.Red("Ошибка при чтении файла: %v", err)
		} else {
			fmt.Println(string(data))
		}
	default:
		color.Red("Неверный выбор")
	}
}
func runApp(s storage.Storage, api api.ApiClient) {
	api.PrintKey()
}
func createBin(binList *bins.BinList, fileName string, storage storage.Storage) {
	id := promtData("Enter Id")
	private, err := strconv.ParseBool(promtData("Enter private (true/false)"))
	if err != nil {
		color.Red("Cannot convert to bool")
		return
	}
	name := promtData("Enter name")
	bin := bins.CreateBin(id, name, private)
	binList.Bins = append(binList.Bins, bin)
	err = storage.Save(*binList, fileName)
		if err != nil {
		color.Red("Ошибка при сохранении: %v", err)
	}
}

func promtData(s string) string {
	fmt.Print(s + ": ")
	var str string
	fmt.Scanln(&str)
	return str
}
