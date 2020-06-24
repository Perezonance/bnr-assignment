package handlers

import (
	"fmt"
	"github.com/Perezonance/bnr-assignment/src/pkg/server"
	"net/http"

	"github.com/Perezonance/bnr-assignment/src/pkg/util"
)

type (
	Handler struct {
		s server.Server
	}
)

func NewHandler(server server.Server) *Handler{
	return &Handler{s:server}
}

func (h *Handler)UserHandler(w http.ResponseWriter, r *http.Request) {
	util.InfoLog(fmt.Sprintf("Received a %v request at UserHandler", r.Method))
	switch r.Method {
	case http.MethodGet:
		h.s.GetUser(w, r)
	case http.MethodPut:
		h.s.PostUser(w, r)
	case http.MethodPost:
		h.s.PostUser(w, r)
	case http.MethodDelete:
		h.s.DeleteUser(w, r)
	default:
		util.InfoLog(fmt.Sprintf("request type %v not handled", r.Method))
		util.WriteRes(http.StatusNotFound, http.StatusText(http.StatusNotFound), w)
	}
	return
}

func (h *Handler)PostHandler(w http.ResponseWriter, r *http.Request) {
	util.InfoLog(fmt.Sprintf("Received a %v request at PostHandler", r.Method))
	switch r.Method {
	case http.MethodGet:
		h.s.GetPost(w, r)
	case http.MethodPut:
		h.s.PostPost(w, r)
	case http.MethodPost:
		h.s.PostPost(w, r)
	case http.MethodDelete:
		h.s.DeletePost(w, r)
	default:
		util.InfoLog(fmt.Sprintf("request type %v not handled", r.Method))
		util.WriteRes(http.StatusNotFound, http.StatusText(http.StatusNotFound), w)
	}
	return
}