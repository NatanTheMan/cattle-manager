package main

import (
	"github.com/NatanTheMan/cattle-manager/config"
	"github.com/NatanTheMan/cattle-manager/router"
)
var (logger * config.Logger)

func main() {
  logger=config.GetLogger("main")
	err := config.Init()
	if err != nil {
    logger.Errorf("config initialization error: %v", err)
		return
	}
	router.Initialize()
}
