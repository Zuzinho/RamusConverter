package converter

import (
	"io"
	"ramus/converter/ramustypes/arrow"
	"ramus/converter/ramustypes/box"
	"regexp"
)

func Convert(reader io.Reader) (mainBox *box.Box, err error) {
	buf, err := io.ReadAll(reader)
	if err != nil {
		return
	}

	diagReg, _ := regexp.Compile("DIAGRAM GRAPHIC.*;\\s*CREATION DATE.*;\\s*REVISION DATE.*;\\s*TITLE.*;\\s*STATUS.*;")

	boxReg, _ := regexp.Compile("BOX \\d+ ;\\s*NAME.*;\\s*BOX COORDINATES.*;\\s*DETAIL REFERENCE.*;")

	arrowReg, _ := regexp.Compile("ARROWSEG \\d+ ;\\s*SOURCE.*;\\s*PATH.*;\\s*LABEL.*;\\s*.*\\s*SINK.*;")

	sources := diagReg.Split(string(buf), -1)

	for _, srcStr := range sources[1:] {
		boxes := make(box.Boxes, 0)
		arrows := make(arrow.Arrows, 0)
		src := []byte(srcStr)

		boxesSrc := boxReg.FindAll(src, -1)
		for _, boxSrc := range boxesSrc {
			newBox, err := box.NewBox(boxSrc)
			if err != nil {
				return nil, err
			}

			boxes = append(boxes, newBox)
		}

		arrowsSrc := arrowReg.FindAll(src, -1)
		for _, arrowSrc := range arrowsSrc {
			newArrow, err := arrow.NewArrow(arrowSrc)
			if err != nil {
				return nil, err
			}

			arrows = append(arrows, newArrow)
		}

		arrows.ConnectBranches()

		boxes.AddArrows(arrows)

		if len(boxes) == 1 {
			if boxes[0].Id == 0 {
				mainBox = boxes[0]
				continue
			}
		}
		mainBox.AddChildBoxes(boxes...)
	}

	return
}
