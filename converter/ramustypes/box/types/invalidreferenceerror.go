package types

import "fmt"

type InvalidReferenceError struct {
	error
	reference string
}

func NewInvalidReferenceError(reference string) InvalidReferenceError {
	return InvalidReferenceError{
		reference: reference,
	}
}

func (err InvalidReferenceError) Error() string {
	return fmt.Sprintf("want [A-Z][0-9]+ reference struct - have '%s'", err.reference)
}
