// AIzaSyCPdQtZiqCo0mnK9cBSFwo8S3f9vGmNrUE
package utils

import (
	"encoding/json"
	"fmt"
	"go-address-microservice/models"
	"math"
	"net/http"

	"github.com/labstack/gommon/log"
)

func GeocodeAddress(address string) (res models.MapsResult) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://maps.googleapis.com/maps/api/geocode/json", nil)

	if err != nil {
		log.Error(err)
	}

	// if you appending to existing query this works fine
	params := req.URL.Query()
	params.Add("address", address)
	params.Add("key", "AIzaSyCPdQtZiqCo0mnK9cBSFwo8S3f9vGmNrUE")

	req.URL.RawQuery = params.Encode()

	fmt.Println(req.URL.String())

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Printf("%v\n", res.Results[0].Geometry.Location)

	return res

}

func distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}
