// Package main provides an application to query weather conditions at
// Airports by supplying their ICAO airport code
//
//  Examples:
//
//    ./ark --help
//    Usage of ./ark:
//      -airport-code string
//					the four letter ICAO airport code
//
//    ./ark -airport-code KPTW
//    Weather Conditions:
//			Location: Pottstown, Pottstown Limerick Airport, PA
//			Temperature: 31.0 F (-0.6 C)
//
//  References:
//    - https://en.wikipedia.org/wiki/ICAO_airport_code
package main

import (
	"github.com/indiebrain/ark-go/cli"
	"github.com/indiebrain/ark-go/noaa"
)

func main() {
	arguments := cli.ParseInput()
	observationUrl := noaa.URL(arguments.AirportCode)
	cli.PrintObservation(noaa.Fetch(observationUrl))
}
