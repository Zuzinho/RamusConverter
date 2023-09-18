package ramustypes

import (
	"fmt"
	"strings"
)

type IsMainProcess bool

func (is IsMainProcess) Process() string {
	if is {
		return "Процесс"
	}
	return "Подпроцесс"
}

type Box struct {
	Id              int
	Name            string
	Reference       string
	Boxes           []*Box
	InputArrows     []*InputArrow
	OutputArrows    []*OutputArrow
	ControlArrows   []*ControlArrow
	MechanismArrows []*MechanismArrow
}

func NewBox(id int, name string, reference string) *Box {
	name = strings.ReplaceAll(name, "<CR>", " ")
	name = strings.ReplaceAll(name, "'", "")
	return &Box{
		Id:              id,
		Name:            name,
		Reference:       reference,
		InputArrows:     make([]*InputArrow, 0),
		OutputArrows:    make([]*OutputArrow, 0),
		ControlArrows:   make([]*ControlArrow, 0),
		MechanismArrows: make([]*MechanismArrow, 0),
	}
}

func (box *Box) AddArrow(arrow Arrow, tp string) {
	switch tp {
	case "I":
		iArr := InputArrow(arrow)
		box.InputArrows = append(box.InputArrows, &iArr)
	case "O":
		oArr := OutputArrow(arrow)
		box.OutputArrows = append(box.OutputArrows, &oArr)
	case "M":
		mArr := MechanismArrow(arrow)
		box.MechanismArrows = append(box.MechanismArrows, &mArr)
	case "C":
		cArr := ControlArrow(arrow)
		box.ControlArrows = append(box.ControlArrows, &cArr)
	}
}

func (box *Box) AddBox(childBox ...*Box) {
	box.Boxes = append(box.Boxes, childBox...)
}

func (box *Box) ProcessType() string {
	if box.Reference == "A0" {
		return "Процесс"
	}

	return "Подпроцесс"
}

const BoxPrefix = "{LWI I 4 255 255 255}"

func (box *Box) StringAsList() string {
	tp := box.ProcessType()
	controlStr := "Управляющим потоком (потоком Управления) является:"
	mechString := fmt.Sprintf("%s осуществляется (механизм):", tp)
	inputStr := fmt.Sprintf("Входами %sа являются:", tp)
	outputString := fmt.Sprintf("Выходом %sа является:", tp)

	for i, arr := range box.ControlArrows {
		controlStr += " " + Arrow(*arr).StringAsList()
		if i == len(box.ControlArrows)-1 {
			controlStr += "."
		} else {
			controlStr += ","
		}
	}

	for i, arr := range box.MechanismArrows {
		mechString += " " + Arrow(*arr).StringAsList()
		if i == len(box.MechanismArrows)-1 {
			mechString += "."
		} else {
			mechString += ","
		}
	}

	for i, arr := range box.InputArrows {
		inputStr += " " + Arrow(*arr).StringAsList()
		if i == len(box.InputArrows)-1 {
			inputStr += "."
		} else {
			inputStr += ","
		}
	}

	for i, arr := range box.OutputArrows {
		outputString += " " + Arrow(*arr).StringAsList()
		if i == len(box.OutputArrows)-1 {
			outputString += "."
		} else {
			outputString += ","
		}
	}

	return fmt.Sprintf("**%s \"%s\"**:\n\n"+
		"+ %s\n\n"+
		"+ %s\n\n"+
		"+ %s\n\n"+
		"+ %s\n", tp, box.Name, controlStr, mechString, inputStr, outputString)
}

func (box *Box) StringAsNotationElementTable() string {
	var control, mech, input, output strings.Builder

	for i, arrow := range box.InputArrows {
		input.WriteString(fmt.Sprintf("%s I%d\t", arrow.Label, i+1))
	}

	for i, arrow := range box.OutputArrows {
		output.WriteString(fmt.Sprintf("%s O%d\t", arrow.Label, i+1))
	}

	for i, arrow := range box.MechanismArrows {
		mech.WriteString(fmt.Sprintf("%s M%d\t", arrow.Label, i+1))
	}

	for i, arrow := range box.ControlArrows {
		control.WriteString(fmt.Sprintf("%s C%d\t", arrow.Label, i+1))
	}

	return fmt.Sprintf("|%s %s|%s|%s|%s|%s|", box.Name, box.Reference,
		input.String(), output.String(), mech.String(), control.String())
}

type StringFuncType func(*Box)

func (box *Box) StringAsConnectionTypeTable() string {
	builder := strings.Builder{}

	controlMap := make(connectionTypeMap)

	var stringFunc StringFuncType
	stringFunc = func(box *Box) {
		if box == nil {
			return
		}

		for _, arrow := range box.ControlArrows {
			controlMap.addBox(arrow.Label, box, "C")
		}

		for _, childBox := range box.Boxes {
			stringFunc(childBox)
		}
	}

	stringFunc(box)

	ioMap := make(connectionTypeMap)

	stringFunc = func(box *Box) {
		if box == nil {
			return
		}

		for _, arrow := range box.InputArrows {
			ioMap.addBox(arrow.Label, box, "I")
		}
		for _, arrow := range box.OutputArrows {
			ioMap.addBox(arrow.Label, box, "O")
		}

		for _, childBox := range box.Boxes {
			stringFunc(childBox)
		}
	}

	stringFunc(box)

	builder.WriteString(controlMap.String("C").String())
	builder.WriteString(ioMap.String("IO").String())

	return builder.String()
}

func (box *Box) StringAsObjectTypeTable() string {
	builder := strings.Builder{}

	for _, arrow := range box.InputArrows {
		builder.WriteString(fmt.Sprintf("\n|Вход|%s|", arrow.Label))
	}

	for _, arrow := range box.OutputArrows {
		builder.WriteString(fmt.Sprintf("\n|Выход|%s|", arrow.Label))
	}

	inMap := make(connectionTypeMap)
	var stringFunc StringFuncType
	stringFunc = func(box *Box) {
		if box == nil {
			return
		}

		for _, arrow := range box.InputArrows {
			inMap.addBox(arrow.Label, box, "I")
		}
		for _, arrow := range box.OutputArrows {
			inMap.addBox(arrow.Label, box, "O")
		}

		for _, childBox := range box.Boxes {
			stringFunc(childBox)
		}
	}

	stringFunc(box)

	for k, v := range inMap {
		if v.output && v.input {
			builder.WriteString(fmt.Sprintf("\n|Внутренний поток|%s|", k))
		}
	}

	return builder.String()
}

type mapValueType struct {
	boxes  []*Box
	input  bool
	output bool
}

type connectionTypeMap map[string]mapValueType

func (mp connectionTypeMap) addBox(label string, box *Box, tp string) {
	if _, ok := mp[label]; ok {
		v := mp[label]
		v.boxes = append(v.boxes, box)
		mp[label] = v
	} else {
		mp[label] = mapValueType{boxes: []*Box{box}}
	}

	v := mp[label]
	switch tp {
	case "I":
		v.input = true
	case "O":
		v.output = true
	}
	mp[label] = v
}

func (mp connectionTypeMap) String(tp string) *strings.Builder {
	builder := strings.Builder{}

	var connType string
	if tp == "C" {
		connType = "Управление"
	} else tp == "IO" {
		connType = "Выход-вход"
	}

	for k, v := range mp {
		if tp == "C" || v.input && v.output {
			diagsBuilder := strings.Builder{}
			for _, box := range v.boxes {
				diagsBuilder.WriteString(fmt.Sprintf("%s %s ", box.Name, box.Reference))
			}

			builder.WriteString(fmt.Sprintf("\n|%s|%s|%s|", diagsBuilder.String(), k, connType))
		}
	}

	return &builder
}
