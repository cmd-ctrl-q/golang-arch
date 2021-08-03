package main

import (
	"fmt"

	"github.com/cmd-ctrl-q/golang-arch/internal/models"
	"github.com/cmd-ctrl-q/golang-arch/internal/storage"
	"github.com/cmd-ctrl-q/golang-arch/internal/storage/harddrive"
	"github.com/cmd-ctrl-q/golang-arch/internal/storage/mongo"
)

func main() {

	mongoStorage := mongo.NewMongo()
	hdStorage := harddrive.NewHD()

	storage.Put(mongoStorage, 1, models.User{First: "alice"})
	storage.Put(hdStorage, 1, models.User{First: "bob"})
	storage.Put(mongoStorage, 2, models.User{First: "candice"})
	storage.Put(hdStorage, 2, models.User{First: "dameon"})

	fmt.Println(storage.Get(mongoStorage, 1))
	fmt.Println(storage.Get(hdStorage, 2))

	fmt.Println(storage.Get(mongoStorage, 1))
	fmt.Println(storage.Get(hdStorage, 2))
}
