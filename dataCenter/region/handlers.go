package region

import "net/http"

// CreateRegion make new region on "/region"
func CreateRegion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Create Region Handler"))

	}
}
