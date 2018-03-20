// Package cli provides the textual interface between ark and the the
// command line
package cli

import (
	"../conditions"
	"../configuration"
	"flag"
	"fmt"
)

var airportCodeArgument string

func init() {
	flag.StringVar(&airportCodeArgument, "airport-code", "", "the four letter ICAO airport code")
}

// Creates the configuration object used by the internals of ark from
// the supplied command line arguments
func ParseInput()(configuration.Configuration) {
	flag.Parse()
	return configuration.Configuration{ AirportCode: airportCodeArgument }
}

// Creates a textual report of the supplied conditions
func FormatConditions(conditions conditions.Conditions)(string) {
	return fmt.Sprintf(
		"Weather Conditions:\n\tLocation: %s\n\tTemperature: %s",
		conditions.Location,
		conditions.Temperature)
}
