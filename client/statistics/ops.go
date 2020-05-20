package statistics

import (
	"bytes"
	"io/ioutil"
)

func HandleUploadFile(fileName string) (*bytes.Reader, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	dataBytes := bytes.NewReader(data)
	return dataBytes, nil
}
