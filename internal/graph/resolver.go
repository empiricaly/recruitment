package graph

import (
	"github.com/empiricaly/recruitment/internal/admin"
	"github.com/empiricaly/recruitment/internal/mturk"
	"github.com/empiricaly/recruitment/internal/storage"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you
// require here.

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	MTurk     *mturk.Session
	Store     *storage.Conn
	Admins    []admin.User
	SecretKey string
}
