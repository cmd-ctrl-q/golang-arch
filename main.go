/*

Assertion with errors.As()

Assert errors to determine which type they are in order to
better diagnose the problem and return descriptive logs.

Also can access fields or methods assocaited with another
error type. Assert if this error is also of this other type,
which will give us access to the methods and fields.

`errors.As(err error, target interface{})`

*/
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {

	var perr *os.PathError
	f, err := os.Open("file-01.txt")
	log.Println(err)
	if err != nil {
		switch {
		case errors.Is(err, os.ErrPermission) && errors.As(err, &perr):
			err = fmt.Errorf("you do not have permission to open file : %w and path is %s", err, perr.Path)
		case errors.Is(err, os.ErrNotExist) && errors.As(err, &perr):
			err = fmt.Errorf("file does not exist : %w and path is %s", err, perr.Path)
		case errors.As(err, &perr):
			err = fmt.Errorf("original error %s : path is %s", err, perr.Path)
		case err != nil:
			err = fmt.Errorf("file could not be opened : %s", err)
		}
		log.Println(err)
	} else {
		fmt.Println(f)
		f.Close()
	}
}
