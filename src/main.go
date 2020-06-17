package main

import (
	"github.com/Perezonance/bnr-assignment/src/pkg/server"
	"github.com/Perezonance/bnr-assignment/src/pkg/storage"
	"github.com/Perezonance/bnr-assignment/src/pkg/util"

	"fmt"
	"net/http"
)

//TODO: Remove hard coded values and pull from ENV instead
const (
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
	util.ErrorLog("", start(true))
}

func start(mock bool) error {
	var db storage.Persistence

	if !mock {
		util.InfoLog("Setting up connection to AWS DynamoDB")
		sessionConf := storage.AWSSessionConfig{
			AWSRegion:             sessAWSRegion,
			AWSCredentialsProfile: sessAWSCredentialsProfile,
			AWSConfigProfile:      sessAWSConfigProfile,
			AWSConfigFile:         sessAWSConfigFile,
		}

		sess, err := storage.NewAWSSession(sessionConf)
		if err != nil {
			util.ErrorLog("Failed to establish an AWS Session", err)
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
		util.InfoLog("Setting up a Mock DynamoDB")
		mockConf := storage.DynamoMockConfig{
			UserTableName: dynamoUserTable,
			PostTableName: dynamoPostTable,
		}

		mock := storage.NewMockDynamo(mockConf)

		db = mock
	}

	s, err := server.NewServer(db)
	if err != nil {
		util.ErrorLog("Failed to establish the server", err)
		return err
	}

	util.InfoLog("Starting up server...")
	http.HandleFunc(root + "/user", s.UserHandler)
	http.HandleFunc(root + "/post", s.PostHandler)
	util.InfoLog(fmt.Sprintf("Listening on %v\n", fullAddr))

	return http.ListenAndServe(fullAddr, nil)
}