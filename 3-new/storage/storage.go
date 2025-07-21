package storage

import (
	"encoding/json"

	"main.go/bins"
	"main.go/file"
)
type Storage interface{
	Save(bins.BinList, string) error
	Load(string) (bins.BinList, error)
}
type storage1 struct{}
func NewStorage() Storage{
	return &storage1{}
}
func (s* storage1)Save(binList bins.BinList, fileName string) error {
	data, err := json.Marshal(binList)
	if err != nil {
		return err
	}
	file.WriteFile(data, fileName)
	return nil
}

func (s * storage1)Load(fileName string) (bins.BinList, error) {
	data, err := file.ReadFile(fileName)
	if err != nil {
		return bins.BinList{}, err
	}
	var list bins.BinList
	err = json.Unmarshal(data, &list)
	return list, err
}