package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/Pashakrut94/DataProject/dataCenter/region"

	"github.com/gorilla/mux"
)

var (
	port     = flag.String("port", ":8080", "Local port")
	pgUser   = flag.String("pg_user", "Pasha", "PostgreSQL name")
	pgPwd    = flag.String("pg_pwd", "pwd0123456789", "PostgreSQL password")
	pgHost   = flag.String("pg_host", "localhost", "PostgreSQL host")
	pgPort   = flag.String("pg_port", "54320", "PostgreSQL port")
	pgDBname = flag.String("pg_dbname", "mydb", "PostgreSQL name of DB")
)

func main() {
	flag.Parse()

	router := mux.NewRouter()

	router.Handle("/region", region.CreateRegion())

	http.Handle("/", router)
	fmt.Printf("Server starts at: %s\n", *port)
	http.ListenAndServe(*port, nil)
}
