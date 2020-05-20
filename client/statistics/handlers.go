package statistics

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Pashakrut94/DataProject/client/handlers"
	"github.com/pkg/errors"
)

func GetTotal() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var url = "http://localhost:8080/total"

		resp, err := http.Get(url)
		if err != nil {
			handlers.HandleResponseError(w, errors.Wrap(err, err.Error()).Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			handlers.HandleResponseError(w, errors.Wrap(err, err.Error()).Error(), http.StatusInternalServerError)
		}

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		fmt.Println("response Body:", string(body))

		handlers.HandlerResponseBody(w, body)

	}
}

func GetRegion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isoCode := r.URL.Query().Get("code")

		var url = "http://localhost:8080/region?code=" + isoCode

		resp, err := http.Get(url)
		if err != nil {
			handlers.HandleResponseError(w, errors.Wrap(err, err.Error()).Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			handlers.HandleResponseError(w, errors.Wrap(err, err.Error()).Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		fmt.Println("response Body:", string(body))

		handlers.HandlerResponseBody(w, body)

	}
}

func UploadFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var url = "http://localhost:8080/region"

		regions, err := HandleUploadFile(r)
		if err != nil {
			handlers.HandleResponseError(w, errors.Wrap(err, err.Error()).Error(), http.StatusInternalServerError)
			return
		}

		for i := 0; i < len(regions.Regions); i++ {
			regions.Regions[i].Country = "Russia"

			regionByte, err := json.Marshal(regions.Regions[i])
			if err != nil {
				handlers.HandleResponseError(w, errors.Wrap(err, err.Error()).Error(), http.StatusInternalServerError)
				return
			}

			region := bytes.NewReader(regionByte)

			resp, err := http.Post(url, "application/json", region)
			if err != nil {
				handlers.HandleResponseError(w, errors.Wrap(err, err.Error()).Error(), http.StatusInternalServerError)
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				handlers.HandleResponseError(w, errors.Wrap(err, err.Error()).Error(), http.StatusInternalServerError)
				return
			}

			handlers.HandlerResponseBody(w, body)
		}

	}
}
