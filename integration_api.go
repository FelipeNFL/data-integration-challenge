package main 

import (
	"os"
	"log"
	"context"
	"net/http"
	"github.com/gorilla/mux"
	"yawoen/core"
	"yawoen/services"
)



func main() {
	log.SetOutput(os.Stderr)

	port := ":8080"

	r := mux.NewRouter()
    r.HandleFunc("/company", services.UpdateCompany).Methods(http.MethodPut)
    r.HandleFunc("/health", services.Health).Methods(http.MethodGet)
	
	r.Use(core.SetDbMiddleware)
	r.Use(core.ParseBodyMiddleware)
	
	log.Print("Running Yawoen Integration API! Listen on"+port)
    log.Fatal(http.ListenAndServe(port, r))
}

