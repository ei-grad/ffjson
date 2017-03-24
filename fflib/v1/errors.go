package v1

import "fmt"

type UnknownFieldsError struct {
	Keys []string // JSON keys that could not be mapped to fields
}

func NewUnknownFieldsError(keys []string) *UnknownFieldsError {
	return &UnknownFieldsError{Keys: keys}
}

func (e *UnknownFieldsError) Error() string {
	return fmt.Sprintf("json: object has unknown keys %q", e.Keys)
}
