package objecttypes

import (
	"fmt"
	"ramus/converter/ramustypes/arrow"
	"ramus/converter/ramustypes/box"
	"ramus/stringcreater"
	"strings"
)

func CreateTable(mainBox *box.Box) *strings.Builder {
	builder := strings.Builder{}

	builder.WriteString(tableTitle)

	fillBuilderByNotationEl(&builder, mainBox.InputArrows, entry)
	fillBuilderByNotationEl(&builder, mainBox.OutputArrows, exit)

	var stringFunc stringcreater.StringFuncType
	stringFunc = func(box *box.Box) {
		for _, arrow := range box.InputArrows {
			if !arrow.IsInterior() {
				continue
			}

			builder.WriteString(fmt.Sprintf("|%s|%s|\n", interior, streamCellPattern.replaceStreamLabel(arrow.Label)))
		}

		for _, childBox := range box.Boxes {
			stringFunc(childBox)
		}
	}

	stringFunc(mainBox)

	return &builder
}

func fillBuilderByNotationEl(builder *strings.Builder, arrows arrow.Arrows, element notationElement) {
	for _, arrow := range arrows {
		builder.WriteString(fmt.Sprintf("|%s|%s|\n", element,
			streamCellPattern.replaceStreamLabel(arrow.Label)))
	}
}
