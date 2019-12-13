package main

import (
	"github.com/namsral/flag"
)

var (
	nrConfig string
)

func main() {
	flag.StringVar(&nrConfig, "newrelic", "", "New Relic configuration key")
	flag.Parse()
	NewRouter()
}
