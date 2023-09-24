package processdescription

import (
	"fmt"
	"ramus/converter/ramustypes/arrow"
	"ramus/converter/ramustypes/box"
	"strings"
)

func CreateTable(mainBox *box.Box) *strings.Builder {
	builder := strings.Builder{}

	builder.WriteString(tableTitle)

	for _, box := range mainBox.Boxes {
		mechBuilder := createBuilder(box.MechanismArrows)
		inputBuilder := createBuilder(box.InputArrows)
		outputBuilder := createBuilder(box.OutputArrows)

		builder.WriteString(fmt.Sprintf("|%s|%s|%s|%s|%s|\n",
			processNamePattern.replaceProcessName(box.Name),
			processDescPattern.replaceProcessName(box.Name),
			mechBuilder.String(),
			inputBuilder.String(),
			outputBuilder.String()))
	}

	return &builder
}

func createBuilder(arrows arrow.Arrows) *strings.Builder {
	builder := strings.Builder{}

	for _, arrow := range arrows {
		builder.WriteString(fmt.Sprintf("%s ", streamPattern.replaceStreamLabel(arrow.Label)))
	}

	return &builder
}
