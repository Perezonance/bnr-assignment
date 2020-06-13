package storage

import "github.com/Perezonance/bnr-assignment/src/pkg/models"

type Persistence interface {
	GetUser(id string) (models.User, error)
	PostUser(user models.User) error
	DeleteUser(user models.User) error
	GetPost(id string) (models.Post, error)
	PostPost(user models.Post) error
	DeletePost(user models.Post) error
}
