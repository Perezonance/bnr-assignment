package storage

import "github.com/Perezonance/bnr-assignment/src/pkg/models"

type Persistence interface {
	GetUser(id float64) (models.User, error)
	PostUser(user models.User) error
	DeleteUser(user models.User) error
	GetPost(id float64) (models.Post, error)
	PostPost(user models.Post) error
	DeletePost(user models.Post) error
}
