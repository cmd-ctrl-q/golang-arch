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

	f, err := os.Open("file-01.txt")
	if errors.Is(err, os.ErrPermission) {
		var pErr *os.PathError
		if errors.As(err, &pErr) {
			err = fmt.Errorf("permission denied : %w", pErr)
		} else {
			err = fmt.Errorf("permission denied : %w", err)
		}
		log.Println(err)
	} else if errors.Is(err, os.ErrNotExist) {
		err = fmt.Errorf("error not exist : %w", err)
		log.Println(err)
	} else if !errors.Is(err, nil) {
		err = fmt.Errorf("file could not be opened : %w", err)
		log.Println(err)
	}
	defer f.Close()
}
