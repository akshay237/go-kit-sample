package main

import (
	"errors"
	"math/rand"
)

var (
	ErrAlreadyDeleted = errors.New("Task already deleted")
	ErrNotExist       = errors.New("Task not exists")
)

type model struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Service interface {
	add(name string) (*model, error)
	remove(id int) error
	getInfo(id int) (*model, error)
	getAll() ([]*model, error)
	update(id int, name string) (*model, error)
}

type ToDoSvc struct {
	itemMap map[int]model
	ids     []int
}

func NewService() Service {
	return &ToDoSvc{
		itemMap: make(map[int]model),
		ids:     make([]int, 0),
	}
}

func (t *ToDoSvc) add(name string) (*model, error) {
	id := rand.Intn(100000)
	_, isok := t.itemMap[id]
	if isok {
		id = rand.Intn(100000)
	}
	modelI := model{
		Id:   id,
		Name: name,
	}
	t.itemMap[id] = modelI
	return &modelI, nil
}

func (t *ToDoSvc) remove(id int) error {
	_, isok := t.itemMap[id]
	if !isok {
		return ErrAlreadyDeleted
	}
	delete(t.itemMap, id)
	return nil
}

func (t *ToDoSvc) getInfo(id int) (*model, error) {
	taskInfo, isok := t.itemMap[id]
	if !isok {
		return nil, ErrNotExist
	}
	return &taskInfo, nil
}

func (t *ToDoSvc) getAll() ([]*model, error) {
	resp := []*model{}
	for _, v := range t.itemMap {
		resp = append(resp, &v)
	}
	return resp, nil
}

func (t *ToDoSvc) update(id int, name string) (*model, error) {
	taskInfo, isok := t.itemMap[id]
	if !isok {
		return nil, ErrNotExist
	}
	taskInfo.Name = name
	t.itemMap[id] = taskInfo
	return &taskInfo, nil
}
