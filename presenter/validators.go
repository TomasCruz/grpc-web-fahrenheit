package presenter

import (
	"net/http"
	"strconv"

	"github.com/TomasCruz/grpc-web-fahrenheit/model"
)

// Float64InputValidator validates float64 input, passing control to f if input is ok
func Float64InputValidator(route string, f func(float64) http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		length := len(route)
		float64String := r.URL.Path[length:]
		if float64String[len(float64String)-1] == '/' {
			float64String = float64String[:len(float64String)-1]
		}

		f64, err := strconv.ParseFloat(float64String, 64)
		if err != nil {
			errorResponse(rw, model.ErrInvalidInput, http.StatusNotAcceptable)
			return
		}

		f(f64)(rw, r)
	}
}
