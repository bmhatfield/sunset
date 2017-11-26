package sunset

import (
	"log"
	"time"
)

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

// AutoUpdatingTime performs the coordinated requests to get SunriseSunsetResults
// on a regular schedule (once per day) and delivers the results via the result channel.
func AutoUpdatingTime() <-chan *SunriseSunsetResults {
	c := make(chan *SunriseSunsetResults, 1)

	go func() {
		for {
			results, err := Time()

			if err != nil {
				log.Println("Unable to refresh sunset time", err)
				time.Sleep(30 * time.Second)
				continue
			}

			select {
			case c <- results:
				log.Println("Updated Sunset Time Published", results.Sunset.Local())
			default:
				log.Println("Skipping publishing updated sunset time, channel full")
			}

			nextSolarNoon := results.SolarNoon.Local().Add(24 * time.Hour)
			untilNextSolarNoon := nextSolarNoon.Sub(time.Now())

			log.Printf("Will next update sunset time in %0.2f seconds\n", untilNextSolarNoon.Seconds())
			time.Sleep(untilNextSolarNoon)
		}
	}()

	return c
}
