package main

import (
	"tiddlybench-cli/internal/dispatch"
	logger "tiddlybench-cli/internal/logger"
)

func main() {
	log := logger.GetInstance()

	dispatch.Run(log)
}
