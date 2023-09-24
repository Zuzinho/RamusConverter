package box

import (
	"log"
	"ramus/converter/ramustypes/arrow"
	arrowtypes "ramus/converter/ramustypes/arrow/types"
	boxtypes "ramus/converter/ramustypes/box/types"
	"strings"
)

type Box struct {
	Id              int
	Name            string
	Reference       *boxtypes.Reference
	Boxes           Boxes
	InputArrows     arrow.Arrows
	OutputArrows    arrow.Arrows
	ControlArrows   arrow.Arrows
	MechanismArrows arrow.Arrows
}

func NewBox(src []byte) (*Box, error) {
	log.Println(string(src))

	cnv := newConverter(src)

	id, err := cnv.convertBoxId()
	if err != nil {
		return nil, err
	}

	name := cnv.convertName()
	name = strings.ReplaceAll(name, "<CR>", " ")

	reference, err := cnv.convertReference()
	if err != nil {
		return nil, err
	}

	return &Box{
		Id:              id,
		Name:            name,
		Reference:       reference,
		Boxes:           make(Boxes, 0),
		InputArrows:     make(arrow.Arrows, 0),
		OutputArrows:    make(arrow.Arrows, 0),
		ControlArrows:   make(arrow.Arrows, 0),
		MechanismArrows: make(arrow.Arrows, 0),
	}, nil
}

func (box *Box) AddArrow(arr *arrow.Arrow, tp arrowtypes.Type) {
	switch tp {
	case arrowtypes.INPUT:
		box.InputArrows = append(box.InputArrows, arr)
	case arrowtypes.OUTPUT:
		box.OutputArrows = append(box.OutputArrows, arr)
	case arrowtypes.CONTROL:
		box.ControlArrows = append(box.ControlArrows, arr)
	case arrowtypes.MECHANISM:
		box.MechanismArrows = append(box.MechanismArrows, arr)
	}
}

func (box *Box) AddChildBoxes(childBoxes ...*Box) {
	if len(childBoxes) == 0 {
		return
	}

	refIds := childBoxes[0].Reference.ReferenceIds

	parentBox := box

	for _, id := range refIds {
		parentBox = parentBox.Boxes.BoxById(id)
	}

	parentBox.Boxes = append(parentBox.Boxes, childBoxes...)
}

func (box *Box) Process() string {
	if box.Id == 0 {
		return "Процесс"
	}

	return "Подпроцесс"
}
