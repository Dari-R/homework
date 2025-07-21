package api

import (
	"fmt"

	"main.go/config"
)
type ApiClient interface {
	PrintKey()
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
