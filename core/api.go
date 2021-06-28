package core 

import (
	"os"
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type FunctionEndpoint func(http.ResponseWriter, *http.Request)

type Endpoint struct {
	Path string
	Function FunctionEndpoint
	MethodHttp string
}

func RunAPI(endpoints []Endpoint, middlewares []mux.MiddlewareFunc, port string, apiName string) {
	log.SetOutput(os.Stderr)

	r := mux.NewRouter()

	for _, endpoint := range endpoints {
	    r.HandleFunc(endpoint.Path, endpoint.Function).Methods(endpoint.MethodHttp)
	}

	for _, middleware := range middlewares {
		r.Use(middleware)
	}

	log.Print(fmt.Sprintf("Running %s! Listen on %s", apiName, port))
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))	
}