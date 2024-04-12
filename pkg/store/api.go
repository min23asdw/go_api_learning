package store

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store Store
}

// constucer
func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()

	subrouter := router.PathPrefix("/api/v1").Subrouter()
	// regis services
	tasksService := NewTasksService(s.store)
	tasksService.RegisterRoutes(subrouter)

	log.Println("Starting API server at", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
