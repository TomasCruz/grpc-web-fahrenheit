package presenter

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/TomasCruz/grpc-web-fahrenheit/model"
)

//C2FHandler is converting degrees celsius to degrees fahrenheit
// Status OK (200) is returned with the body containing celsius/fahrenheit pair.
// NotAcceptable (406) if path parameter can't be parsed to a float64
// FailedDependency (424) is returned if client returned an error
// Status InternalServerError (500) is returned in case of general errors.
func C2FHandler(c float64) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			f   float64
			err error
		)

		if f, err = model.C2F(c); err != nil {
			errorResponse(w, err, http.StatusFailedDependency)
			return
		}

		data, err := json.Marshal(DegreePair{C: c, F: f})
		if err != nil {
			err = errors.New("Something terrible happened")
			errorResponse(w, err, http.StatusInternalServerError)
			return
		}

		w.Write(data)
		return
	}
}

//F2CHandler is converting degrees fahrenheit to degrees celsius
// Status OK (200) is returned with the body containing celsius/fahrenheit pair.
// NotAcceptable (406) if path parameter can't be parsed to a float64
// FailedDependency (424) is returned if client returned an error
// Status InternalServerError (500) is returned in case of general errors.
func F2CHandler(f float64) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			c   float64
			err error
		)

		if c, err = model.F2C(f); err != nil {
			errorResponse(w, err, http.StatusFailedDependency)
			return
		}

		data, err := json.Marshal(DegreePair{C: c, F: f})
		if err != nil {
			err = errors.New("Something terrible happened")
			errorResponse(w, err, http.StatusInternalServerError)
			return
		}

		w.Write(data)
		return
	}
}
