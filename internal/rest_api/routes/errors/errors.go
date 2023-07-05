package errors

import "fmt"

type ErrFormKeyNotExist struct {
	KeyName string
}

func (e *ErrFormKeyNotExist) Error() string {
	return fmt.Sprintf("Form key '%s' not exist", e.KeyName)
}

type ErrInvalidFilePath struct {
	Path string
}

func (e *ErrInvalidFilePath) Error() string {
	return fmt.Sprintf("The path \"%s\" could not be processed.", e.Path)
}
