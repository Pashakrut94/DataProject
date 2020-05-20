package statistics

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Pashakrut94/DataProject/client/handlers"
	"github.com/pkg/errors"
)

func GetTotal() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://localhost:8080/total")
		if err != nil {
			handlers.HandleResponseError(w, errors.Wrap(err, err.Error()).Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

	}
}

func UploadFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var url = "http://localhost:8080/download"

		data, err := ioutil.ReadFile("data.json")
		if err != nil {
			log.Fatalln("File reading error", err)
			return
		}

		dataBytes := bytes.NewReader(data)

		resp, err := http.Post(url, "application/json", dataBytes)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))

	}
}
