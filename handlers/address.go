package handlers

import (
	"go-address-microservice/models"
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

	return c.JSON(http.StatusOK, "test")
}
