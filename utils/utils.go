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

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}
