package box

import (
	"ramus/converter/ramustypes/box/types"
	"regexp"
	"strconv"
	"strings"
)

type converter struct {
	source []byte
}

func newConverter(src []byte) converter {
	return converter{
		source: src,
	}
}

func (cnv converter) convertBoxId() (boxId int, err error) {
	boxReg, _ := regexp.Compile("BOX \\d+ ;")
	digReg, _ := regexp.Compile("\\d+")

	line := boxReg.Find(cnv.source)
	boxId, err = strconv.Atoi(string(digReg.Find(line)))

	return
}

func (cnv converter) convertName() string {
	nameLineReg, _ := regexp.Compile("NAME '\\{.*}.*' ;")
	nameReg, _ := regexp.Compile("'\\{.*}.*'")
	nameInfoReg, _ := regexp.Compile("\\{.*}")

	line := nameLineReg.Find(cnv.source)
	nameValue := nameReg.Find(line)
	name := string(nameInfoReg.ReplaceAll(nameValue, []byte("")))

	name = strings.ReplaceAll(name, "'", "")

	return name
}

func (cnv converter) convertReference() (*types.Reference, error) {
	refLineReg, _ := regexp.Compile("DETAIL REFERENCE N [A-Z]\\d+ ;")
	refReg, _ := regexp.Compile("[A-Z]\\d+")

	line := refLineReg.Find(cnv.source)
	reference := refReg.Find(line)

	return types.NewReference(reference)
}
