package models

type Response struct {
	AddressOne   Address
	AddressTwo   Address
	DistanceInKm float64
}

func NewResponse(firstAddress Address, secondAddress Address, distance float64) Response {
	var newResp Response
	newResp.AddressOne = firstAddress
	newResp.AddressTwo = secondAddress
	newResp.DistanceInKm = distance

	return newResp
}
