package main

import (
	"context"
	"encoding/json"
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
	fmt.Println("about to start the server at port 8080")
	http.ListenAndServe(":8080", nil)
}

func assignRoutes(router *mux.Router) *mux.Router {
	service := service.MakeServiceImpl()
	// var logger log.Logger
	// {
	//  logger = log.NewLogfmtLogger(os.Stderr)
	//  logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	//  logger = log.With(logger, "caller", log.DefaultCaller)
	// }

	// options := []httpTransport.ServerOption{
	//  httpTransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	//  httpTransport.ServerErrorEncoder(encodeError),
	// }

	listPilotsHandler := httpTransport.NewServer(
		endpoint.MakeListPilotsEndpoint(service),
		endpoint.MakeDecoder(endpoint.ListPilotsRequest{}),
		endpoint.EncodeResponse,
		// options...,
	)

	// router.Methods("GET").Path("/pilots").Handler(httpTransport.NewServer(
	//  endpoint.MakeListPilotsEndpoint(service),
	//  endpoint.MakeDecoder(endpoint.ListPilotsRequest{}),
	//  endpoint.EncodeResponse,
	//  options...,
	// ))
	router.Handle("/pilots", listPilotsHandler).Methods("GET")

	return router
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
func codeFrom(err error) int {
	fmt.Printf("%#v\n", err)
	// switch err {
	// // case ErrNotFound:
	// //   return http.StatusNotFound
	// case decoders.ErrInvalidCustomerID:
	//  return http.StatusBadRequest
	// default:
	//  return http.StatusInternalServerError
	// }
	return http.StatusInternalServerError
}
