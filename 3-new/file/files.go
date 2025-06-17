package file

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func ReadFile(fileName string) ([]byte, error) {
	if filepath.Ext(fileName) != ".json" {
		return nil, errors.New("FILE MUST BE IN JSON FORMAT")
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
	extension := filepath.Ext(fileName)
	if extension == ".json" {
		_, err = file.Write(data)
		if err != nil {
			color.Red("Could not write file")
		}
		color.Green("Done")
	}else{
		color.Red("File must be in JSON format")
	}
}
