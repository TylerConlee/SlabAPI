package main

import (
	logger "github.com/tylerconlee/SlabAPI/log"
)

var (
	log = logger.Log
)

func main() {
	log.Info("Deskmate - Zendesk Assistant by Tyler Conlee")

	NewRouter()
}
