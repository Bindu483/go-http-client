package commons

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse response for error
type ErrorResponse struct {
	Data    []interface{} `json:"data"`
	Status  string        `json:"status"`
	Message string        `json:"message"`
}

// HandleError handles error and send response
func HandleError(rw http.ResponseWriter, err error) {
	// build default response
	var response *ErrorResponse
	response = &ErrorResponse{Data: make([]interface{}, 0), Message: "somethingWentWrong",
		Status: http.StatusText(http.StatusInternalServerError)}
	rw.Header().Set("Content-Type", "application/json")
	// set header, message and status
	switch err {
	case InvalidRequest:
		rw.WriteHeader(http.StatusBadRequest)
		response.Message = "invalidRequest"
		response.Status = http.StatusText(http.StatusBadRequest)
	default:
		rw.WriteHeader(http.StatusInternalServerError)
	}

	// send response
	json.NewEncoder(rw).Encode(response)
	return
}
