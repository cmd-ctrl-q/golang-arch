/*

See http package
Error wraps around type error.
Therefore Error is a wrapper and implement Unwrap to unwrap
base errors from the wrapper.
```
type Error interface {
	error
	Timeout() bool // is the error a timeout?
	Temporary() bool // is the error temporary?
}
```

### Is() - used for unwrapping errors
Old way of using Is(),
`if err == os.ErrExist`

New way, uses Is()
`if errors.Is(err, os.ErrExist)`

### As() - used for converting errors and accessing fields of a struct
Old way of using As() - ie checking error types - uses assertion
```
ifperr, ok := err.(*os.PathError); ok {
	fmt.Println(perr.Path)
}
```

New way, uses As()
```
var perr *os.PathError
if errors.As(err, &perr) {
	fmt.Println(perr.Path)
}
```
*/
package main

import (
	"errors"
	"fmt"
	"net"
)

type ErrFile struct {
	Filename string
	Base     error
}

func (e ErrFile) Error() string {
	return fmt.Sprintf("file %s : %v", e.Filename, e.Base)
}

func (e ErrFile) Unwrap() error {
	return e.Base
}

var ErrNotExist = fmt.Errorf("file does not exist")

func openFile(filename string) (string, error) {
	return "", ErrFile{
		Filename: filename,
		Base:     ErrNotExist,
	}
}

func processFile(filename string) error {
	_, err := openFile(filename)
	if err != nil {
		return fmt.Errorf("error while opening file: %w", err)
	}

	return nil
}

func main() {

	// Using 'Is'
	err := processFile("test.txt")
	if err != nil {
		if errors.Is(err, ErrNotExist) {
			fmt.Println("this is an ErrNotExist")
		}
		fmt.Println(err)
	}

	// Using 'As'
	// to get more specific error types
	err = processFile("test.txt")
	if err != nil {
		// check if file error
		var fErr ErrFile
		// As() extracts the base type from the struct error.
		// works because the error in processFile() is dealt with
		// fmt.Errorf()
		if errors.As(err, &fErr) {
			fmt.Printf("was unable to do something with file %s\n", fErr.Filename)
		}

		// check if network error
		var netErr net.Error
		if errors.As(err, &netErr) {
			// check if network error is temporary
			if netErr.Temporary() {
				// retry
				fmt.Println("Retrying...")
			}
		}
		// some other error
		fmt.Println(err)
	}
}
