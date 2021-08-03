package main

import "fmt"

type user struct {
	first string
}

type mongoDB map[int]user

// func (r receiver) name(params) {}
func (m mongoDB) save(n int, u user) {
	m[n] = u
}

func (m mongoDB) retrieve(n int) user {
	return m[n]
}

type harddrive map[int]user

func (hd harddrive) save(n int, u user) {
	hd[n] = u
}

func (hd harddrive) retrieve(n int) user {
	return hd[n]
}

// mongoDB and harddrive implement accessor
type accessor interface {
	save(n int, u user)
	retrieve(n int) user
}

func put(a accessor, n int, u user) {
	a.save(n, u)
}

func get(a accessor, n int) user {
	return a.retrieve(n)
}

func main() {

	mongoStorage := mongoDB{}
	harddriveStorage := harddrive{}

	put(mongoStorage, 1, user{first: "alice"})
	put(harddriveStorage, 1, user{first: "bob"})
	put(mongoStorage, 2, user{first: "candice"})
	put(harddriveStorage, 2, user{first: "dameon"})

	fmt.Println(get(mongoStorage, 1))
	fmt.Println(get(harddriveStorage, 2))

	fmt.Println(get(mongoStorage, 1))
	fmt.Println(get(harddriveStorage, 2))
}
