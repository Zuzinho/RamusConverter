package table

import (
	"ramus/converter/ramustypes/box"
	"ramus/stringcreater/table/connectiontypes"
	"ramus/stringcreater/table/notationelements"
	"ramus/stringcreater/table/objecttypes"
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
