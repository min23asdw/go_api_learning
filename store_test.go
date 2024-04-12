package main

import (
	"github.com/min23asdw/go_api_learning/store"
)

// Mocks

type MockStore struct{}

// func (s *MockStore) CreateProject(p *Project) error {
// 	return nil
// }

// func (s *MockStore) GetProject(id string) (*Project, error) {
// 	return &Project{Name: "Super cool project"}, nil
// }

// func (s *MockStore) DeleteProject(id string) error {
// 	return nil
// }

func (s *MockStore) CreateUser() error {
	return nil
}

// func (s *MockStore) GetUserByID(id string) (*User, error) {
// 	return &User{}, nil
// }

func (s *MockStore) CreateTask(t *store.Task_model) (*store.Task_model, error) {
	return &store.Task_model{}, nil
}

// GetTask implements Store.
func (s *MockStore) GetTask(id string) (*store.Task_model, error) {
	return &store.Task_model{}, nil
}
