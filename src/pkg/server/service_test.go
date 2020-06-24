package server

import (
	"github.com/Perezonance/bnr-assignment/src/pkg/storage"
	"net/http"
	"testing"
)

var (
	mockDB *storage.DynamoMock
	mockServer Server
)

func init() {
	mockConfig := storage.DynamoMockConfig{
		UserTableName: "test-users",
		PostTableName: "test_posts",
	}

	mockDB = storage.NewMockDynamo(mockConfig)

	mockServer, _ = NewServer(mockDB)
}

func TestServer_getUser(t *testing.T) {
	type (
		input struct {
			w http.ResponseWriter
			r *http.Request
		}
		expected struct {
			mockState storage.DynamoMock
		}
		testCase struct {
			i input
			e expected
		}
	)
	var testCases []testCase

	for i, v := range testCases {
		mockServer.getUser(v.i.w, v.i.r)
		if *mockDB == v.e.mockState {

		}
	}
}

func TestServer_postUser(t *testing.T) {
	//TODO: UNIMPLEMENTED
}

func TestServer_deleteUser(t *testing.T) {
	//TODO: UNIMPLEMENTED
}

func TestServer_getPost(t *testing.T) {
	//TODO: UNIMPLEMENTED
}

func TestServer_postPost(t *testing.T) {
	//TODO: UNIMPLEMENTED
}

func TestServer_deletePost(t *testing.T) {
	//TODO: UNIMPLEMENTED
}