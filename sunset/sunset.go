package sunset

import "time"

// Time performs the coordinated requests to get SunriseSunsetResults
func Time() (*SunriseSunsetResults, error) {
	empty := &SunriseSunsetResults{}

	ip, err := GetIP()

	if err != nil {
		return empty, err
	}

	geo, err := GetGeo(ip)

	if err != nil {
		return empty, err
	}

	sunsetTime, err := GetSunset(geo, time.Now())

	return sunsetTime, err
}
