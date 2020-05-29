package main

import (
	"tiddly-cli/internal/apicall"
	"tiddly-cli/internal/config"
	logger "tiddly-cli/internal/logger"
)

func main() {
	log := logger.GetInstance()
	log.Info("Running tiddly-cli")

	config := config.New("user", "pass", "url")
	apicall := apicall.New(log, config)
	s := apicall.GetAllTiddlers()

	log.Info(s)
}
