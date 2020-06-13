package server

import (
	"github.com/Perezonance/bnr-assignment/src/pkg/models"
	"github.com/Perezonance/bnr-assignment/src/pkg/storage"

	"encoding/json"
	"fmt"
	"net/http"
)

type (
	Server struct {
		db *storage.DynamoDB
	}
)

/////////////////////////////////// User Service Functions ///////////////////////////////////

func NewServer(db *storage.DynamoDB) (Server, error) {
	//TODO: Server instance id setup
	return Server{db:db}, nil
}

//TODO: need to allow both an array or a single json object as request payload
func (s *Server)getUser(w http.ResponseWriter, r *http.Request) {
	writeRes(200, "UNIMPLEMENTED", w)
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
			fmt.Printf("Input not valid")
			writeRes(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), w)
		}
	}
	users = append(users, reqUser.User, reqUsers.Payload)
	for _, u := range users {
		go func(){
			err := s.db.PostUser(u)
			if err != nil {
				//TODO: ERROR HANDLING
			}
		}()
	}
	writeRes(http.StatusOK, http.StatusText(http.StatusOK), w)
}

func (s *Server)deleteUser(w http.ResponseWriter, r *http.Request) {
	writeRes(200, "UNIMPLEMENTED", w)
}

/////////////////////////////////// Post Service Functions ///////////////////////////////////

func (s *Server)getPost(w http.ResponseWriter, r *http.Request) {
	writeRes(200, "UNIMPLEMENTED", w)
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
			fmt.Printf("Input not valid")
			writeRes(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), w)
		}
	}
	posts = append(posts, reqPost.Post, reqPosts.Payload)
	for _, p := range posts {
		go func(){
			err := s.db.PostPost(p)
			if err != nil {
				//TODO: ERROR HANDLING
			}
		}()
	}
	writeRes(http.StatusOK, http.StatusText(http.StatusOK), w)
}

func (s *Server)deletePost(w http.ResponseWriter, r *http.Request) {
	writeRes(200, "UNIMPLEMENTED", w)
}