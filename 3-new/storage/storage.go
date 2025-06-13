package storage

import (
	"encoding/json"

	"main.go/bins"
	"main.go/file"
)

func Save(binList bins.BinList, fileName string) error {
	data, err := json.Marshal(binList)
	if err != nil {
		return err
	}
	file.WriteFile(data, fileName)
	return nil
}

func Load(fileName string) (bins.BinList, error) {
	data, err := file.ReadFile(fileName)
	if err != nil {
		return bins.BinList{}, err
	}
	var list bins.BinList
	err = json.Unmarshal(data, &list)
	return list, err
}