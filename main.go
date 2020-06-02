package main

import (
	"tiddly-cli/internal/dispatch"
	logger "tiddly-cli/internal/logger"
)

func main() {
	log := logger.GetInstance()

	dispatch.Dispatch(log)
}
