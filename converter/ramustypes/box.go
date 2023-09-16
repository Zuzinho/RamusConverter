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

func (box *Box) AddBox(childBox *Box) {
	box.Boxes = append(box.Boxes, childBox)
}

const BoxPrefix = "{LWI I 4 255 255 255}"

func (box *Box) StringAsList() string {
	controlStr := "Управляющим потоком (потоком Управления) является:"
	mechString := "Процесс осуществляется (механизм):"
	inputStr := "Входами процесса являются:"
	outputString := "Выходом процесса является:"

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

	return fmt.Sprintf("**Процесс \"%s\"**:\n\n"+
		"+ %s\n\n"+
		"+ %s\n\n"+
		"+ %s\n\n"+
		"+ %s\n", box.Name, controlStr, mechString, inputStr, outputString)
}

func (box *Box) StringAsTable() string {
	var control, mech, input, output strings.Builder

	for i, arrow := range box.InputArrows {
		input.WriteString(fmt.Sprintf("%s I%d\t", arrow.Label, i))
	}

	for i, arrow := range box.OutputArrows {
		output.WriteString(fmt.Sprintf("%s O%d\t", arrow.Label, i))
	}

	for i, arrow := range box.MechanismArrows {
		mech.WriteString(fmt.Sprintf("%s M%d\t", arrow.Label, i))
	}

	for i, arrow := range box.ControlArrows {
		control.WriteString(fmt.Sprintf("%s C%d\t", arrow.Label, i))
	}

	return fmt.Sprintf("|%s %s|%s|%s|%s|%s|", box.Name, box.Reference,
		input.String(), output.String(), mech.String(), control.String())
}
