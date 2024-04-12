package store

import (
	"database/sql"

	"github.com/min23asdw/go_api_learning/pkg/models"
)

type Store interface {
	//user
	CreateUser(u *models.User) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	CreateTask(*models.Task) (*models.Task, error)
	GetTask(id string) (*models.Task, error)
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

func (s *Storage) CreateUser(u *models.User) (*models.User, error) {
	rows, err := s.db.Exec("INSERT INTO users (email, firstName, lastName, password) VALUES (?, ?, ?, ?)", u.Email, u.FirstName, u.LastName, u.Password)
	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	u.ID = id
	return u, nil
}

func (s *Storage) CreateTask(t *models.Task) (*models.Task, error) {
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

func (s *Storage) GetTask(id string) (*models.Task, error) {
	var t models.Task
	err := s.db.QueryRow("SELECT id, name, status, project_id, assigned_to, createdAt FROM tasks WHERE id = ?", id).Scan(&t.ID, &t.Name, &t.Status, &t.ProjectID, &t.AssignedToID, &t.CreatedAt)
	return &t, err
}
func (s *Storage) GetUserByID(id string) (*models.User, error) {
	var u models.User
	err := s.db.QueryRow("SELECT id, email, firstName, lastName, createdAt FROM users WHERE id = ?", id).Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.CreatedAt)
	return &u, err
}
