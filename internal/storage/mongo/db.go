package mongo

import (
	"github.com/cmd-ctrl-q/golang-arch/internal/models"
	"github.com/cmd-ctrl-q/golang-arch/internal/storage"
)

type MongoDB map[int]models.User

// func (r receiver) name(params) {}
func (m MongoDB) Save(n int, u models.User) {
	m[n] = u
}

func (m MongoDB) Retrieve(n int) models.User {
	return m[n]
}

func NewMongo() storage.Accessor {
	return MongoDB{}
}
