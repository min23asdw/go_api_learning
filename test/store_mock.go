package main_test

import (
	"github.com/min23asdw/go_api_learning/pkg/models"
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

func (s *MockStore) GetUserByID(id string) (*models.User, error) {
	return &models.User{}, nil
}

func (s *MockStore) CreateTask(t *models.Task) (*models.Task, error) {
	return &models.Task{}, nil
}

// GetTask implements Store.
func (s *MockStore) GetTask(id string) (*models.Task, error) {
	return &models.Task{}, nil
}
