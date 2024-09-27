package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"service1/server/handler"
)

type Server struct {
	router *mux.Router
}

func New() *Server {
	r := mux.NewRouter()
	r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		fmt.Fprint(w, `{"status": "OK"}`)
	})

	r.HandleFunc("/user/{id}", handler.UserGet).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", handler.UserUpdate).Methods(http.MethodPost)
	r.HandleFunc("/user", handler.UserDelete).Methods(http.MethodDelete)

	return &Server{
		router: r,
	}
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
