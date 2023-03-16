package api

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jaredmyers/apifun/go_api/services"
	"github.com/jaredmyers/apifun/go_api/storage"
)

var listenAddr string = ":8000"
var server *Server

func init() {
	store, err := storage.NewMockStore()
	if err != nil {
		log.Fatal(err)
	}

	// still figuring out how/when/why to isolate the
	// service layers for mocking (if necessary at all?)
	userService := services.NewUserService(store)
	server = NewServer(listenAddr, userService)
	server.RegisterRoutes()
}

func TestGetUsers_Get(t *testing.T) {

	req, _ := http.NewRequest("GET", "/users", nil)
	resp := executeRequest(req)

	expected := http.StatusOK
	actual := resp.Code
	if expected != actual {
		t.Errorf("Expected %d, got %d\n", expected, actual)
	}

	//log.Printf("Got: %v\n", resp.Body)
}

func TestGetUsers_Post(t *testing.T) {

	req, _ := http.NewRequest("POST", "/users", nil)
	resp := executeRequest(req)

	expected := http.StatusMethodNotAllowed
	actual := resp.Code
	if expected != actual {
		t.Errorf("Expected %d, got %d\n", expected, actual)
	}
}

// -- Utility --

func executeRequest(req *http.Request) *httptest.ResponseRecorder {

	rr := httptest.NewRecorder()
	server.Router.ServeHTTP(rr, req)

	return rr
}
