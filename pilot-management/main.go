package main

import (
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"pilot-management/endpoint"
	"pilot-management/service"
)

func assignRoutes(router *mux.Router) *mux.Router {
	service := service.MakeServiceImpl()

	listPilotsHandler := httpTransport.NewServer(
		endpoint.MakeListPilotsEndpoint(service),
		endpoint.MakeDecoder(endpoint.ListPilotsRequest{}),
		endpoint.EncodeResponse)

	router.Handle("/pilots", listPilotsHandler).Methods("GET")

	return router
}

func main() {
	router := mux.NewRouter()
	assignRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
