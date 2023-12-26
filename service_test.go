package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	type Args struct {
		model *model
	}

	testcases := []struct {
		name string
		args Args
		err  error
	}{
		{
			name: "Add test case 1",
			args: Args{
				model: &model{
					Name: "finish go-kit sample",
				},
			},
			err: nil,
		},
		{
			name: "Add test case 2",
			args: Args{
				model: &model{
					Name: "go-kit sample is in progress",
				},
			},
			err: nil,
		},
	}

	svc := NewService()
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := svc.add(tc.args.model.Name)
			assert.Equal(t, tc.err, err)
		})
	}
}

func TestRemove(t *testing.T) {
	type Args struct {
		id int
	}

	testcases := []struct {
		name string
		args Args
		err  error
	}{
		{
			name: "remove task 1",
			args: Args{
				id: 1,
			},
			err: nil,
		},
	}

	svc := NewService()
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := svc.remove(tc.args.id)
			assert.Equal(t, tc.err, err)
		})
	}
}

func TestUpdate(t *testing.T) {
	type Args struct {
		id   int
		name string
	}

	testcases := []struct {
		name string
		args Args
		err  error
	}{
		{
			name: "update task 1",
			args: Args{
				id:   1,
				name: "update task",
			},
			err: nil,
		},
	}

	svc := NewService()
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			task, err := svc.update(tc.args.id, tc.args.name)
			isok := assert.Equal(t, tc.err, err)
			if isok {
				fmt.Println("Task: ", task)
			}
		})
	}
}

func TestGetInfo(t *testing.T) {
	type Args struct {
		id int
	}

	testcases := []struct {
		name string
		args Args
		err  error
	}{
		{
			name: "get task 1",
			args: Args{
				id: 1,
			},
			err: nil,
		},
	}

	svc := NewService()
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			task, err := svc.getInfo(tc.args.id)
			isok := assert.Equal(t, tc.err, err)
			if isok {
				fmt.Println("Task: ", task)
			}
		})
	}
}

func TestGetAll(t *testing.T) {

	testcases := []struct {
		name string
		err  error
	}{
		{
			name: "get all tasks",
			err:  nil,
		},
	}

	svc := NewService()
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			task, err := svc.getAll()
			isok := assert.Equal(t, tc.err, err)
			if isok {
				fmt.Println("Tasks: ", task)
			}
		})
	}
}
