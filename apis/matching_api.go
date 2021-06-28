package main 

import (
	"os"
	"net/http"
	"github.com/gorilla/mux"
	"yawoen/core"
	"yawoen/services"
)

func main() {
	getCompany := core.Endpoint{ Path: "/company", Function: services.GetCompanyByNameAndZipCode, MethodHttp: http.MethodGet }
	health := core.Endpoint{ Path: "/health", Function: services.Health, MethodHttp: http.MethodGet }
	endpoints := []core.Endpoint{ getCompany , health }

	middlewares := []mux.MiddlewareFunc{ core.SetDbMiddleware, core.ParseBodyMiddleware }

	port := os.Getenv("API_PORT")

	apiName := "Yawoen Matching API"

	core.RunAPI(endpoints, middlewares, port, apiName)
}

