package storage

import (
	"encoding/json"

	"main.go/bins"
	"main.go/file"
)

type storage1 struct {
    filename string
}

type Storage interface {
	Save(bins.BinList) error
	Load() (bins.BinList, error)
}

func NewStorage(fileName string) Storage {
	return &storage1{filename: fileName}
}

func (s *storage1) Save(binList bins.BinList) error {
	data, err := json.Marshal(binList)
	if err != nil {
		return err
	}
	file.WriteFile(data, s.filename)
	return nil
}

func (s *storage1) Load() (bins.BinList, error) {
	data, err := file.ReadFile(s.filename)
	if err != nil {
		return bins.BinList{}, err
	}
	var list bins.BinList
	err = json.Unmarshal(data, &list)
	return list, err
}
