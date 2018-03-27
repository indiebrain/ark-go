package main

import(
	"bytes"
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"net/http"
	"os"
)

// Observation is the internal representation of a weather observation
// as described by the NOAA observation API
type Observation struct {
	Location string `xml:"location"`
	Temperature string `xml:"temperature_string"`
	Weather string `xml:"weather"`
	ObservedAt string `xml:"observation_time"`
}

const NOAA_API_BASE_URL = "http://w1.weather.gov/xml/current_obs/"

// Fetch obtains the observation data from the NOAA api for the given airport code
func Fetch(airportCode string)(Observation) {
	url := URL(airportCode)
	response, error := http.Get(url)
	if error != nil {
		fmt.Fprintf(os.Stderr, "Failed to Fetch '%s': %s", url, error)
		os.Exit(1)
	}
	return ParseObservation(response)
}

func URL(airportCode string)(string) {
	var buffer bytes.Buffer
	buffer.WriteString(NOAA_API_BASE_URL)
	buffer.WriteString(airportCode)
	buffer.WriteString(".xml")

	return buffer.String()
}

func ParseObservation(response *http.Response)(Observation) {
	defer response.Body.Close()

	return decode(response)
}


func decode(response *http.Response)(Observation) {
	observation := Observation{}
	decoder := xml.NewDecoder(response.Body)
	decoder.CharsetReader = charset.NewReaderLabel
	error := decoder.Decode(&observation)
	if error != nil {
		fmt.Fprintf(os.Stderr, "Failed to decode response '%s': %s", response.Body, error)
		os.Exit(1)
	}

	return observation
}
