package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/Pashakrut94/DataProject/client/statistics"
	"github.com/gorilla/mux"
)

var (
	port = flag.String("port", ":5050", "Local port client")
)

func main() {

	router := mux.NewRouter()

	router.Handle("/upload", statistics.UploadFile())
	router.Handle("/statistics", statistics.GetTotal())
	router.Handle("/stats", statistics.GetRegion()) // Ставлю statistics, и ухожу почемут на /statistics (не парсит r.FormValue?)

	http.Handle("/", router)
	fmt.Printf("Client starts at: %s\n", *port)
	http.ListenAndServe(*port, nil)

}
