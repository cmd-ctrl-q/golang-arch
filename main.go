package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string `json:"first"`
}

func main() {

	// tcp
	// handlers
	// encoder
	http.HandleFunc("/encoder", func(w http.ResponseWriter, r *http.Request) {
		// encode data to json then display it on web page.
		// make some data
		names := []person{
			{First: "alice"},
			{First: "bob"},
			{First: "cora"},
		}
		// encode and send to web page
		err := json.NewEncoder(w).Encode(&names)
		if err != nil {
			log.Println("error encoding data into json", err)
			return
		}
	})

	// decoder
	http.HandleFunc("/decoder", func(w http.ResponseWriter, r *http.Request) {

		// create object that the json data will decode into
		var people []person

		// get the json data and decode it
		err := json.NewDecoder(r.Body).Decode(&people)
		if err != nil {
			log.Println("error decoding json data into object", err)
			return
		}
		fmt.Println("people:\n", people)
	})

	http.ListenAndServe(":8080", nil)
}
