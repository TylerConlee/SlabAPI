package main

import (
	"fmt"
	"os"

	logger "github.com/tylerconlee/SlabAPI/log"
	"go.uber.org/zap"

	"github.com/namsral/flag"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var (
	nrConfig string

	nrApp *newrelic.Application

	log = logger.Log
)

func main() {
	log.Info("Slab - Zendesk Assistant by Tyler Conlee")
	flag.StringVar(&nrConfig, "newrelic", "", "New Relic configuration key")
	flag.Parse()
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("Fatal error", zap.String("Error", err.Error()))
	}
	if nrConfig != "" {
		appname := fmt.Sprintf("slabAPI %s", hostname)
		nrApp, err = newrelic.NewApplication(
			newrelic.ConfigAppName(appname),
			newrelic.ConfigLicense(nrConfig),
		)
		if err != nil {
			log.Fatal("Fatal error", zap.String("Error", err.Error()))
		}
	}
	NewRouter()
}
