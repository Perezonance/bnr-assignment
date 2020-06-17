package server

import (
	"fmt"
	"net/http"

	"github.com/Perezonance/bnr-assignment/src/pkg/util"
)

func (s *Server)UserHandler(w http.ResponseWriter, r *http.Request) {
	util.InfoLog(fmt.Sprintf("Received a %v request at UserHandler", r.Method))
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
		util.InfoLog(fmt.Sprintf("request type %v not handled", r.Method))
		writeRes(http.StatusNotFound, http.StatusText(http.StatusNotFound), w)
	}
	return
}

func (s *Server)PostHandler(w http.ResponseWriter, r *http.Request) {
	util.InfoLog(fmt.Sprintf("Received a %v request at PostHandler", r.Method))
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
		util.InfoLog(fmt.Sprintf("request type %v not handled", r.Method))
		writeRes(http.StatusNotFound, http.StatusText(http.StatusNotFound), w)
	}
	return
}