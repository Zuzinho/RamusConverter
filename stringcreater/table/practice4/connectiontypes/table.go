package connectiontypes

import (
	"fmt"
	"ramus/converter/ramustypes/box"
	"ramus/stringcreater"
	"strings"
)

func CreateTable(mainBox *box.Box) *strings.Builder {
	builder := strings.Builder{}

	builder.WriteString(tableTitle)

	mp := make(map[string]box.Boxes)

	stringFunc := func(tp connectionType) {
		for k, v := range mp {
			cellBuilder := strings.Builder{}

			for i := len(v) - 1; i >= 0; i-- {
				box := v[i]
				cellBuilder.WriteString(processCellPattern.replaceProcessNameReferenceId(box.Name, box.Reference.String(), box.Id) + " ")
			}
			builder.WriteString(fmt.Sprintf("|%s|%s|%s|\n",
				cellBuilder.String(), k, tp))
		}

	}

	var fillFunc stringcreater.StringFuncType
	fillFunc = func(parentBox *box.Box) {
		for i := len(parentBox.Boxes) - 1; i >= 0; i-- {
			fillFunc(parentBox.Boxes[i])
		}

		for _, arrow := range parentBox.ControlArrows {
			boxes, ok := mp[arrow.Label]
			if ok {
				boxes = append(boxes, parentBox)
			} else {
				boxes = box.Boxes{parentBox}
			}

			mp[arrow.Label] = boxes
		}
	}

	fillFunc(mainBox)
	stringFunc(CONTROL)

	mp = make(map[string]box.Boxes)

	fillFunc = func(parentBox *box.Box) {
		for _, arrow := range parentBox.InputArrows {
			if !arrow.IsInterior() {
				continue
			}

			boxes, ok := mp[arrow.Label]
			if ok {
				boxes = append(boxes, parentBox)
			} else {
				boxes = box.Boxes{parentBox}
			}

			mp[arrow.Label] = boxes
		}

		for _, arrow := range parentBox.OutputArrows {
			boxes, ok := mp[arrow.Label]
			if !ok {
				continue
			}

			sinkBox := boxes[0]
			if len(sinkBox.Reference.String()) != len(parentBox.Reference.String()) {
				continue
			}

			boxes = append(boxes, parentBox)
			mp[arrow.Label] = boxes
		}

		for i := len(parentBox.Boxes) - 1; i >= 0; i-- {
			fillFunc(parentBox.Boxes[i])
		}
	}

	fillFunc(mainBox)
	stringFunc(IOPUT)

	return &builder
}
