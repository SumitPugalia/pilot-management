package endpoint

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Data   interface{} `json:"data"`
	Errors []error     `json:"errors"`
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	fmt.Println(response)
	// if e, ok := response.(errorer); ok && e.error() != nil {
	// 	// Not a Go kit transport error, but a business-logic error.
	// 	// Provide those as HTTP errors.
	// 	encodeErrorResponse(ctx, e.error(), w)
	// 	return nil
	// }
	// w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// return json.NewEncoder(w).Encode(response)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(response.(Response).Errors))
	return json.NewEncoder(w).Encode(response)
}

func codeFrom(err []error) int {
	fmt.Println(err)
	// if err != [] {
	// 	switch err[0] {
	// 		// case 1:
	// 		// 	return http.StatusBadRequest
	// 		default:
	// 			return http.StatusInternalServerError

	// 	}
	// }
	return http.StatusOK
}
