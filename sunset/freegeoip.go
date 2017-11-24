package sunset

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
Example Good:
	curl https://freegeoip.net/json/73.238.152.142
		{
			"ip":"73.238.152.142",
			"country_code":"US",
			"country_name":"United States",
			"region_code":"MA",
			"region_name":"Massachusetts",
			"city":"Cambridge",
			"zip_code":"02139",
			"time_zone":"America/New_York",
			"latitude":42.3646,
			"longitude":-71.1028,
			"metro_code":506
		}

Example Bad:
	curl https://freegeoip.net/json/
		{
			"ip":"2601:184:4780:18be:ba27:ebff:fea5:253c",
			"country_code":"US",
			"country_name":"United States",
			"region_code":"",
			"region_name":"",
			"city":"",
			"zip_code":"",
			"time_zone":"",
			"latitude":37.751,
			"longitude":-97.822,
			"metro_code":0
		}
*/

// FreeGeoIPResponse respresents a freegeoip.net JSON API Response
type FreeGeoIPResponse struct {
	IP          string  `json:"IP"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"region_name"`
	City        string  `json:"city"`
	ZipCode     string  `json:"zip_code"`
	TimeZone    string  `json:"time_zone"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	MetroCode   int     `json:"metro_code"`
}

// GetGeo returns a response struct
func GetGeo(ipify *IPifyResponse) (*FreeGeoIPResponse, error) {
	geoip := &FreeGeoIPResponse{}

	resp, err := http.Get(fmt.Sprintf("https://freegeoip.net/json/%s", ipify.IP))

	if err != nil {
		return geoip, err
	}

	if resp.StatusCode != 200 {
		return geoip, fmt.Errorf("Unexpected FreeGeoIP Response Code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return geoip, err
	}

	err = json.Unmarshal(body, geoip)

	if err != nil {
		return geoip, err
	}

	if geoip.RegionCode == "" && geoip.RegionName == "" {
		return geoip, fmt.Errorf("Incomplete FreeGeoIP Response data: %+v", geoip)
	}

	return geoip, nil
}
