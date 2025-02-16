package request

import (
	"http_server/pkg/response"
	"net/http"
)

func HandleBody[T any](writer *http.ResponseWriter, reader *http.Request) (*T, error) {
	body, decodeErr := Decode[T](reader.Body)
	if decodeErr != nil {
		response.Json(*writer, decodeErr, http.StatusBadRequest)
		return nil, decodeErr
	}

	validateErr := IsValid(body)
	if validateErr != nil {
		response.Json(*writer, validateErr, http.StatusBadRequest)
		return nil, validateErr
	}
	return &body, nil
}
