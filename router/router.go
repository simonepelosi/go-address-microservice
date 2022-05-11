package router

import (
	"go-address-microservice/api"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	//set main routes
	api.MainGroup(e)

	return e
}
