package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	addressesJSON    = []byte(`{ "addresses": [ "Via di S. Basilio, 15, 00187 Roma RM, Italy", "Via Forestella, 13, 03026 Pofi FR, Italy", "Address C", "Address D"]}`)
	addressesBADJSON = []byte(`{ "addresses": []}`)
	resultOK         = "{\"AddressOne\":{\"Name\":\"Via di S. Basilio, 15, 00187 Roma RM, Italy\",\"Lat\":41.9053319,\"Lng\":12.489857},\"AddressTwo\":{\"Name\":\"Via Forestella, 13, 03026 Pofi FR, Italy\",\"Lat\":41.5532695,\"Lng\":13.4022403},\"DistanceInKm\":85}\n"
)

func TestGetClosest(t *testing.T) {

	t.Run("returns closest addresses on POST", func(t *testing.T) {

		e := echo.New()
		request := newPostRequest(addressesJSON)
		response := httptest.NewRecorder()
		c := e.NewContext(request, response)

		if assert.NoError(t, CheckAddresses(c)) {
			assert.Equal(t, http.StatusOK, response.Code)
			assert.Equal(t, resultOK, response.Body.String())
		}

	})

	t.Run("returns Interval Server Error on POST", func(t *testing.T) {

		e := echo.New()
		request := newPostRequest(addressesBADJSON)
		response := httptest.NewRecorder()
		c := e.NewContext(request, response)

		assert.Error(t, CheckAddresses(c))
	})
}

func newPostRequest(addresses []byte) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/check-closest", bytes.NewBuffer(addresses))
	req.Header.Set("Content-Type", "application/json")
	return req
}
