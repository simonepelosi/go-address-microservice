package handlers

import (
	"go-address-microservice/models"
	"go-address-microservice/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// {addresses: [“Address A”, “Address B”, “Address C”, "Address D"]}
func CheckAddresses(c echo.Context) (err error) {
	addresses := new(models.Addresses)

	if err = c.Bind(addresses); err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	log.Printf("addresses  to check: %s", addresses.AddressesList)

	utils.GeocodeAddress(addresses.AddressesList[0])
	return c.JSON(http.StatusOK, addresses)
}
