package main

import (
	"github.com/Perezonance/bnr-assignment/src/pkg/server"
	"github.com/Perezonance/bnr-assignment/src/pkg/storage"

	"fmt"
	"net/http"
)

const (
	//Temp structured logging solution
	errLog 		= "ERROR LOG >> "
	infoLog 	= "INFO  LOG >> "

	root = "/api/v1"
	port = "8080"
	addr = "0.0.0.0"
	fullAddr = addr + ":" + port

	//AWS Session constants
	sessAWSRegion = "us-east-2"
	sessAWSCredentialsProfile = "perezonance-dynamo"
	sessAWSConfigProfile = "perezonance-dynamo"
	sessAWSConfigFile = ""

	//Dynamo constants
	dynamoUserTable = "bnr-users"
	dynamoPostTable = "bnr-posts"
	dynamoUserTableEndpoint = ""
	dynamoPostTableEndpoint = ""
)

func main() {
	fmt.Println(start(true))
}

func start(mock bool) error {
	var db storage.Persistence

	if !mock {
		sessionConf := storage.AWSSessionConfig{
			AWSRegion:             sessAWSRegion,
			AWSCredentialsProfile: sessAWSCredentialsProfile,
			AWSConfigProfile:      sessAWSConfigProfile,
			AWSConfigFile:         sessAWSConfigFile,
		}

		sess, err := storage.NewAWSSession(sessionConf)
		if err != nil {
			fmt.Printf(errLog + "Failed to establish an AWS session:\n%v\n", err)
			return err
		}

		dynamoConf := storage.DynamoConfig{
			UserTable:        	dynamoUserTable,
			PostTable:        	dynamoPostTable,
			UserTableEndpoint: 	dynamoUserTableEndpoint,
			PostTableEndpoint: 	dynamoPostTableEndpoint,
			AWSSession:     	sess,
		}

		db = storage.NewDynamo(dynamoConf)
	} else {
		mockConf := storage.DynamoMockConfig{
			UserTableName: dynamoUserTable,
			PostTableName: dynamoPostTable,
		}

		mock := storage.NewMockDynamo(mockConf)

		db = mock
	}
	s, err := server.NewServer(db)
	if err != nil {
		fmt.Printf(errLog + "Failed to establish the server:\n%v\n", err)
		return err
	}

	fmt.Printf(infoLog + "Starting up server...\n")

	http.HandleFunc(root + "/user", s.UserHandler)
	http.HandleFunc(root + "/post", s.PostHandler)

	fmt.Printf(infoLog + "Listening on %v\n", fullAddr)

	return http.ListenAndServe(fullAddr, nil)
}