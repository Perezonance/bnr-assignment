package primitives

import "github.com/pkg/errors"

var (
	ErrUserNotFound = errors.New("User not found")
	ErrPostNotFound = errors.New("Post not found")
)
