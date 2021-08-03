package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("file-01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read %d bytes: %q\n%s", count, data[:count], string(data))

	file2, err := os.Create("file-02.txt")
	if err != nil {
		log.Fatal(err)
	}

	// copy content from file into file2
	n, err := io.Copy(file2, file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bytes written into file2: ", n)
}
