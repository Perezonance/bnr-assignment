package storage

import (
	"fmt"
	"github.com/Perezonance/bnr-assignment/src/pkg/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const (
	logRoot = "user-management-server>> "
)

type (
	DynamoClient struct {
		db 			*dynamodb.DynamoDB
		userTable 	string
		postTable 	string
	}

	DynamoConfig struct {
		UserTable			string
		PostTable 			string
		UserTableEndpoint 	string
		PostTableEndpoint 	string
		AwsRegion 			string
	}
)

func NewDynamo(c DynamoConfig) (*DynamoClient, error){
	config := &aws.Config{
		Credentials: credentials.NewSharedCredentials("", "perezonance-dynamo"),
		Region:      aws.String(c.AwsRegion),
	}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:  			*config,
		Profile:           	"perezonance-dynamo",
		SharedConfigState: 	session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	return &DynamoClient{
		db: svc,
		userTable: c.UserTable,
		postTable: c.PostTable,
	}, nil
}

//TODO:Implement Business Logic

/////////////////////////////////// User Services ///////////////////////////////////

func (d *DynamoClient)GetUser(id float64) (models.User, error){
	fmt.Printf(logRoot + "Searching %v table for user with id:%v\n", d.userTable, id)

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
		//TODO: ERROR HANDLING
		fmt.Printf("Error getting item from dynamo table:\n%v\n", err)
		return u, err
	}

	err = dynamodbattribute.UnmarshalMap(res.Item, &u)
	if err != nil {
		//TODO: ERROR HANDLING
		fmt.Printf("Error Unmarshaling item retrieved from dynamodb table:\n%v\n", err)
		return u, err
	}

	return u, nil
}

func (d *DynamoClient)PostUser(user models.User) error {
	fmt.Printf(logRoot + "Inserting user into %v table:%v\n", d.userTable, user)

	attrVal, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		//TODO: ERROR HANDLING
		fmt.Printf("Error creating dynamodb attributeValue:\n%v\n", err)
		return err
	}

	item := &dynamodb.PutItemInput{
		Item: attrVal,
		TableName: aws.String(d.userTable),
	}

	_, err = d.db.PutItem(item)
	if err != nil {
		//TODO: ERROR HANDLING
		fmt.Printf("Error putting item in db table:\n%v\n", err)
		return err
	}

	return nil
}

func (d *DynamoClient)DeleteUser(user models.User) error {
	fmt.Printf(logRoot + "Deleting user from %v table:%v\n", d.userTable, user)

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
		//TODO: ERROR HANDLING
		fmt.Printf("Error deleting item from table:\n%v\n", err)
		return err
	}

	return nil
}

/////////////////////////////////// Post Services ///////////////////////////////////

func (d *DynamoClient)GetPost(id float64) (models.Post, error){
	fmt.Printf(logRoot + "Searching %v table for post with id:%v\n", d.postTable, id)

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
		//TODO: ERROR HANDLING
		fmt.Printf("Error getting item from dynamo table:\n%v\n", err)
		return p, err
	}

	err = dynamodbattribute.UnmarshalMap(res.Item, &p)
	if err != nil {
		//TODO: ERROR HANDLING
		fmt.Printf("Error Unmarshaling item retrieved from dynamodb table:\n%v\n", err)
		return p, err
	}

	return p, nil
}

func (d *DynamoClient)PostPost(post models.Post) error {
	fmt.Printf(logRoot + "Inserting post into %v table:%v\n", d.postTable, post)

	attrVal, err := dynamodbattribute.MarshalMap(post)
	if err != nil {
		//TODO: ERROR HANDLING
		fmt.Printf("Error creating dynamodb attributeValue:\n%v\n", err)
		return err
	}

	item := &dynamodb.PutItemInput{
		Item: attrVal,
		TableName: aws.String(d.postTable),
	}

	_, err = d.db.PutItem(item)
	if err != nil {
		//TODO: ERROR HANDLING
		fmt.Printf("Error putting item in db table:\n%v\n", err)
		return err
	}

	return nil
}

func (d *DynamoClient)DeletePost(post models.Post) error {
	fmt.Printf(logRoot + "Deleting post from %v table:%v\n", d.postTable, post)

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
		//TODO: ERROR HANDLING
		fmt.Printf("Error deleting item from table:\n%v\n", err)
		return err
	}

	return nil
}