package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ReadJSON(reader io.Reader, data any) error {
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&data)
	if err != nil {
		return err
	}

	return nil
}

func WriteJSON(data any, w http.ResponseWriter, status ...int) error {
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(status) == 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(status[0])
	}

	_, err = w.Write(resp)
	if err != nil {
		return err
	}

	return nil
}
