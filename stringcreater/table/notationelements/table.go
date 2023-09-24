package notationelements

import (
	"fmt"
	"ramus/converter/ramustypes/arrow"
	"ramus/converter/ramustypes/arrow/types"
	"ramus/converter/ramustypes/box"
	"ramus/stringcreater"
	"strings"
)

func CreateTable(mainBox *box.Box) *strings.Builder {
	builder := strings.Builder{}
	builder.WriteString(tableTitle)

	var stringFunc stringcreater.StringFuncType
	stringFunc = func(box *box.Box) {
		processCell := processCellPattern.replaceProcessNameReferenceId(box.Name, box.Reference.String(), box.Id)

		var inputBuilder, outputBuilder, mechBuilder, controlBuilder strings.Builder

		fillStreamBuilder(&inputBuilder, box.InputArrows, box.Id)
		fillStreamBuilder(&outputBuilder, box.OutputArrows, box.Id)
		fillStreamBuilder(&mechBuilder, box.MechanismArrows, box.Id)
		fillStreamBuilder(&controlBuilder, box.ControlArrows, box.Id)

		builder.WriteString(fmt.Sprintf("|%s|%s|%s|%s|%s|\n",
			processCell, inputBuilder.String(), outputBuilder.String(), mechBuilder.String(), controlBuilder.String()))

		for _, childBox := range box.Boxes {
			stringFunc(childBox)
		}
	}

	stringFunc(mainBox)

	return &builder
}

func fillStreamBuilder(builder *strings.Builder, arrows arrow.Arrows, boxId int) {
	for _, arrow := range arrows {
		var info *types.IOPutInfo
		if source := arrow.Source; source.Type == types.BOX {
			if info = source.Info; info.BoxId == boxId {
			}
		}
		if sink := arrow.Sink; sink.Type == types.BOX {
			if info = sink.Info; info.BoxId == boxId {
			}
		}

		builder.WriteString(streamCellPattern.replaceStreamLabelIOPutInfo(arrow.Label, info.String()) + " ")
	}
}
