package converter

import (
	"io"
	"log"
	"ramus/converter/ramustypes"
	"regexp"
	"strings"
)

func convert(rd io.Reader) ([][]*ramustypes.Box, error) {
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

	mainBoxes := make([][]*ramustypes.Box, 0)

	splitDiags := regDiag.Split(string(buf), -1)
	for _, diag := range splitDiags {
		boxesSrcs := regBox.FindAll([]byte(diag), -1)
		boxes := make([]*ramustypes.Box, 0)
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
						boxes[i] = box.AddArrow(*arrow, info.Type)
						log.Println(info, boxes[i])
					}
				}
			}
		}

		mainBoxes = append(mainBoxes, boxes)
	}

	return mainBoxes, nil
}

func ConvertAsList(rd io.Reader) (*strings.Reader, error) {
	boxes, err := convert(rd)
	if err != nil {
		return nil, err
	}

	builder := strings.Builder{}

	for _, boxArr := range boxes {
		for _, box := range boxArr {
			builder.WriteString(box.StringAsList() + "\n")
		}
	}

	return strings.NewReader(builder.String()), nil
}

func ConvertAsTable(rd io.Reader) (*strings.Reader, error) {
	boxes, err := convert(rd)
	if err != nil {
		return nil, err
	}

	builder := strings.Builder{}

	builder.WriteString("| **Наименование диаграммы/код** | **Вход** | **Выход** | **Механизм** | **Управление** |\n" +
		"|--------------------------------|----------|-----------|--------------|----------------|\n")

	for _, boxArr := range boxes {
		for _, box := range boxArr {
			builder.WriteString(box.StringAsTable() + "\n")
		}
	}

	return strings.NewReader(builder.String()), nil
}
