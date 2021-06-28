package main 

import (
	"os"
	"net/http"
	"github.com/gorilla/mux"
	"yawoen/core"
	"yawoen/services"
)

func main() {
	updateCompany := core.Endpoint{ Path: "/company", Function: services.UpdateCompany, MethodHttp: http.MethodPut }
	health := core.Endpoint{ Path: "/health", Function: services.Health, MethodHttp: http.MethodGet }
	endpoints := []core.Endpoint{ updateCompany , health }

	middlewares := []mux.MiddlewareFunc{ core.SetDbMiddleware, core.ParseBodyMiddleware }

	port := os.Getenv("API_PORT")

	apiName := "Yawoen Integration API"

	core.RunAPI(endpoints, middlewares, port, apiName)
}

