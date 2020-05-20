package statistics

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Region struct {
	Region    string `json:region`
	ISOCode   string `json:isocode`
	Infected  int    `json:infected`
	Recovered int    `json:recovered`
	Deceased  int    `json:deceased`
	Country   string `json:country`
}
type Regions struct {
	Regions []Region `json:"infectedByRegion"`
}

// Upload from file data to temporary file and return regions in suitable format
func HandleUploadFile(r *http.Request) (Regions, error) {
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("myFile")
	if err != nil {
		return Regions{}, err
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile("temp-files", "upload-*.json")
	if err != nil {
		return Regions{}, err
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return Regions{}, err
	}

	tempFile.Write(fileBytes)

	buf := bytes.NewBuffer(nil)

	_, err = io.Copy(buf, tempFile)
	if err != nil {
		return Regions{}, err
	}

	bytesData := buf.Bytes()

	var regions Regions

	err = json.Unmarshal(bytesData, &regions)
	if err != nil {
		return Regions{}, err
	}
	return regions, nil
}
