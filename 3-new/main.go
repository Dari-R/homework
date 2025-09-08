package main

import (
	"flag"
	"fmt"
	"os"

	"main.go/api"
	"main.go/config"
)

func main() {
	create := flag.Bool("create", false, "Create a new bin")
	update := flag.Bool("update", false, "Update a bin")
	deleteFlag := flag.Bool("delete", false, "Delete a bin")
	get := flag.Bool("get", false, "Get a bin")
	list := flag.Bool("list", false, "List all bins")

	file := flag.String("file", "", "Path to JSON file")
	id := flag.String("id", "", "Bin ID")
	name := flag.String("name", "", "Bin name")

	flag.Parse()
	fmt.Println(*name)
	cfg := config.NewConfig()
	apiClient := api.NewApi(cfg)
	switch {
	case *create:
		apiClient.Create(*file, *name)
	case *update:
		apiClient.Update(*id, *name)
	case *deleteFlag:
		apiClient.Delete(*id)
	case *get:
		apiClient.Get(*id)
	case *list:
		apiClient.List()
	default:
		fmt.Println("Ничего не выбрано. Используй --create, --update, --delete, --get или --list.")
		flag.Usage()
		os.Exit(1)
	}
}

// func promtData(s string) string {
// 	fmt.Print(s + ": ")
// 	var str string
// 	fmt.Scanln(&str)
// 	return str
// }



// func createBin(binList *bins.BinList, fileName string, storage storage.Storage) {
// 	id := promtData("Enter Id")
// 	private, err := strconv.ParseBool(promtData("Enter private (true/false)"))
// 	if err != nil {
// 		color.Red("Cannot convert to bool")
// 		return
// 	}
// 	name := promtData("Enter name")
// 	bin := bins.CreateBin(id, name, private)
// 	binList.Bins = append(binList.Bins, bin)
// 	err = storage.Save(*binList)
// 	if err != nil {
// 		color.Red("Ошибка при сохранении: %v", err)
// 	}
// }
