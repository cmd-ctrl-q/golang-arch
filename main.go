package main

import (
	"errors"
	"fmt"
)

func foo() error {
	return errors.New("foo")
}

func bar() error {
	// %w means the error is wrapped
	return fmt.Errorf("bar: %w", foo())
}

func baz() error {
	return fmt.Errorf("baz: %w", bar())
}

func main() {

	err := baz()     // baz is most recent error
	fmt.Println(err) // baz: bar: foo

	baseErr := errors.Unwrap(err)
	fmt.Println(baseErr) // bar: foo

	baseErr = errors.Unwrap(baseErr)
	fmt.Println(baseErr) // foo
}
