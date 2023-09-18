package converter

import (
	"fmt"
	"io"
	"ramus/converter/ramustypes"
	"regexp"
	"strings"
)

func convert(rd io.Reader) (*ramustypes.Box, error) {
	buf, err := io.ReadAll(rd)
	if err != nil {
		return nil, err
	}

	regDiag, err := regexp.Compile("DIAGRAM GRAPHIC A-?\\d* ;\\s*CREATION DATE.*;\\s*REVISION DATE.*;\\s*TITLE.*;\\s*STATUS WORKING ;")
	if err != nil {
		return nil, err
	}

	regBox, err := regexp.Compile("BOX \\d* ;\\s*NAME.*;\\s*BOX COORDINATES.*;\\s*DETAIL REFERENCE.*;")
	if err != nil {
		return nil, err
	}

	regArrow, err := regexp.Compile("ARROWSEG \\d* ;\\s*SOURCE.*;\\s*PATH.*;\\s*LABEL.*\\s*.*;\\s*SINK.*;")
	if err != nil {
		return nil, err
	}

	mainBoxes := make([]*ramustypes.Box, 0)

	splitDiags := regDiag.Split(string(buf), -1)
	for _, diag := range splitDiags {
		boxes := make([]*ramustypes.Box, 0)
		boxesSrcs := regBox.FindAll([]byte(diag), -1)
		for _, boxesSrc := range boxesSrcs {
			box, err := convertBox(boxesSrc)
			if err != nil {
				return nil, err
			}

			boxes = append(boxes, box)
		}

		arrowSrcs := regArrow.FindAll([]byte(diag), -1)
		for _, arrowSrc := range arrowSrcs {
			arrow, infos, err := convertArrow(arrowSrc)
			if err != nil {
				return nil, err
			}

			for _, info := range infos {
				for i := range boxes {
					box := boxes[i]
					if box.Id == info.BoxId {
						box.AddArrow(*arrow, info.Type)
					}
				}
			}
		}

		mainBoxes = append(mainBoxes, boxes...)
	}

	mainBox := convertArrToBox(mainBoxes)

	return mainBox, nil
}

func convertArrToBox(boxes []*ramustypes.Box) *ramustypes.Box {
	if len(boxes) == 0 {
		return nil
	}

	mainBox := boxes[0]

	for i := 1; i < len(boxes); i++ {
		box := boxes[i]
		if len(box.Reference) == 2 {
			mainBox.AddBox(box)
			continue
		}

		parentBox := mainBox
		for j := 1; j < len(box.Reference)-1; j++ {
			id := int(box.Reference[j]) - 49
			parentBox = parentBox.Boxes[id]
		}
		parentBox.AddBox(box)
	}

	return mainBox
}

type stringFuncType func(*ramustypes.Box)

func ConvertAsList(rd io.Reader) (*strings.Reader, error) {
	mainBox, err := convert(rd)
	if err != nil {
		return nil, err
	}

	builder := strings.Builder{}

	var stringFunc stringFuncType
	stringFunc = func(box *ramustypes.Box) {
		if box == nil {
			return
		}

		resBuilder := strings.Builder{}
		if len(box.Boxes) > 0 {
			resBuilder.WriteString(fmt.Sprintf("%s \"%s\" предполагает выполнение "+
				"следующих подпроцессов:", box.ProcessType(), box.Name))
			for i, childBox := range box.Boxes {
				resBuilder.WriteString(fmt.Sprintf(" \"%s\"", childBox.Name))
				if i == len(box.Boxes)-1 {
					resBuilder.WriteString(".")
				} else {
					resBuilder.WriteString(",")
				}
			}
		}

		builder.WriteString(box.StringAsList() + "\n\n" + resBuilder.String() + "\n\n")

		for _, childBox := range box.Boxes {
			stringFunc(childBox)
		}
	}

	stringFunc(mainBox)

	return strings.NewReader(builder.String()), nil
}

func ConvertAsTables(rd io.Reader) (*strings.Reader, error) {
	mainBox, err := convert(rd)
	if err != nil {
		return nil, err
	}

	builder := strings.Builder{}

	notElBuilder := convertAsNotationElementTable(mainBox)
	connTpBuilder := convertAsConnectionTypeTable(mainBox)
	tpObjBuilder := convertAsTypeObjectTable(mainBox)

	builder.WriteString(notElBuilder.String() + "\n\n" + connTpBuilder.String() + "\n\n" + tpObjBuilder.String())

	return strings.NewReader(builder.String()), nil
}

func convertAsNotationElementTable(mainBox *ramustypes.Box) *strings.Builder {
	builder := strings.Builder{}

	builder.WriteString("| **Наименование диаграммы/код** | **Вход** | **Выход** | **Механизм** | **Управление** |\n" +
		"|--------------------------------|----------|-----------|--------------|----------------|\n")

	var stringFunc stringFuncType
	stringFunc = func(box *ramustypes.Box) {
		if box == nil {
			return
		}

		builder.WriteString(box.StringAsNotationElementTable() + "\n")

		for _, childBox := range box.Boxes {
			stringFunc(childBox)
		}
	}

	stringFunc(mainBox)

	return &builder
}

type connectionTypeMap map[string][]*ramustypes.Box

func (mp connectionTypeMap) addBox(label string, box *ramustypes.Box) {
	if _, ok := mp[label]; ok {
		mp[label] = append(mp[label], box)
	} else {
		mp[label] = append(mp[label], box)
	}
}

func convertAsConnectionTypeTable(mainBox *ramustypes.Box) *strings.Builder {
	builder := strings.Builder{}

	builder.WriteString("| **Наименование диаграммы/код** | **Наименование потока** | **Тип связи** |\n" +
		"|--------------------------------|-------------------------|---------------|")

	controlMap := make(connectionTypeMap)

	var stringFunc stringFuncType
	stringFunc = func(box *ramustypes.Box) {
		if box == nil {
			return
		}

		for _, arrow := range box.ControlArrows {
			controlMap.addBox(arrow.Label, box)
		}

		for _, childBox := range box.Boxes {
			stringFunc(childBox)
		}
	}

	stringFunc(mainBox)

	ioMap := make(connectionTypeMap)

	stringFunc = func(box *ramustypes.Box) {
		if box == nil {
			return
		}

		for _, arrow := range box.InputArrows {
			ioMap.addBox(arrow.Label, box)
		}
		for _, arrow := range box.OutputArrows {
			ioMap.addBox(arrow.Label, box)
		}

		for _, childBox := range box.Boxes {
			stringFunc(childBox)
		}
	}

	stringFunc(mainBox)

	for k, v := range controlMap {
		diagsBuilder := strings.Builder{}
		for _, box := range v {
			diagsBuilder.WriteString(fmt.Sprintf("%s %s ", box.Name, box.Reference))
		}

		builder.WriteString(fmt.Sprintf("\n|%s|%s|Управление|", diagsBuilder.String(), k))
	}

	for k, v := range ioMap {
		if len(v) < 2 {
			continue
		}

		diagsBuilder := strings.Builder{}
		for _, box := range v {
			diagsBuilder.WriteString(fmt.Sprintf("%s %s ", box.Name, box.Reference))
		}

		builder.WriteString(fmt.Sprintf("\n|%s|%s|Выход-Вход|", diagsBuilder.String(), k))
	}

	return &builder
}

func convertAsTypeObjectTable(mainBox *ramustypes.Box) *strings.Builder {
	builder := strings.Builder{}

	builder.WriteString("| **Элемент нотации IDEF0** | **Наименование преобразуемого объекта** |\n" +
		"|---------------------------|-----------------------------------------|")

	builder.WriteString(mainBox.StringAsObjectTypeTable())

	return &builder
}
