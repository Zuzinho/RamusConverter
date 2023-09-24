package table

import (
	"ramus/converter/ramustypes/box"
	"ramus/stringcreater/table/practice4"
	"strings"
)

func TablesByPractice(practiceNumber string) (func(*box.Box) *strings.Builder, error) {
	switch practiceNumber {
	case "4":
		return practice4.CreateTable, nil
	default:
		return nil, NewInvalidPracticeError(practiceNumber)
	}
}
