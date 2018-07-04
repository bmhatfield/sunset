package sunset

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Coordinates represent geographic coordinates
type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

/*
Example:
	https://api.sunrise-sunset.org/json?lat=36.7201600&lng=-4.4203400&formatted=0
	{
		"results": {
			"sunrise":"2017-11-24T07:05:28+00:00",
			"sunset":"2017-11-24T17:03:40+00:00",
			"solar_noon":"2017-11-24T12:04:34+00:00",
			"day_length":35892,
			"civil_twilight_begin":"2017-11-24T06:37:28+00:00",
			"civil_twilight_end":"2017-11-24T17:31:41+00:00",
			"nautical_twilight_begin":"2017-11-24T06:05:44+00:00",
			"nautical_twilight_end":"2017-11-24T18:03:24+00:00",
			"astronomical_twilight_begin":"2017-11-24T05:34:45+00:00",
			"astronomical_twilight_end":"2017-11-24T18:34:24+00:00"
		},
		"status": "OK"
	}

*/

// SunriseSunsetResults represents the results struct from sunrise-sunset.org
type SunriseSunsetResults struct {
	Sunrise                   time.Time `json:"sunrise"`
	Sunset                    time.Time `json:"sunset"`
	SolarNoon                 time.Time `json:"solar_noon"`
	DayLength                 int       `json:"day_length"`
	CivilTwilightBegin        time.Time `json:"civil_twilight_begin"`
	CivilTwilightEnd          time.Time `json:"civil_twilight_end"`
	NauticalTwilightBegin     time.Time `json:"nautical_twilight_begin"`
	NauticalTwilightEnd       time.Time `json:"nautical_twilight_end"`
	AstronomicalTwilightBegin time.Time `json:"astronomical_twilight_begin"`
	AstronomicalTwilightEnd   time.Time `json:"astronomical_twilight_end"`
}

// SunriseSunsetResponseContainer represents the response container from sunrise-sunset.org
type SunriseSunsetResponseContainer struct {
	Results SunriseSunsetResults `json:"results"`
	Status  string               `json:"status"`
}

// GetSunset returns a response struct
func GetSunset(geo *Coordinates, date time.Time) (*SunriseSunsetResults, error) {
	results := &SunriseSunsetResults{}

	if date.IsZero() {
		date = time.Now()
	}

	formatString := "https://api.sunrise-sunset.org/json?formatted=0&lat=%f&lng=%f&date=%s"
	requestURI := fmt.Sprintf(formatString, geo.Latitude, geo.Longitude, date.Format("2006-01-02"))

	resp, err := http.Get(requestURI)

	if err != nil {
		return results, err
	}

	if resp.StatusCode != 200 {
		return results, fmt.Errorf("Unexpected Sunrise-Sunset Response Code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return results, err
	}

	container := &SunriseSunsetResponseContainer{}
	err = json.Unmarshal(body, container)

	if err != nil {
		return results, err
	}

	if container.Status != "OK" {
		return results, fmt.Errorf("Invalid Sunrise-Sunset Response Status: %s", container.Status)
	}

	return &container.Results, nil
}
