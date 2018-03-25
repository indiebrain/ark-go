package cli

import(
	"../noaa"
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
	observation := noaa.Observation{
		Location: "Location",
		Temperature: "Temperature",
	}
	expected := fmt.Sprintf(
		"Weather Conditions:\n\tLocation: %s\n\tTemperature: %s",
		observation.Location,
		observation.Temperature)

	actual := FormatObservation(observation)

	if(expected != actual) {
		t.Errorf(
			"Failed to FormatObservation: expccted \n'%s'\n but got\n'%s'",
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
