package arrow

import (
	"log"
	"ramus/converter/ramustypes/arrow/types"
	"strings"
)

type Arrow struct {
	Id          int
	Source      *types.IOPutStruct
	Label       string
	Sink        *types.IOPutStruct
	SourceArrow *Arrow
}

func NewArrow(src []byte) (*Arrow, error) {
	log.Println(string(src))

	cnv := newConverter(src)

	id, err := cnv.convertArrowId()
	if err != nil {
		return nil, err
	}

	source, err := cnv.convertSourceSink("SOURCE")
	if err != nil {
		return nil, err
	}

	label := cnv.convertLabel()
	if err != nil {
		return nil, err
	}
	label = strings.ReplaceAll(label, "<CR>", " ")

	sink, err := cnv.convertSourceSink("SINK")
	if err != nil {
		return nil, err
	}

	return &Arrow{
		Id:     id,
		Source: source,
		Label:  label,
		Sink:   sink,
	}, nil
}

func (arrow *Arrow) IsInterior() bool {
	if arrow.Source.Type == types.BOX && arrow.Sink.Type == types.BOX {
		return true
	}

	if arrow.Source.Type != types.BRANCH {
		return false
	}

	var isInteriorBranch func(*Arrow) bool
	isInteriorBranch = func(srcArrow *Arrow) bool {
		if srcArrow.SourceArrow == nil {
			return srcArrow.Source.Type == types.BOX
		}

		return isInteriorBranch(srcArrow.SourceArrow)
	}

	return isInteriorBranch(arrow)
}

func (arrow *Arrow) IOPutStructsByBoxType() []*types.IOPutStruct {
	putStructs := make([]*types.IOPutStruct, 0)

	if arrow.Source.Type == types.BOX {
		putStructs = append(putStructs, arrow.Source)
	}
	if arrow.Sink.Type == types.BOX {
		putStructs = append(putStructs, arrow.Sink)
	}

	return putStructs
}
