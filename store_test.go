package main

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

func (s *MockStore) CreateTask(t *Task_model) (*Task_model, error) {
	return &Task_model{}, nil
}

// GetTask implements Store.
func (s *MockStore) GetTask(id string) (*Task_model, error) {
	return &Task_model{}, nil
}
