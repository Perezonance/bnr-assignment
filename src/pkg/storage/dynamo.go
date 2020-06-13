package storage

import (
	"fmt"

	"github.com/Perezonance/bnr-assignment/src/pkg/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	awsDynamoEndpoint 	string = ""
	awsDynamoRegion		string = ""

	userTable string = ""
	postTable string = ""

	logRoot = "user-management-server>> "
)

type (
	DynamoDB struct {
		db *dynamodb.DynamoDB
	}
)

func NewDynamo() (*DynamoDB, error){
	config := &aws.Config{
		Endpoint: aws.String(awsDynamoEndpoint),
		Region: aws.String(awsDynamoRegion),
	}

	sess := session.Must(session.NewSession(config))

	svc := dynamodb.New(sess)

	return &DynamoDB{db: svc}, nil
}

//TODO:Implement Business Logic

/////////////////////////////////// User Services ///////////////////////////////////

func (d *DynamoDB)GetUser(id float64) (models.User, error){
	fmt.Printf(logRoot + "Searching %v table for user with id:%v\n", userTable, id)
	fmt.Println("UNIMPLEMENTED")
	return models.User{}, nil
}

func (d *DynamoDB)PostUser(user models.User) error {
	fmt.Printf(logRoot + "Inserting user into %v table:%v\n", userTable, user)
	fmt.Println("UNIMPLEMENTED")
	return nil
}
func (d *DynamoDB)DeleteUser(user models.User) error {
	fmt.Printf(logRoot + "Deleting user from %v table:%v\n", userTable, user)
	fmt.Println("UNIMPLEMENTED")
	return nil
}

/////////////////////////////////// Post Services ///////////////////////////////////

func (d *DynamoDB)GetPost(id float64) (models.Post, error){
	fmt.Printf(logRoot + "Searching %v table for post with id:%v\n", postTable, id)
	fmt.Println("UNIMPLEMENTED")
	return models.Post{}, nil
}

func (d *DynamoDB)PostPost(user models.Post) error {
	fmt.Printf(logRoot + "Inserting post into %v table:%v\n", postTable, user)
	fmt.Println("UNIMPLEMENTED")
	return nil
}
func (d *DynamoDB)DeletePost(user models.Post) error {
	fmt.Printf(logRoot + "Deleting post from %v table:%v\n", postTable, user)
	fmt.Println("UNIMPLEMENTED")
	return nil
}