package types

import (
	"fmt"
)

type IOPutType string

const (
	BORDER IOPutType = "BORDER"
	BOX    IOPutType = "BOX"
	BRANCH IOPutType = "BRANCH"
)

func NewIOPutType(ioPutType string) (IOPutType, error) {
	switch tp := IOPutType(ioPutType); tp {
	case BORDER, BOX, BRANCH:
		return tp, nil
	default:
		return "", NewUnknownIOPutTypeError(tp)
	}
}

type IOPutInfo struct {
	BoxId     int
	ArrowType Type
	ArrowId   int
}

func NewIOPutInfo(info string) (*IOPutInfo, error) {
	if len(info) != 3 {
		return nil, NewInvalidIOPutInfoError(info)
	}

	boxId := int(info[0]) - 48
	arrowTp := string(info[1])
	arrowId := int(info[2]) - 48

	tp, err := NewArrowType(arrowTp)
	if err != nil {
		return nil, err
	}

	return &IOPutInfo{
		BoxId:     boxId,
		ArrowType: tp,
		ArrowId:   arrowId,
	}, nil
}

func (putInfo *IOPutInfo) String() string {
	return fmt.Sprintf("%s%d", putInfo.ArrowType, putInfo.ArrowId)
}

type IOPutStruct struct {
	Type          IOPutType
	Info          *IOPutInfo
	BranchArrowId int
}

func NewIOPutStruct(tp string, info string, branchArrowId int) (*IOPutStruct, error) {
	putType, err := NewIOPutType(tp)
	if err != nil {
		return nil, err
	}

	if putType == BORDER {
		return &IOPutStruct{Type: putType}, nil
	}

	if putType == BRANCH {
		return &IOPutStruct{Type: putType, BranchArrowId: branchArrowId}, nil
	}

	infoStruct, err := NewIOPutInfo(info)
	if err != nil {
		return nil, err
	}

	return &IOPutStruct{
		Type: putType,
		Info: infoStruct,
	}, nil
}
