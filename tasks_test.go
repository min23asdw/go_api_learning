package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/min23asdw/go_api_learning/pkg/store"
	"github.com/min23asdw/go_api_learning/pkg/utils"
)

func TestCreateTask(t *testing.T) {
	ms := &MockStore{}
	service := store.NewTasksService(ms)
	t.Run("should return error if name is empty", func(t *testing.T) {
		payload := &store.Task_model{
			Name: "",
		}

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/tasks", service.HandleCreateTask)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}

		var response utils.ErrorResponse
		err = json.NewDecoder(rr.Body).Decode(&response)
		if err != nil {
			t.Fatal(err)
		}

		if response.Error != store.ErrNameRequired.Error() {
			t.Errorf("expected error message %s, got %s", response.Error, store.ErrNameRequired.Error())
		}
	})
	t.Run("should create a task", func(t *testing.T) {
		payload := &store.Task_model{
			Name:         "A",
			ProjectID:    1,
			AssignedToID: 42,
		}

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/tasks", service.HandleCreateTask)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusAccepted, rr.Code)
		}

		// var response ErrorResponse
		// err = json.NewDecoder(rr.Body).Decode(&response)
		// if err != nil {
		// 	t.Fatal(err)
		// }

		// if response.Error != errNameRequired.Error() {
		// 	t.Errorf("expected error message %s, got %s", response.Error, errNameRequired.Error())
		// }
	})

}

func TestGetTask(t *testing.T) {
	ms := &MockStore{}
	service := store.NewTasksService(ms)

	t.Run("get task by id", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodGet, "/tasks/42", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/tasks/{id}", service.HandleGetTask)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusAccepted, rr.Code)
		}

	})

}
