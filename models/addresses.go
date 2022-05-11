package models

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"github.com/labstack/gommon/log"
)

// {addresses: [“Address A”, “Address B”, “Address C”, "Address D"]}
type WrappedAddresses struct {
	AddressesList []string `json:"addresses"`
}

func (wrapped WrappedAddresses) Unwrap() Addresses {
	var addresses = Addresses{}
	for _, value := range wrapped.AddressesList {
		var result, err = GeocodeAddress(value)

		if err != nil {
			continue
		}

		addresses.AddressesList = append(addresses.AddressesList, Address{Name: value, Lat: result.Results[0].Geometry.Location.Lat, Lng: result.Results[0].Geometry.Location.Lng})

	}

	return addresses
}

type Addresses struct {
	AddressesList []Address
}

type Address struct {
	Name string
	Lat  float64
	Lng  float64
}

func (addresses Addresses) CheckNearest() (Address, Address, float64) {
	minDistance := math.Inf(+1)
	var firstAddress Address
	var secondAddress Address

	for index_x, x := range addresses.AddressesList {
		for _, y := range addresses.AddressesList[index_x+1:] {
			computedDistance := Distance(x.Lat, x.Lng, y.Lat, y.Lng)
			if computedDistance < minDistance {
				firstAddress = x
				secondAddress = y
				minDistance = computedDistance
			}

		}

	}

	return firstAddress, secondAddress, minDistance
}

// Address utils

func GeocodeAddress(address string) (MapsResult, error) {
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

	var res = MapsResult{}
	json.NewDecoder(resp.Body).Decode(&res)

	if res.Status != "OK" {
		return res, fmt.Errorf(res.Status)
	}

	return res, nil

}

func Hsin(theta float64) float64 {
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
	h := Hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*Hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}
