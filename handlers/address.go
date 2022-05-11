package handlers

import (
	"fmt"
	"go-address-microservice/models"

	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// {addresses: [“Address A”, “Address B”, “Address C”, "Address D"]}
func CheckAddresses(c echo.Context) (err error) {
	wrappedAddresses := new(models.WrappedAddresses)
	if err = c.Bind(wrappedAddresses); err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	log.Printf("addresses  to check: %s", wrappedAddresses.AddressesList)
	addresses := wrappedAddresses.Unwrap()

	fmt.Printf("%v", addresses)
	return c.JSON(http.StatusOK, addresses)
}
