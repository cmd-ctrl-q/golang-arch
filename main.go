package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

	// *** tcp ***

	// encoder
	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		// encode data into json
		p1 := person{First: "alice"}
		err = json.NewEncoder(w).Encode(p1.First)
		if err != nil {
			log.Println("Could not encode data", err)
		}
	})

	// decoder
	// https://curlbuilder.com/
	// curl -XGET -H "Content-type: application/json" -d '{"First": "Bob"}' 'localhost:8080/decode'
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var p1 person
		err := json.NewDecoder(r.Body).Decode(&p1)
		if err != nil {
			log.Println("Could not decode json data", err)
		}

		log.Println("Person: ", p1)
	})

	http.ListenAndServe(":8080", nil)
}
