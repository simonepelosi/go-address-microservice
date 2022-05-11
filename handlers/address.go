package handlers

import (
	"go-address-microservice/models"

	"net/http"

	"github.com/labstack/echo/v4"
)

// {addresses: [“Address A”, “Address B”, “Address C”, "Address D"]}
func CheckAddresses(c echo.Context) (err error) {
	wrappedAddresses := new(models.WrappedAddresses)
	if err = c.Bind(wrappedAddresses); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error reading response body!")
	}

	addresses := wrappedAddresses.Unwrap()

	x, y, distance, err := addresses.CheckNearest()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Can't find two near points!")
	}
	response := models.NewResponse(x, y, distance)

	return c.JSON(http.StatusOK, response)
}
