package file

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func ReadFile(fileName string) ([]byte, error) {
	if filepath.Ext(fileName) != ".json" {
		return nil, errors.New("FILE MUSST BE IN JSON FORMAT")
	}
	data, err := os.ReadFile(fileName)
	if err != nil  {
		return nil, err
	} 
	return data, nil
}

func WriteFile(data []byte, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		color.Red("Could not create file")
		return
	}
	defer file.Close()
	extention := filepath.Ext(fileName)
	if extention == ".json" {
		_, err = file.Write(data)
		if err != nil {
			color.Red("Could not write file")
		}
		color.Green("Done")
	}else{
		color.Red("File musst be in JSON format")
	}
}
