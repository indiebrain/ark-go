package cli

import(
	"../../src/cli"
	"../../src/conditions"
	"bytes"
	"os"
	"fmt"
	"testing"
)


const AirportCode = "KPTW"

func TestParseInput(t *testing.T) {
	os.Args = []string{ "ark", airportCodeFlag(AirportCode) }

	configuration := cli.ParseInput()

	if(configuration.AirportCode != AirportCode) {
		t.Errorf(
			"Failed to parse airport code: expected %s, but got %s",
			AirportCode,
			configuration.AirportCode)
	}
}

func TestFormatConditions(t *testing.T) {
	conditions := conditions.Conditions{
		Location: "Location",
		Temperature: "Temperature",
	}
	expected := fmt.Sprintf(
		"Weather Conditions:\n\tLocation: %s\n\tTemperature: %s",
		conditions.Location,
		conditions.Temperature)

	actual := cli.FormatConditions(conditions)

	if(expected != actual) {
		t.Errorf(
			"Failed to print conditions: expccted \n'%s'\n but got\n'%s'",
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
