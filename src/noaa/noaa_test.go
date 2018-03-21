package noaa

import (
	"../conditions"
	"../configuration"
	"net/http/httptest"
	"net/http"
	"fmt"
	"testing"
)

func TestURL(t *testing.T) {
	configuration := configuration.Configuration{ AirportCode: "AIRPORT_CODE" }
	expectedURL := "http://w1.weather.gov/xml/current_obs/AIRPORT_CODE.xml"

	url := URL(configuration)

	if(expectedURL != url) {
		t.Errorf(
			"Failed to form URL:\nexpected:\t'%s'\n but got:\t'%s'",
			expectedURL,
			url)
	}
}

func TestFetch(t *testing.T) {
	server := mockServer()
	defer server.Close()
	expectedConditions := conditions.Conditions{
		Location: "Pottstown, Pottstown Limerick Airport, PA",
		Temperature: "30.0 F (-1.1 C)"}
	conditions := Fetch(server.URL)

	if(conditions != expectedConditions) {
		t.Errorf(
			"Failed to Fetch:\nexpected:\t'%s'\nbut got:\t'%s'",
			expectedConditions,
			conditions)
	}
}

func mockServer() *httptest.Server {
	f := func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(writer, observationXML)

	}

	return httptest.NewServer(http.HandlerFunc(f))
}

var observationXML = `<?xml version="1.0" encoding="ISO-8859-1"?>
<?xml-stylesheet href="latest_ob.xsl" type="text/xsl"?>
<current_observation version="1.0"
	 xmlns:xsd="http://www.w3.org/2001/XMLSchema"
	 xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	 xsi:noNamespaceSchemaLocation="http://www.weather.gov/view/current_observation.xsd">
	<credit>NOAA's National Weather Service</credit>
	<credit_URL>http://weather.gov/</credit_URL>
	<image>
		<url>http://weather.gov/images/xml_logo.gif</url>
		<title>NOAA's National Weather Service</title>
		<link>http://weather.gov</link>
	</image>
	<suggested_pickup>15 minutes after the hour</suggested_pickup>
	<suggested_pickup_period>60</suggested_pickup_period>
	<location>Pottstown, Pottstown Limerick Airport, PA</location>
	<station_id>KPTW</station_id>
	<latitude>40.238</latitude>
	<longitude>-75.5549</longitude>
	<observation_time>Last Updated on Mar 20 2018, 1:54 pm EDT</observation_time>
				<observation_time_rfc822>Tue, 20 Mar 2018 13:54:00 -0400</observation_time_rfc822>
	<weather>Light Snow</weather>
	<temperature_string>30.0 F (-1.1 C)</temperature_string>
	<temp_f>30.0</temp_f>
	<temp_c>-1.1</temp_c>
	<relative_humidity>58</relative_humidity>
	<wind_string>Northeast at 11.5 MPH (10 KT)</wind_string>
	<wind_dir>Northeast</wind_dir>
	<wind_degrees>50</wind_degrees>
	<wind_mph>11.5</wind_mph>
	<wind_kt>10</wind_kt>
	<pressure_string>1010.5 mb</pressure_string>
	<pressure_mb>1010.5</pressure_mb>
	<pressure_in>29.84</pressure_in>
	<dewpoint_string>17.1 F (-8.3 C)</dewpoint_string>
	<dewpoint_f>17.1</dewpoint_f>
	<dewpoint_c>-8.3</dewpoint_c>
	<windchill_string>20 F (-7 C)</windchill_string>
				<windchill_f>20</windchill_f>
				<windchill_c>-7</windchill_c>
	<visibility_mi>4.00</visibility_mi>
	<icon_url_base>http://forecast.weather.gov/images/wtf/small/</icon_url_base>
	<two_day_history_url>http://www.weather.gov/data/obhistory/KPTW.html</two_day_history_url>
	<icon_url_name>sn.png</icon_url_name>
	<ob_url>http://www.weather.gov/data/METAR/KPTW.1.txt</ob_url>
	<disclaimer_url>http://weather.gov/disclaimer.html</disclaimer_url>
	<copyright_url>http://weather.gov/disclaimer.html</copyright_url>
	<privacy_policy_url>http://weather.gov/notice.html</privacy_policy_url>
</current_observation>
<`
