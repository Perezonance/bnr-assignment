package storage

import (
	"testing"
)

func TestNewDynamo(t *testing.T) {
	type (
		input struct {
			testConf DynamoConfig
		}
		expected struct {
			client DynamoClient
		}
		test struct {
			in input
			exp expected
		}
	)
	var testCases []test

	for i, v := range testCases {
		t.Logf("Test case %v", i)
		if *NewDynamo(v.in.testConf) != v.exp.client {
			t.Logf("Test case #%v failed", i)
			t.Fail()
		}
	}
}

func TestDynamoClient_GetUser(t *testing.T) {
	//UNIMPLEMENTED
}

func TestDynamoClient_PostUser(t *testing.T) {
	//UNIMPLEMENTED
}

func TestDynamoClient_DeleteUser(t *testing.T) {
	//UNIMPLEMENTED
}

func TestDynamoClient_GetPost(t *testing.T) {
	//UNIMPLEMENTED
}

func TestDynamoClient_PostPost(t *testing.T) {
	//UNIMPLEMENTED
}

func TestDynamoClient_DeletePost(t *testing.T) {
	//UNIMPLEMENTED
}