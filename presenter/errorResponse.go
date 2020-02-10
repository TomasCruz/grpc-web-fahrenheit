package presenter

import (
	"encoding/json"
	"net/http"
)

func errorResponse(w http.ResponseWriter, err error, errCode int) {
	if err != nil {
		// if there is also an error with JSON marshalling, return 500
		data, err := json.Marshal(ErrStruct{Msg: err.Error()})
		if err != nil {
			errCode = http.StatusInternalServerError
		}

		w.WriteHeader(errCode)
		w.Write(data)
		return
	}

	w.WriteHeader(errCode)
}
