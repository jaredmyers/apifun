package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jaredmyers/apifun/go_api/models"
	"github.com/jaredmyers/apifun/go_api/services"
)

type Server struct {
	listenAddr  string
	Router      *chi.Mux
	userService services.UserServicer
}

func NewServer(listenAddr string, userService services.UserServicer) *Server {
	return &Server{
		listenAddr:  listenAddr,
		Router:      chi.NewRouter(),
		userService: userService,
	}
}

// Register connects handlers to router
func (s *Server) RegisterRoutes() {
	s.Router.HandleFunc("/users", errorHandler(s.handleGetUsers))
	s.Router.HandleFunc("/user", errorHandler(s.handleUser))
	s.Router.HandleFunc("/user/{id}", errorHandler(s.handleUserById))
}

func (s *Server) Run() {
	log.Println("API Server running on port:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, s.Router)
}

// -- Route method switches
func (s *Server) handleUser(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.getUser(w, r)
	case "POST":
		return s.postUser(w, r)
	default:
		// may return nil here instead
		statusCode := http.StatusMethodNotAllowed
		return ErrParams{StatusCode: statusCode, StatusText: http.StatusText(statusCode)}
	}
}
func (s *Server) handleUserById(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.getUserById(w, r)
	case "PUT":
		return s.putUserById(w, r)
	case "DELETE":
		return s.deleteUserById(w, r)
	default:
		// may return nil here instead
		statusCode := http.StatusMethodNotAllowed
		return ErrParams{StatusCode: statusCode, StatusText: http.StatusText(statusCode)}
	}
}

// ---- tester -----
func (s *Server) handleGetUsers(w http.ResponseWriter, r *http.Request) error {

	if r.Method != http.MethodGet {
		return ErrParams{StatusCode: http.StatusMethodNotAllowed, StatusText: "Method Not Allowed"}
	}

	users, err := s.userService.GetUsers()
	if err != nil {
		return err
	}
	return WriteJson(w, r, http.StatusOK, &models.GetUsersResponse{Users: users})
}

// ----

func (s *Server) getUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *Server) postUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *Server) getUserById(w http.ResponseWriter, r *http.Request) error {
	log.Println("FROM handleUserIdGet")

	userSuppliedID := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(userSuppliedID)
	if err != nil {
		sC := http.StatusBadRequest
		return ErrParams{StatusCode: sC, StatusText: http.StatusText(sC)}
	}

	log.Println(userID)
	res, err := s.userService.GetUser(userID)
	if err != nil {
		return err
	}
	return WriteJson(w, r, http.StatusOK, res)
}
func (s *Server) putUserById(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *Server) deleteUserById(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// ---- Utility ----

func WriteJson(w http.ResponseWriter, r *http.Request, status int, v any) error {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	log.Println(r.URL, r.Method, status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiPlainError struct {
	Error string `json:"error"`
}

type ApiError struct {
	Ep ErrParams `json:"error"`
}

type ErrParams struct {
	StatusCode int    `json:"code"`
	StatusText string `json:"message"`
}

func (e ErrParams) Error() string {
	return e.StatusText
}

func errorHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {

			errParams, ok := err.(ErrParams)
			if ok {
				WriteJson(w, r, errParams.StatusCode, ApiError{errParams})
			} else {
				// may change entirely
				// converts any internal server err to 500
				// redirects actually err to stdout for now
				log.Println(err.Error())
				sC := http.StatusInternalServerError
				WriteJson(w, r, sC, ApiPlainError{http.StatusText(sC)})
			}
		}
	}
}
