package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/min23asdw/go_api_learning/store"
)

type APIServer struct {
	addr  string
	store store.Store
}

// constucer
func NewAPIServer(addr string, store store.Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()

	subrouter := router.PathPrefix("/api/v1").Subrouter()
	// regis services
	tasksService := store.NewTasksService(s.store)
	tasksService.RegisterRoutes(subrouter)

	log.Println("Starting API server at", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
