package table

import "fmt"

type InvalidPracticeError struct {
	error
	practiceNumber string
}

func NewInvalidPracticeError(practiceNumber string) InvalidPracticeError {
	return InvalidPracticeError{
		practiceNumber: practiceNumber,
	}
}

func (err InvalidPracticeError) Error() string {
	return fmt.Sprintf("unknown practice number - %s", err.practiceNumber)
}
