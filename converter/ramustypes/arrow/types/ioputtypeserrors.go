package types

import "fmt"

type UnknownIOPutTypeError struct {
	error
	ioPutType IOPutType
}

func NewUnknownIOPutTypeError(ioPutType IOPutType) UnknownIOPutTypeError {
	return UnknownIOPutTypeError{
		ioPutType: ioPutType,
	}
}

func (err UnknownIOPutTypeError) Error() string {
	return fmt.Sprintf("unknown in/out -put type '%s'", err.ioPutType)
}

type InvalidIOPutInfoError struct {
	error
	info string
}

func NewInvalidIOPutInfoError(info string) InvalidIOPutInfoError {
	return InvalidIOPutInfoError{
		info: info,
	}
}

func (err InvalidIOPutInfoError) Error() string {
	return fmt.Sprintf("want [0-9](IOCM)[0-9] info struct - have %s", err.info)
}
