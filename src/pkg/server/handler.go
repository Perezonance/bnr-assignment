package server

import (
	"fmt"
	"net/http"
)

func (s *Server)UserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received a %v request at UserHandler\n", r.Method)
	switch r.Method {
	case http.MethodGet:
		s.getUser(w, r)
	case http.MethodPut:
		s.postUser(w, r)
	case http.MethodPost:
		s.postUser(w, r)
	case http.MethodDelete:
		s.deleteUser(w, r)
	default:
		fmt.Printf("ERROR LOG > %v type of Request not handled..", r.Method)
		writeRes(http.StatusNotFound, http.StatusText(http.StatusNotFound), w)
	}
	return
}

func (s *Server)PostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received a %v request at PostHandler\n", r.Method)
	switch r.Method {
	case http.MethodGet:
		s.getPost(w, r)
	case http.MethodPut:
		s.postPost(w, r)
	case http.MethodPost:
		s.postPost(w, r)
	case http.MethodDelete:
		s.deletePost(w, r)
	default:
		fmt.Printf("ERROR LOG > %v type of Request not handled..", r.Method)
		writeRes(http.StatusNotFound, http.StatusText(http.StatusNotFound), w)
	}
	return
}