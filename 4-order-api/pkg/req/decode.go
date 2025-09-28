package req

import (
	"encoding/json"
	"io"
)

func Decoder[T any](body io.ReadCloser) (T, error) {
	var payload T
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}
