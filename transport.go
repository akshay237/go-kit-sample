package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
)

var (
	ErrBadRequest = errors.New("Bad Request")
	ErrInvalidId  = errors.New("Invalid id")
)

const (
	CharsetUTF8                    = "charset=UTF-8"
	HeaderContentType              = "Content-Type"
	MIMEApplicationJSON            = "application/json"
	MIMEApplicationJSONCharsetUTF8 = MIMEApplicationJSON + "; " + CharsetUTF8
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
	res, ok := response.(*addResponse)
	if !ok {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	w.Header().Set(HeaderContentType, MIMEApplicationJSONCharsetUTF8)
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(res.model)
}

func encodeGetResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(*getResponse)
	if !ok {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	w.Header().Set(HeaderContentType, MIMEApplicationJSONCharsetUTF8)
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(res.model)
}

func encodeUpdateResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(*updateResponse)
	if !ok {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	w.Header().Set(HeaderContentType, MIMEApplicationJSONCharsetUTF8)
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(res.model)
}

func encodeRemoveResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(*removeResponse)
	if !ok {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	w.Header().Set(HeaderContentType, MIMEApplicationJSONCharsetUTF8)
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(res.err)
}

func encodeGetAllResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(*getAllResponse)
	if !ok {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	w.Header().Set(HeaderContentType, MIMEApplicationJSONCharsetUTF8)
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(res.model)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case ErrBadRequest:
		w.WriteHeader(http.StatusBadRequest)
	case ErrInvalidId:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func makeHandler(s Service) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
	}

	addHandler := httptransport.NewServer(
		makeAddEndpoint(s),
		decodeAddRequest,
		encodeAddResponse,
		options...,
	)

	updateHandler := httptransport.NewServer(
		makeUpdateEndpoint(s),
		decodeUpdateRequest,
		encodeUpdateResponse,
		options...,
	)

	removeHandler := httptransport.NewServer(
		makeRemoveEndpoint(s),
		decodeRemoveRequest,
		encodeRemoveResponse,
		options...,
	)

	getHandler := httptransport.NewServer(
		makeGetEndpoint(s),
		decodeGetRequest,
		encodeGetResponse,
		options...,
	)

	getAllHandler := httptransport.NewServer(
		makeGetAllEndpoint(s),
		decodeGetAllRequest,
		encodeGetAllResponse,
		options...,
	)

	r := chi.NewRouter()
	r.Route("/items", func(r chi.Router) {
		r.Get("/getall", getAllHandler.ServeHTTP)
		r.Get("/get", getHandler.ServeHTTP)
		r.Post("/add", addHandler.ServeHTTP)
		r.Post("/update/{ID}", updateHandler.ServeHTTP)
		r.Get("/remove/{ID}", removeHandler.ServeHTTP)
	})
	return r
}
