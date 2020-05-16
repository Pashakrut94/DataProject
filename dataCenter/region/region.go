package main

import (
	"encoding/json"
	"net/http"
	
	"dataCenter/handlers"

	"github.com/pkg/errors"
)

type Region struct {
	Region    string `json:region`
	ISOCode   string `json:isoCode`
	Infected  int    `json:infected`
	Recovered int    `json:recovered`
	Deceased  int    `json:deceased`
}

func CreateRegion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, pretty := r.URL.Query()["pretty"]
		var region Region
		if err := json.NewDecoder(r.Body).Decode(&region); err != nil {
			handlers.HandleResponseError(w, errors.Wrap(err, "error parsing signup request").Error(), http.StatusBadRequest)
			return
		}

		handlers.HandleResponse(w, region, pretty)

	}
}
