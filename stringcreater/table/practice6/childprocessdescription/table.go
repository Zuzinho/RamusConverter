package childprocessdescription

import (
	"fmt"
	"log"
	"ramus/converter/ramustypes/arrow"
	"ramus/converter/ramustypes/arrow/types"
	"ramus/converter/ramustypes/box"
	"ramus/stringcreater"
	"strconv"
	"strings"
)

func CreateTable(mainBox *box.Box) *strings.Builder {
	builder := strings.Builder{}

	builder.WriteString(tableTitle)

	var stringFunc stringcreater.StringFuncType
	stringFunc = func(box *box.Box) {
		mechBuilder := createBuilder(box.MechanismArrows)
		inputBuilder := createBuilder(box.InputArrows)
		outputBuilder := createBuilder(box.OutputArrows)

		entryId := sourceBoxId(box.InputArrows[0])
		exitId := sinkBoxId(box.OutputArrows[0])

		entryBuilder := strings.Builder{}
		exitBuilder := strings.Builder{}
		if entryId == -1 {
			entryBuilder.WriteString(entry)
		} else {
			entryBuilder.WriteString(strconv.Itoa(entryId - box.Id))
		}

		if exitId == -1 {
			exitBuilder.WriteString(exit)
		} else {
			exitBuilder.WriteString(strconv.Itoa(exitId - box.Id))
		}

		log.Println(entryBuilder.String())
		log.Println(exitBuilder.String())

		builder.WriteString(fmt.Sprintf("|%s|%s|%s|%s|%s|%s|%s|\n",
			processNamePattern.replaceProcessName(box.Name),
			processDescPattern.replaceProcessName(box.Name),
			mechBuilder.String(),
			inputBuilder.String(),
			entryBuilder.String(),
			outputBuilder.String(),
			exitBuilder.String(),
		))

		for _, childBox := range box.Boxes {
			stringFunc(childBox)
		}
	}

	stringFunc(mainBox)

	return &builder
}

func createBuilder(arrows arrow.Arrows) *strings.Builder {
	builder := strings.Builder{}

	for _, arrow := range arrows {
		builder.WriteString(fmt.Sprintf("%s ", streamPattern.replaceStreamLabel(arrow.Label)))
	}

	return &builder
}

func sourceBoxId(arrow *arrow.Arrow) int {
	if arrow.SourceArrow != nil {
		return sourceBoxId(arrow.SourceArrow)
	}

	if arrow.Source.Type == types.BORDER {
		return -1
	}

	return arrow.Source.Info.BoxId
}

func sinkBoxId(arrow *arrow.Arrow) int {
	if len(arrow.SinkArrows) > 0 {
		return sinkBoxId(arrow.SinkArrows[0])
	}

	if arrow.Sink.Type == types.BORDER {
		return -1
	}

	return arrow.Sink.Info.BoxId
}
