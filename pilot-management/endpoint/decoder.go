package endpoint

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	ErrBadRouting = errors.New("bad routing")
)

type ListPilotsRequest struct{}

type GetPilotRequest struct {
	Id string `json:"id"`
}

type DeletePilotRequest struct {
	Id string `json:"id"`
}

type CreatePilotRequest struct {
	UserId     string `json:"userId"`
	CodeName   string `json:"codeName"`
	SupplierId string `json:"supplierId"`
	MarketId   string `json:"marketId"`
	ServiceId  string `json:"serviceId"`
}

type UpdatePilotRequest struct {
	Id         string `json:"id"`
	UserId     string `json:"userId"`
	CodeName   string `json:"codeName"`
	SupplierId string `json:"supplierId"`
	MarketId   string `json:"marketId"`
	ServiceId  string `json:"serviceId"`
}

type ChangeStatePilotRequest struct {
	Id    string `json:"id"`
	State string `json:"state"`
}

func DecodeListPilotsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request ListPilotsRequest
	return request, nil
}

func DecodeGetPilotRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return GetPilotRequest{Id: id}, nil
}

func DecodeCreatePilotRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req CreatePilotRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

func DecodeUpdatePilotRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}

	var req UpdatePilotRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	req.Id = id
	return req, nil
}

func DecodeChangeStatePilotRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	state, okk := vars["state"]
	if !ok || !okk {
		return nil, ErrBadRouting
	}

	var req ChangeStatePilotRequest
	req.Id = id
	req.State = state
	return req, nil
}

func DecodeDeletePilotRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return DeletePilotRequest{Id: id}, nil
}
