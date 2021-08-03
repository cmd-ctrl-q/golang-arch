package storage

import "github.com/cmd-ctrl-q/golang-arch/internal/models"

// Accessor is used to Save and Retrieve users
type Accessor interface {
	Save(n int, u models.User)
	Retrieve(n int) models.User
}

func Put(a Accessor, n int, u models.User) {
	a.Save(n, u)
}

func Get(a Accessor, n int) models.User {
	return a.Retrieve(n)
}
