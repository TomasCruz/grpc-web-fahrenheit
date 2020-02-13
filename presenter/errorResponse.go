package presenter

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/TomasCruz/grpc-web-fahrenheit/callstack"
)

func errorResponse(w http.ResponseWriter, err error, errCode int) {
	if err != nil {
		callstack.LogErrStack(err)
		// if there is also an error with JSON marshalling, return 500
		data, err := json.Marshal(ErrStruct{Msg: errors.Unwrap(err).Error()})
		if err != nil {
			errCode = http.StatusInternalServerError
		}

		w.WriteHeader(errCode)
		w.Write(data)
		return
	}

	w.WriteHeader(errCode)
}
