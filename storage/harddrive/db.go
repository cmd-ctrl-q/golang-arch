package harddrive

import (
	"github.com/cmd-ctrl-q/golang-arch/internal/models"
	"github.com/cmd-ctrl-q/golang-arch/storage"
)

type HardDrive map[int]models.User

func (hd HardDrive) Save(n int, u models.User) {
	hd[n] = u
}

func (hd HardDrive) Retrieve(n int) models.User {
	return hd[n]
}

func NewHD() storage.Accessor {
	return HardDrive{}
}
