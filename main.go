package main

import (
	"github.com/EdersonSouza02/ApiRestGolang-Oppotunity.git/config"
	"github.com/EdersonSouza02/ApiRestGolang-Oppotunity.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize config
	err := config.Init()

	if err != nil {
		logger.Errorf("Config inicialization error:%v", err)
		return
	}
	// Initialize Router
	router.Inicialize()

}
