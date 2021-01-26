package main

import (
	logger "github.com/tylerconlee/SlabAPI/log"
)

var (
	log = logger.Log
)

func main() {
	log.Info("Slab - Zendesk Assistant by Tyler Conlee")

	NewRouter()
}
