package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("file-01.txt")
	if errors.Is(err, os.ErrNotExist) {
		err = fmt.Errorf("error not exist : %w", err)
	}
	if errors.Is(err, os.ErrPermission) {
		err = fmt.Errorf("permission denied : %w", err)
	}
	if !errors.Is(err, nil) {
		err = fmt.Errorf("another error : %w", err)
	}
	defer f.Close()
	fmt.Println("Stack error:\n", err)
}
