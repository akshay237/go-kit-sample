package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var (
	ErrBadRequest = errors.New("Bad Request")
	ErrInvalidId  = errors.New("Invalid id")
)

// decode the incoming requests
func decodeAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req addRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrBadRequest
	}
	return req, nil
}

func decodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req getRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrBadRequest
	}
	return req, nil
}

func decodeGetAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return struct{}{}, nil
}

func decodeUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id, err := strconv.Atoi(chi.URLParam(r, "ID"))
	if err != nil {
		return nil, ErrInvalidId
	}
	name := chi.URLParam(r, "name")
	return updateRequest{
		Id:   id,
		Name: name,
	}, nil
}

func decodeRemoveRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id, err := strconv.Atoi(chi.URLParam(r, "ID"))
	if err != nil {
		return nil, ErrInvalidId
	}
	return removeRequest{
		Id: id,
	}, nil
}

// encode the responses
func encodeAddResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(*addResponse)
	return json.NewEncoder(w).Encode(res)
}

func encodeGetResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(*getResponse)
	return json.NewEncoder(w).Encode(res)
}

func encodeUpdateResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(*updateResponse)
	return json.NewEncoder(w).Encode(res)
}
