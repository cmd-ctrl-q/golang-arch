/*
Errors
Go erros are very *explicit* and gives a deterministic
and sequential understanding of the flow of the coee.

Many languages, you might write the code then add the error
handling later, but with Go, you handle the errors as you
write the functionality.

In some language, you may not know where to check for errors,
but with Go, it is explicit.

Golang builds documentation from comments.

Package /builtin
- dont have to import
- everything in it is available
- documentation for everything built into the code base.
- panic(), error type, recover(), print(), etc.


Dealing with errors
- panic(error)
	- shuts program down
	- use when your program cannot function without something.
	- often times, ignore recover
- returning
	- wrap errors in more detail
	- fmt.Errorf()
- handling errors with log (errors should only be handled once)
	- log.Print()
	- log.Fatal()
	- log.Panic()
- printing errors
	- fmt.Print()
- special code based on the error
	- io.EOF(),
	- return custom http status code
	- show instructions to user to resolve error


*/

package main

import (
	"errors"
	"fmt"
	"time"
)

// ************** Creating Custom errors **************

type ErrFileNotFound struct {
	Filename string
	When     time.Time // when error occurs
}

// valid error function because it implements the method Error(),
// which is of type error
func (e ErrFileNotFound) Error() string {
	return fmt.Sprintf("File %s was not found at %v", e.Filename, e.When)
}

// check if custom error struct is type of some error
func (e ErrFileNotFound) Is(other error) bool {
	_, ok := other.(ErrFileNotFound)
	return ok
}

var ErrorNotExist = fmt.Errorf("file does not exist")

// func openFile(filename string) (string, error) {
// 	return "", ErrNotExist
// }

func (e ErrFile) Error() string {
	return fmt.Sprintf("File %s was not found at", e.Filename)
}

// optional function for an error that is adding more detail
// to another error.
func (e ErrFile) Unwrap() error {
	// return the error youre adding more detail to
	return e.Base
}

func openFile(filename string) (string, error) {
	return "", ErrFile{
		Filename: filename,
		Base:     ErrNotExist,
	}
}

// more complex way of wrapping another error (base error)
// but needs an Unwrap() function to return the base error
type ErrFile struct {
	Filename string
	Base     error
}

var ErrNotExist = fmt.Errorf("file does not exist in file: %v", "fileName.txt")

// errors.New errors allows you to better compare errors
var ErrUserNotExist = errors.New("user does not exist")

func main() {
	err := ErrUserNotExist
	if err == ErrUserNotExist {
		fmt.Println("Youn need to register first")
	} else {
		fmt.Println("Unknown error")
	}

	fileErr := ErrFileNotFound{
		Filename: "test.txt",
		When:     time.Now(),
	}

	// built in
	fmt.Println(errors.Is(fileErr, ErrFileNotFound{})) // true
	fmt.Println(fileErr)

	// also built in
	// false without the custom Is()
	fmt.Println(errors.Is(fileErr, ErrFileNotFound{})) // false

	// also built in
	// true with or without the custom Is()
	fmt.Println(errors.Is(fileErr, ErrFileNotFound{
		Filename: "test.txt",
		When:     fileErr.When,
	})) // true

	// custom
	fmt.Println(fileErr.Is(ErrFileNotFound{})) // true

	// ************** Wrapping errors **************

	// Should not return raw errors, should wrap them with more details.

	// simple: use Errorf()
	_, err = openFile("test.txt")
	if err != nil {
		// simple and most common approach of wrapping errors
		wrapped := fmt.Errorf("unable to open file %v : %w", "test.txt", err)
		fmt.Println(wrapped)
		if errors.Is(wrapped, ErrNotExist) {
			fmt.Println("wrapped is of type ErrNotExist")
		} else {
			fmt.Println("wrapped is NOT of type ErrNotExist")
		}
	}
	// by using Errorf(), the Is() method still works.

	// could also make custom error struct that takes in the base error,
	_, err = openFile("test.txt")
	if err != nil {
		// even tho there is a wrapper which adds more detail to the error,
		// fundamentally, err is of type ErrNotExist
		if errors.Is(err, ErrNotExist) {
			fmt.Println("This is an ErrNotExist")
		}
		fmt.Println(err)
	}
}
