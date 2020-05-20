package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	"github.com/Pashakrut94/DataProject/dataCenter/region"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var (
	port     = flag.String("port", ":8080", "Local port server")
	pgUser   = flag.String("pg_user", "Pasha", "PostgreSQL name")
	pgPwd    = flag.String("pg_pwd", "12345678", "PostgreSQL password")
	pgHost   = flag.String("pg_host", "localhost", "PostgreSQL host")
	pgPort   = flag.String("pg_port", "54320", "PostgreSQL port")
	pgDBname = flag.String("pg_dbname", "mydb", "PostgreSQL name of DB")
)

func main() {
	flag.Parse()
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", *pgUser, *pgPwd, *pgHost, *pgPort, *pgDBname)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	regionRepo := region.NewRegionRepo(db)

	router := mux.NewRouter()

	router.Handle("/region", region.CreateRegion(*regionRepo)).Methods("POST")
	router.Handle("/region", region.GetRegion(*regionRepo)).Methods("GET")
	router.Handle("/total", region.GetTotal(*regionRepo)).Methods("GET")

	router.Handle("/download", region.DownloadFile(*regionRepo)).Methods("POST")

	http.Handle("/", router)
	fmt.Printf("Server starts at: %s\n", *port)
	http.ListenAndServe(*port, nil)
}
