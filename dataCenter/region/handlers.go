package region

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Pashakrut94/DataProject/dataCenter/handlers"
	"github.com/pkg/errors"
)

func CreateRegion(repo RegionRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var region Region
		if err := json.NewDecoder(r.Body).Decode(&region); err != nil {
			handlers.HandleResponseError(w, errors.Wrap(err, "error parsing signup request").Error(), http.StatusBadRequest)
			return
		}
		region, err := HandleCreateRegion(repo, region)
		if err != nil {
			handlers.HandleResponseError(w, errors.Wrap(err, "error creating region").Error(), http.StatusInternalServerError)
			return
		}
		handlers.HandleResponse(w, region)
	}
}

func GetRegion(repo RegionRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isoCode := r.FormValue("code")
		region, err := HandleGetRegion(repo, isoCode)
		switch errors.Cause(err) {
		case ErrNotFound:
			// http: superfluous response.WriteHeader call: from w.WriteHeader(statusCode)
			handlers.HandleResponseError(w, errors.Wrap(err, "no content").Error(), http.StatusNotFound)
			return
		case nil:
			handlers.HandleResponse(w, region)
			return
		default:
			handlers.HandleResponseError(w, errors.Wrap(err, err.Error()).Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetTotal(repo RegionRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		total, err := HandleGetTotal(repo)
		switch errors.Cause(err) {
		case ErrEmptyDB:
			// http: superfluous response.WriteHeader call: from w.WriteHeader(statusCode)
			handlers.HandleResponseError(w, errors.Wrap(err, "no data for aggregate").Error(), http.StatusNoContent)
			return
		case nil:
			handlers.HandleResponse(w, total)
			return
		default:
			handlers.HandleResponseError(w, errors.Wrap(err, err.Error()).Error(), http.StatusInternalServerError)
			return
		}
	}
}

func DownloadFile(repo RegionRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			handlers.HandleResponseError(w, errors.Wrap(err, err.Error()).Error(), http.StatusInternalServerError)
			return
		}
		regions, err := ParseFileFromRequest(body)
		if err != nil {
			handlers.HandleResponseError(w, errors.Wrap(err, err.Error()).Error(), http.StatusInternalServerError)
			return
		}

		for i := 0; i < len(regions.Regions); i++ {
			regions.Regions[i].Country = "Russia"

			_, err := HandleCreateRegion(repo, regions.Regions[i])
			if err != nil {
				handlers.HandleResponseError(w, errors.Wrap(err, "error creating region").Error(), http.StatusInternalServerError)
				return
			}
		}
		handlers.HandleResponse(w, "success: file downloaded")
	}
}
