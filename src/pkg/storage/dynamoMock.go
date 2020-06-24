package storage

import (
	"fmt"

	"github.com/Perezonance/bnr-assignment/src/pkg/models"
	"github.com/Perezonance/bnr-assignment/src/pkg/primitives"
	"github.com/Perezonance/bnr-assignment/src/pkg/util"
)

type (
	DynamoMock struct {
		Users map[float64]models.User
		Posts map[float64]models.Post

		userTable string
		postTable string
	}

	DynamoMockConfig struct {
		UserTableName string
		PostTableName string
	}
)

func NewMockDynamo(c DynamoMockConfig) *DynamoMock{
	return &DynamoMock{
		Users:     make(map[float64]models.User),
		Posts:     make(map[float64]models.Post),
		userTable: c.UserTableName,
		postTable: c.PostTableName,
	}
}

/////////////////////////////////// User Dynamo Methods ///////////////////////////////////

func (d *DynamoMock)GetUser(id float64) (models.User, error){
	util.InfoLog(fmt.Sprintf("Searching %v table for user with id:%v", d.userTable, id))

	var (
		user = 	models.User{}
		err  =	primitives.ErrUserNotFound
	)
	user = d.Users[id]
	if(user != models.User{}) {
		err = nil
	}
	if err == primitives.ErrUserNotFound {
		util.ErrorLog(fmt.Sprintf("Failed to find a user with id %v", id), err)
	}
	return user, err
}

func (d *DynamoMock)PostUser(user models.User) error {
	util.InfoLog(fmt.Sprintf("Inserting user into %v table:\n%v", d.userTable, user))
	d.Users[user.Id] = user
	return nil
}

func (d *DynamoMock)DeleteUser(user models.User) error {
	util.InfoLog(fmt.Sprintf("Deleting user from %v table:%v\n", d.userTable, user))
	delete(d.Users, user.Id)
	return nil
}

/////////////////////////////////// Post Dynamo Methods ///////////////////////////////////

func (d *DynamoMock)GetPost(id float64) (models.Post, error){
	util.InfoLog(fmt.Sprintf("Searching %v table for post with id:%v\n", d.postTable, id))

	var (
		post = models.Post{}
		err = primitives.ErrPostNotFound
	)

	post = d.Posts[id]
	if(post != models.Post{}) {
		err = nil
	}
	return post, err
}

func (d *DynamoMock)PostPost(post models.Post) error {
	util.InfoLog(fmt.Sprintf("Inserting post into %v table:%v\n", d.postTable, post))
	d.Posts[post.Id] = post
	return nil
}

func (d *DynamoMock)DeletePost(post models.Post) error {
	util.InfoLog(fmt.Sprintf("Deleting post from %v table:%v\n", d.postTable, post))
	delete(d.Posts, post.Id)
	return nil
}