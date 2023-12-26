package main

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type addRequest struct {
	Name string `json:"string"`
}

type updateRequest struct {
	Id   int    `json:"id"`
	Name string `json:"string"`
}

type getRequest struct {
	Id int `json:"id"`
}

type removeRequest struct {
	Id int `json:"id"`
}

type addResponse struct {
	model *model
	err   error
}

type getResponse struct {
	model *model
	err   error
}

type updateResponse struct {
	model *model
	err   error
}

type getAllResponse struct {
	model []*model
	err   error
}

type removeResponse struct {
	err error
}

func makeAddEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		input, isok := req.(addRequest)
		if !isok {
			return nil, errors.New("Invalid Request Body")
		}
		info, err := s.add(input.Name)
		return &addResponse{
			model: info,
			err:   err,
		}, nil
	}
}

func makeGetEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		input, isok := req.(getRequest)
		if !isok {
			return nil, errors.New("Invalid Request Body")
		}
		info, err := s.getInfo(input.Id)
		return &getResponse{
			model: info,
			err:   err,
		}, nil
	}
}

func makeUpdateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		input, isok := req.(updateRequest)
		if !isok {
			return nil, errors.New("Invalid Request Body")
		}
		info, err := s.update(input.Id, input.Name)
		return &updateResponse{
			model: info,
			err:   err,
		}, nil
	}
}

func makeRemoveEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		input, isok := req.(removeRequest)
		if !isok {
			return nil, errors.New("Invalid Request Body")
		}
		err := s.remove(input.Id)
		return &removeResponse{
			err: err,
		}, nil
	}
}

func makeGetAllEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		info, err := s.getAll()
		return &getAllResponse{
			model: info,
			err:   err,
		}, nil
	}
}
