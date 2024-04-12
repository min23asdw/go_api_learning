package main

import "database/sql"

type Store interface {
	//user
	CreateUser() error
	CreateTask(*Task_model) (*Task_model, error)
	GetTask(id string) (*Task_model, error)
}

// sql.DB for commu with database
type Storage struct {
	db *sql.DB
}

//connstuct

func NewStore(db *sql.DB) *Storage {

	return &Storage{
		db: db,
	}
}

func (s *Storage) CreateUser() error {
	return nil
}

func (s *Storage) CreateTask(t *Task_model) (*Task_model, error) {
	rows, err := s.db.Exec("INSERT INTO tasks (name, status, project_id, assigned_to) VALUES (?, ?, ?, ?)", t.Name, t.Status, t.ProjectID, t.AssignedToID)

	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	t.ID = id
	return t, nil
}

func (s *Storage) GetTask(id string) (*Task_model, error) {
	var t Task_model
	err := s.db.QueryRow("SELECT id, name, status, project_id, assigned_to, createdAt FROM tasks WHERE id = ?", id).Scan(&t.ID, &t.Name, &t.Status, &t.ProjectID, &t.AssignedToID, &t.CreatedAt)
	return &t, err
}