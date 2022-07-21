package main

import (
	"github.com/bshadmehr-com/auth/app"
	"github.com/bshadmehr-com/libs/logger"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}
