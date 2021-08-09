package fileWriter

import (
	"fmt"
	"io"
	"os"
)

const (
	// Error descriptions
	ErrCreatingFile  = "failed while creating file"
	ErrWritingToFile = "failed while writing to file"
	ErrClosingFile   = "failed while closing file"
)

type WriteFileError struct {
	Op  string
	err error
}

func (e *WriteFileError) Error() string {
	if e != nil {
		return e.err.Error()
	}
	return ""
}

func (e *WriteFileError) Unwrap() error {
	return e.err
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
				err: fmt.Errorf("failed while creating or accessing file: %w", err),
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
			err: fmt.Errorf("failed while writing to file: %w", err),
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
			err: fmt.Errorf("failed while closing a file: %w", err),
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
