package main

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
}

func NewService() Service {
	return &ToDoSvc{}
}

func (t *ToDoSvc) add(name string) (*model, error) {
	return nil, nil
}

func (t *ToDoSvc) remove(id int) error {
	return nil
}

func (t *ToDoSvc) getInfo(id int) (*model, error) {
	return nil, nil
}

func (t *ToDoSvc) getAll() ([]*model, error) {
	return nil, nil
}

func (t *ToDoSvc) update(id int, name string) (*model, error) {
	return nil, nil
}
