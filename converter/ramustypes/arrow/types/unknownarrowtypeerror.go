package types

import "fmt"

type UnknownArrowTypeError struct {
	error
	tp string
}

func NewUnknownArrowTypeError(tp string) UnknownArrowTypeError {
	return UnknownArrowTypeError{
		tp: tp,
	}
}

func (err UnknownArrowTypeError) Error() string {
	return fmt.Sprintf("unknown arrow type '%s'", err.tp)
}
