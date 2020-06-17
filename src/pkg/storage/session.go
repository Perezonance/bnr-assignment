package storage

import (
	"github.com/Perezonance/bnr-assignment/src/pkg/util"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
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
		util.ErrorLog("Failed to setup session with given configuration", err)
		return nil, err
	}

	s := session.Must(sess, err)

	return s, nil
}