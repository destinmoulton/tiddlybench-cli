package main

import (
	"tikli/internal/dispatch"
	logger "tikli/internal/logger"
)

func main() {
	log := logger.GetInstance()

	dispatch.Dispatch(log)
}
