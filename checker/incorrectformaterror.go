package checker

import "fmt"

type IncorrectFormatErr struct {
	error
	format string
}

func NewIncorrectFormatErr(format string) *IncorrectFormatErr {
	return &IncorrectFormatErr{
		format: format,
	}
}

func (err IncorrectFormatErr) String() string {
	return fmt.Sprintf("wanted '.idl' format - '.%s' got", err.format)
}
