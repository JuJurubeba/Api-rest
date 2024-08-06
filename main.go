package main

import (
	"fmt"

	router "github.com/JuJurubeba/Api-rest/Router"
	"github.com/JuJurubeba/Api-rest/config"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		fmt.Println(err)
		logger.Errorf("config initialization error: %v", err)
		return
	}

	//initialize router
	router.Initialize()

}
