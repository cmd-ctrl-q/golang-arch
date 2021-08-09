/*
Stop writing so many `if err != nil` statements.
*/
package main

import (
	"errors"
	"fmt"

	"github.com/cmd-ctrl-q/golang-arch/fileWriter"
)

func main() {

	f := fileWriter.NewWriteFile("file.txt")
	f.WriteString("line 1\n")
	f.WriteString("line 2\n")
	f.WriteString("line 3\n")
	f.WriteString("line 4\n")
	f.Close()

	err := f.Err()
	if err != nil {

		var fErr *fileWriter.WriteFileError
		if errors.As(err, &fErr) {

			if fErr.Op == "NewWriteFile-Create" {
				fmt.Println("Error occured at NewWriteFile()")
			}
		} else {
			fmt.Println("Could not assert err as &fErr")
		}
	} else {
		fmt.Println("No error!")
	}

}
