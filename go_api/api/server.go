package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jaredmyers/apifun/go_api/storage"
)

type Server struct {
	listenAddr  string
	userService *storage.UserService
}

func NewServer(listenAddr string, userService *storage.UserService) *Server {
	return &Server{
		listenAddr:  listenAddr,
		userService: userService,
	}
}

func (s *Server) Run() {
	router := chi.NewRouter()

	router.HandleFunc("/users", makeHandleFunc(s.handleGetUsers))
	router.HandleFunc("/user", makeHandleFunc(s.handleUser))
	router.HandleFunc("/user/{id}", makeHandleFunc(s.handleUserId))

	log.Println("API Server running on port:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

// -- Route method switches
func (s *Server) handleUser(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleUserGet(w, r)
	case "POST":
		return s.handleUserPost(w, r)
	default:

	}
	return WriteJson(w, r, http.StatusMethodNotAllowed, ApiError{Error: fmt.Errorf("not allowed").Error()})
}
func (s *Server) handleUserId(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleUserIdGet(w, r)
	case "PUT":
		return s.handleUserIdPut(w, r)
	case "DELETE":
		return s.handleUserIdDelete(w, r)
	default:

	}
	return WriteJson(w, r, http.StatusMethodNotAllowed, ApiError{Error: fmt.Errorf("not allowed").Error()})
}

// ---- tester
func (s *Server) handleGetUsers(w http.ResponseWriter, r *http.Request) error {
	log.Println("running handleUserGet")
	s.userService.Store.GetUsers()
	return nil
}

// ----

func (s *Server) handleUserGet(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *Server) handleUserPost(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *Server) handleUserIdGet(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *Server) handleUserIdPut(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *Server) handleUserIdDelete(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// ----

func WriteJson(w http.ResponseWriter, r *http.Request, status int, v any) error {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	log.Println(r.URL, r.Method, status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, r, http.StatusBadRequest, ApiError{err.Error()})
		}

	}

}
