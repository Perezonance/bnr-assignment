package server

import (
	"fmt"
	"github.com/Perezonance/bnr-assignment/src/pkg/models"
	"github.com/Perezonance/bnr-assignment/src/pkg/storage"
	"github.com/Perezonance/bnr-assignment/src/pkg/util"

	"encoding/json"
	"net/http"
)

type (
	Server struct {
		db storage.Persistence
	}
)

/////////////////////////////////// User Service Functions ///////////////////////////////////

func NewServer(p storage.Persistence) (Server, error) {
	return Server{db:p}, nil
}

func (s *Server)getUser(w http.ResponseWriter, r *http.Request) {
	var (
		users []models.User
		ids []float64
		reqIds models.RequestUsersById
		reqId models.RequestUserById
	)
	err := json.NewDecoder(r.Body).Decode(&reqIds)
	if err != nil {
		//If not multiple User Ids, check single
		err := json.NewDecoder(r.Body).Decode(&reqId)
		if err != nil {
			//Input invalid
			util.InfoLog("Request input invalid")
			writeRes(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), w)
		}
	}
	defer func() {
		err := r.Body.Close()
		if err != nil {
			util.ErrorLog("Failed to close reader stream of request body", err)
			writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
			return
		}
	}()

	ids = append(reqIds.Users, reqId.Id)

	for _, i := range ids {
		go func(){
			user, err := s.db.GetUser(i)
			if err != nil {
				util.ErrorLog("Failed to getUser", err)
				writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
				return
			}
			users = append(users, user)
		}()
	}
	raw, err := json.Marshal(users)
	if err != nil {
		util.ErrorLog("Failed to marshal users into response", err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}
	res := string(raw)
	writeRes(http.StatusOK, res, w)
}

func (s *Server)postUser(w http.ResponseWriter, r *http.Request) {
	var (
		users []models.User
		reqUsers models.RequestUsersPayload
		reqUser models.RequestUserPayload
	)
	err := json.NewDecoder(r.Body).Decode(&reqUsers)
	if err != nil {
		//If not multiple Users, then check if single user
		err := json.NewDecoder(r.Body).Decode(&reqUser)
		if err != nil {
			//Input not valid
			util.InfoLog("Request input invalid")
			writeRes(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), w)
		}
	}
	defer func() {
		err := r.Body.Close()
		if err != nil {
			util.ErrorLog("Failed to close reader stream of request body", err)
			writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
			return
		}
	}()

	users = append(reqUsers.Payload, reqUser.User)
	for _, u := range users {
		go func(){
			util.DebugLog(fmt.Sprintf("User being inserted:\n%v", u))
			err := s.db.PostUser(u)
			if err != nil {
				util.ErrorLog("Failed to PostUser", err)
				writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
				return
			}
		}()
	}
	writeRes(http.StatusOK, http.StatusText(http.StatusOK), w)
}

func (s *Server)deleteUser(w http.ResponseWriter, r *http.Request) {
	var (
		users []models.User
		ids []float64
		reqIds models.RequestUsersById
		reqId models.RequestUserById
	)
	err := json.NewDecoder(r.Body).Decode(&reqIds)
	if err != nil {
		//If not multiple User Ids, check single
		err := json.NewDecoder(r.Body).Decode(&reqId)
		if err != nil {
			//Input invalid
			util.InfoLog("Request input invalid")
			writeRes(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), w)
		}
	}
	defer func() {
		err := r.Body.Close()
		if err != nil {
			util.ErrorLog("Failed to close reader stream of request body", err)
			writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
			return
		}
	}()

	ids = append(reqIds.Users, reqId.Id)

	for _, i := range ids {
		go func(){
			u, err := s.db.GetUser(i)
			if err != nil {
				util.ErrorLog("Failed to GetUser", err)
				writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
				return
			}
			err = s.db.DeleteUser(u)
			if err != nil {
				util.ErrorLog("Failed to DeleteUser", err)
				writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
				return
			}
			//TODO: There might be an issue with having many goroutines access this slice. May have to use a channel...
			users = append(users, u)
		}()
	}
	raw, err := json.Marshal(users)
	if err != nil {
		util.ErrorLog("Failed to marshal deleted users into response", err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}
	res := string(raw)
	writeRes(http.StatusOK, res, w)
}

/////////////////////////////////// Post Service Functions ///////////////////////////////////

func (s *Server)getPost(w http.ResponseWriter, r *http.Request) {
	var (
		posts []models.Post
		ids []float64
		reqIds models.RequestPostsById
		reqId models.RequestPostById
	)
	err := json.NewDecoder(r.Body).Decode(&reqIds)
	if err != nil {
		//If not multiple User Ids, check single
		err := json.NewDecoder(r.Body).Decode(&reqId)
		if err != nil {
			//Input invalid
			util.InfoLog("Request input invalid")
			writeRes(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), w)
		}
	}
	defer func() {
		err := r.Body.Close()
		if err != nil {
			util.ErrorLog("Failed to close reader stream of request body", err)
			writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
			return
		}
	}()

	ids = append(reqIds.Posts, reqId.Id)

	for _, i := range ids {
		go func(){
			post, err := s.db.GetPost(i)
			if err != nil {
				util.ErrorLog("Failed to GetPost", err)
				writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
				return
			}
			posts = append(posts, post)
		}()
	}
	raw, err := json.Marshal(posts)
	if err != nil {
		util.ErrorLog("Failed to marshal posts into response", err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}
	res := string(raw)
	writeRes(http.StatusOK, res, w)
}

func (s *Server)postPost(w http.ResponseWriter, r *http.Request) {
	var (
		posts []models.Post
		reqPosts models.RequestPostsPayload
		reqPost models.RequestPostPayload
	)
	err := json.NewDecoder(r.Body).Decode(&reqPosts)
	if err != nil {
		//If not multiple Posts, then check if single post
		err := json.NewDecoder(r.Body).Decode(&reqPost)
		if err != nil {
			//Input not valid
			util.InfoLog("Request input invalid")
			writeRes(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), w)
		}
	}
	defer func() {
		err := r.Body.Close()
		if err != nil {
			util.ErrorLog("Failed to close reader stream of request body", err)
			writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
			return
		}
	}()

	posts = append(reqPosts.Payload, reqPost.Post)
	for _, p := range posts {
		go func(){
			err := s.db.PostPost(p)
			if err != nil {
				util.ErrorLog("Failed to PostPost", err)
				writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
				return
			}
		}()
	}
	writeRes(http.StatusOK, http.StatusText(http.StatusOK), w)
}

func (s *Server)deletePost(w http.ResponseWriter, r *http.Request) {
	var (
		posts []models.Post
		ids []float64
		reqIds models.RequestPostsById
		reqId models.RequestPostById
	)
	err := json.NewDecoder(r.Body).Decode(&reqIds)
	if err != nil {
		//If not multiple User Ids, check single
		err := json.NewDecoder(r.Body).Decode(&reqId)
		if err != nil {
			//Input invalid
			util.InfoLog("Request input invalid")
			writeRes(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), w)
		}
	}
	defer func() {
		err := r.Body.Close()
		if err != nil {
			util.ErrorLog("Failed to close reader stream of request body", err)
			writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
			return
		}
	}()

	ids = append(reqIds.Posts, reqId.Id)

	for _, i := range ids {
		go func(){
			u, err := s.db.GetPost(i)
			if err != nil {
				util.ErrorLog("Failed to GetPost", err)
				writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
				return
			}
			err = s.db.DeletePost(u)
			if err != nil {
				util.ErrorLog("Failed to DeleteUser", err)
				writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
				return
			}
			//TODO: There might be an issue with having many goroutines access this slice. May have to use a channel...
			posts = append(posts, u)
		}()
	}
	raw, err := json.Marshal(posts)
	if err != nil {
		util.ErrorLog("Failed to marshal deleted posts into response", err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}

	res := string(raw)
	writeRes(http.StatusOK, res, w)
}