package endpoint

import (
	"context"
	"encoding/json"
	"net/http"
)

type Response struct {
	Data   interface{} `json:"data"`
	Errors []error     `json:"errors"`
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
