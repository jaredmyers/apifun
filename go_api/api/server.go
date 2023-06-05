package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	m "github.com/jaredmyers/apifun/go_api/models"
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
	s.Router.HandleFunc("/user", errorHandler(s.handleUser))
	s.Router.HandleFunc("/user/{id}", errorHandler(s.handleUserById))
	s.Router.HandleFunc("/users", errorHandler(s.handleGetUsers))
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
		return s.registerUser(w, r)
	default:
		// may return nil here instead
		sC := http.StatusMethodNotAllowed
		return ApiErrParams{sC, http.StatusText(sC)}
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
		sC := http.StatusMethodNotAllowed
		return ApiErrParams{sC, http.StatusText(sC)}
	}
}

// ---- tester -----
func (s *Server) handleGetUsers(w http.ResponseWriter, r *http.Request) error {

	if r.Method != http.MethodGet {
		sC := http.StatusMethodNotAllowed
		return ApiErrParams{sC, http.StatusText(sC)}
	}

	// go out to userservice -> db
	users, err := s.userService.GetUsers()
	if err != nil {
		return err
	}
	return WriteJson(w, r, http.StatusOK, &m.GetUsersResponse{Users: users})
}

// ----

func (s *Server) getUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *Server) registerUser(w http.ResponseWriter, r *http.Request) error {
	log.Println("registering a user")

	var req m.RegisterUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	err := s.userService.RegisterUser(&req)
	if err != nil {
		return err
	}

	return WriteJson(w, r, http.StatusOK, &m.RegisterUserResponse{Status: "Registration Success"})
}

func (s *Server) getUserById(w http.ResponseWriter, r *http.Request) error {
	log.Println("FROM handleUserIdGet")

	start := time.Now()
	userSuppliedID := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(userSuppliedID)
	if err != nil {
		sC := http.StatusBadRequest
		return ApiErrParams{StatusCode: sC, StatusText: http.StatusText(sC)}
	}

	log.Println(userID)
	resp, err := s.userService.GetUser(userID)
	if err != nil {
		return err
	}

	log.Println(time.Since(start))
	return WriteJson(w, r, http.StatusOK, resp)
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

	// simple log to stdout serverside for now
	log.Println(r.URL, r.Method, status)

	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Ep ApiErrParams `json:"error"`
}

type ApiErrParams struct {
	StatusCode int    `json:"code"`
	StatusText string `json:"message"`
}

func (e ApiErrParams) Error() string {
	return e.StatusText
}

func errorHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {

			// errors which originated within the api package layer
			errParams, ok := err.(ApiErrParams)
			if ok {
				WriteJson(w, r, errParams.StatusCode, ApiError{errParams})
				return
			}

			// in the case of an error at the service or storage layer
			// differentiate between a real internal error and a bad request
			errorResp, ok := err.(m.InternalErrResp)
			if ok {
				log.Println(err.Error()) // log to stdout for now
				sC := http.StatusInternalServerError
				switch errorResp.Code() {
				case m.CodeNotFound:
					sC = http.StatusNotFound
				case m.CodeInternalError:
					sC = http.StatusInternalServerError
				case m.CodeInvalidArgument:
					sC = http.StatusUnprocessableEntity
					errParams := ApiErrParams{sC, err.Error()}
					WriteJson(w, r, sC, ApiError{errParams})
					return
				}

				errParams := ApiErrParams{sC, http.StatusText(sC)}
				WriteJson(w, r, sC, ApiError{errParams})
				return
			}
		}
	}
}

/*
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
				test, ok := err.(m.ErrorResponse)
				if ok {
					log.Println("this is running:", test.Code)
					switch test.Code {
					case 1:
						sC := http.StatusBadRequest
						WriteJson(w, r, sC, ApiPlainError{http.StatusText(sC)})
					default:
						sC := http.StatusInternalServerError
						WriteJson(w, r, sC, ApiPlainError{http.StatusText(sC)})
					}
				}
			}
		}
	}
}
*/
