package endpoint

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
)

var (
	ErrBadRequest = errors.New("bad request")
)

type ListPilotsRequest struct{}

type GetPilotRequest struct {
	Id guuid.UUID `json:"id"`
}

type DeletePilotRequest struct {
	Id guuid.UUID `json:"id"`
}

type CreatePilotRequest struct {
	UserId     string `json:"userId"`
	CodeName   string `json:"codeName"`
	SupplierId string `json:"supplierId"`
	MarketId   string `json:"marketId"`
	ServiceId  string `json:"serviceId"`
}

type UpdatePilotRequest struct {
	Id         guuid.UUID `json:"id"`
	UserId     string     `json:"userId"`
	CodeName   string     `json:"codeName"`
	SupplierId string     `json:"supplierId"`
	MarketId   string     `json:"marketId"`
	ServiceId  string     `json:"serviceId"`
}

type ChangeStatePilotRequest struct {
	Id    guuid.UUID `json:"id"`
	State string     `json:"state"`
}

func DecodeListPilotsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request ListPilotsRequest
	return request, nil
}

func DecodeGetPilotRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRequest
	}
	uuid, err := guuid.Parse(id)

	if err != nil {
		return nil, ErrBadRequest
	}

	return GetPilotRequest{Id: uuid}, nil
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
		return nil, ErrBadRequest
	}
	uuid, err := guuid.Parse(id)

	if err != nil {
		return nil, ErrBadRequest
	}
	var req UpdatePilotRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	req.Id = uuid
	return req, nil
}

func DecodeChangeStatePilotRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	state, okk := vars["state"]
	if !(ok || okk) {
		return nil, ErrBadRequest
	}

	uuid, err := guuid.Parse(id)

	if err != nil {
		return nil, ErrBadRequest
	}

	var req ChangeStatePilotRequest
	req.Id = uuid
	req.State = state
	return req, nil
}

func DecodeDeletePilotRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRequest
	}

	uuid, err := guuid.Parse(id)

	if err != nil {
		return nil, ErrBadRequest
	}
	return DeletePilotRequest{Id: uuid}, nil
}
