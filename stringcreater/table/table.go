package table

import (
	"ramus/converter/ramustypes/box"
	"ramus/stringcreater/table/practice4"
	"ramus/stringcreater/table/practice6"
	"strings"
)

func TablesByPractice(practiceNumber string) (func(*box.Box) *strings.Builder, error) {
	switch practiceNumber {
	case "4":
		return practice4.CreateTable, nil
	case "6":
		return practice6.CreateTable, nil
	default:
		return nil, NewInvalidPracticeError(practiceNumber)
	}
}
