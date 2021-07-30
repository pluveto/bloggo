package bloggo

import "github.com/pluveto/bloggo/pkg/errcode"

var (
	ErrConflictEmail    = errcode.New(614001, "Email is conflict with another user's.")
	ErrConflictUsername = errcode.New(614002, "Username is conflict with another user's.")
)
