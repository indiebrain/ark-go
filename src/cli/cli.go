// Package cli provides the textual interface between ark and the the
// command line
package cli

import (
	"flag"
)

type Configuration struct {
	AirportCode string
}

var airportCodeArgument string

func init() {
	flag.StringVar(&airportCodeArgument, "airport-code", "", "the four letter ICAO airport code")
}

// Creates the configuration object used by the internals of ark from the supplied command line arguments
func ParseInput()(Configuration) {
	flag.Parse()
	return Configuration{ AirportCode: airportCodeArgument }
}
