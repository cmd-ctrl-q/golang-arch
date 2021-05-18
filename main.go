package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
}

func main() {

	p1 := person{First: "alice"}
	p2 := person{First: "bob"}

	xp := []person{p1, p2}

	// marshal []person object into json
	bs, err := json.Marshal(&xp)
	if err != nil {
		log.Panic("error marshalling []person into json", err)
	}

	log.Println(string(bs))

	// unmarshal json into a []person object
	people := []person{}
	err = json.Unmarshal(bs, &people) // creat object
	if err != nil {
		log.Panic("error unmarshalling json into []person object", err)
	}

	fmt.Println(people)
}
