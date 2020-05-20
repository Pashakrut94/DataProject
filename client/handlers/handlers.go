package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type FormatResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func HandleResponseError(w http.ResponseWriter, msg string, statusCode int) {
	c := FormatResponse{Error: msg}
	data, err := json.Marshal(&c)
	if err != nil {
		http.Error(w, errors.Wrap(err, "error marshaling response").Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, errors.Wrap(err, "error writing data").Error(), http.StatusInternalServerError)
		return
	}
}

func HandleResponse(w http.ResponseWriter, payload interface{}) {
	c := FormatResponse{Data: payload}
	var (
		data []byte
		err  error
	)

	data, err = json.MarshalIndent(&c, "", " ")

	if err != nil {
		http.Error(w, errors.Wrap(err, "error marshaling response").Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, errors.Wrap(err, "error writing data").Error(), http.StatusInternalServerError)
		return
	}
}

func HandlerResponseBody(w http.ResponseWriter, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(body)
	if err != nil {
		http.Error(w, errors.Wrap(err, "error writing data").Error(), http.StatusInternalServerError)
		return
	}
}
