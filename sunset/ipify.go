package sunset

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
Example:
	'https://api.ipify.org?format=json'
	{"ip":"73.238.152.142"}
*/

// IPifyResponse represents an IPify API response
type IPifyResponse struct {
	IP string `json:"ip"`
}

// GetIP returns a response struct
func GetIP() (*IPifyResponse, error) {
	ipify := &IPifyResponse{}

	resp, err := http.Get("https://api.ipify.org?format=json")

	if err != nil {
		return ipify, err
	}

	if resp.StatusCode != 200 {
		return ipify, fmt.Errorf("Unexpected IPify Response Code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return ipify, err
	}

	err = json.Unmarshal(body, ipify)

	return ipify, err
}
