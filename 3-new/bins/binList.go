package bins

import (
	"encoding/json"

	"github.com/fatih/color"
	"main.go/file"
)

type BinList struct {
	Bins []Bin `json:"bins"`
}

func NewBinList(fileName string) (*BinList, error) {
	data, err := file.ReadFile(fileName)
	if err != nil {
		return &BinList{
			Bins: []Bin{},
		}, err
	}
	var binList BinList
	err = json.Unmarshal(data, &binList)
	if err != nil {
		color.Red("не удалось разобрать файл")
		return &BinList{
			Bins: []Bin{},
		}, err
	}
	return &binList, nil
}


func (binsList *BinList) ToBytes() ([]byte, error) {
	data, err := json.Marshal(binsList)
	if err != nil {
		return nil, err
	}
	return data, nil
}
