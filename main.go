/*
Stop writing so many `if err != nil` statements.
*/
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	// Error descriptions
	ErrDefault       = "file error"
	ErrCreatingFile  = "error when creating file"
	ErrWritingToFile = "error when writing to file"
	ErrClosingFile   = "error when closing file"
)

type WriteFileError struct {
	Op  string
	Err error
}

func (e *WriteFileError) Error() string {
	if e != nil {
		return e.Err.Error()
	}
	return ""
}

func (e *WriteFileError) Unwrap() error {
	return e.Err
}

type writeFile struct {
	f   *os.File
	err *WriteFileError
}

// NewWriteFile creates a new file to write to
func NewWriteFile(filename string) *writeFile {
	f, err := os.Create(filename)

	if err != nil {
		return &writeFile{
			f: nil,
			err: &WriteFileError{
				Op:  "NewWriteFile-Create",
				Err: fmt.Errorf("failed while creating or accessing file: %w", err),
			},
		}
	}

	return &writeFile{
		f:   f,
		err: nil,
	}
}

// WriteString writes text to a file,
// Since there is no return, it is mutating the error if there is an error.
func (w *writeFile) WriteString(text string) {
	if w.err != nil {
		return
	}

	_, err := io.WriteString(w.f, text)
	if err != nil {
		w.err = &WriteFileError{
			Op:  "WriteString",
			Err: fmt.Errorf("failed while writing to file: %w", err),
		}
	}
}

// Close closes a file
func (w *writeFile) Close() {
	// cant close if the file wasnt open
	if w.f == nil {
		return
	}

	// close file
	err := w.f.Close()
	if err != nil {
		w.err = &WriteFileError{
			Op:  "Close",
			Err: fmt.Errorf("failed while closing a file: %w", err),
		}
	}
}

// Err returns a file error
func (w *writeFile) Err() error {
	if w.err == nil {
		return nil
	}

	return w.err
}

func main() {

	f := NewWriteFile("file.txt")
	f.WriteString("line 1\n")
	f.WriteString("line 2\n")
	f.WriteString("line 3\n")
	f.WriteString("line 4\n")
	f.Close()

	err := f.Err()
	if err != nil {

		var fErr *WriteFileError
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
