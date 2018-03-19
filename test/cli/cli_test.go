package cli

import(
	"testing"
	"os"
	"bytes"
	"../../src/cli"
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

func airportCodeFlag(code string)(string) {
	var buffer bytes.Buffer
	buffer.WriteString("-airport-code=")
	buffer.WriteString(code)
	return buffer.String()
}
