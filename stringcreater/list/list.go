package list

import (
	"fmt"
	"ramus/converter/ramustypes/arrow"
	"ramus/converter/ramustypes/box"
	"ramus/stringcreater"
	"strings"
)

func CreateList(mainBox *box.Box) *strings.Builder {
	builder := strings.Builder{}

	var stringFunc stringcreater.StringFuncType
	stringFunc = func(box *box.Box) {
		process := box.Process()

		insideBuilder, controlBuilder, mechBuilder, inputBuilder, outputBuilder := getProcessPatterns(process, box.Name)

		fillInsideBuilder(insideBuilder, box.Boxes)
		fillStreamBuilder(controlBuilder, box.ControlArrows)
		fillStreamBuilder(mechBuilder, box.MechanismArrows)
		fillStreamBuilder(inputBuilder, box.InputArrows)
		fillStreamBuilder(outputBuilder, box.OutputArrows)

		if len(box.Boxes) > 0 {
			builder.WriteString(insideBuilder.String() + "\n\n")
		}

		title := processTitle.replaceProcessTypeName(process, box.Name)
		builder.WriteString(fmt.Sprintf("**%s**\n\n+ %s\n\n+ %s\n\n+ %s\n\n+ %s\n\n\n\n",
			title, controlBuilder.String(), mechBuilder.String(), inputBuilder.String(), outputBuilder.String()))

		for _, childBox := range box.Boxes {
			stringFunc(childBox)
		}
	}

	stringFunc(mainBox)

	return &builder
}

func fillStreamBuilder(builder *strings.Builder, arrows arrow.Arrows) string {
	arrowsLen := len(arrows)

	for i, arrow := range arrows {
		builder.WriteString(fmt.Sprintf(" **%s**", insertedStream.replaceStreamLabel(arrow.Label)))
		builder.WriteString(separator(i, arrowsLen))
	}

	return builder.String()
}

func fillInsideBuilder(builder *strings.Builder, boxes box.Boxes) string {
	boxesLen := len(boxes)

	for i, box := range boxes {
		builder.WriteString(fmt.Sprintf(" %s", insertedProcess.replaceProcessName(box.Name)))
		builder.WriteString(separator(i, boxesLen))
	}

	return builder.String()
}

func separator(i, len int) string {
	if i == len-1 {
		return "."
	}

	return ","
}
