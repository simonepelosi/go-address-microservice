package models

type Response struct {
	FirstAddress  Address
	SecondAddress Address
	Distance      float64
}

func NewResponse(firstAddress Address, secondAddress Address, distance float64) Response {
	var newResp Response
	newResp.FirstAddress = firstAddress
	newResp.SecondAddress = secondAddress
	newResp.Distance = distance

	return newResp
}
