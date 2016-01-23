package geocoding

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	//	"strings"
)

var geolocationAPI = "http://maps.googleapis.com/maps/api/geocode/json"

type Geolocation struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat  float64 `json:"lat"`
				Long float64 `json:"lng"`
			}
		}
		FormattedAddress string `json:"formatted_address"`
	}
}

func FetchLatLong(address string) (Geolocation, error) {
	url := url.Values{}
	url.Set("address", address)

	reqUrl := fmt.Sprintf("%s?address=%s", geolocationAPI, url.Encode())

	var g Geolocation
	respBody, err := doGetRequest(reqUrl, "")
	if err != nil {
		return g, err
	}

	if err := json.Unmarshal(respBody, &g); err != nil {
		return g, fmt.Errorf("Unmarshal error: %v", err)
	}

	return g, nil
}

func doGetRequest(reqUrl string, query string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest error: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client.Do error: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll error: %v", err)
	}

	return body, nil
}
