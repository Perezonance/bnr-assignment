package storage

import (
	"fmt"
	"github.com/Perezonance/bnr-assignment/src/pkg/models"
	"github.com/Perezonance/bnr-assignment/src/pkg/util"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type (
	DynamoClient struct {
		db 			*dynamodb.DynamoDB
		userTable 	string
		postTable 	string
	}

	DynamoConfig struct {
		UserTable				string
		PostTable 				string
		UserTableEndpoint 		string
		PostTableEndpoint 		string
		AWSSession 				*session.Session
	}
)

func NewDynamo(c DynamoConfig) *DynamoClient{
	return &DynamoClient{
		db:        dynamodb.New(c.AWSSession),
		userTable: c.UserTable,
		postTable: c.PostTable,
	}
}

/////////////////////////////////// User Services ///////////////////////////////////

func (d *DynamoClient)GetUser(id float64) (models.User, error){
	util.InfoLog(fmt.Sprintf("Searching %v table for user with id:%v", d.userTable, id))

	u := models.User{}

	res, err := d.db.GetItem(&dynamodb.GetItemInput{
		Key: 			map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(fmt.Sprintf("%v", id)),
			},
		},
		TableName:      aws.String(d.userTable),
	})
	if err != nil {
		util.ErrorLog("Failed to get item from table", err)
		return u, err
	}

	err = dynamodbattribute.UnmarshalMap(res.Item, &u)
	if err != nil {
		util.ErrorLog("Failed to unmarshal item retrieved from table", err)
		return u, err
	}

	return u, nil
}

func (d *DynamoClient)PostUser(user models.User) error {
	util.InfoLog(fmt.Sprintf("Inserting user into %v table:\n%v", d.userTable, user))

	attrVal, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		util.ErrorLog("Error creating dynamodb attributeValue", err)
		return err
	}

	item := &dynamodb.PutItemInput{
		Item: attrVal,
		TableName: aws.String(d.userTable),
	}

	_, err = d.db.PutItem(item)
	if err != nil {
		util.ErrorLog("Error putting item into db table", err)
		return err
	}
	return nil
}

func (d *DynamoClient)DeleteUser(user models.User) error {
	util.InfoLog(fmt.Sprintf("Deleting user from %v table:%v\n", d.userTable, user))

	delInput := &dynamodb.DeleteItemInput{
		Key: 		map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(fmt.Sprintf("%v", user.Id)),
			},
		},
		TableName:  aws.String(d.userTable),
	}

	_, err := d.db.DeleteItem(delInput)
	if err != nil {
		util.ErrorLog("Failed to delete item from table", err)
		return err
	}
	return nil
}

/////////////////////////////////// Post Services ///////////////////////////////////

func (d *DynamoClient)GetPost(id float64) (models.Post, error){
	util.InfoLog(fmt.Sprintf("Searching %v table for post with id:%v\n", d.postTable, id))

	p := models.Post{}

	res, err := d.db.GetItem(&dynamodb.GetItemInput{
		Key: 			map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(fmt.Sprintf("%v", id)),
			},
		},
		TableName: aws.String(d.postTable),
	})
	if err != nil {
		util.ErrorLog("Failed to get item from table", err)
		return p, err
	}

	err = dynamodbattribute.UnmarshalMap(res.Item, &p)
	if err != nil {
		util.ErrorLog("Failed to unmarshal retrieved item", err)
		return p, err
	}
	return p, nil
}

func (d *DynamoClient)PostPost(post models.Post) error {
	util.InfoLog(fmt.Sprintf("Inserting post into %v table:%v\n", d.postTable, post))

	attrVal, err := dynamodbattribute.MarshalMap(post)
	if err != nil {
		util.ErrorLog("Failed to create dynamodb attributeValue", err)
		return err
	}

	item := &dynamodb.PutItemInput{
		Item: attrVal,
		TableName: aws.String(d.postTable),
	}

	_, err = d.db.PutItem(item)
	if err != nil {
		util.ErrorLog("Failed to put item in table", err)
		return err
	}
	return nil
}

func (d *DynamoClient)DeletePost(post models.Post) error {
	util.InfoLog(fmt.Sprintf("Deleting post from %v table:%v\n", d.postTable, post))

	delInput := &dynamodb.DeleteItemInput{
		Key: 		map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(fmt.Sprintf("%v", post.Id)),
			},
		},
		TableName:  aws.String(d.postTable),
	}

	_, err := d.db.DeleteItem(delInput)
	if err != nil {
		util.ErrorLog("Failed to delete item from table", err)
		return err
	}
	return nil
}