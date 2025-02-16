package request

import (
	"encoding/json"
	"io"
)

func Decode[T any](body io.ReadCloser) (T, error) {
	var payload T
	decodeErr := json.NewDecoder(body).Decode(&payload)
	if decodeErr != nil {
		return payload, decodeErr
	}
	return payload, nil
}
