package main

import (
	"fmt"
	"net/http"
	"pilot-management/endpoint"
	"pilot-management/service"

	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	assignRoutes(router)
	http.Handle("/", router)
	fmt.Println("About to start the server at port 8080")
	http.ListenAndServe(":8080", nil)
}

func assignRoutes(router *mux.Router) *mux.Router {
	service := service.MakeServiceImpl()

	listPilotsHandler := httpTransport.NewServer(
		endpoint.MakeListPilotsEndpoint(service),
		endpoint.DecodeListPilotsRequest,
		endpoint.EncodeResponse,
	)

	getPilotHandler := httpTransport.NewServer(
		endpoint.MakeGetPilotEndpoint(service),
		endpoint.DecodeGetPilotRequest,
		endpoint.EncodeResponse,
	)

	CreatePilotHandler := httpTransport.NewServer(
		endpoint.MakeCreatePilotEndpoint(service),
		endpoint.DecodeCreatePilotRequest,
		endpoint.EncodeResponse,
	)

	UpdatePilotHandler := httpTransport.NewServer(
		endpoint.MakeUpdatePilotEndpoint(service),
		endpoint.DecodeUpdatePilotRequest,
		endpoint.EncodeResponse,
	)

	DeletePilotHandler := httpTransport.NewServer(
		endpoint.MakeDeletePilotEndpoint(service),
		endpoint.DecodeDeletePilotRequest,
		endpoint.EncodeResponse,
	)

	StatePilotHandler := httpTransport.NewServer(
		endpoint.MakeStatePilotEndpoint(service),
		endpoint.DecodeStatePilotRequest,
		endpoint.EncodeResponse,
	)

	router.Handle("/supply/pilots", listPilotsHandler).Methods("GET")
	router.Handle("/supply/pilots/{id}", getPilotHandler).Methods("GET")
	router.Handle("/supply/pilots", CreatePilotHandler).Methods("POST")
	router.Handle("/supply/pilots/{id}", UpdatePilotHandler).Methods("PUT")
	router.Handle("/supply/pilots/{id}", DeletePilotHandler).Methods("DELETE")
	router.Handle("/supply/pilots/{id}/{state}", StatePilotHandler).Methods("PUT")
	return router
}
