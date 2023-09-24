package practice4

import (
	"ramus/converter/ramustypes/box"
	"ramus/stringcreater/table/practice4/connectiontypes"
	"ramus/stringcreater/table/practice4/notationelements"
	"ramus/stringcreater/table/practice4/objecttypes"
	"strings"
)

func CreateTable(mainBox *box.Box) *strings.Builder {
	builder := strings.Builder{}

	notationTable := notationelements.CreateTable(mainBox)
	connectionTable := connectiontypes.CreateTable(mainBox)
	objectTable := objecttypes.CreateTable(mainBox)

	builder.WriteString(notationTable.String() + "\n" +
		connectionTable.String() + "\n" +
		objectTable.String() + "\n")

	return &builder
}
