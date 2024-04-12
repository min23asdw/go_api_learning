package store

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/min23asdw/go_api_learning/pkg/utils"
)

var ErrNameRequired = errors.New("name is required")
var ErrProjectIDRequired = errors.New("project id is required")
var ErrUserIDRequired = errors.New("user id is required")

type TasksService struct {
	store Store
}

func NewTasksService(s Store) *TasksService {
	return &TasksService{store: s}
}

func (s *TasksService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", s.HandleCreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", s.HandleGetTask).Methods("GET")
}

func (s *TasksService) HandleCreateTask(w http.ResponseWriter, r *http.Request) {

	// body in binary
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var task *Task_model
	//  parse !!!  and check it can parse
	err = json.Unmarshal(body, &task)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	// check it not emtry
	if err := validateTaskPayload(task); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	t, err := s.store.CreateTask(task)
	// can't CreateTask
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "Error creating task"})
		return
	}
	//everything OK
	utils.WriteJSON(w, http.StatusCreated, t)
}

func (s *TasksService) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	//just for safety
	if id == "" {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "id is required"})
		return
	}

	t, err := s.store.GetTask(id)

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "task not found"})
		return
	}

	//everything OK
	utils.WriteJSON(w, http.StatusOK, t)
}

func validateTaskPayload(task *Task_model) error {
	if task.Name == "" {
		return ErrNameRequired
	}

	if task.ProjectID == 0 {
		return ErrProjectIDRequired
	}

	if task.AssignedToID == 0 {
		return ErrUserIDRequired
	}

	return nil
}
