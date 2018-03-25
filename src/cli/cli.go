// Package cli provides the textual interface between ark and the the
// command line
package cli

import (
	"../noaa"
	"flag"
	"fmt"
)

type Arguments struct {
	AirportCode string
}

var airportCodeArgument string

func init() {
	flag.StringVar(&airportCodeArgument, "airport-code", "", "the four letter ICAO airport code")
}

// Creates the configuration object used by the internals of ark from
// the supplied command line arguments
func ParseInput()(Arguments) {
	flag.Parse()
	return Arguments{ AirportCode: airportCodeArgument }
}

// Creates a textual report of the supplied observation data
func FormatObservation(observation noaa.Observation)(string) {
	return fmt.Sprintf(
		"Weather Conditions:\n\tLocation: %s\n\tTemperature: %s",
		observation.Location,
		observation.Temperature)
}

func PrintObservation(observation noaa.Observation) {
	fmt.Println(FormatObservation(observation))
}
