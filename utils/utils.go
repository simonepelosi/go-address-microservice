// AIzaSyCPdQtZiqCo0mnK9cBSFwo8S3f9vGmNrUE
package utils

import (
	"encoding/json"
	"fmt"
	"go-address-microservice/models"
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

	fmt.Printf("%v", res)

	return res

}
