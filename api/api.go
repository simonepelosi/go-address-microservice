package api

import (
	"go-address-microservice/handlers"

	"github.com/labstack/echo/v4"
)

func MainGroup(e *echo.Echo) {
	// Route / to handler function
	e.POST("/check-closest", handlers.CheckAddresses)
}
