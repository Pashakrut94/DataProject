package region

import (
	"encoding/json"
)

type Regions struct {
	Regions []Region `json:"infectedByRegion"`
}

func ParseFileFromRequest(body []byte) (Regions, error) {
	var regions Regions

	err := json.Unmarshal(body, &regions)
	if err != nil {
		return Regions{}, err
	}
	return regions, nil
}
