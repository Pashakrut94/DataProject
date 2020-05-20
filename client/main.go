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

	router.Handle("/upload", statistics.UploadFile()).Methods("POST")
	// router.Handle("/statistics", statistics.GetTotal()).Methods("GET")

	http.Handle("/", router)
	fmt.Printf("Server starts at: %s\n", *port)
	http.ListenAndServe(*port, nil)

}
