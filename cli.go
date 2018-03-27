package main

import (
	"flag"
	"fmt"
	"os"
)

// Arguments a container for internal ark state read from the command
// line
type Arguments struct {
	AirportCode string
}

var airportCodeArgument string

func init() {
	flag.StringVar(
		&airportCodeArgument,
		"airport-code",
		"",
		"the four letter ICAO airport code")
}

// ParseInput creates translates user input supplied on the command
// line into state usable by the internals of ark
func ParseInput()(Arguments) {
	flag.Parse()
	if("" == airportCodeArgument) {
		flag.Usage()
		os.Exit(1)
	}
	return Arguments{ AirportCode: airportCodeArgument }
}

// FormatObservation creates a textual report of the supplied
// observation data
func FormatObservation(observation Observation)(string) {
	return fmt.Sprintf(
		"Conditions at %s:\n\tWeather: %s\n\tTemperature: %s\n\tLast Observation: %s\n",
		observation.Location,
		observation.Weather,
		observation.Temperature,
		observation.ObservedAt)
}

// PrintObservation writes a textual representation of an observation
// to output
func PrintObservation(observation Observation) {
	fmt.Println(FormatObservation(observation))
}
