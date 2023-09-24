package box

import (
	"ramus/converter/ramustypes/arrow"
)

type Boxes []*Box

func (boxes Boxes) AddArrows(arrows arrow.Arrows) {
	for _, arrow := range arrows {
		putStructs := arrow.IOPutStructsByBoxType()

		for _, putStruct := range putStructs {
			for _, box := range boxes {
				if putStruct.Info.BoxId != box.Id {
					continue
				}

				box.AddArrow(arrow, putStruct.Info.ArrowType)
				break
			}
		}
	}
}

func (boxes Boxes) BoxById(id int) *Box {
	for _, box := range boxes {
		if box.Id == id {
			return box
		}
	}

	return nil
}
