package main

import (
	"fmt"
	"os"

	logger "github.com/tylerconlee/SlabAPI/log"

	"github.com/namsral/flag"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var (
	nrConfig string

	nrApp *newrelic.Application

	log = logger.Log
)

func main() {
	flag.StringVar(&nrConfig, "newrelic", "", "New Relic configuration key")
	flag.Parse()
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal()
	}
	if nrConfig != "" {
		appname := fmt.Sprintf("slabAPI %s", hostname)
		nrApp, err = newrelic.NewApplication(
			newrelic.ConfigAppName(appname),
			newrelic.ConfigLicense(nrConfig),
		)
		if err != nil {

		}
	}
	NewRouter()
}
