package region

import (
	"encoding/json"
	"net/http"

	"github.com/Pashakrut94/DataProject/dataCenter/handlers"
	"github.com/pkg/errors"
)

func CreateRegion(repo RegionRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, pretty := r.URL.Query()["pretty"]
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
		handlers.HandleResponse(w, region, pretty)
	}
}

func GetRegion(repo RegionRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isoCode := r.FormValue("code")
		region, err := HandleGetRegion(repo, isoCode)
		switch errors.Cause(err) {
		case ErrNotFound:
			// http: superfluous response.WriteHeader call: from w.WriteHeader(statusCode)
			handlers.HandleResponseError(w, errors.Wrap(err, "no content").Error(), http.StatusNoContent)
			return
		case nil:
			handlers.HandleResponse(w, region, false)
			return
		default:
			handlers.HandleResponseError(w, errors.Wrap(err, err.Error()).Error(), http.StatusInternalServerError)
			return
		}
	}
}
