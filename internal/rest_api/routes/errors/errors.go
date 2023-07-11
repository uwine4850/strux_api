package errors

import "fmt"

type ErrFormKeyNotExist struct {
	KeyName string
}

func (e *ErrFormKeyNotExist) Error() string {
	return fmt.Sprintf("Form key '%s' not exist", e.KeyName)
}
