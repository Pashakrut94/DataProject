package statistics

import (
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

		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		fmt.Println("response Body:", string(body))

		handlers.HandlerResponseBody(w, body)

	}
}

func GetRegion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// /statistics?code=RU-KDA перебрасывает сразу на /statistics и идет по другому ендпоинту получается..
		// /stats?code=RU-KDA все работает как надо, в чем проблема?
		isocode := r.FormValue("code")

		var url = "http://localhost:8080/region?code=" + isocode

		resp, err := http.Get(url)
		if err != nil {
			handlers.HandleResponseError(w, errors.Wrap(err, err.Error()).Error(), http.StatusInternalServerError)
			return
		}

		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		fmt.Println("response Body:", string(body))

		handlers.HandlerResponseBody(w, body)

	}
}

func UploadFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			fileName = "data.json"
			url      = "http://localhost:8080/download"
		)

		dataBytes, err := HandleUploadFile(fileName)
		if err != nil {
			handlers.HandleResponseError(w, errors.Wrap(err, err.Error()).Error(), http.StatusInternalServerError)
			return
		}

		resp, err := http.Post(url, "application/json", dataBytes)
		if err != nil {
			handlers.HandleResponseError(w, errors.Wrap(err, err.Error()).Error(), http.StatusInternalServerError)
			return
		}

		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		fmt.Println("response Body:", string(body))

		handlers.HandlerResponseBody(w, body)

	}
}
