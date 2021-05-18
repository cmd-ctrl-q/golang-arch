package main

import (
	"encoding/json"
	"log"
)

type person struct {
	First string
}

func main() {

	p1 := person{First: "alice"}
	p2 := person{First: "bob"}

	xp := []person{p1, p2}

	bs, err := json.Marshal(&xp)
	if err != nil {
		log.Panic(err)
	}

	log.Println(string(bs))
}
