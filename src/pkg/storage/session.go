package storage

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

const (
	errLog = "ERROR LOG >> "
)

type (
	AWSSessionConfig struct {
		AWSRegion 				string
		AWSCredentialsProfile 	string
		AWSConfigProfile 		string
		AWSConfigFile			string
	}

)

func NewAWSSession(c AWSSessionConfig) (*session.Session, error){
	cred := credentials.NewSharedCredentials(c.AWSConfigFile, c.AWSCredentialsProfile)

	config := aws.Config{
		Credentials:                       cred,
		Region:                            aws.String(c.AWSRegion),
	}

	sess, err := session.NewSessionWithOptions(session.Options{
		Config:                  config,
		Profile:                 c.AWSConfigProfile,
		SharedConfigState:       0,
	})
	if err != nil {
		//TODO: ERROR HANDLING
		fmt.Printf("%vFailed to setup session with given configuration:\n%v\n", errLog, err)
		return nil, err
	}

	s := session.Must(sess, err)

	return s, nil
}