package presenter

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/TomasCruz/grpc-web-fahrenheit/model"
)

// HealthHandler displays health status of the service.
// Status OK (200) is returned if service's client can be used,
// else status FailedDependency (424) is returned.
// Status InternalServerError (500) is returned in case of general errors.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	var (
		err          error
		status       bool
		statusString string
	)

	if status, err = model.Health(); err != nil {
		errorResponse(w, err, http.StatusFailedDependency)
		return
	} else if status {
		statusString = "Up"
	}

	data, err := json.Marshal(Health{Status: statusString})
	if err != nil {
		err = errors.New("Something terrible happened with JSON marshalling")
		errorResponse(w, err, http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
