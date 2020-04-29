package main

import (
	"fmt"
	"os"

	"github.com/namsral/flag"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var (
	nrConfig string

	nrApp *newrelic.Application
)

func main() {
	flag.StringVar(&nrConfig, "newrelic", "", "New Relic configuration key")
	flag.Parse()
	hostname, err := os.Hostname()
	if err != nil {

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
