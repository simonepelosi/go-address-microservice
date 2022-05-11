package main

import (
	"go-address-microservice/router"
)

func main() {
	e := router.New()
	e.Logger.Fatal(e.Start(":8000"))
}
