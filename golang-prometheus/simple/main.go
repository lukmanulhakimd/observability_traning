package main

import (
	"fmt"
	"github.com/gorilla/mux"
	// "github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)


func main() {
	router := mux.NewRouter()

	// Prometheus endpoint
	// router.Path("/metrics").Handler(promhttp.Handler())

	// Serving static files
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Serving requests on port 9000")
	err := http.ListenAndServe(":9000", router)
	log.Fatal(err)
}
