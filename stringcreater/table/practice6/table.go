package practice6

import (
	"ramus/converter/ramustypes/box"
	"ramus/stringcreater/table/practice6/childprocessdescription"
	"ramus/stringcreater/table/practice6/processdescription"
	"strings"
)

func CreateTable(mainBox *box.Box) *strings.Builder {
	builder := strings.Builder{}

	prTable := processdescription.CreateTable(mainBox)
	childPrTable := childprocessdescription.CreateTable(mainBox)

	builder.WriteString(prTable.String() + "\n" +
		childPrTable.String() + "\n")

	return &builder
}
