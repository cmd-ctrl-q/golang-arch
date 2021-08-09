/*
Stop writing so many `if err != nil` statements.
*/
package main

import "fmt"

type person struct {
	pseudonym string
}

func (p person) String() string {
	return fmt.Sprintf("my name is %s", p.pseudonym)
}

func main() {

	p := person{
		pseudonym: "batman",
	}

	fmt.Println(p) // my name is batman
	fmt.Print(p)   // my name is batman
}
