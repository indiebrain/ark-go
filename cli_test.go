package main

import(
	"bytes"
	"os"
	"fmt"
	"testing"
)

const AirportCode = "KPTW"

func TestParseInput(t *testing.T) {
	os.Args = []string{ "ark", airportCodeFlag(AirportCode) }

	configuration := ParseInput()

	if(configuration.AirportCode != AirportCode) {
		t.Errorf(
			"Failed to parse airport code: expected %s, but got %s",
			AirportCode,
			configuration.AirportCode)
	}
}

func TestFormatObservation(t *testing.T) {
	observation := Observation{
		Location: "Location",
		Temperature: "Temperature",
		Weather: "Weather",
		ObservedAt: "ObservedAt"}
	expected := fmt.Sprintf("Conditions at %s:\n\tWeather: %s\n\tTemperature: %s\n\tLast Observation: %s\n",
		observation.Location,
		observation.Weather,
		observation.Temperature,
		observation.ObservedAt)

	actual := FormatObservation(observation)

	if(expected != actual) {
		t.Errorf(
			"Failed to FormatObservation: expected \n'%s'\n but got\n'%s'",
			expected,
			actual)
	}
}

func airportCodeFlag(code string)(string) {
	var buffer bytes.Buffer
	buffer.WriteString("-airport-code=")
	buffer.WriteString(code)
	return buffer.String()
}
