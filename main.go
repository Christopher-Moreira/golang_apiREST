package main

import (
	"github.com/Christopher-Moreira/golang_apiREST/config"
	"github.com/Christopher-Moreira/golang_apiREST/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config Initialization Error: %v", err)
		return
	}
	// Initialize Router
	router.Initialize()
}
