package main

import "github.com/aN0mad/golog/log"

func main() {

	// Enable debug logging
	log.EnableDebug()

	// Do logging things
	log.Debug("Debug message")
	log.Info("Info message")
	log.Error("Error message")
	log.Fatal("Fatal message")
}
