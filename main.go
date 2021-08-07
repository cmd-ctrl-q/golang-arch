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
